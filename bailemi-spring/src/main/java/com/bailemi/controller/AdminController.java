package com.bailemi.controller;

import com.bailemi.dto.PageResponse;
import com.bailemi.dto.UserQueryRequest;
import com.bailemi.dto.UserResponse;
import com.bailemi.dto.UserStatusRequest;
import com.bailemi.service.UserService;
import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.web.bind.annotation.*;

import java.util.Map;

@RestController
@RequestMapping("/api/v1/admin")
@RequiredArgsConstructor
public class AdminController {

    private final UserService userService;

    @GetMapping("/users")
    @PreAuthorize("hasRole('ADMIN')")
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
    public ResponseEntity<?> updateUserStatus(
            @PathVariable Long userId,
            @RequestBody UserStatusRequest request) {
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
    public ResponseEntity<?> banUser(@PathVariable Long userId) {
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
    public ResponseEntity<?> unbanUser(@PathVariable Long userId) {
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
}
