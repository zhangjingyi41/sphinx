-- 用户表
CREATE TABLE `users` (
    `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '用户ID',
    `username` VARCHAR(64) NOT NULL COMMENT '用户名',
    `password` VARCHAR(128) NOT NULL COMMENT '密码',
    `phone` VARCHAR(20) NOT NULL COMMENT '手机号',
    `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态：1-正常，0-禁用',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_phone` (`phone`),
    UNIQUE KEY `idx_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

-- 第三方账号关联表
CREATE TABLE `external_identities` (
    `user_id` BIGINT NOT NULL COMMENT '用户ID',
    `provider` VARCHAR(32) NOT NULL COMMENT '第三方平台',
    `external_user_id` VARCHAR(128) NOT NULL COMMENT '第三方用户ID',
    `access_token` VARCHAR(512) NOT NULL COMMENT '访问令牌',
    `refresh_token` VARCHAR(512) DEFAULT NULL COMMENT '刷新令牌',
    `expired_at` TIMESTAMP NOT NULL COMMENT '令牌过期时间',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`user_id`, `provider`),
    UNIQUE KEY `idx_provider_external_id` (`provider`, `external_user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='第三方账号关联表';

-- OAuth客户端表
CREATE TABLE `oauth_clients` (
    `client_id` VARCHAR(64) NOT NULL COMMENT '客户端ID',
    `client_secret` VARCHAR(128) NOT NULL COMMENT '客户端密钥',
    `redirect_uri` VARCHAR(512) NOT NULL COMMENT '回调地址',
    `scope` VARCHAR(512) DEFAULT NULL COMMENT '权限范围',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`client_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='OAuth客户端表';

-- OAuth授权码表
CREATE TABLE `oauth_authorization_codes` (
    `code` VARCHAR(128) NOT NULL COMMENT '授权码',
    `client_id` VARCHAR(64) NOT NULL COMMENT '客户端ID',
    `user_id` BIGINT NOT NULL COMMENT '用户ID',
    `redirect_uri` VARCHAR(512) NOT NULL COMMENT '回调地址',
    `scope` VARCHAR(512) DEFAULT NULL COMMENT '权限范围',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `expires_at` TIMESTAMP NOT NULL COMMENT '过期时间',
    PRIMARY KEY (`code`),
    KEY `idx_client_user` (`client_id`, `user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='OAuth授权码表';

-- OAuth令牌表
CREATE TABLE `oauth_tokens` (
    `access_token` VARCHAR(128) NOT NULL COMMENT '访问令牌',
    `client_id` VARCHAR(64) NOT NULL COMMENT '客户端ID',
    `user_id` BIGINT NOT NULL COMMENT '用户ID',
    `expires_at` TIMESTAMP NOT NULL COMMENT '过期时间',
    `scope` VARCHAR(512) DEFAULT NULL COMMENT '权限范围',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`access_token`),
    KEY `idx_client_user` (`client_id`, `user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='OAuth令牌表';