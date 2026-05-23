package com.bailemi.controller;

import com.bailemi.dto.PageResponse;
import com.bailemi.dto.UserQueryRequest;
import com.bailemi.dto.UserResponse;
import com.bailemi.entity.User;
import com.bailemi.service.UserService;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.http.MediaType;
import org.springframework.security.test.context.support.WithMockUser;
import org.springframework.test.web.servlet.MockMvc;

import java.time.LocalDateTime;
import java.util.Arrays;
import java.util.List;

import static org.mockito.ArgumentMatchers.any;
import static org.mockito.ArgumentMatchers.eq;
import static org.mockito.Mockito.when;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.*;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.*;

@SpringBootTest
@AutoConfigureMockMvc
class AdminControllerTest {

    @Autowired
    private MockMvc mockMvc;

    @MockBean
    private UserService userService;

    private UserResponse testUserResponse;

    @BeforeEach
    void setUp() {
        List<String> roles = Arrays.asList("USER");
        testUserResponse = new UserResponse(
            1L,
            "testuser",
            "test@example.com",
            "13800138000",
            null,
            1,
            LocalDateTime.now(),
            LocalDateTime.now(),
            roles
        );
    }

    @Test
    @WithMockUser(roles = "ADMIN")
    void queryUsers_WithAdminRole_ShouldReturnSuccess() throws Exception {
        List<UserResponse> users = Arrays.asList(testUserResponse);
        PageResponse<UserResponse> pageResponse = new PageResponse<>(
            users,
            1L,
            1,
            10,
            1
        );

        when(userService.queryUsers(any(UserQueryRequest.class))).thenReturn(pageResponse);

        mockMvc.perform(get("/api/v1/admin/users")
                .param("page", "1")
                .param("pageSize", "10"))
            .andExpect(status().isOk())
            .andExpect(jsonPath("$.code").value(0))
            .andExpect(jsonPath("$.message").value("success"))
            .andExpect(jsonPath("$.data.items[0].username").value("testuser"));
    }

    @Test
    void queryUsers_WithoutAuthentication_ShouldReturnUnauthorized() throws Exception {
        mockMvc.perform(get("/api/v1/admin/users"))
            .andExpect(status().isUnauthorized());
    }

    @Test
    @WithMockUser(roles = "USER")
    void queryUsers_WithUserRole_ShouldReturnForbidden() throws Exception {
        mockMvc.perform(get("/api/v1/admin/users"))
            .andExpect(status().isForbidden());
    }

    @Test
    @WithMockUser(roles = "ADMIN")
    void getUser_WithValidId_ShouldReturnUser() throws Exception {
        when(userService.getUserById(1L)).thenReturn(testUserResponse);

        mockMvc.perform(get("/api/v1/admin/users/1"))
            .andExpect(status().isOk())
            .andExpect(jsonPath("$.code").value(0))
            .andExpect(jsonPath("$.data.username").value("testuser"));
    }

    @Test
    @WithMockUser(roles = "ADMIN")
    void getUser_WithInvalidId_ShouldReturnBadRequest() throws Exception {
        when(userService.getUserById(999L))
            .thenThrow(new RuntimeException("用户不存在"));

        mockMvc.perform(get("/api/v1/admin/users/999"))
            .andExpect(status().isBadRequest())
            .andExpect(jsonPath("$.code").value(400))
            .andExpect(jsonPath("$.message").value("用户不存在"));
    }

    @Test
    @WithMockUser(roles = "ADMIN")
    void banUser_ShouldReturnSuccess() throws Exception {
        UserResponse bannedUser = new UserResponse(
            1L, "testuser", "test@example.com", null, null, 0,
            LocalDateTime.now(), LocalDateTime.now(), List.of("USER")
        );

        when(userService.updateUserStatus(1L, 0)).thenReturn(bannedUser);

        mockMvc.perform(post("/api/v1/admin/users/1/ban"))
            .andExpect(status().isOk())
            .andExpect(jsonPath("$.code").value(0))
            .andExpect(jsonPath("$.message").value("用户已封禁"));
    }

    @Test
    @WithMockUser(roles = "ADMIN")
    void unbanUser_ShouldReturnSuccess() throws Exception {
        when(userService.updateUserStatus(1L, 1)).thenReturn(testUserResponse);

        mockMvc.perform(post("/api/v1/admin/users/1/unban"))
            .andExpect(status().isOk())
            .andExpect(jsonPath("$.code").value(0))
            .andExpect(jsonPath("$.message").value("用户已解封"));
    }

    @Test
    @WithMockUser(roles = "ADMIN")
    void deleteUser_ShouldReturnSuccess() throws Exception {
        mockMvc.perform(delete("/api/v1/admin/users/1"))
            .andExpect(status().isOk())
            .andExpect(jsonPath("$.code").value(0))
            .andExpect(jsonPath("$.message").value("用户已删除"));
    }
}