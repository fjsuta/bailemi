package com.bailemi.dto;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class AuthResponse {
    private Long userId;
    private String username;
    private String email;
    private String phone;
    private String avatarUrl;
    private String accessToken;
    private String refreshToken;
    private Long expiresIn;
}
