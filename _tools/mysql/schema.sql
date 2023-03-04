CREATE DATABASE IF NOT EXISTS `cinemarch`;
USE `cinemarch`;

CREATE TABLE IF NOT EXISTS `cinemas` (
  `id` VARCHAR(128) NOT NULL,
  `name` VARCHAR(128) NOT NULL,
  `prefecture` VARCHAR(128) NOT NULL,
  `address` VARCHAR(128) NOT NULL,
  `web_site` VARCHAR(128) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_cinema` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `movies` (
  `id` VARCHAR(128) NOT NULL,
  `title` VARCHAR(128) NOT NULL,
  `release_date` DATETIME NOT NULL,
  `release_status` VARCHAR(128) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_movie` (`title`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `screen_movies` (
  `id` VARCHAR(128) NOT NULL,
  `cinema_id` VARCHAR(128) NOT NULL,
  `movie_id` VARCHAR(128) NOT NULL,
  `screen_type` VARCHAR(128) NOT NULL,
  `translate_type` VARCHAR(128) DEFAULT NULL,
  `three_d` TINYINT(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `unqi_screen_movie` (`cinema_id`, `movie_id`, `screen_type`, `translate_type`, `three_d`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `screen_schedules` (
  `id` VARCHAR(128) NOT NULL,
  `screen_movie_id` VARCHAR(128) NOT NULL,
  `start_time` DATETIME NOT NULL,
  `end_time` DATETIME NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_screen_schedule` (`screen_movie_id`, `start_time`, `end_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
