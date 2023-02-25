CREATE TABLE `member` (
`id` bigint NOT NULL COMMENT '用户ID',
`account` varchar(20) NOT NULL COMMENT '用户账户',
`password` varchar(64) NOT NULL COMMENT '用户密码',
`nick_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '用户昵称',
`mobile` varchar(255) DEFAULT NULL COMMENT '手机号码',
`real_name` varchar(255) DEFAULT NULL COMMENT '真实姓名',
`avatar` varchar(255) DEFAULT NULL COMMENT '用户头像',
`sex` tinyint DEFAULT '0' COMMENT '性别 0:男 1:女',
`last_login_time` datetime DEFAULT NULL COMMENT '最后一次登录时间',
`deleted` tinyint NOT NULL DEFAULT '0' COMMENT '逻辑删除 0:未删除 1:已删除',
`create_time` datetime DEFAULT NULL COMMENT '创建时间',
`modify_time` datetime DEFAULT NULL COMMENT '修改时间',
PRIMARY KEY (`id`),
UNIQUE KEY `unq_account` (`account`) COMMENT '用户账户不能重复'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户表';