drop database if exists `eilieili`;
create database `eilieili` character set utf8mb4;
use `eilieili`;

drop table if exists `vote`;
drop table if exists `account_content`;
drop table if exists `aution`;
drop table if exists `account`;
drop table if exists `content`;
drop table if exists `bidwinner`;

create table `account`
(
   `account_id`           int(10) unsigned not null primary key auto_increment,
   `email`                varchar(256)  NOT NULL DEFAULT '' COMMENT 'email地址',
   `username`             varchar(256) NOT NULL DEFAULT '' COMMENT '用户名',
   `identity_id`          varchar(256)  NOT NULL DEFAULT '' COMMENT '验证身份id',
   `address`              varchar(256)  NOT NULL DEFAULT '' COMMENT 'ether地址'
)ENGINE= InnoDB DEFAULT CHARSET = utf8mb4;
CREATE UNIQUE INDEX account_email_uindex ON eilieili.account (email);
CREATE UNIQUE INDEX account_name_uindex ON eilieili.account (username);
alter table account comment '账户表';

create table `content` 
(
   `content_id`           int primary key not null auto_increment, 
   `title`                varchar(256) NOT NULL DEFAULT '' COMMENT '名称', 
   `content`              varchar(256) NOT NULL DEFAULT '' COMMENT '资产', 
   `content_hash`         varchar(256) NOT NULL DEFAULT '' COMMENT '资产hash',
   `price`                int(100) unsigned NOT NULL DEFAULT '0' COMMENT '价格', 
   `weight`               int(100) unsigned NOT NULL DEFAULT '0' COMMENT '权重',
   `ts`                   int(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
   INDEX(`content_hash`)
)ENGINE= InnoDB DEFAULT CHARSET = utf8mb4;

create table `account_content`
(
   `content_hash`         varchar(256) NOT NULL DEFAULT '' COMMENT '资产hash',
   `token_id`             int unsigned unique NOT NULL DEFAULT '0' COMMENT '资产_id',
   `address`              varchar(256) NOT NULL DEFAULT '' COMMENT 'ether地址',
   `ts`                   int(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
   INDEX(`content_hash`)
)ENGINE= InnoDB DEFAULT CHARSET = utf8mb4;


create table `auction`
(
   `content_hash`         varchar(256) NOT NULL DEFAULT '' COMMENT '资产hash',
   `address`              varchar(256) NOT NULL DEFAULT '' COMMENT 'ether地址',
   `token_id`             int unsigned unique NOT NULL DEFAULT '0' COMMENT '资产_id',
   `percent`              int unsigned NOT NULL DEFAULT '0' COMMENT '权重',
   `price`                int unsigned NOT NULL DEFAULT '0' COMMENT '价格',
   `status`               int unsigned NOT NULL DEFAULT '0' COMMENT '状态',
   `ts`                   int(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
   INDEX(`content_hash`)
)ENGINE= InnoDB DEFAULT CHARSET = utf8mb4;

create table `bidwinner` 
(
   `id`                   int primary key not null auto_increment, 
   `token_id`             int unsigned not null unique DEFAULT '0' COMMENT '资产_id', 
   `price`                int unsigned NOT NULL DEFAULT '0' COMMENT '价格',
   `address`              varchar(256) NOT NULL DEFAULT '' COMMENT 'ether地址',
   `ts`                   int(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间'
)ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

-- 记录投票信息
create table `vote`
(
   `vote_id`              int primary key auto_increment,
   `address`              varchar(256) NOT NULL DEFAULT '' COMMENT 'ether地址',
   -- content_hash         varchar(256), (content_hash不唯一: 原资产=已分割资产)
   `token_id`             int unsigned NOT NULL DEFAULT '0' COMMENT '资产_id',
   `vote_time`            int(20) unsigned NOT NULL DEFAULT '0' COMMENT '投票时间',
   `comment`              varchar(256) NOT NULL DEFAULT '' COMMENT '备注'
)ENGINE= InnoDB DEFAULT CHARSET = utf8mb4;

alter table `vote` comment '投票表，一个账户一个图片，只能投一票，一票代表30pxc';
-- CREATE UNIQUE INDEX vote_uindex ON `vote` (address);

DROP TABLE IF EXISTS `voteCount`;
-- 投票票数
create table `voteCount`
(
   `id`              	int primary key auto_increment,
   `address`            varchar(256) NOT NULL DEFAULT '' COMMENT 'ether地址',
   `token_id`           int unsigned unique NOT NULL DEFAULT '0' COMMENT '资产_id',
   `vote_time`          int(20) unsigned NOT NULL DEFAULT '0' COMMENT '投票时间',
   `amount`             int(20) unsigned NOT NULL DEFAULT '0' COMMENT '投票票数'
)ENGINE= InnoDB DEFAULT CHARSET = utf8mb4;

alter table `vote` comment '投票表，一个账户一个图片，只能投一票，一票代表30pxc';

DROP TABLE IF EXISTS `lt_blackip`;

CREATE TABLE `lt_blackip`
(
   `id` int(10)unsigned primary key NOT NULL AUTO_INCREMENT,
   `ip` varchar(50)NOT NULL DEFAULT '' COMMENT 'IP地址',
   `blacktime` int(10)unsigned NOT NULL DEFAULT '0' COMMENT '黑名单限制到期时间',
   `sys_created` int(10)unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
   `sys_updated` int(10)unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
   PRIMARY KEY(`id`),
   UNIQUE KEY `ip`(`ip`)
)ENGINE = InnoDB AUTO_INCREMENT = 2 DEFAULT CHARSET = utf8mb4;

DROP TABLE IF EXISTS `lt_userday`;

CREATE TABLE `lt_userday`
(
   `id` int(10)unsigned primary key NOT NULL AUTO_INCREMENT,
   `uid` int(10)unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
   `day` int(10)unsigned NOT NULL DEFAULT '0' COMMENT '日期，如：20180725',
   `num` int(10)unsigned NOT NULL DEFAULT '0' COMMENT '次数',
   `sys_created` int(10)unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
   `sys_updated` int(10)unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
   PRIMARY KEY(`id`),
   UNIQUE KEY `uid_day`(`uid`, `day`)
)ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

DROP TABLE IF EXISTS `lt_user`;

CREATE TABLE `lt_user`
(
   `id` int(10) unsigned primary key NOT NULL AUTO_INCREMENT,
   `username` varchar(50)NOT NULL DEFAULT '' COMMENT '用户名',
   `blacktime` int(10)unsigned NOT NULL DEFAULT '0' COMMENT '黑名单限制到期时间',
   `realname` varchar(50)NOT NULL DEFAULT '' COMMENT '联系人',
   `mobile` varchar(50)NOT NULL DEFAULT '' COMMENT '手机号',
   `address` varchar(255)NOT NULL DEFAULT '' COMMENT '联系地址',
   `sys_created` int(10)unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
   `sys_updated` int(10)unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
   `sys_ip` varchar(50)NOT NULL DEFAULT '' COMMENT 'IP地址',
   PRIMARY KEY(`id`)
)ENGINE = InnoDB AUTO_INCREMENT = 2 DEFAULT CHARSET = utf8mb4;

delete from `account`;
delete from `vote`;
delete from `auction`;
delete from `account_content`;
delete from `content`;
delete from `bidwinner`;
delete from `lt_userday`;
delete from `lt_user`;
delete from `lt_blackip`;