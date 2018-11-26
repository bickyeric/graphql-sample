-- +mig Up
CREATE TABLE IF NOT EXISTS `tax` (
  `id` INT UNSIGNED NOT NULL,
  `outlet_id` INT UNSIGNED NOT NULL,
  `store_id` INT UNSIGNED NOT NULL,
  `amount` INT,
  PRIMARY KEY (`id`),
  INDEX `fk_tax_outlet1_idx` (`outlet_id` ASC, `store_id` ASC),
  CONSTRAINT `fk_tax_outlet1`
    FOREIGN KEY (`outlet_id` , `store_id`)
    REFERENCES `outlet` (`id` , `store_id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);
-- +mig Down
DROP TABLE `tax`;
