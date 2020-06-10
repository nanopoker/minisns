CREATE TABLE `minisns` . `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `userid` bigint(20) NOT NULL,
  `username` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'Username',
  `salt` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'random string used to encrypt password',
  `password` varchar(96) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'encrypted password',
  `nickname` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'Nickname',
  `email` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'Email',
  `createtime` datetime NOT NULL COMMENT 'timestamp when created',
  `updatetime` datetime NOT NULL COMMENT 'timestamp when updated',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_username` (`username`),
  UNIQUE KEY `uniq_userid` (`userid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci