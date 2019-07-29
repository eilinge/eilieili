drop database if exists `eilieili`;
create database `eilieili` character set utf8;
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
   `email`                varchar(50)  NOT NULL DEFAULT '' COMMENT 'email地址',
   `username`             varchar(30)  NOT NULL DEFAULT '' COMMENT '用户名',
   `identity_id`          varchar(100)  NOT NULL DEFAULT '' COMMENT '验证身份id',
   `address`              varchar(256)  NOT NULL DEFAULT '' COMMENT 'ether地址'
);
CREATE UNIQUE INDEX account_email_uindex ON eilieili.account (email);
CREATE UNIQUE INDEX account_name_uindex ON eilieili.account (username);
alter table account comment '账户表';


create table `content` 
(
   `content_id`           int primary key not null auto_increment, 
   `title`                varchar(100) NOT NULL DEFAULT '' COMMENT '名称', 
   `content`              varchar(256) NOT NULL DEFAULT '' COMMENT '资产', 
   `content_hash`         varchar(100) NOT NULL DEFAULT '' COMMENT '资产hash',
   `price`                int(100) unsigned NOT NULL DEFAULT '0' COMMENT '价格', 
   `weight`               int(100) unsigned NOT NULL DEFAULT '0' COMMENT '权重',
   `ts`                  timestamp not null DEFAULT '2019-08-16' COMMENT '当前时间'
);

create table `account_content`
(
   `content_hash`         varchar(100) NOT NULL DEFAULT '' COMMENT '资产hash',
   `token_id`             int unsigned unique NOT NULL DEFAULT '0' COMMENT '资产_id',
   `address`              varchar(100) NOT NULL DEFAULT '' COMMENT 'ether地址',
   `ts`                   timestamp not null DEFAULT '2019-08-16' COMMENT '当前时间'
);


create table `auction`
(
   `content_hash`         varchar(256) NOT NULL DEFAULT '' COMMENT '资产hash',
   `address`              varchar(100) NOT NULL DEFAULT '' COMMENT 'ether地址',
   `token_id`             int unsigned unique NOT NULL DEFAULT '0' COMMENT '资产_id',
   `percent`              int unsigned NOT NULL DEFAULT '0' COMMENT '权重',
   `price`                int unsigned NOT NULL DEFAULT '0' COMMENT '价格',
   `status`               int unsigned NOT NULL DEFAULT '0' COMMENT '状态',
   `ts`                   timestamp NOT NULL DEFAULT '2019-08-16' COMMENT '当前时间'
);

create table `bidwinner` 
(
   `id`                   int primary key not null auto_increment, 
   `token_id`             int unsigned not null unique DEFAULT '0' COMMENT '资产_id', 
   `price`                int unsigned NOT NULL DEFAULT '0' COMMENT '价格',
   `address`              varchar (100) NOT NULL DEFAULT '' COMMENT 'ether地址',
   `ts`                   timestamp not null DEFAULT '2019-08-16' COMMENT '当前时间'
);

-- 记录投票信息
create table `vote`
(
   `vote_id`              int primary key auto_increment,
   `address`              varchar (100) NOT NULL DEFAULT '' COMMENT 'ether地址',
   -- content_hash         varchar(256), (content_hash不唯一: 原资产=已分割资产)
   `token_id`             int unsigned unique NOT NULL DEFAULT '0' COMMENT '资产_id',
   `vote_time`            timestamp NOT NULL DEFAULT '0' COMMENT '当前时间',
   `comment`              varchar(100) NOT NULL DEFAULT '' COMMENT '备注'
);

alter table `vote` comment '投票表，一个账户一个图片，只能投一票，一票代表30pxc';
CREATE UNIQUE INDEX vote_uindex ON `vote` (address);

delete from `account`;

delete from `vote`;
delete from `aution`;
delete from `account_content`;
delete from `content`;
delete from `bidwinner`;