-- +mig Up
CREATE TABLE IF NOT EXISTS `variant` (
  `id` INT UNSIGNED NOT NULL,
  `item_id` INT UNSIGNED NOT NULL,
  `price` INT,
  `sku` VARCHAR(45),
  `stock` INT,
  `track_stock` BOOLEAN,
  `alert_at` INT,
  `alert` BOOLEAN,
  `cost` INT,
  `track_cogs` BOOLEAN,
  PRIMARY KEY (`id`, `item_id`),
  INDEX `fk_variant_item1_idx` (`item_id` ASC),
  CONSTRAINT `fk_variant_item1`
    FOREIGN KEY (`item_id`)
    REFERENCES `item` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);
-- +mig Down
DROP TABLE `variant`;
