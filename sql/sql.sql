DROP DATABASE IF EXISTS `default`;
DROP DATABASE IF EXISTS `booksys`;

set names utf8;
CREATE DATABASE IF NOT EXISTS `default` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;
CREATE DATABASE IF NOT EXISTS `booksys` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;

USE booksys;

-- --------------------------------------------------
--  Table Structure for `t_admin_entity`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `t_admin_entity` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user` varchar(50) NOT NULL COMMENT '账号',
    `password` varchar(300) NOT NULL COMMENT '密码',
    `sex` tinyint NOT NULL default 0 COMMENT '性别',
    `age` tinyint NOT NULL default 0 COMMENT '年龄',
    `phone` varchar(32) NOT NULL default '' COMMENT '手机号',
    `name` varchar(100) NOT NULL default '' COMMENT '名字',
    `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
    `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
    `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
    PRIMARY KEY(`id`)
) ENGINE=InnoDB COMMENT='admin entity' DEFAULT CHARSET=utf8;
CREATE UNIQUE INDEX `t_admin_entity_user` ON `t_admin_entity` (`user`);

CREATE TABLE IF NOT EXISTS `t_book_entity` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
    `isBorrow` tinyint NOT NULL COMMENT '0未借出 1已借出',
    `type` int NOT NULL COMMENT '类型',
    `name` varchar(100) NOT NULL COMMENT '书名',
    `author` varchar(100) NOT NULL COMMENT '作者',
    `press` varchar(100) NOT NULL COMMENT '出版社',
    `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
    `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
    `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
    PRIMARY KEY(`id`)
) ENGINE=InnoDB COMMENT='admin entity''' DEFAULT CHARSET=utf8;
CREATE INDEX `t_book_entity_type` ON `t_book_entity` (`type`);
