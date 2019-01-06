INSERT INTO `companies` (`name`, `type`) VALUES ('名古屋鉄道', 4);

INSERT INTO `railway_lines` (`name`, `type`, `operation_company`) VALUES ('名古屋本線', 12, '名古屋鉄道');

INSERT INTO `stations` (`station_name`, `center_latlong`, `railway_line_name`) VALUES ('矢作橋', GeomFromText('POINT(34.81476250 137.34562500)'), '名古屋本線');
