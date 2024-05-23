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

 Date: 19/05/2024 15:37:37
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
  `role_id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `role_name` varchar(191) DEFAULT NULL COMMENT '角色名',
  `status` varchar(4) DEFAULT NULL COMMENT '状态 0禁用 1正常',
  `remark` varchar(255) DEFAULT NULL,
  `admin` tinyint(1) DEFAULT NULL,
  `data_scope` varchar(128) DEFAULT NULL,
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `role_code` varchar(128) DEFAULT NULL,
  `sort` bigint DEFAULT NULL,
  PRIMARY KEY (`role_id`),
  KEY `idx_sys_role_create_by` (`create_by`),
  KEY `idx_sys_role_update_by` (`update_by`),
  KEY `idx_sys_role_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
BEGIN;
INSERT INTO `sys_role` VALUES (1, '系统管理员', '1', '', 0, '', 1, 0, '2024-05-17 23:05:13.672', '2024-05-17 23:05:13.672', NULL, 'admin', 1);
INSERT INTO `sys_role` VALUES (2, '普通用户', '1', '', 0, '', 0, 1, '2024-05-17 23:33:35.387', '2024-05-18 18:03:28.945', NULL, 'common', 3);
INSERT INTO `sys_role` VALUES (3, '开发者', '1', '', 0, '', 0, 1, '2024-05-17 23:35:35.162', '2024-05-18 18:03:29.422', NULL, 'developer', 4);
INSERT INTO `sys_role` VALUES (4, '安全管理员', '1', '', 0, '', 0, 1, '2024-05-18 12:50:47.000', '2024-05-18 18:03:28.471', NULL, 'security', 2);
INSERT INTO `sys_role` VALUES (5, '测试员', '1', '', 0, '', 1, 0, '2024-05-18 12:59:56.689', '2024-05-18 12:59:56.689', '2024-05-18 13:01:30.703', 'test', 1);
INSERT INTO `sys_role` VALUES (6, '测试员', '1', '', 0, '', 0, 1, '2024-05-18 13:01:45.311', '2024-05-18 18:03:30.191', NULL, 'test', 6);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
