use devdb;

-- https://engineering.linecorp.com/ja/blog/mysql-workbench-index-behavior-with-visual-explain/

CREATE TABLE IF NOT EXISTS `cv` (
    `id` INT(10) NOT NULL AUTO_INCREMENT,
    `ad_id` INT(10) NOT NULL,
    `user_id` INT(10) NOT NULL,
    `status` tinyINT(4) NOT NULL,
    `created_at` INT(10) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `ad` (
    `id` INT(10) NOT NULL AUTO_INCREMENT,
    `title` VARCHAR(256) NOT NULL,
    `description` TEXT NOT NULL,
    `type` TINYINT(10) NOT NULL,
    `start_at` INT(10) NOT NULL,
    `end_at` INT(10) NOT NULL,
    `updated_at` INT(10) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;