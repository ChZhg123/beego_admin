CREATE TABLE `admin_role` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL COMMENT '角色名称',
  `power` varchar(500) DEFAULT NULL COMMENT '权限',
  `desc` varchar(255) DEFAULT NULL COMMENT '描述',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0 未删除 1已删除',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) COMMENT='角色表' ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;