CREATE TABLE IF NOT EXISTS `companies` (
  `name`               VARCHAR(128) NOT NULL,
  `type`               TINYINT UNSIGNED NOT NULL,
  PRIMARY KEY (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `railway_lines` (
  `name`               VARCHAR(128) NOT NULL,
  `type`               TINYINT UNSIGNED NOT NULL,
  `operation_company`  VARCHAR(128) NOT NULL,
  PRIMARY KEY (`name`),
  FOREIGN KEY (`operation_company`) REFERENCES `companies`(`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `stations` (
  `id`                  INT NOT NULL AUTO_INCREMENT,
  `station_name`        VARCHAR(128) NOT NULL,
  `center_latlong`      GEOMETRY NOT NULL,
  `railway_line_name`   VARCHAR(128) NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`railway_line_name`) REFERENCES `railway_lines`(`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
