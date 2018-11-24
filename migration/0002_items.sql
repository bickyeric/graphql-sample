-- +mig Up
CREATE TABLE IF NOT EXISTS `items` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL,
  `price` INT UNSIGNED NOT NULL,
  `description` TEXT NULL,
  `image` VARCHAR(255) NULL DEFAULT "",
  `created_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX `name_index` (`name` ASC));
-- +mig Down
DROP TABLE `items`;