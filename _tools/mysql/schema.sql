CREATE DATABASE IF NOT EXISTS `cinemarch`;
USE `cinemarch`;

CREATE TABLE IF NOT EXISTS `cinemas` (
  `id` VARCHAR(64) NOT NULL,
  `name` VARCHAR(64) NOT NULL,
  `prefecture` VARCHAR(16) NOT NULL,
  `address` VARCHAR(255) NOT NULL,
  `web_site_url` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_cinema` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `movies` (
  `id` VARCHAR(64) NOT NULL,
  `title` VARCHAR(64) NOT NULL,
  `release_date` DATETIME NOT NULL,
  `release_status` VARCHAR(16) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_movie` (`title`, `release_date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `screen_movies` (
  `id` VARCHAR(64) NOT NULL,
  `cinema_id` VARCHAR(64) NOT NULL,
  `movie_id` VARCHAR(64) NOT NULL,
  `screen_type` VARCHAR(32),
  `subtitle_or_dubbing` VARCHAR(32),
  `three_d` BOOLEAN,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unqi_screen_movie` (`cinema_id`, `movie_id`, `screen_type`, `subtitle_or_dubbing`, `three_d`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `screen_schedules` (
  `id` VARCHAR(64) NOT NULL,
  `screen_movie_id` VARCHAR(64) NOT NULL,
  `start_time` DATETIME NOT NULL,
  `end_time` DATETIME NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
