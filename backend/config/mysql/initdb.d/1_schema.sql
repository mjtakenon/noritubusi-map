CREATE TABLE IF NOT EXISTS `buildings` (
  `id`                     INT NOT NULL AUTO_INCREMENT,
  `name`                   VARCHAR(128) NOT NULL,
  `latlong`                GEOMETRY NOT NULL,
  `connected_railways_num`  TINYINT UNSIGNED NOT NULL,
  PRIMARY KEY (`id`),
  SPATIAL INDEX (`latlong`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `railways` (
  `id`                     INT NOT NULL AUTO_INCREMENT,
  `name`                   VARCHAR(128) NOT NULL,
  PRIMARY KEY (`id`),
  INDEX (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `stations` (
  `id`                     INT NOT NULL AUTO_INCREMENT,
  `name`                   VARCHAR(128) NOT NULL,
  `railway_id`             INT NOT NULL,
  `building_id`            INT NOT NULL,
  `num_in_railway`         TINYINT UNSIGNED NOT NULL,
  PRIMARY KEY (`id`),
  -- FOREIGN KEY(列名) REFERENCES 親テーブル名(親列名)
  FOREIGN KEY (`building_id`) REFERENCES buildings(`id`),
  FOREIGN KEY (`railway_id`) REFERENCES railways(`id`),
  INDEX (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
