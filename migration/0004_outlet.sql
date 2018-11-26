-- +mig Up
CREATE TABLE IF NOT EXISTS `outlet` (
  `id` INT UNSIGNED NOT NULL,
  `name` VARCHAR(255),
  `address` VARCHAR(255),
  `phone_number` VARCHAR(30),
  `city` VARCHAR(100),
  `state` VARCHAR(150),
  `status` BOOLEAN,
  `zip` VARCHAR(15),
  `store_id` INT UNSIGNED NOT NULL,
  PRIMARY KEY (`id`, `store_id`),
  INDEX `fk_outlet_store1_idx` (`store_id` ASC),
  CONSTRAINT `fk_outlet_store1`
    FOREIGN KEY (`store_id`)
    REFERENCES `store` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);
-- +mig Down
DROP TABLE `outlet`;
