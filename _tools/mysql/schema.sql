CREATE DATABASE IF NOT EXISTS `cinema-search`;
USE `cinema-search`;

CREATE TABLE IF NOT EXISTS `cinemas` (
  `id` VARCHAR(255) NOT NULL,
  `name` VARCHAR(255) NOT NULL,
  `address` VARCHAR(255) NOT NULL,
  `url` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `screening_movies` (
  `id` VARCHAR(255) NOT NULL,
  `cinema_id` VARCHAR(255) NOT NULL,
  `movie_id` VARCHAR(255) NOT NULL,
  `start_time` DATETIME NOT NULL,
  `end_time` DATETIME NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `screening_movie` (`cinema_id`, `movie_id`, `start_time`, `end_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `movies` (
  `id` VARCHAR(255) NOT NULL,
  `title` VARCHAR(255) NOT NULL,
  `release_date` DATETIME NOT NULL,
  `release_status` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `screening_movie_screening_types` (
  `id` VARCHAR(255) NOT NULL,
  `screening_movie_id` VARCHAR(255) NOT NULL,
  `screening_type_id` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `movie_type` (`screening_movie_id`, `screening_type_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `screening_types` (
  `id` VARCHAR(255) NOT NULL,
  `name` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
