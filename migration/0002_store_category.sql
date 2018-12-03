-- +mig Up
CREATE TABLE IF NOT EXISTS `store_category` (
  `id` INT UNSIGNED NOT NULL,
  `name` VARCHAR(45) NOT NULL,
  `type_id` INT UNSIGNED NOT NULL,
  PRIMARY KEY (`type_id`, `id`),
  INDEX `fk_store_type_idx` (`type_id` ASC),
  CONSTRAINT `fk_store_type`
    FOREIGN KEY (`type_id`)
    REFERENCES `store_type` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);
-- +mig Down
DROP TABLE `store_category`;
