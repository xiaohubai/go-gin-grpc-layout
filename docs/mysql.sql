CREATE TABLE `user` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `uid` bigint unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `user_name` varchar(20) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(255) NOT NULL DEFAULT '123456' COMMENT '密码',
  `salt` varchar(20) NOT NULL DEFAULT 'abcdef' COMMENT '加盐',
  `role_id` tinyint NOT NULL DEFAULT '0' COMMENT '角色Id',
  `phone` varchar(11) NOT NULL DEFAULT '' COMMENT '手机号',
  `email` varchar(255) NOT NULL DEFAULT '' COMMENT '邮箱',
  `created_at` bigint unsigned NOT NULL COMMENT '创建时间',
  `updated_at` bigint unsigned NOT NULL COMMENT '更新时间',
  `extra_info` json DEFAULT NULL COMMENT '补充信息',
  PRIMARY KEY (`id`),
  KEY `idk_uid` (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户表';


CREATE TABLE `role` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `role_id` tinyint NOT NULL DEFAULT '0' COMMENT '角色ID',
  `role_name` varchar(20) NOT NULL DEFAULT '' COMMENT '角色名称',
  `created_at` bigint unsigned NOT NULL COMMENT '创建时间',
  `updated_at` bigint unsigned NOT NULL COMMENT '更新时间',
  `extra_info` json DEFAULT NULL COMMENT '补充信息',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='角色表';