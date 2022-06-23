CREATE DATABASE IF NOT EXISTS radiofy DEFAULT COLLATE utf8mb4_general_ci;
USE radiofy;

CREATE TABLE IF NOT EXISTS `user` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `create_date` DATETIME DEFAULT(CURDATE()),
    `spotify_id` VARCHAR(50) NOT NULL,
    `token` VARCHAR(150) NOT NULL,
    `mail` VARCHAR(100),
    `country` VARCHAR(4),
    `product` VARCHAR(10),
    PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `club` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `owner_id` INT NOT NULL,
    `name` VARCHAR(50) NOT NULL,
    `description` VARCHAR(150),
    PRIMARY KEY (`id`),
    FOREIGN KEY (`owner_id`) REFERENCES `user`(`id`)
);

CREATE TABLE IF NOT EXISTS `listener` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `user_id` INT NOT NULL,
    `club_id` INT,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_id`) REFERENCES `user`(`id`),
    FOREIGN KEY (`club_id`) REFERENCES `club`(`id`)
);

CREATE TABLE IF NOT EXISTS `club_song` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `club_id` INT NOT NULL,
    `song_id` VARCHAR(150),
    PRIMARY KEY (`id`),
    FOREIGN KEY (`club_id`) REFERENCES `club`(`id`)
);