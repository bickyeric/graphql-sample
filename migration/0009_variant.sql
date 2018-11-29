-- +mig Up
CREATE TABLE IF NOT EXISTS `variant` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `item_id` INT UNSIGNED NOT NULL,
  `price` INT NOT NULL,
  `sku` VARCHAR(45) NOT NULL,
  `stock` INT NOT NULL,
  `track_stock` BOOLEAN NOT NULL DEFAULT 0,
  `alert_at` INT NOT NULL,
  `alert` BOOLEAN NOT NULL DEFAULT 0,
  `cost` INT NOT NULL,
  `track_cogs` BOOLEAN NOT NULL DEFAULT 0,
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
