-- +mig Up
CREATE TABLE IF NOT EXISTS `customers` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `first_name` VARCHAR(50) NOT NULL,
  `last_name` VARCHAR(50) NOT NULL,
  `gender` TINYINT(1) NULL DEFAULT 1,
  `phone_number` VARCHAR(45) NULL DEFAULT "",
  `address` VARCHAR(255) NULL DEFAULT "",
  `city` VARCHAR(45) NULL DEFAULT "",
  `zip` VARCHAR(20) NULL DEFAULT "",
  `comment` VARCHAR(45) NULL DEFAULT "",
  PRIMARY KEY (`id`));
-- +mig Down
DROP TABLE `customers`;
