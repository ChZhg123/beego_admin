/*
 Navicat Premium Data Transfer

 Source Server         : 本地
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : localhost:3306
 Source Schema         : beego_blog

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 21/10/2021 14:46:12
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin_info
-- ----------------------------
DROP TABLE IF EXISTS `admin_info`;
CREATE TABLE `admin_info`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `username` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `nickname` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户昵称',
  `phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '手机号',
  `password` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
  `role_id` int(11) NULL DEFAULT 0 COMMENT '角色ID',
  `status` tinyint(4) NOT NULL DEFAULT 0 COMMENT '0 未删除 1已删除',
  `login_time` datetime(0) NULL DEFAULT NULL COMMENT '登录时间',
  `login_ip` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '登录IP',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `username`(`username`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 19 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '管理员表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_info
-- ----------------------------
INSERT INTO `admin_info` VALUES (1, 'admin', '超级管理员', '15063332154', '14e1b600b1fd579f47433b88e8d85291', 1, 0, '2021-10-20 15:00:15', '127.0.0.1:57486', '2021-10-12 14:13:28', '2021-10-20 15:00:15');
INSERT INTO `admin_info` VALUES (17, 'ceshi', '测试', '15063337229', '9db06bcff9248837f86d1a6bcf41c9e7', 1, 0, NULL, '', '2021-10-13 16:02:06', '2021-10-13 16:42:54');
INSERT INTO `admin_info` VALUES (18, 'ceshi1', '测试1', '15063337229', '14e1b600b1fd579f47433b88e8d85291', 1, 1, NULL, '', '2021-10-13 16:07:20', '2021-10-13 17:06:07');

-- ----------------------------
-- Table structure for admin_power
-- ----------------------------
DROP TABLE IF EXISTS `admin_power`;
CREATE TABLE `admin_power`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '模块名称',
  `url` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '模块路径',
  `pid` bigint(20) NOT NULL DEFAULT 0 COMMENT '上级id',
  `sort` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
  `icon` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '图标',
  `create_time` datetime(0) NOT NULL COMMENT '创建时间',
  `update_time` datetime(0) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_power
-- ----------------------------
INSERT INTO `admin_power` VALUES (1, '系统管理', '', 0, 0, 'fa-cog', '2021-07-09 18:08:35', '2021-07-09 18:08:35');
INSERT INTO `admin_power` VALUES (2, '管理员列表', 'admin/system/userList', 1, 8, 'fa-user', '2021-07-09 18:08:35', '2021-07-09 18:08:35');
INSERT INTO `admin_power` VALUES (3, '角色列表', 'admin/system/roleList', 1, 9, 'fa-users', '2021-07-09 18:08:35', '2021-07-09 18:08:35');

-- ----------------------------
-- Table structure for admin_role
-- ----------------------------
DROP TABLE IF EXISTS `admin_role`;
CREATE TABLE `admin_role`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色名称',
  `power` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '权限',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
  `status` tinyint(4) NOT NULL DEFAULT 0 COMMENT '0 未删除 1已删除',
  `create_time` datetime(0) NOT NULL COMMENT '创建时间',
  `update_time` datetime(0) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_role
-- ----------------------------
INSERT INTO `admin_role` VALUES (1, '管理员', '1,2,3', '超级管理员(有所有权限)', 0, '2021-07-09 18:08:39', '2021-07-12 16:03:50');
INSERT INTO `admin_role` VALUES (2, 'role按时', '', '打打阿萨德发', 1, '2021-10-13 17:30:17', '2021-10-13 17:30:38');
INSERT INTO `admin_role` VALUES (3, '阿士大夫', '1,2', '', 0, '2021-10-15 15:15:32', '2021-10-20 16:10:39');

-- ----------------------------
-- Table structure for schema_migrations
-- ----------------------------
DROP TABLE IF EXISTS `schema_migrations`;
CREATE TABLE `schema_migrations`  (
  `version` bigint(20) NOT NULL,
  `dirty` tinyint(1) NOT NULL,
  PRIMARY KEY (`version`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of schema_migrations
-- ----------------------------
INSERT INTO `schema_migrations` VALUES (20210924102752, 0);

SET FOREIGN_KEY_CHECKS = 1;
