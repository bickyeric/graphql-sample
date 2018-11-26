-- +mig Up
CREATE TABLE IF NOT EXISTS `store_category` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  `type_id` INT UNSIGNED NOT NULL,
  PRIMARY KEY (`id`, `type_id`),
  INDEX `fk_store_type_idx` (`type_id` ASC),
  CONSTRAINT `fk_store_type`
    FOREIGN KEY (`type_id`)
    REFERENCES `store_type` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);
-- +mig Down
DROP TABLE `store_category`;
