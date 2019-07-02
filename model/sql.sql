CREATE TABLE `links` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `url` varchar(200) COLLATE utf8mb4_bin NOT NULL COMMENT '长连接',
  `keyword` varchar(50) COLLATE utf8mb4_bin NOT NULL COMMENT '短链接码',
  `status` tinyint(1) NOT NULL COMMENT '1系统分配 2用户自定义',
  `create_time` bigint(20) NOT NULL,
  `update_time` bigint(20) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `links_UN` (`url`),
  UNIQUE KEY `links_keyword` (`keyword`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
