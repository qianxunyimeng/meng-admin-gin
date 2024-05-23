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

 Date: 19/05/2024 15:38:03
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept` (
  `dept_id` bigint NOT NULL AUTO_INCREMENT,
  `parent_id` bigint DEFAULT NULL,
  `dept_path` varchar(255) DEFAULT NULL,
  `dept_name` varchar(128) DEFAULT NULL,
  `sort` tinyint DEFAULT NULL,
  `leader` varchar(128) DEFAULT NULL,
  `phone` varchar(11) DEFAULT NULL,
  `email` varchar(64) DEFAULT NULL,
  `status` tinyint DEFAULT NULL COMMENT '状态 1:正常 0 停用',
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`dept_id`),
  KEY `idx_sys_dept_create_by` (`create_by`),
  KEY `idx_sys_dept_update_by` (`update_by`),
  KEY `idx_sys_dept_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of sys_dept
-- ----------------------------
BEGIN;
INSERT INTO `sys_dept` VALUES (1, 0, '', '组织机构', 1, '', '', '', 1, 0, 1, '2024-05-14 00:03:10.208', '2024-05-18 10:26:05.325', NULL);
INSERT INTO `sys_dept` VALUES (2, 1, '', '北京分公司', 1, '', '', '', 1, 0, 1, '2024-05-14 21:55:27.390', '2024-05-14 22:50:36.471', '2024-05-14 22:54:42.301');
INSERT INTO `sys_dept` VALUES (3, 1, '', '武汉分公司', 2, '', '', '', 1, 0, 1, '2024-05-14 21:56:20.846', '2024-05-14 22:55:28.358', NULL);
INSERT INTO `sys_dept` VALUES (4, 1, '', '成都分公司', 5, '', '', '', 1, 0, 1, '2024-05-14 22:04:18.174', '2024-05-18 10:23:11.685', NULL);
INSERT INTO `sys_dept` VALUES (5, 1, '', '深圳分公司', 4, '', '', '', 1, 0, 1, '2024-05-14 22:07:25.836', '2024-05-14 22:55:48.270', NULL);
INSERT INTO `sys_dept` VALUES (6, 1, '', '北京总公司', 1, '', '', '', 1, 1, 0, '2024-05-14 22:51:20.561', '2024-05-14 22:51:20.561', NULL);
INSERT INTO `sys_dept` VALUES (7, 1, '', '上海分公司', 3, '', '', '', 1, 0, 1, '2024-05-14 22:52:20.446', '2024-05-14 22:55:34.561', NULL);
INSERT INTO `sys_dept` VALUES (8, 6, '', '研发部', 3, '', '', '', 1, 0, 1, '2024-05-14 23:01:41.475', '2024-05-15 20:51:40.791', NULL);
INSERT INTO `sys_dept` VALUES (9, 6, '', '市场部', 4, '', '', '', 1, 0, 1, '2024-05-14 23:01:54.538', '2024-05-15 20:51:50.748', NULL);
INSERT INTO `sys_dept` VALUES (10, 6, '', '财务部', 2, '', '', '', 1, 0, 1, '2024-05-14 23:02:14.266', '2024-05-15 20:51:20.453', NULL);
INSERT INTO `sys_dept` VALUES (11, 6, '', '运营部', 6, '', '', '', 1, 0, 1, '2024-05-14 23:02:22.820', '2024-05-15 20:52:14.822', NULL);
INSERT INTO `sys_dept` VALUES (12, 3, '', '研发部', 1, '', '', '', 1, 0, 1, '2024-05-15 20:50:17.423', '2024-05-15 20:50:17.423', NULL);
INSERT INTO `sys_dept` VALUES (13, 6, '', '人力资源', 5, '', '', '', 1, 0, 1, '2024-05-15 20:50:32.452', '2024-05-15 20:52:05.529', NULL);
INSERT INTO `sys_dept` VALUES (14, 6, '', '销售部', 1, '张三', '', '', 1, 0, 1, '2024-05-15 20:50:54.289', '2024-05-15 21:25:57.596', NULL);
INSERT INTO `sys_dept` VALUES (15, 7, '', '人力资源', 1, '', '', '', 1, 0, 1, '2024-05-15 20:52:31.008', '2024-05-15 20:52:31.008', NULL);
INSERT INTO `sys_dept` VALUES (16, 5, '', '人力资源', 1, '', '', '', 1, 0, 1, '2024-05-15 20:52:39.253', '2024-05-15 20:52:39.253', NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
