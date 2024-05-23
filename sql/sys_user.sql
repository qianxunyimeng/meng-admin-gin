/*
 Navicat Premium Data Transfer

 Source Server         : 本机mysql
 Source Server Type    : MySQL
 Source Server Version : 80027
 Source Host           : 127.0.0.1:3306
 Source Schema         : meng-admin

 Target Server Type    : MySQL
 Target Server Version : 80027
 File Encoding         : 65001

 Date: 19/05/2024 15:37:21
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
  `user_id` bigint NOT NULL AUTO_INCREMENT COMMENT '编码',
  `password` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户登录密码',
  `salt` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '加盐',
  `nick_name` varchar(128) COLLATE utf8mb4_general_ci DEFAULT '系统用户' COMMENT '用户昵称',
  `sex` varchar(10) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '性别',
  `avatar` varchar(191) COLLATE utf8mb4_general_ci DEFAULT 'https://pic.qianxun.shop/i/2024/03/29/66059924431c9.webp' COMMENT '用户头像',
  `phone` varchar(11) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户手机号',
  `email` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户邮箱',
  `status` varchar(191) COLLATE utf8mb4_general_ci DEFAULT '1' COMMENT '用户是否被冻结 1正常 0冻结',
  `role_id` mediumint DEFAULT NULL COMMENT '角色ID',
  `dept_id` mediumint DEFAULT NULL COMMENT '部门',
  `post_id` mediumint DEFAULT NULL COMMENT '岗位',
  `remark` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注',
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `user_name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户登录名',
  PRIMARY KEY (`user_id`),
  KEY `idx_sys_user_create_by` (`create_by`),
  KEY `idx_sys_user_update_by` (`update_by`),
  KEY `idx_sys_user_deleted_at` (`deleted_at`),
  KEY `idx_sys_user_user_name` (`user_name`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
BEGIN;
INSERT INTO `sys_user` VALUES (1, '$2a$10$0GXVLaqxyrzqMfjooIPaeuvHQX0A2TA5vwjUJRrsG0uJtsTSu8Ci.', '', '千寻', '', 'https://pic.qianxun.shop/i/2024/03/29/66059924431c9.webp', '15812341234', 'sql668@qianxun.com', '1', 1, 1, 0, '', 1, 0, '2024-05-19 15:26:48.881', '2024-05-19 15:26:48.881', NULL, 'admin');
INSERT INTO `sys_user` VALUES (2, '$2a$10$5DuTJR8IKPHDT1SflJmpS.Y.OrncDype1u3/GtFPgf2WpS2JjMLri', '', '普通用户', '', 'https://pic.qianxun.shop/i/2024/03/29/66059924431c9.webp', '15812341235', 'xxx@123.com', '1', 2, 1, 0, '', 1, 0, '2024-05-19 15:28:17.475', '2024-05-19 15:28:17.475', NULL, 'sql');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
