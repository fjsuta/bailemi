package com.bailemi.service;

import com.bailemi.dto.AuthResponse;
import com.bailemi.dto.LoginRequest;
import com.bailemi.dto.RegisterRequest;
import com.bailemi.entity.Role;
import com.bailemi.entity.User;
import com.bailemi.repository.RoleRepository;
import com.bailemi.repository.UserRepository;
import com.bailemi.security.JwtTokenProvider;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.crypto.password.PasswordEncoder;

import java.util.Optional;
import java.util.Set;

import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.ArgumentMatchers.anyString;
import static org.mockito.Mockito.*;

@ExtendWith(MockitoExtension.class)
class UserServiceTest {

    @Mock
    private UserRepository userRepository;

    @Mock
    private RoleRepository roleRepository;

    @Mock
    private PasswordEncoder passwordEncoder;

    @Mock
    private JwtTokenProvider jwtTokenProvider;

    @Mock
    private AuthenticationManager authenticationManager;

    @InjectMocks
    private UserService userService;

    private User testUser;
    private Role testRole;

    @BeforeEach
    void setUp() {
        testRole = new Role();
        testRole.setId(1L);
        testRole.setName("USER");
        testRole.setDescription("普通用户");

        testUser = new User();
        testUser.setId(1L);
        testUser.setUsername("testuser");
        testUser.setPassword("encodedPassword");
        testUser.setEmail("test@example.com");
        testUser.setPhone("13800138000");
        testUser.setStatus(1);
        testUser.setDeleted(0);
        testUser.setRoles(Set.of(testRole));
    }

    @Test
    void register_WithValidRequest_ShouldCreateUser() {
        RegisterRequest request = new RegisterRequest();
        request.setUsername("newuser");
        request.setPassword("password123");
        request.setEmail("new@example.com");

        when(userRepository.existsByUsername("newuser")).thenReturn(false);
        when(userRepository.existsByEmail("new@example.com")).thenReturn(false);
        when(passwordEncoder.encode("password123")).thenReturn("encodedPassword");
        when(roleRepository.findByName("USER")).thenReturn(Optional.of(testRole));
        when(userRepository.save(any(User.class))).thenAnswer(invocation -> {
            User user = invocation.getArgument(0);
            user.setId(2L);
            return user;
        });
        when(jwtTokenProvider.generateAccessToken(any())).thenReturn("accessToken");
        when(jwtTokenProvider.generateRefreshToken(any())).thenReturn("refreshToken");
        when(jwtTokenProvider.getAccessTokenExpiration()).thenReturn(86400000L);

        AuthResponse response = userService.register(request);

        assertNotNull(response);
        assertEquals("newuser", response.getUsername());
        assertEquals("new@example.com", response.getEmail());
        assertEquals("accessToken", response.getAccessToken());
        assertEquals("refreshToken", response.getRefreshToken());

        verify(userRepository).save(any(User.class));
    }

    @Test
    void register_WithExistingUsername_ShouldThrowException() {
        RegisterRequest request = new RegisterRequest();
        request.setUsername("existinguser");
        request.setPassword("password123");

        when(userRepository.existsByUsername("existinguser")).thenReturn(true);

        RuntimeException exception = assertThrows(RuntimeException.class, 
            () -> userService.register(request));

        assertEquals("用户名已存在", exception.getMessage());
        verify(userRepository, never()).save(any(User.class));
    }

    @Test
    void register_WithExistingEmail_ShouldThrowException() {
        RegisterRequest request = new RegisterRequest();
        request.setUsername("newuser");
        request.setPassword("password123");
        request.setEmail("existing@example.com");

        when(userRepository.existsByUsername("newuser")).thenReturn(false);
        when(userRepository.existsByEmail("existing@example.com")).thenReturn(true);

        RuntimeException exception = assertThrows(RuntimeException.class, 
            () -> userService.register(request));

        assertEquals("邮箱已被注册", exception.getMessage());
        verify(userRepository, never()).save(any(User.class));
    }

    @Test
    void login_WithValidCredentials_ShouldReturnAuthResponse() {
        LoginRequest request = new LoginRequest();
        request.setAccount("testuser");
        request.setPassword("password123");

        Authentication authentication = mock(Authentication.class);
        when(authentication.getPrincipal()).thenReturn(testUser);
        when(authenticationManager.authenticate(any(UsernamePasswordAuthenticationToken.class)))
            .thenReturn(authentication);
        when(jwtTokenProvider.generateAccessToken(any())).thenReturn("accessToken");
        when(jwtTokenProvider.generateRefreshToken(any())).thenReturn("refreshToken");
        when(jwtTokenProvider.getAccessTokenExpiration()).thenReturn(86400000L);
        when(userRepository.save(any(User.class))).thenReturn(testUser);

        AuthResponse response = userService.login(request);

        assertNotNull(response);
        assertEquals("testuser", response.getUsername());
        assertEquals("accessToken", response.getAccessToken());
        assertEquals("refreshToken", response.getRefreshToken());
    }

    @Test
    void login_WithBannedUser_ShouldThrowException() {
        testUser.setStatus(0);
        
        LoginRequest request = new LoginRequest();
        request.setAccount("banneduser");
        request.setPassword("password123");

        Authentication authentication = mock(Authentication.class);
        when(authentication.getPrincipal()).thenReturn(testUser);
        when(authenticationManager.authenticate(any(UsernamePasswordAuthenticationToken.class)))
            .thenReturn(authentication);

        RuntimeException exception = assertThrows(RuntimeException.class, 
            () -> userService.login(request));

        assertEquals("账号已被封禁", exception.getMessage());
    }

    @Test
    void updateUserStatus_WithValidId_ShouldUpdateStatus() {
        when(userRepository.findById(1L)).thenReturn(Optional.of(testUser));
        when(userRepository.save(any(User.class))).thenReturn(testUser);

        var response = userService.updateUserStatus(1L, 0);

        assertNotNull(response);
        verify(userRepository).save(any(User.class));
    }

    @Test
    void updateUserStatus_WithInvalidId_ShouldThrowException() {
        when(userRepository.findById(999L)).thenReturn(Optional.empty());

        RuntimeException exception = assertThrows(RuntimeException.class, 
            () -> userService.updateUserStatus(999L, 0));

        assertEquals("用户不存在", exception.getMessage());
    }

    @Test
    void getUserById_WithValidId_ShouldReturnUser() {
        when(userRepository.findById(1L)).thenReturn(Optional.of(testUser));

        var response = userService.getUserById(1L);

        assertNotNull(response);
        assertEquals("testuser", response.getUsername());
        assertEquals("test@example.com", response.getEmail());
    }

    @Test
    void deleteUser_WithValidId_ShouldSoftDeleteUser() {
        when(userRepository.findById(1L)).thenReturn(Optional.of(testUser));
        when(userRepository.save(any(User.class))).thenReturn(testUser);

        userService.deleteUser(1L);

        assertEquals(1, testUser.getDeleted());
        verify(userRepository).save(testUser);
    }
}