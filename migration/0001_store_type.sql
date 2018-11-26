-- +mig Up
CREATE TABLE IF NOT EXISTS `store_type` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(50) NOT NULL,
  PRIMARY KEY (`id`)
);
-- +mig Down
DROP TABLE `store_type`;
