DROP TABLE stations;
DROP TABLE buildings;
DROP TABLE railways;

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
  `type`                   TINYINT UNSIGNED NOT NULL,
  `company_name`           VARCHAR(128) NOT NULL,
  `service_provider_type`  TINYINT UNSIGNED NOT NULL,
  PRIMARY KEY (`id`),
  INDEX (`company_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `stations` (
  `id`                     INT NOT NULL AUTO_INCREMENT,
  `name`                   VARCHAR(128) NOT NULL,
  `railway_id`             INT NOT NULL,
  `building_id`            INT NOT NULL,
  PRIMARY KEY (`id`),
  -- FOREIGN KEY(列名) REFERENCES 親テーブル名(親列名)
  FOREIGN KEY (`building_id`) REFERENCES buildings(`id`), 
  FOREIGN KEY (`railway_id`) REFERENCES railways(`id`) ,
  INDEX (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO `buildings` (`name`, `latlong`) VALUES ('豊橋', GeomFromText('POINT(34.76289000 137.38115000)'));
INSERT INTO `buildings` (`name`, `latlong`) VALUES ('新所原', GeomFromText('POINT(34.72326000 137.48445500)'));

INSERT INTO `railways` (`name`, `type`, `company_name`, `service_provider_type`) VALUES ('東海道新幹線', 11 , '東海旅客鉄道', 1); 
INSERT INTO `railways` (`name`, `type`, `company_name`, `service_provider_type`) VALUES ('東海道線', 11 , '東海旅客鉄道', 1); 
INSERT INTO `railways` (`name`, `type`, `company_name`, `service_provider_type`) VALUES ('飯田線', 11 , '東海旅客鉄道', 1); 
INSERT INTO `railways` (`name`, `type`, `company_name`, `service_provider_type`) VALUES ('名古屋本線', 11 , '名古屋鉄道', 1); 
INSERT INTO `railways` (`name`, `type`, `company_name`, `service_provider_type`) VALUES ('天竜浜名湖線', 11 , '天竜浜名湖鉄道', 1); 

INSERT INTO `stations` (`name`, `railway_id`, `building_id`) VALUES ('豊橋', 1, 1);
INSERT INTO `stations` (`name`, `railway_id`, `building_id`) VALUES ('豊橋', 2, 1);
INSERT INTO `stations` (`name`, `railway_id`, `building_id`) VALUES ('豊橋', 3, 1);
INSERT INTO `stations` (`name`, `railway_id`, `building_id`) VALUES ('豊橋', 4, 1);
INSERT INTO `stations` (`name`, `railway_id`, `building_id`) VALUES ('新所原', 2, 2);
INSERT INTO `stations` (`name`, `railway_id`, `building_id`) VALUES ('新所原', 5, 2);
