-- +mig Up
CREATE TABLE IF NOT EXISTS `modifier` (
  `id` INT UNSIGNED NOT NULL PRIMARY KEY,
  `name` VARCHAR(255) NOT NULL,
  `outlet_id` INT UNSIGNED NOT NULL,
  `store_id` INT UNSIGNED NOT NULL,
  INDEX `fk_modifier_outlet1_idx` (`outlet_id` ASC, `store_id` ASC),
  CONSTRAINT `fk_modifier_outlet1`
    FOREIGN KEY (`store_id`, `outlet_id`)
    REFERENCES `outlet` (`store_id`, `id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
);
-- +mig Down
DROP TABLE `modifier`;
