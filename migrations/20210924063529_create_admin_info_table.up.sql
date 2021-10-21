CREATE TABLE `admin_info` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `username` varchar(20) NOT NULL COMMENT '用户名',
  `nickname` varchar(20) NOT NULL COMMENT '用户昵称',
  `phone` varchar(20) DEFAULT NULL COMMENT '手机号',
  `password` varchar(128) NOT NULL DEFAULT '' COMMENT '密码',
  `role_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '角色ID',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0 未删除 1已删除',
  `login_time` datetime DEFAULT NULL COMMENT '登录时间',
  `login_ip` varchar(50) DEFAULT NULL DEFAULT '' COMMENT '登录IP',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) COMMENT='管理员表' ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;