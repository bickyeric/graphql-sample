-- +mig Up
CREATE TABLE IF NOT EXISTS `modifier` (
  `id` INT UNSIGNED NOT NULL,
  `name` VARCHAR(255) NULL,
  PRIMARY KEY (`id`)
);
-- +mig Down
DROP TABLE `modifier`;
