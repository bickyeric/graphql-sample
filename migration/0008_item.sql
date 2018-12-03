-- +mig Up
CREATE TABLE IF NOT EXISTS `item` (
  `id` INT UNSIGNED NOT NULL,
  `outlet_id` INT UNSIGNED NOT NULL,
  `store_id` INT UNSIGNED NOT NULL,
  `name` VARCHAR(255) NOT NULL,
  `category_id` INT UNSIGNED NULL,
  `description` VARCHAR(255) NOT NULL,
  `modifier_id` INT UNSIGNED NULL,
  `image` VARCHAR(255) NULL,
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
    FOREIGN KEY (`store_id`, `outlet_id`)
    REFERENCES `outlet` (`store_id`, `id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);
-- +mig Down
DROP TABLE `item`;
