-- +mig Up
CREATE TABLE IF NOT EXISTS `discount` (
  `id` INT UNSIGNED NOT NULL,
  `outlet_id` INT UNSIGNED NOT NULL,
  `store_id` INT UNSIGNED NOT NULL,
  `amount` INT NOT NULL,
  `type` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_discount_outlet1_idx` (`outlet_id` ASC, `store_id` ASC),
  CONSTRAINT `fk_discount_outlet1`
    FOREIGN KEY (`store_id`, `outlet_id`)
    REFERENCES `outlet` (`store_id`, `id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);
-- +mig Down
DROP TABLE `discount`;
