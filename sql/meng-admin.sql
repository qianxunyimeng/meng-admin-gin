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

 Date: 31/03/2024 13:11:46
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_api
-- ----------------------------
DROP TABLE IF EXISTS `sys_api`;
CREATE TABLE `sys_api` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `handle` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'handle',
  `title` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '标题',
  `path` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '地址',
  `action` varchar(16) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '请求类型',
  `type` varchar(16) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '接口类型',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`id`),
  KEY `idx_sys_api_create_by` (`create_by`),
  KEY `idx_sys_api_update_by` (`update_by`),
  KEY `idx_sys_api_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of sys_api
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept` (
  `dept_id` bigint NOT NULL AUTO_INCREMENT,
  `parent_id` bigint DEFAULT NULL,
  `dept_path` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `dept_name` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `sort` tinyint DEFAULT NULL,
  `leader` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `phone` varchar(11) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `email` varchar(64) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` tinyint DEFAULT NULL,
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`dept_id`),
  KEY `idx_sys_dept_create_by` (`create_by`),
  KEY `idx_sys_dept_update_by` (`update_by`),
  KEY `idx_sys_dept_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of sys_dept
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu` (
  `menu_id` bigint NOT NULL AUTO_INCREMENT,
  `menu_name` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `title` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `icon` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `path` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `paths` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `menu_type` varchar(1) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `action` varchar(16) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `permission` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `parent_id` smallint DEFAULT NULL,
  `no_cache` tinyint(1) DEFAULT NULL,
  `breadcrumb` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `component` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `sort` tinyint DEFAULT NULL,
  `visible` varchar(1) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `is_frame` varchar(1) COLLATE utf8mb4_general_ci DEFAULT '0',
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`menu_id`),
  KEY `idx_sys_menu_create_by` (`create_by`),
  KEY `idx_sys_menu_update_by` (`update_by`),
  KEY `idx_sys_menu_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_menu_api_rule
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu_api_rule`;
CREATE TABLE `sys_menu_api_rule` (
  `sys_menu_menu_id` bigint NOT NULL,
  `sys_api_id` bigint NOT NULL COMMENT '主键编码',
  PRIMARY KEY (`sys_menu_menu_id`,`sys_api_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of sys_menu_api_rule
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_post`;
CREATE TABLE `sys_post` (
  `post_id` bigint NOT NULL AUTO_INCREMENT,
  `post_name` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `post_code` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `sort` tinyint DEFAULT NULL,
  `status` tinyint DEFAULT NULL,
  `remark` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`post_id`),
  KEY `idx_sys_post_create_by` (`create_by`),
  KEY `idx_sys_post_update_by` (`update_by`),
  KEY `idx_sys_post_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of sys_post
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
  `role_id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `role_name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '角色名',
  `status` varchar(4) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `role_key` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `role_sort` bigint DEFAULT NULL,
  `remark` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `admin` tinyint(1) DEFAULT NULL,
  `data_scope` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`role_id`),
  KEY `idx_sys_role_create_by` (`create_by`),
  KEY `idx_sys_role_update_by` (`update_by`),
  KEY `idx_sys_role_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_role_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_dept`;
CREATE TABLE `sys_role_dept` (
  `role_id` bigint unsigned NOT NULL COMMENT '角色ID',
  `dept_id` bigint NOT NULL,
  PRIMARY KEY (`role_id`,`dept_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of sys_role_dept
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu` (
  `role_id` bigint unsigned NOT NULL COMMENT '角色ID',
  `menu_id` bigint NOT NULL,
  PRIMARY KEY (`role_id`,`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of sys_role_menu
-- ----------------------------
BEGIN;
COMMIT;

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
