package com.bailemi.dto;

import lombok.Data;

@Data
public class UserQueryRequest {
    private String username;
    private String startDate;
    private String endDate;
    private Integer page = 1;
    private Integer pageSize = 10;
}
