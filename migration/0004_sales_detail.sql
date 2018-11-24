-- +mig Up
CREATE TABLE IF NOT EXISTS `sales_detail` (
  `item_id` INT UNSIGNED NOT NULL,
  `sales_id` INT UNSIGNED NOT NULL,
  `price` INT NULL,
  `quantity_purchased` INT NULL DEFAULT 0,
  PRIMARY KEY (`item_id`, `sales_id`),
  INDEX `fk_item_has_sales_sales1_idx` (`sales_id` ASC),
  INDEX `fk_item_has_sales_item_idx` (`item_id` ASC),
  CONSTRAINT `fk_item_has_sales_item`
    FOREIGN KEY (`item_id`)
    REFERENCES `items` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_item_has_sales_sales1`
    FOREIGN KEY (`sales_id`)
    REFERENCES `sales` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE);
-- +mig Down
DROP TABLE `sales_detail`;
