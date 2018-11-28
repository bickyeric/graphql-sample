-- +mig Up
CREATE TABLE IF NOT EXISTS `item_category` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45),
  `outlet_id` INT UNSIGNED NOT NULL,
  `store_id` INT UNSIGNED NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_item_category_outlet1_idx` (`outlet_id` ASC, `store_id` ASC),
  CONSTRAINT `fk_item_category_outlet1`
    FOREIGN KEY (`outlet_id` , `store_id`)
    REFERENCES `outlet` (`id` , `store_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
);
-- +mig Down
DROP TABLE `item_category`;
