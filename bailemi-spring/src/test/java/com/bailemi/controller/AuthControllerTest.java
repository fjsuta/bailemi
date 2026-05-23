package com.bailemi.controller;

import com.bailemi.dto.AuthResponse;
import com.bailemi.dto.LoginRequest;
import com.bailemi.dto.RegisterRequest;
import com.bailemi.service.UserService;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.http.MediaType;
import org.springframework.test.web.servlet.MockMvc;

import java.time.LocalDateTime;

import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.when;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.post;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.*;

@SpringBootTest
@AutoConfigureMockMvc
class AuthControllerTest {

    @Autowired
    private MockMvc mockMvc;

    @Autowired
    private ObjectMapper objectMapper;

    @MockBean
    private UserService userService;

    @Test
    void register_WithValidRequest_ShouldReturnSuccess() throws Exception {
        RegisterRequest request = new RegisterRequest();
        request.setUsername("testuser");
        request.setPassword("password123");
        request.setEmail("test@example.com");

        AuthResponse response = new AuthResponse(
            1L,
            "testuser",
            "test@example.com",
            null,
            null,
            "accessToken",
            "refreshToken",
            86400000L
        );

        when(userService.register(any(RegisterRequest.class))).thenReturn(response);

        mockMvc.perform(post("/api/v1/auth/register")
                .contentType(MediaType.APPLICATION_JSON)
                .content(objectMapper.writeValueAsString(request)))
            .andExpect(status().isOk())
            .andExpect(jsonPath("$.code").value(0))
            .andExpect(jsonPath("$.message").value("注册成功"))
            .andExpect(jsonPath("$.data.username").value("testuser"))
            .andExpect(jsonPath("$.data.accessToken").value("accessToken"));
    }

    @Test
    void register_WithExistingUsername_ShouldReturnBadRequest() throws Exception {
        RegisterRequest request = new RegisterRequest();
        request.setUsername("existinguser");
        request.setPassword("password123");

        when(userService.register(any(RegisterRequest.class)))
            .thenThrow(new RuntimeException("用户名已存在"));

        mockMvc.perform(post("/api/v1/auth/register")
                .contentType(MediaType.APPLICATION_JSON)
                .content(objectMapper.writeValueAsString(request)))
            .andExpect(status().isBadRequest())
            .andExpect(jsonPath("$.code").value(400))
            .andExpect(jsonPath("$.message").value("用户名已存在"));
    }

    @Test
    void register_WithInvalidRequest_ShouldReturnBadRequest() throws Exception {
        RegisterRequest request = new RegisterRequest();
        request.setUsername("");  // Empty username
        request.setPassword("pass");  // Too short password

        mockMvc.perform(post("/api/v1/auth/register")
                .contentType(MediaType.APPLICATION_JSON)
                .content(objectMapper.writeValueAsString(request)))
            .andExpect(status().isBadRequest());
    }

    @Test
    void login_WithValidCredentials_ShouldReturnSuccess() throws Exception {
        LoginRequest request = new LoginRequest();
        request.setAccount("testuser");
        request.setPassword("password123");

        AuthResponse response = new AuthResponse(
            1L,
            "testuser",
            "test@example.com",
            null,
            null,
            "accessToken",
            "refreshToken",
            86400000L
        );

        when(userService.login(any(LoginRequest.class))).thenReturn(response);

        mockMvc.perform(post("/api/v1/auth/login")
                .contentType(MediaType.APPLICATION_JSON)
                .content(objectMapper.writeValueAsString(request)))
            .andExpect(status().isOk())
            .andExpect(jsonPath("$.code").value(0))
            .andExpect(jsonPath("$.message").value("登录成功"))
            .andExpect(jsonPath("$.data.username").value("testuser"))
            .andExpect(jsonPath("$.data.accessToken").value("accessToken"));
    }

    @Test
    void login_WithInvalidCredentials_ShouldReturnUnauthorized() throws Exception {
        LoginRequest request = new LoginRequest();
        request.setAccount("wronguser");
        request.setPassword("wrongpassword");

        when(userService.login(any(LoginRequest.class)))
            .thenThrow(new RuntimeException("用户名或密码错误"));

        mockMvc.perform(post("/api/v1/auth/login")
                .contentType(MediaType.APPLICATION_JSON)
                .content(objectMapper.writeValueAsString(request)))
            .andExpect(status().isBadRequest())
            .andExpect(jsonPath("$.code").value(401))
            .andExpect(jsonPath("$.message").value("用户名或密码错误"));
    }

    @Test
    void login_WithEmptyAccount_ShouldReturnBadRequest() throws Exception {
        LoginRequest request = new LoginRequest();
        request.setAccount("");
        request.setPassword("password123");

        mockMvc.perform(post("/api/v1/auth/login")
                .contentType(MediaType.APPLICATION_JSON)
                .content(objectMapper.writeValueAsString(request)))
            .andExpect(status().isBadRequest());
    }
}