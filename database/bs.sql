/*
Navicat MySQL Data Transfer

Source Server         : local
Source Server Version : 50517
Source Host           : localhost:3306
Source Database       : ibs

Target Server Type    : MYSQL
Target Server Version : 50517
File Encoding         : 65001

Date: 2015-04-16 10:30:06
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `tb_admin`
-- ----------------------------
DROP TABLE IF EXISTS `tb_admin`;
CREATE TABLE `tb_admin` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT NULL,
  `password` varchar(50) DEFAULT NULL,
  `isEnable` int(11) DEFAULT '1',
  `createTime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of tb_admin
-- ----------------------------
INSERT INTO `tb_admin` VALUES ('1', 'admin', 'admin', '1', '2015-01-05 00:00:00');
INSERT INTO `tb_admin` VALUES ('2', 'user', 'user', '1', '2015-01-08 13:40:29');
INSERT INTO `tb_admin` VALUES ('3', '111', '111', '1', '2015-04-16 10:05:11');

-- ----------------------------
-- Table structure for `tb_resource`
-- ----------------------------
DROP TABLE IF EXISTS `tb_resource`;
CREATE TABLE `tb_resource` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `parentId` int(11) DEFAULT NULL,
  `name` varchar(50) DEFAULT NULL,
  `controller` varchar(50) DEFAULT NULL,
  `action` varchar(50) DEFAULT NULL,
  `sort` int(11) DEFAULT NULL COMMENT '排序索引',
  `url` varchar(500) DEFAULT NULL,
  `isShow` int(11) DEFAULT NULL COMMENT '是否在菜单中显示',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of tb_resource
-- ----------------------------
INSERT INTO `tb_resource` VALUES ('1', '-1', '用户', '', '', '0', '', '1');
INSERT INTO `tb_resource` VALUES ('2', '-1', '角色', '', '', '0', '', '1');
INSERT INTO `tb_resource` VALUES ('3', '-1', '资源', '', '', '0', '', '1');
INSERT INTO `tb_resource` VALUES ('4', '-1', '权限', '', '', '0', '', '1');
INSERT INTO `tb_resource` VALUES ('5', '1', '添加用户', 'AdminController', 'add', '0', '/admin/add', '1');
INSERT INTO `tb_resource` VALUES ('6', '1', '管理用户', 'AdminController', 'index', '0', '/admin/index', '1');
INSERT INTO `tb_resource` VALUES ('7', '2', '添加角色', 'RoleController', 'add', '0', '/role/add', '1');
INSERT INTO `tb_resource` VALUES ('8', '2', '管理角色', 'RoleController', 'index', '0', '/role/index', '1');
INSERT INTO `tb_resource` VALUES ('9', '3', '添加资源', 'ResourceController', 'add', '0', '/resource/add', '1');
INSERT INTO `tb_resource` VALUES ('10', '3', '管理资源', 'ResourceController', 'index', '0', '/resource/index', '1');
INSERT INTO `tb_resource` VALUES ('11', '4', '用户角色', 'AuthorityController', 'userRole', '0', '/authority/userRole', '1');
INSERT INTO `tb_resource` VALUES ('12', '4', '角色资源', 'AuthorityController', 'roleResource', '0', '/authority/roleResource', '1');
INSERT INTO `tb_resource` VALUES ('13', '3', 'Home', 'AdminController', 'home', '0', '/admin/home', '0');
INSERT INTO `tb_resource` VALUES ('14', '4', 'addUsersInRole', 'AuthorityController', 'addUsersInRole', '0', '/authority/addUsersInRole', '0');
INSERT INTO `tb_resource` VALUES ('15', '4', 'deleteUsersFromRole', 'AuthorityController', 'deleteUsersFromRole', '0', '/authority/deleteUsersFromRole', '0');
INSERT INTO `tb_resource` VALUES ('16', '4', 'addResourcesInRole', 'AuthorityController', 'addResourcesInRole', '0', '/authority/addResourcesInRole', '0');
INSERT INTO `tb_resource` VALUES ('17', '4', 'sideBar', 'AuthorityController', 'sideBar', '0', '/authority/sideBar', '0');
INSERT INTO `tb_resource` VALUES ('18', '1', '添加用户', 'AdminController', 'addDo', '0', '/admin/add', '0');
INSERT INTO `tb_resource` VALUES ('19', '2', '添加角色', 'RoleController', 'addDo', '0', '/role/add', '0');
INSERT INTO `tb_resource` VALUES ('20', '3', '添加资源', 'ResourceController', 'addDo', '0', '/resource/add', '0');
INSERT INTO `tb_resource` VALUES ('21', '1', '编辑用户', 'AdminController', 'edit', '0', '/admin/edit', '0');
INSERT INTO `tb_resource` VALUES ('22', '1', '删除用户', 'AdminController', 'delete', '0', '/admin/delete', '0');
INSERT INTO `tb_resource` VALUES ('23', '2', '编辑角色', 'RoleController', 'edit', '0', '/role/edit', '0');
INSERT INTO `tb_resource` VALUES ('24', '3', '编辑资源', 'ResourceController', 'edit', '0', '/resource/edit', '0');
INSERT INTO `tb_resource` VALUES ('25', '1', '编辑用户', 'AdminController', 'editDo', '0', '/admin/edit', '0');
INSERT INTO `tb_resource` VALUES ('26', '2', '编辑角色', 'RoleController', 'editDo', '0', '/role/edit', '0');
INSERT INTO `tb_resource` VALUES ('27', '3', '编辑资源', 'ResourceController', 'editDo', '0', '/resource/edit', '0');
INSERT INTO `tb_resource` VALUES ('28', '2', '删除角色', 'RoleController', 'delete', '0', '/role/delete', '0');
INSERT INTO `tb_resource` VALUES ('29', '3', '删除资源', 'ResourceController', 'delete', '0', '/resource/delete', '0');
INSERT INTO `tb_resource` VALUES ('30', '4', '用户角色', 'AuthorityController', 'userRoleDo', '0', '/authority/user-role', '0');
INSERT INTO `tb_resource` VALUES ('31', '4', '角色资源', 'AuthorityController', 'roleResourceDo', '0', '/authority/role-resource', '0');

-- ----------------------------
-- Table structure for `tb_role`
-- ----------------------------
DROP TABLE IF EXISTS `tb_role`;
CREATE TABLE `tb_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  `createTime` datetime DEFAULT NULL,
  `isBase` int(11) DEFAULT NULL COMMENT '是否基础角色，基础角色不允许删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of tb_role
-- ----------------------------
INSERT INTO `tb_role` VALUES ('1', '超级管理员', '2015-01-05 00:00:00', '1');
INSERT INTO `tb_role` VALUES ('2', '普通用户', '2015-01-08 13:44:28', '0');

-- ----------------------------
-- Table structure for `tb_role_resource`
-- ----------------------------
DROP TABLE IF EXISTS `tb_role_resource`;
CREATE TABLE `tb_role_resource` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `roleId` int(11) DEFAULT NULL,
  `resourceId` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=161 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of tb_role_resource
-- ----------------------------
INSERT INTO `tb_role_resource` VALUES ('132', '2', '13');
INSERT INTO `tb_role_resource` VALUES ('133', '2', '17');
INSERT INTO `tb_role_resource` VALUES ('134', '1', '5');
INSERT INTO `tb_role_resource` VALUES ('135', '1', '6');
INSERT INTO `tb_role_resource` VALUES ('136', '1', '18');
INSERT INTO `tb_role_resource` VALUES ('137', '1', '21');
INSERT INTO `tb_role_resource` VALUES ('138', '1', '22');
INSERT INTO `tb_role_resource` VALUES ('139', '1', '25');
INSERT INTO `tb_role_resource` VALUES ('140', '1', '7');
INSERT INTO `tb_role_resource` VALUES ('141', '1', '8');
INSERT INTO `tb_role_resource` VALUES ('142', '1', '19');
INSERT INTO `tb_role_resource` VALUES ('143', '1', '23');
INSERT INTO `tb_role_resource` VALUES ('144', '1', '26');
INSERT INTO `tb_role_resource` VALUES ('145', '1', '28');
INSERT INTO `tb_role_resource` VALUES ('146', '1', '9');
INSERT INTO `tb_role_resource` VALUES ('147', '1', '10');
INSERT INTO `tb_role_resource` VALUES ('148', '1', '13');
INSERT INTO `tb_role_resource` VALUES ('149', '1', '20');
INSERT INTO `tb_role_resource` VALUES ('150', '1', '24');
INSERT INTO `tb_role_resource` VALUES ('151', '1', '27');
INSERT INTO `tb_role_resource` VALUES ('152', '1', '29');
INSERT INTO `tb_role_resource` VALUES ('153', '1', '11');
INSERT INTO `tb_role_resource` VALUES ('154', '1', '12');
INSERT INTO `tb_role_resource` VALUES ('155', '1', '14');
INSERT INTO `tb_role_resource` VALUES ('156', '1', '15');
INSERT INTO `tb_role_resource` VALUES ('157', '1', '16');
INSERT INTO `tb_role_resource` VALUES ('158', '1', '17');
INSERT INTO `tb_role_resource` VALUES ('159', '1', '30');
INSERT INTO `tb_role_resource` VALUES ('160', '1', '31');

-- ----------------------------
-- Table structure for `tb_user`
-- ----------------------------
DROP TABLE IF EXISTS `tb_user`;
CREATE TABLE `tb_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT NULL COMMENT '用户名',
  `password` varchar(50) DEFAULT NULL COMMENT '密码',
  `email` varchar(50) DEFAULT NULL,
  `phone` varchar(50) DEFAULT NULL COMMENT '手机号',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of tb_user
-- ----------------------------
INSERT INTO `tb_user` VALUES ('1', 'wlb', '3333333', 'wanglibing@qq.com', '18888888888');
INSERT INTO `tb_user` VALUES ('2', 'wanglibing', '111111', 'iamwanglibing@qq.com', '18810039685');

-- ----------------------------
-- Table structure for `tb_user_role`
-- ----------------------------
DROP TABLE IF EXISTS `tb_user_role`;
CREATE TABLE `tb_user_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `adminId` int(11) DEFAULT NULL,
  `roleId` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of tb_user_role
-- ----------------------------
INSERT INTO `tb_user_role` VALUES ('1', '1', '1');
INSERT INTO `tb_user_role` VALUES ('7', '2', '2');