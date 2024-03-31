/*
 Navicat Premium Data Transfer

 Source Server         : 本机mysql
 Source Server Type    : MySQL
 Source Server Version : 80027
 Source Host           : localhost:3306
 Source Schema         : meng-admin

 Target Server Type    : MySQL
 Target Server Version : 80027
 File Encoding         : 65001

 Date: 30/03/2024 01:27:36
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_users
-- ----------------------------
DROP TABLE IF EXISTS `sys_users`;
CREATE TABLE `sys_users` (
  `user_id` bigint NOT NULL AUTO_INCREMENT COMMENT '编码',
  `username` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户登录名',
  `password` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户登录密码',
  `salt` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '加盐',
  `nick_name` varchar(128) COLLATE utf8mb4_general_ci DEFAULT '系统用户' COMMENT '用户昵称',
  `sex` varchar(10) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '性别',
  `avatar` varchar(191) COLLATE utf8mb4_general_ci DEFAULT 'https://pic.qianxun.shop/i/2024/03/29/66059924431c9.webp' COMMENT '用户头像',
  `phone` varchar(11) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户手机号',
  `email` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户邮箱',
  `status` bigint DEFAULT '1' COMMENT '用户是否被冻结 1正常 2冻结',
  `role_id` mediumint DEFAULT NULL COMMENT '角色ID',
  `dept_id` mediumint DEFAULT NULL COMMENT '部门',
  `post_id` mediumint DEFAULT NULL COMMENT '岗位',
  `remark` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注',
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`user_id`),
  KEY `idx_sys_users_username` (`username`),
  KEY `idx_sys_users_create_by` (`create_by`),
  KEY `idx_sys_users_update_by` (`update_by`),
  KEY `idx_sys_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of sys_users
-- ----------------------------
BEGIN;
INSERT INTO `sys_users` VALUES (1, 'admin', '$2a$10$1q4t9Ooue.T1gH1PuQmx.etuwRLHfw4/yK/NQX7RsobA/.WyDeuBW', '', '系统管理员', '', 'https://pic.qianxun.shop/i/2024/03/29/66059924431c9.webp', '15888888888', 'admin@123.com', 1, 0, 1, 0, '', 1, 0, '2024-03-30 01:24:02.872', '2024-03-30 01:24:02.872', NULL);
INSERT INTO `sys_users` VALUES (2, 'user1', '$2a$10$ZQ72BwtxOLR12b.ybCF/LeUA.09p.eeMC0KwIesNXEd2P5gAShRQC', '', '普通用户', '', 'https://pic.qianxun.shop/i/2024/03/29/66059924431c9.webp', '15888888888', 'user@123.com', 1, 0, 1, 0, '', 1, 0, '2024-03-30 01:25:46.103', '2024-03-30 01:25:46.103', NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
