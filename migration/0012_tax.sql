-- +mig Up
CREATE TABLE IF NOT EXISTS `tax` (
  `id` INT UNSIGNED NOT NULL,
  `outlet_id` INT UNSIGNED NOT NULL,
  `store_id` INT UNSIGNED NOT NULL,
  `amount` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_tax_outlet1_idx` (`outlet_id` ASC, `store_id` ASC),
  CONSTRAINT `fk_tax_outlet1`
    FOREIGN KEY (`store_id`, `outlet_id`)
    REFERENCES `outlet` (`store_id`, `id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);
-- +mig Down
DROP TABLE `tax`;
