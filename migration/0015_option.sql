-- +mig Up
CREATE TABLE IF NOT EXISTS `option` (
  `id` INT UNSIGNED NOT NULL,
  `name` VARCHAR(255),
  `price` INT,
  `modifier_id` INT UNSIGNED NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_modifier1_idx` (`modifier_id` ASC),
  CONSTRAINT `fk_option_modifier1`
    FOREIGN KEY (`modifier_id`)
    REFERENCES `modifier` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);
-- +mig Down
DROP TABLE `option`;
