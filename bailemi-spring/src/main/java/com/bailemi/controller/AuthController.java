package com.bailemi.controller;

import com.bailemi.dto.*;
import com.bailemi.service.UserService;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.media.Content;
import io.swagger.v3.oas.annotations.media.Schema;
import io.swagger.v3.oas.annotations.responses.ApiResponse;
import io.swagger.v3.oas.annotations.responses.ApiResponses;
import io.swagger.v3.oas.annotations.tags.Tag;
import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.security.core.annotation.AuthenticationPrincipal;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.web.bind.annotation.*;

import java.util.Map;

@RestController
@RequestMapping("/api/v1")
@RequiredArgsConstructor
@Tag(name = "认证管理", description = "用户注册、登录相关接口")
public class AuthController {

    private final UserService userService;

    @PostMapping("/auth/register")
    @Operation(summary = "用户注册", description = "新用户注册接口，支持用户名、邮箱、手机号注册")
    @ApiResponses(value = {
        @ApiResponse(responseCode = "200", description = "注册成功",
                content = @Content(schema = @Schema(implementation = AuthResponse.class))),
        @ApiResponse(responseCode = "400", description = "注册失败，用户名或邮箱已存在")
    })
    public ResponseEntity<?> register(@Valid @RequestBody RegisterRequest request) {
        try {
            AuthResponse response = userService.register(request);
            return ResponseEntity.ok(Map.of(
                "code", 0,
                "message", "注册成功",
                "data", response
            ));
        } catch (Exception e) {
            return ResponseEntity.badRequest().body(Map.of(
                "code", 400,
                "message", e.getMessage()
            ));
        }
    }

    @PostMapping("/auth/login")
    @Operation(summary = "用户登录", description = "支持用户名、邮箱或手机号登录")
    @ApiResponses(value = {
        @ApiResponse(responseCode = "200", description = "登录成功",
                content = @Content(schema = @Schema(implementation = AuthResponse.class))),
        @ApiResponse(responseCode = "401", description = "登录失败，账号或密码错误")
    })
    public ResponseEntity<?> login(@Valid @RequestBody LoginRequest request) {
        try {
            AuthResponse response = userService.login(request);
            return ResponseEntity.ok(Map.of(
                "code", 0,
                "message", "登录成功",
                "data", response
            ));
        } catch (Exception e) {
            return ResponseEntity.badRequest().body(Map.of(
                "code", 401,
                "message", e.getMessage()
            ));
        }
    }

    @GetMapping("/user/me")
    @PreAuthorize("isAuthenticated()")
    @Operation(summary = "获取当前用户信息", description = "获取已登录用户的详细信息")
    @ApiResponses(value = {
        @ApiResponse(responseCode = "200", description = "获取成功",
                content = @Content(schema = @Schema(implementation = UserResponse.class))),
        @ApiResponse(responseCode = "401", description = "未认证")
    })
    public ResponseEntity<?> getCurrentUser(@AuthenticationPrincipal UserDetails userDetails) {
        try {
            UserResponse response = userService.getCurrentUser(userDetails.getUsername());
            return ResponseEntity.ok(Map.of(
                "code", 0,
                "message", "success",
                "data", response
            ));
        } catch (Exception e) {
            return ResponseEntity.badRequest().body(Map.of(
                "code", 400,
                "message", e.getMessage()
            ));
        }
    }
}