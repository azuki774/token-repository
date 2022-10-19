grant all privileges on *.* to root@"%";
SET character_set_client = utf8mb4;
CREATE TABLE `oauth2` (
    token_name VARCHAR(64),
    access_token VARCHAR(255),
    refresh_token VARCHAR(255),
    refresh_url VARCHAR(255),
    expired_at TIMESTAMP NULL DEFAULT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp on update current_timestamp,
    PRIMARY KEY (`token_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `oauth2_history` (
    id int NOT NULL AUTO_INCREMENT,
    token_name VARCHAR(64),
    expired_in INT,
    refresh_url VARCHAR(255),
    requested_at TIMESTAMP DEFAULT current_timestamp,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp on update current_timestamp,
    INDEX idx_name (`token_name`),
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
