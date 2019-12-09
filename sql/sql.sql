DROP DATABASE IF EXISTS `default`;
DROP DATABASE IF EXISTS `booksys`;

set names utf8;
CREATE DATABASE IF NOT EXISTS `default` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;
CREATE DATABASE IF NOT EXISTS `booksys` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;

USE booksys;

-- --------------------------------------------------
--  Table Structure for `t_chat_room`
-- --------------------------------------------------
-- CREATE TABLE IF NOT EXISTS `t_chat_room` (
-- `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
-- `isClose` int(11) NOT NULL COMMENT '房间是否关闭',
-- `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
-- `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
-- `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
-- PRIMARY KEY(`id`)
-- ) ENGINE=InnoDB COMMENT='chat room' DEFAULT CHARSET=utf8;

-- --------------------------------------------------
--  Table Structure for `t_chat_history`
-- --------------------------------------------------
-- CREATE TABLE IF NOT EXISTS `t_chat_history` (
-- `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
-- `roomId` bigint(20) NOT NULL COMMENT '房间id',
-- `userId` varchar(256) NOT NULL COMMENT '用户id',
-- `content` varchar(1024) NOT NULL COMMENT "聊天数据",
-- `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
-- `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
-- `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
-- PRIMARY KEY(`id`)
-- ) ENGINE=InnoDB COMMENT='chat history' DEFAULT CHARSET=utf8;