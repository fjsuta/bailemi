package com.bailemi.service;

import com.bailemi.dto.*;
import com.bailemi.entity.Role;
import com.bailemi.entity.User;
import com.bailemi.repository.RoleRepository;
import com.bailemi.repository.UserRepository;
import com.bailemi.security.JwtTokenProvider;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.time.LocalDateTime;
import java.util.List;
import java.util.stream.Collectors;

@Slf4j
@Service
@RequiredArgsConstructor
public class UserService implements UserDetailsService {

    private final UserRepository userRepository;
    private final RoleRepository roleRepository;
    private final PasswordEncoder passwordEncoder;
    private final JwtTokenProvider jwtTokenProvider;
    private final AuthenticationManager authenticationManager;

    @Override
    public UserDetails loadUserByUsername(String username) throws UsernameNotFoundException {
        User user = userRepository.findByUsername(username)
                .orElseThrow(() -> new UsernameNotFoundException("用户不存在: " + username));
        
        if (user.getDeleted() == 1) {
            throw new UsernameNotFoundException("用户已被删除: " + username);
        }
        
        return user;
    }

    @Transactional
    public AuthResponse register(RegisterRequest request) {
        if (userRepository.existsByUsername(request.getUsername())) {
            throw new RuntimeException("用户名已存在");
        }

        if (request.getEmail() != null && userRepository.existsByEmail(request.getEmail())) {
            throw new RuntimeException("邮箱已被注册");
        }

        if (request.getPhone() != null && userRepository.existsByPhone(request.getPhone())) {
            throw new RuntimeException("手机号已被注册");
        }

        if (request.getVerifyCode() != null && !request.getVerifyCode().isEmpty()) {
            if (!validateVerifyCode(request.getVerifyCode(), request.getUsername())) {
                throw new RuntimeException("验证码错误或已过期");
            }
        }

        User user = new User();
        user.setUsername(request.getUsername());
        user.setPassword(passwordEncoder.encode(request.getPassword()));
        user.setEmail(request.getEmail());
        user.setPhone(request.getPhone());
        user.setStatus(1);
        user.setDeleted(0);

        Role userRole = roleRepository.findByName("USER")
                .orElseGet(() -> {
                    Role role = new Role();
                    role.setName("USER");
                    role.setDescription("普通用户");
                    return roleRepository.save(role);
                });

        user.getRoles().add(userRole);
        userRepository.save(user);

        log.info("新用户注册: username={}, email={}", user.getUsername(), user.getEmail());

        String accessToken = jwtTokenProvider.generateAccessToken(user);
        String refreshToken = jwtTokenProvider.generateRefreshToken(user);

        return new AuthResponse(
                user.getId(),
                user.getUsername(),
                user.getEmail(),
                user.getPhone(),
                user.getAvatarUrl(),
                accessToken,
                refreshToken,
                jwtTokenProvider.getAccessTokenExpiration()
        );
    }

    public AuthResponse login(LoginRequest request) {
        Authentication authentication = authenticationManager.authenticate(
                new UsernamePasswordAuthenticationToken(request.getAccount(), request.getPassword())
        );

        User user = (User) authentication.getPrincipal();

        if (user.getStatus() == 0) {
            log.warn("登录失败: 账号已被封禁, account={}", request.getAccount());
            throw new RuntimeException("账号已被封禁");
        }

        if (user.getDeleted() == 1) {
            log.warn("登录失败: 账号已被删除, account={}", request.getAccount());
            throw new RuntimeException("账号已被删除");
        }

        String accessToken = jwtTokenProvider.generateAccessToken(user);
        String refreshToken = jwtTokenProvider.generateRefreshToken(user);

        user.setLastLoginIp(getClientIp());
        user.setLastLoginAt(LocalDateTime.now());
        userRepository.save(user);

        log.info("用户登录成功: username={}, ip={}", user.getUsername(), user.getLastLoginIp());

        return new AuthResponse(
                user.getId(),
                user.getUsername(),
                user.getEmail(),
                user.getPhone(),
                user.getAvatarUrl(),
                accessToken,
                refreshToken,
                jwtTokenProvider.getAccessTokenExpiration()
        );
    }

    public UserResponse getCurrentUser(String username) {
        User user = userRepository.findByUsername(username)
                .orElseThrow(() -> new UsernameNotFoundException("用户不存在"));

        return convertToResponse(user);
    }

    public PageResponse<UserResponse> queryUsers(UserQueryRequest request) {
        if (request.getPageSize() > 100) {
            request.setPageSize(100);
        }

        LocalDateTime startDate = request.getStartDate() != null ? 
                LocalDateTime.parse(request.getStartDate() + "T00:00:00") : null;
        LocalDateTime endDate = request.getEndDate() != null ? 
                LocalDateTime.parse(request.getEndDate() + "T23:59:59") : null;

        var pageable = org.springframework.data.domain.PageRequest.of(
                request.getPage() - 1, 
                request.getPageSize()
        );

        var page = userRepository.findUsersWithFilters(
                request.getUsername(),
                startDate,
                endDate,
                pageable
        );

        List<UserResponse> users = page.getContent().stream()
                .map(this::convertToResponse)
                .collect(Collectors.toList());

        return new PageResponse<>(
                users,
                page.getTotalElements(),
                request.getPage(),
                request.getPageSize(),
                page.getTotalPages()
        );
    }

    @Transactional
    public UserResponse updateUserStatus(Long userId, Integer status) {
        User user = userRepository.findById(userId)
                .orElseThrow(() -> new RuntimeException("用户不存在"));

        String action = status == 1 ? "解封" : "封禁";
        log.info("管理员操作: {}用户, userId={}", action, userId);

        user.setStatus(status);
        userRepository.save(user);

        return convertToResponse(user);
    }

    public UserResponse getUserById(Long userId) {
        User user = userRepository.findById(userId)
                .orElseThrow(() -> new RuntimeException("用户不存在"));

        return convertToResponse(user);
    }

    @Transactional
    public UserResponse updateUser(Long userId, com.bailemi.dto.UserUpdateRequest request) {
        User user = userRepository.findById(userId)
                .orElseThrow(() -> new RuntimeException("用户不存在"));

        if (request.getEmail() != null && !request.getEmail().equals(user.getEmail())) {
            if (userRepository.existsByEmail(request.getEmail())) {
                throw new RuntimeException("邮箱已被使用");
            }
            user.setEmail(request.getEmail());
        }

        if (request.getPhone() != null && !request.getPhone().equals(user.getPhone())) {
            if (userRepository.existsByPhone(request.getPhone())) {
                throw new RuntimeException("手机号已被使用");
            }
            user.setPhone(request.getPhone());
        }

        if (request.getAvatarUrl() != null) {
            user.setAvatarUrl(request.getAvatarUrl());
        }

        userRepository.save(user);
        log.info("管理员操作: 更新用户信息, userId={}", userId);
        return convertToResponse(user);
    }

    @Transactional
    public UserResponse resetUserPassword(Long userId, String newPassword) {
        User user = userRepository.findById(userId)
                .orElseThrow(() -> new RuntimeException("用户不存在"));

        user.setPassword(passwordEncoder.encode(newPassword));
        userRepository.save(user);
        
        log.warn("管理员操作: 重置用户密码, userId={}", userId);
        return convertToResponse(user);
    }

    @Transactional
    public void deleteUser(Long userId) {
        User user = userRepository.findById(userId)
                .orElseThrow(() -> new RuntimeException("用户不存在"));

        log.warn("管理员操作: 删除用户(软删除), userId={}, username={}", userId, user.getUsername());
        user.setDeleted(1);
        userRepository.save(user);
    }

    private boolean validateVerifyCode(String verifyCode, String identifier) {
        return false;
    }

    private String getClientIp() {
        return "127.0.0.1";
    }

    private UserResponse convertToResponse(User user) {
        List<String> roles = user.getRoles().stream()
                .map(Role::getName)
                .collect(Collectors.toList());

        return new UserResponse(
                user.getId(),
                user.getUsername(),
                user.getEmail(),
                user.getPhone(),
                user.getAvatarUrl(),
                user.getStatus(),
                user.getCreatedAt(),
                user.getLastLoginAt(),
                roles
        );
    }
}