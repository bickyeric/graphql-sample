-- +mig Up
CREATE TABLE IF NOT EXISTS `item_category` (
  `id` INT UNSIGNED NOT NULL,
  `name` VARCHAR(45) NOT NULL,
  `outlet_id` INT UNSIGNED NOT NULL,
  `store_id` INT UNSIGNED NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_item_category_outlet1_idx` (`outlet_id` ASC, `store_id` ASC),
  CONSTRAINT `fk_item_category_outlet1`
    FOREIGN KEY (`store_id`, `outlet_id`)
    REFERENCES `outlet` (`store_id`, `id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
);
-- +mig Down
DROP TABLE `item_category`;
