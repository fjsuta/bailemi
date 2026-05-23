package com.bailemi.controller;

import com.bailemi.dto.PageResponse;
import com.bailemi.dto.UserQueryRequest;
import com.bailemi.dto.UserResponse;
import com.bailemi.dto.UserStatusRequest;
import com.bailemi.service.UserService;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.Parameter;
import io.swagger.v3.oas.annotations.media.Content;
import io.swagger.v3.oas.annotations.media.Schema;
import io.swagger.v3.oas.annotations.responses.ApiResponse;
import io.swagger.v3.oas.annotations.responses.ApiResponses;
import io.swagger.v3.oas.annotations.tags.Tag;
import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.web.bind.annotation.*;

import java.util.Map;

@RestController
@RequestMapping("/api/v1/admin")
@RequiredArgsConstructor
@Tag(name = "管理员用户管理", description = "管理员用户管理相关接口，需要 ADMIN 权限")
public class AdminController {

    private final UserService userService;

    @GetMapping("/users")
    @PreAuthorize("hasRole('ADMIN')")
    @Operation(summary = "查询用户列表", description = "分页查询用户列表，支持按用户名和注册时间筛选")
    @ApiResponses(value = {
        @ApiResponse(responseCode = "200", description = "查询成功",
                content = @Content(schema = @Schema(implementation = PageResponse.class))),
        @ApiResponse(responseCode = "403", description = "权限不足")
    })
    public ResponseEntity<?> queryUsers(UserQueryRequest request) {
        try {
            PageResponse<UserResponse> response = userService.queryUsers(request);
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

    @PutMapping("/users/{userId}/status")
    @PreAuthorize("hasRole('ADMIN')")
    @Operation(summary = "更新用户状态", description = "更新用户账号状态（正常/封禁）")
    @ApiResponses(value = {
        @ApiResponse(responseCode = "200", description = "更新成功",
                content = @Content(schema = @Schema(implementation = UserResponse.class))),
        @ApiResponse(responseCode = "400", description = "更新失败")
    })
    public ResponseEntity<?> updateUserStatus(
            @Parameter(description = "用户ID") @PathVariable Long userId,
            @Valid @RequestBody UserStatusRequest request) {
        try {
            UserResponse response = userService.updateUserStatus(userId, request.getStatus());
            String message = request.getStatus() == 1 ? "用户已解封" : "用户已封禁";
            return ResponseEntity.ok(Map.of(
                "code", 0,
                "message", message,
                "data", response
            ));
        } catch (Exception e) {
            return ResponseEntity.badRequest().body(Map.of(
                "code", 400,
                "message", e.getMessage()
            ));
        }
    }

    @PostMapping("/users/{userId}/ban")
    @PreAuthorize("hasRole('ADMIN')")
    @Operation(summary = "封禁用户", description = "封禁指定用户账号")
    @ApiResponses(value = {
        @ApiResponse(responseCode = "200", description = "封禁成功",
                content = @Content(schema = @Schema(implementation = UserResponse.class))),
        @ApiResponse(responseCode = "400", description = "封禁失败")
    })
    public ResponseEntity<?> banUser(
            @Parameter(description = "用户ID") @PathVariable Long userId) {
        try {
            UserResponse response = userService.updateUserStatus(userId, 0);
            return ResponseEntity.ok(Map.of(
                "code", 0,
                "message", "用户已封禁",
                "data", response
            ));
        } catch (Exception e) {
            return ResponseEntity.badRequest().body(Map.of(
                "code", 400,
                "message", e.getMessage()
            ));
        }
    }

    @PostMapping("/users/{userId}/unban")
    @PreAuthorize("hasRole('ADMIN')")
    @Operation(summary = "解封用户", description = "解封指定用户账号")
    @ApiResponses(value = {
        @ApiResponse(responseCode = "200", description = "解封成功",
                content = @Content(schema = @Schema(implementation = UserResponse.class))),
        @ApiResponse(responseCode = "400", description = "解封失败")
    })
    public ResponseEntity<?> unbanUser(
            @Parameter(description = "用户ID") @PathVariable Long userId) {
        try {
            UserResponse response = userService.updateUserStatus(userId, 1);
            return ResponseEntity.ok(Map.of(
                "code", 0,
                "message", "用户已解封",
                "data", response
            ));
        } catch (Exception e) {
            return ResponseEntity.badRequest().body(Map.of(
                "code", 400,
                "message", e.getMessage()
            ));
        }
    }

    @GetMapping("/users/{userId}")
    @PreAuthorize("hasRole('ADMIN')")
    @Operation(summary = "获取用户详情", description = "获取指定用户的详细信息")
    @ApiResponses(value = {
        @ApiResponse(responseCode = "200", description = "获取成功",
                content = @Content(schema = @Schema(implementation = UserResponse.class))),
        @ApiResponse(responseCode = "400", description = "用户不存在")
    })
    public ResponseEntity<?> getUser(
            @Parameter(description = "用户ID") @PathVariable Long userId) {
        try {
            UserResponse response = userService.getUserById(userId);
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

    @PutMapping("/users/{userId}")
    @PreAuthorize("hasRole('ADMIN')")
    @Operation(summary = "更新用户信息", description = "更新用户的邮箱、手机号、头像等信息")
    @ApiResponses(value = {
        @ApiResponse(responseCode = "200", description = "更新成功",
                content = @Content(schema = @Schema(implementation = UserResponse.class))),
        @ApiResponse(responseCode = "400", description = "更新失败")
    })
    public ResponseEntity<?> updateUser(
            @Parameter(description = "用户ID") @PathVariable Long userId,
            @Valid @RequestBody com.bailemi.dto.UserUpdateRequest request) {
        try {
            UserResponse response = userService.updateUser(userId, request);
            return ResponseEntity.ok(Map.of(
                "code", 0,
                "message", "用户信息已更新",
                "data", response
            ));
        } catch (Exception e) {
            return ResponseEntity.badRequest().body(Map.of(
                "code", 400,
                "message", e.getMessage()
            ));
        }
    }

    @PutMapping("/users/{userId}/password")
    @PreAuthorize("hasRole('ADMIN')")
    @Operation(summary = "重置用户密码", description = "管理员重置用户密码")
    @ApiResponses(value = {
        @ApiResponse(responseCode = "200", description = "密码已重置",
                content = @Content(schema = @Schema(implementation = UserResponse.class))),
        @ApiResponse(responseCode = "400", description = "重置失败")
    })
    public ResponseEntity<?> resetUserPassword(
            @Parameter(description = "用户ID") @PathVariable Long userId,
            @Valid @RequestBody com.bailemi.dto.UserPasswordResetRequest request) {
        try {
            UserResponse response = userService.resetUserPassword(userId, request.getNewPassword());
            return ResponseEntity.ok(Map.of(
                "code", 0,
                "message", "密码已重置",
                "data", response
            ));
        } catch (Exception e) {
            return ResponseEntity.badRequest().body(Map.of(
                "code", 400,
                "message", e.getMessage()
            ));
        }
    }

    @DeleteMapping("/users/{userId}")
    @PreAuthorize("hasRole('ADMIN')")
    @Operation(summary = "删除用户", description = "软删除指定用户账号")
    @ApiResponses(value = {
        @ApiResponse(responseCode = "200", description = "删除成功"),
        @ApiResponse(responseCode = "400", description = "删除失败")
    })
    public ResponseEntity<?> deleteUser(
            @Parameter(description = "用户ID") @PathVariable Long userId) {
        try {
            userService.deleteUser(userId);
            return ResponseEntity.ok(Map.of(
                "code", 0,
                "message", "用户已删除"
            ));
        } catch (Exception e) {
            return ResponseEntity.badRequest().body(Map.of(
                "code", 400,
                "message", e.getMessage()
            ));
        }
    }
}