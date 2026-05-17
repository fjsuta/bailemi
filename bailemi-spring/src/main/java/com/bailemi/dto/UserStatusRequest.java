package com.bailemi.dto;

import lombok.Data;

@Data
public class UserStatusRequest {
    private Integer status; // 1: 解封, 0: 封禁
}
