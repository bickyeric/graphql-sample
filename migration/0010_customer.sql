-- +mig Up
CREATE TABLE IF NOT EXISTS `customer` (
  `id` INT UNSIGNED NOT NULL,
  `store_id` INT UNSIGNED NOT NULL,
  `phone_number` VARCHAR(15),
  `name` VARCHAR(255),
  `email` VARCHAR(255),
  `birthday` TIMESTAMP,
  `address` VARCHAR(255),
  `city` VARCHAR(45),
  `state` VARCHAR(45),
  `zip` VARCHAR(15),
  PRIMARY KEY (`id`),
  INDEX `fk_customer_store1_idx` (`store_id` ASC),
  CONSTRAINT `fk_customer_store1`
    FOREIGN KEY (`store_id`)
    REFERENCES `store` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);
-- +mig Down
DROP TABLE `customer`;
