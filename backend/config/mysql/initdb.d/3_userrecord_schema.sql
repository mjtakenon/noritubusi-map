CREATE TABLE IF NOT EXISTS `users` (
  `id`  VARCHAR(128) NOT NULL,
  `hashed_password`  VARCHAR(128) NOT NULL,
  `create_time`  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `change_time`  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `create_time` (`create_time`),
  KEY `change_time` (`change_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `user_records` (
  `id`               INT NOT NULL AUTO_INCREMENT,
  `user_id`          VARCHAR(128) NOT NULL,
  `building_id`      INT NOT NULL,
  `create_time`      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`user_id`) REFERENCES users(`id`),
  FOREIGN KEY (`building_id`) REFERENCES buildings(`id`),
  KEY `create_time` (`create_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;