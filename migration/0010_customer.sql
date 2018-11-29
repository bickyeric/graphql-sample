-- +mig Up
CREATE TABLE IF NOT EXISTS `customer` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `store_id` INT UNSIGNED NOT NULL,
  `phone_number` VARCHAR(15) NULL,
  `name` VARCHAR(255) NOT NULL,
  `email` VARCHAR(255) NULL,
  `birthday` TIMESTAMP NULL,
  `address` VARCHAR(255) NULL,
  `city` VARCHAR(45) NULL,
  `state` VARCHAR(45) NULL,
  `zip` VARCHAR(15) NULL,
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
