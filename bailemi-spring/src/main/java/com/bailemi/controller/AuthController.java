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
import lombok.extern.slf4j.Slf4j;
import org.springframework.http.ResponseEntity;
import org.springframework.security.core.annotation.AuthenticationPrincipal;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.web.bind.annotation.*;

import java.util.Map;

@Slf4j
@RestController
@RequestMapping("/api/v1/auth")
@RequiredArgsConstructor
@Tag(name = "认证管理", description = "用户注册、登录相关接口")
public class AuthController {

    private final UserService userService;

    @PostMapping("/register")
    @Operation(
        summary = "用户注册", 
        description = "新用户注册接口，支持用户名和邮箱注册。注册成功后自动登录并返回 JWT Token。"
    )
    @ApiResponses(value = {
        @ApiResponse(
            responseCode = "200", 
            description = "注册成功",
            content = @Content(schema = @Schema(implementation = AuthResponse.class))
        ),
        @ApiResponse(responseCode = "400", description = "注册失败：用户名已存在、邮箱已被注册、验证码错误"),
        @ApiResponse(responseCode = "422", description = "参数验证失败")
    })
    public ResponseEntity<?> register(@Valid @RequestBody RegisterRequest request) {
        try {
            AuthResponse response = userService.register(request);
            
            log.info("用户注册成功: username={}, email={}", 
                response.getUsername(), response.getEmail());
            
            return ResponseEntity.ok(Map.of(
                "code", 0,
                "message", "注册成功",
                "data", response
            ));
        } catch (IllegalArgumentException e) {
            log.warn("用户注册失败: {}", e.getMessage());
            return ResponseEntity.badRequest().body(Map.of(
                "code", 400,
                "message", e.getMessage()
            ));
        } catch (Exception e) {
            log.error("注册异常: ", e);
            return ResponseEntity.badRequest().body(Map.of(
                "code", 400,
                "message", "注册失败，请稍后重试"
            ));
        }
    }

    @PostMapping("/login")
    @Operation(
        summary = "用户登录", 
        description = "支持用户名、邮箱或手机号登录。登录成功后返回 JWT Token 和用户信息。"
    )
    @ApiResponses(value = {
        @ApiResponse(
            responseCode = "200", 
            description = "登录成功",
            content = @Content(schema = @Schema(implementation = AuthResponse.class))
        ),
        @ApiResponse(responseCode = "401", description = "登录失败：用户名或密码错误"),
        @ApiResponse(responseCode = "403", description = "账号已被封禁"),
        @ApiResponse(responseCode = "422", description = "参数验证失败")
    })
    public ResponseEntity<?> login(@Valid @RequestBody LoginRequest request) {
        try {
            AuthResponse response = userService.login(request);
            
            log.info("用户登录成功: username={}", response.getUsername());
            
            return ResponseEntity.ok(Map.of(
                "code", 0,
                "message", "登录成功",
                "data", response
            ));
        } catch (SecurityException e) {
            log.warn("登录失败 - 账号被封禁: {}", e.getMessage());
            return ResponseEntity.status(403).body(Map.of(
                "code", 403,
                "message", e.getMessage()
            ));
        } catch (IllegalArgumentException e) {
            log.warn("登录失败: {}", e.getMessage());
            return ResponseEntity.status(401).body(Map.of(
                "code", 401,
                "message", "用户名或密码错误"
            ));
        } catch (Exception e) {
            log.error("登录异常: ", e);
            return ResponseEntity.status(401).body(Map.of(
                "code", 401,
                "message", "用户名或密码错误"
            ));
        }
    }

    @GetMapping("/user/me")
    @Operation(
        summary = "获取当前用户信息", 
        description = "获取已登录用户的详细信息，需要提供有效的 JWT Token"
    )
    @ApiResponses(value = {
        @ApiResponse(
            responseCode = "200", 
            description = "获取成功",
            content = @Content(schema = @Schema(implementation = UserResponse.class))
        ),
        @ApiResponse(responseCode = "401", description = "未认证或 Token 无效"),
        @ApiResponse(responseCode = "404", description = "用户不存在")
    })
    public ResponseEntity<?> getCurrentUser(@AuthenticationPrincipal UserDetails userDetails) {
        try {
            if (userDetails == null) {
                return ResponseEntity.status(401).body(Map.of(
                    "code", 401,
                    "message", "请先登录"
                ));
            }
            
            UserResponse response = userService.getCurrentUser(userDetails.getUsername());
            
            return ResponseEntity.ok(Map.of(
                "code", 0,
                "message", "success",
                "data", response
            ));
        } catch (Exception e) {
            log.error("获取用户信息异常: ", e);
            return ResponseEntity.badRequest().body(Map.of(
                "code", 400,
                "message", "获取用户信息失败"
            ));
        }
    }

    @PostMapping("/logout")
    @Operation(
        summary = "用户登出", 
        description = "用户登出接口，前端清除 Token 即可，后端记录登出日志"
    )
    @ApiResponses(value = {
        @ApiResponse(responseCode = "200", description = "登出成功"),
        @ApiResponse(responseCode = "401", description = "未认证")
    })
    public ResponseEntity<?> logout(@AuthenticationPrincipal UserDetails userDetails) {
        if (userDetails != null) {
            log.info("用户登出: username={}", userDetails.getUsername());
        }
        
        return ResponseEntity.ok(Map.of(
            "code", 0,
            "message", "登出成功"
        ));
    }

    @GetMapping("/oauth/config")
    @Operation(
        summary = "获取OAuth配置", 
        description = "获取当前支持的第三方登录提供商列表"
    )
    @ApiResponses(value = {
        @ApiResponse(responseCode = "200", description = "获取成功")
    })
    public ResponseEntity<?> getOAuthConfig() {
        return ResponseEntity.ok(Map.of(
            "code", 0,
            "message", "success",
            "data", Map.of(
                "enabled_providers", java.util.List.of("google", "microsoft", "apple", "wechat"),
                "oauth_available", true
            )
        ));
    }
}