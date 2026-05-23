package com.bailemi.config;

import lombok.Data;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.stereotype.Component;

import java.util.Arrays;
import java.util.List;

@Data
@Component
@ConfigurationProperties(prefix = "cors")
public class CorsProperties {
    
    private String allowedOrigins = "http://localhost:5173,http://localhost:3000";
    
    public List<String> getAllowedOriginsList() {
        if (allowedOrigins == null || allowedOrigins.trim().isEmpty()) {
            return Arrays.asList("*");
        }
        return Arrays.asList(allowedOrigins.split(","));
    }
}