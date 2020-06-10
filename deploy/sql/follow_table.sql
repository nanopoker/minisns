CREATE TABLE `minisns` . `follow` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `follower` bigint(20) NOT NULL COMMENT 'follower’s userid from user_tab',
  `followee` bigint(20) NOT NULL COMMENT 'followee’s userid from user_tab',
  `creatime` datetime NOT NULL COMMENT 'Timestamp when created',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_follower_followee` (`follower`,`followee`),
  KEY `idx_follower` (`follower`),
  KEY `idx_followee` (`followee`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci