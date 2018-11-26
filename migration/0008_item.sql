-- +mig Up
CREATE TABLE IF NOT EXISTS `item` (
  `id` INT UNSIGNED NOT NULL,
  `outlet_id` INT UNSIGNED NOT NULL,
  `store_id` INT UNSIGNED NOT NULL,
  `name` VARCHAR(255),
  `category_id` INT UNSIGNED,
  `description` VARCHAR(255),
  `modifier_id` INT UNSIGNED NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_item_category1_idx` (`category_id` ASC),
  INDEX `fk_item_modifier1_idx` (`modifier_id` ASC),
  INDEX `fk_item_outlet1_idx` (`outlet_id` ASC, `store_id` ASC),
  CONSTRAINT `fk_item_category1`
    FOREIGN KEY (`category_id`)
    REFERENCES `item_category` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_item_modifier1`
    FOREIGN KEY (`modifier_id`)
    REFERENCES `modifier` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_item_outlet1`
    FOREIGN KEY (`outlet_id` , `store_id`)
    REFERENCES `outlet` (`id` , `store_id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);
-- +mig Down
DROP TABLE `item`;
