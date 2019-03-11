CREATE TABLE IF NOT EXISTS `user_records` (
  `user_id`             VARCHAR(128) NOT NULL,
  `station_id`         INT UNSIGNED NOT NULL,
  `create_time`      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;