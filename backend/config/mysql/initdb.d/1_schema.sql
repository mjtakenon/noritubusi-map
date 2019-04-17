CREATE TABLE IF NOT EXISTS `buildings` (
  `id`                     INT NOT NULL AUTO_INCREMENT,
  `name`                   VARCHAR(128) NOT NULL,
  `latlong`                GEOMETRY NOT NULL,
  PRIMARY KEY (`id`),
  SPATIAL INDEX (`latlong`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `railways` (
  `id`                     INT NOT NULL AUTO_INCREMENT,
  `name`                   VARCHAR(128) NOT NULL,
  PRIMARY KEY (`id`),
  INDEX (`company_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `stations` (
  `id`                     INT NOT NULL AUTO_INCREMENT,
  `name`                   VARCHAR(128) NOT NULL,
  `railway_id`             INT NOT NULL,
  `building_id`            INT NOT NULL,
  `num_in_railway`         TINYINT UNSIGNED NOT NULL,
  PRIMARY KEY (`id`),
<<<<<<< a60946f400b50c193829c49d6267e347da3afc5c
  -- FOREIGN KEY(列名) REFERENCES 親テーブル名(親列名)
  FOREIGN KEY (`building_id`) REFERENCES buildings(`id`),
  FOREIGN KEY (`railway_id`) REFERENCES railways(`id`),
=======
  FOREIGN KEY (`building_id`) REFERENCES buildings(`id`), 
  FOREIGN KEY (`railway_id`) REFERENCES railways(`id`) ,
>>>>>>> SQL文の生成コードを追加
  INDEX (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
