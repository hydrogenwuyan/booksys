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
`password` varchar(50) NOT NULL COMMENT '密码',
`sex` tinyint NOT NULL COMMENT '性别',
`age` tinyint NOT NULL COMMENT '年龄',
`phone` varchar(32) NOT NULL COMMENT '手机号',
`name` varchar(100) NOT NULL COMMENT '名字',
`updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
`createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
`deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
PRIMARY KEY(`id`)
) ENGINE=InnoDB COMMENT='admin entity''' DEFAULT CHARSET=utf8;
CREATE UNIQUE INDEX `t_admin_entity_user` ON `t_admin_entity` (`user`);
