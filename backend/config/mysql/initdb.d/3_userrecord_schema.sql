CREATE TABLE IF NOT EXISTS `user_records` (
  `id`                     INT NOT NULL AUTO_INCREMENT,
  `user_id`             VARCHAR(128) NOT NULL,
  `station_id`         INT UNSIGNED NOT NULL,
  `create_time`      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `create_time` (`create_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;