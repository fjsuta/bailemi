package com.bailemi.config;

import com.bailemi.entity.Permission;
import com.bailemi.entity.Role;
import com.bailemi.entity.User;
import com.bailemi.repository.PermissionRepository;
import com.bailemi.repository.RoleRepository;
import com.bailemi.repository.UserRepository;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.boot.CommandLineRunner;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Component;

@Component
@RequiredArgsConstructor
@Slf4j
public class DataInitializer implements CommandLineRunner {

    private final UserRepository userRepository;
    private final RoleRepository roleRepository;
    private final PermissionRepository permissionRepository;
    private final PasswordEncoder passwordEncoder;

    @Override
    public void run(String... args) {
        initializeRoles();
        initializeAdminUser();
    }

    private void initializeRoles() {
        if (roleRepository.findByName("ADMIN").isEmpty()) {
            Role adminRole = new Role();
            adminRole.setName("ADMIN");
            adminRole.setDescription("系统管理员");
            roleRepository.save(adminRole);
            log.info("Created ADMIN role");
        }

        if (roleRepository.findByName("USER").isEmpty()) {
            Role userRole = new Role();
            userRole.setName("USER");
            userRole.setDescription("普通用户");
            roleRepository.save(userRole);
            log.info("Created USER role");
        }

        if (roleRepository.findByName("VIP").isEmpty()) {
            Role vipRole = new Role();
            vipRole.setName("VIP");
            vipRole.setDescription("VIP会员");
            roleRepository.save(vipRole);
            log.info("Created VIP role");
        }
    }

    private void initializeAdminUser() {
        if (userRepository.findByUsername("admin").isEmpty()) {
            User admin = new User();
            admin.setUsername("admin");
            admin.setPassword(passwordEncoder.encode("admin123"));
            admin.setEmail("admin@bailemi.com");
            admin.setStatus(1);
            admin.setDeleted(0);

            Role adminRole = roleRepository.findByName("ADMIN").get();
            admin.getRoles().add(adminRole);

            userRepository.save(admin);
            log.info("Created default admin user: admin / admin123");
        }
    }
}
