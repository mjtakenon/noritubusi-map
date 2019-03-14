CREATE TABLE IF NOT EXISTS `stations` (
  `id`                     INT NOT NULL AUTO_INCREMENT,
  `station_name`           VARCHAR(128) NOT NULL,
  `center_latlong`         GEOMETRY NOT NULL,
  `operation_company`      VARCHAR(128) NOT NULL,
  `service_provider_type`  TINYINT UNSIGNED NOT NULL,
  `railway_line_name`      VARCHAR(128) NOT NULL,
  `railway_type`           TINYINT UNSIGNED NOT NULL,
  PRIMARY KEY (`id`),
  SPATIAL INDEX (`center_latlong`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
