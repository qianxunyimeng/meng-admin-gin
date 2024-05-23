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

 Date: 19/05/2024 15:37:52
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu` (
  `menu_id` bigint NOT NULL AUTO_INCREMENT,
  `menu_name` varchar(128) DEFAULT NULL COMMENT '菜单名称',
  `title` varchar(128) DEFAULT NULL,
  `icon` varchar(128) DEFAULT NULL,
  `path` varchar(128) DEFAULT NULL,
  `paths` varchar(128) DEFAULT NULL,
  `menu_type` varchar(1) DEFAULT NULL,
  `action` varchar(16) DEFAULT NULL,
  `permission` varchar(255) DEFAULT NULL,
  `parent_id` smallint DEFAULT NULL,
  `no_cache` tinyint(1) DEFAULT NULL COMMENT '是否禁用页面缓存',
  `breadcrumb` varchar(255) DEFAULT NULL,
  `component` varchar(255) DEFAULT NULL,
  `sort` tinyint DEFAULT NULL COMMENT '显示排序',
  `visible` varchar(1) DEFAULT NULL COMMENT '菜单状态，1:显示 0:隐藏',
  `is_frame` varchar(1) DEFAULT '0',
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `status` varchar(1) DEFAULT NULL COMMENT '菜单状态，1:正常 0:停用',
  `view_type` varchar(1) DEFAULT NULL COMMENT '试图类型，1:普通页面 2:外链页面 3:内嵌页面',
  `is_internally` varchar(1) DEFAULT '0' COMMENT '是否是系统内置数据，内置数据不可删除',
  PRIMARY KEY (`menu_id`),
  KEY `idx_sys_menu_create_by` (`create_by`),
  KEY `idx_sys_menu_update_by` (`update_by`),
  KEY `idx_sys_menu_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=36 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
BEGIN;
INSERT INTO `sys_menu` VALUES (18, '系统管理', '', 'icon-park-outline:config', 'system', '/0/18', 'M', '', '', 0, 0, '', '', 2, '1', '0', 0, 1, '2024-05-05 16:02:46.396', '2024-05-12 15:11:17.729', NULL, '1', '1', '0');
INSERT INTO `sys_menu` VALUES (19, '菜单嵌套', '', 'icon-park-outline:application-menu', 'menu', '/0/19', 'M', '', '', 0, 0, '', '', 3, '1', '', 0, 1, '2024-05-05 16:03:11.659', '2024-05-12 15:11:47.226', NULL, '1', '1', '');
INSERT INTO `sys_menu` VALUES (20, '用户管理', '', 'icon-park-outline:every-user', 'user', '/0/18/20', 'C', '', '', 18, 0, '', '/views/system/user/index', 1, '1', '0', 0, 1, '2024-05-05 16:03:44.708', '2024-05-12 15:13:53.878', NULL, '1', '1', '0');
INSERT INTO `sys_menu` VALUES (21, '角色管理', '', 'icon-park-outline:permissions', 'role', '/0/18/21', 'C', '', '', 18, 0, '', '/views/system/role/index', 2, '1', '0', 0, 1, '2024-05-05 16:04:10.925', '2024-05-12 15:14:24.747', NULL, '1', '1', '0');
INSERT INTO `sys_menu` VALUES (22, '部门管理', '', 'icon-park-outline:network-tree', 'dept', '/0/18/22', 'C', '', '', 18, 0, '', '/views/system/dept/index', 4, '1', '0', 0, 1, '2024-05-05 16:04:26.806', '2024-05-12 15:15:21.498', NULL, '1', '1', '0');
INSERT INTO `sys_menu` VALUES (23, '菜单管理', '', 'ant-design:menu-unfold-outlined', 'menu', '/0/18/23', 'C', '', '', 18, 0, '', '/views/system/menu/index', 3, '1', '0', 0, 1, '2024-05-05 16:04:36.635', '2024-05-12 15:15:35.533', NULL, '1', '1', '0');
INSERT INTO `sys_menu` VALUES (24, 'menu-1', '', '', 'menu1', '/0/19/24', 'M', '', '', 19, 0, '', '', 0, '1', '0', 0, 1, '2024-05-05 16:23:24.700', '2024-05-12 15:12:01.718', NULL, '1', '1', '0');
INSERT INTO `sys_menu` VALUES (25, 'menu-1-1', '', '', 'menu11', '/0/19/24/25', 'C', '', '', 24, 0, '', '/views/menu/menu1/menu11/index', 0, '1', '0', 1, 0, '2024-05-05 16:24:35.741', '2024-05-05 16:24:35.746', NULL, '1', '1', '0');
INSERT INTO `sys_menu` VALUES (26, 'menu-1-2', '', '', 'menu12', '/0/19/24/26', 'M', '', '', 24, 0, '', '', 0, '1', '0', 1, 0, '2024-05-05 16:25:07.413', '2024-05-05 16:25:07.418', NULL, '1', '1', '0');
INSERT INTO `sys_menu` VALUES (27, 'menu-1-3', '', '', 'menu13', '/0/19/24/27', 'C', '', '', 24, 0, '', '/views/menu/menu1/menu13/index', 0, '1', '0', 1, 0, '2024-05-05 16:25:46.664', '2024-05-05 16:25:46.666', NULL, '1', '1', '0');
INSERT INTO `sys_menu` VALUES (28, 'menu-1-2-1', '', '', 'menu121', '/0/19/24/26/28', 'C', '', '', 26, 0, '', '/views/menu/menu1/menu12/menu121/index', 0, '1', '0', 1, 0, '2024-05-05 16:26:44.619', '2024-05-05 16:26:44.627', NULL, '1', '1', '0');
INSERT INTO `sys_menu` VALUES (29, 'menu-1-2-2', '', '', 'menu122', '/0/19/24/26/29', 'C', '', '', 26, 0, '', '/views/menu/menu1/menu12/menu122/index', 0, '1', '0', 1, 0, '2024-05-05 16:27:28.544', '2024-05-05 16:27:28.547', NULL, '1', '1', '0');
INSERT INTO `sys_menu` VALUES (31, '首页', '', 'ant-design:home-outlined', 'home', '/0/31', 'C', '', '', 0, 0, '', '/views/home/index', 1, '1', '0', 0, 1, '2024-05-08 20:07:13.291', '2024-05-15 20:53:28.672', NULL, '1', '1', '0');
INSERT INTO `sys_menu` VALUES (32, '查询', '', '', '', '/0/18/23/32', 'F', '', 'system:menu:query', 23, 0, '', '', 1, '', '0', 0, 1, '2024-05-09 21:41:03.705', '2024-05-09 21:58:12.609', NULL, '1', '1', '0');
INSERT INTO `sys_menu` VALUES (33, '新增', '', '', '', '/0/18/23/33', 'F', '', 'system:menu:add', 23, 0, '', '', 2, '', '0', 0, 1, '2024-05-09 21:53:03.048', '2024-05-09 21:58:17.322', NULL, '1', '1', '0');
INSERT INTO `sys_menu` VALUES (34, '修改', '', '', '', '/0/18/23/34', 'F', '', 'system:menu:update', 23, 0, '', '', 3, '', '0', 0, 1, '2024-05-09 21:53:23.389', '2024-05-09 21:58:22.121', NULL, '1', '1', '0');
INSERT INTO `sys_menu` VALUES (35, '删除', '', '', '', '/0/18/23/35', 'F', '', 'system:menu:delete', 23, 0, '', '', 4, '', '0', 0, 1, '2024-05-09 21:53:43.256', '2024-05-09 21:58:27.193', NULL, '1', '1', '0');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
