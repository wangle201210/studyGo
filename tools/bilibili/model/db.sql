create database bilibili;
CREATE TABLE `user` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `mid` int(11) unsigned NOT NULL COMMENT '用户id',
    `name` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '用户昵称',
    `sex` varchar(4) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '性别',
    `face` varchar(256) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '头像',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `mid_UNIQUE` (`mid`),
    KEY `idx_n` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='用户信息表';

CREATE TABLE `info` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `mid` int(11) unsigned NOT NULL COMMENT '用户id',
    `archive_view` BIGINT unsigned NOT NULL  DEFAULT '0',
    `likes` BIGINT unsigned NOT NULL  DEFAULT '0' COMMENT '获赞数',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `mid_UNIQUE` (`mid`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='用户的一些数据';