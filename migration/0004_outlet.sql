-- +mig Up
CREATE TABLE IF NOT EXISTS `outlet` (
  `id` INT UNSIGNED NOT NULL,
  `name` VARCHAR(255) NOT NULL,
  `address` VARCHAR(255) NOT NULL,
  `phone_number` VARCHAR(30) NOT NULL,
  `city` VARCHAR(100) NOT NULL,
  `state` VARCHAR(150) NOT NULL,
  `status` BOOLEAN NOT NULL DEFAULT 1,
  `zip` VARCHAR(15) NOT NULL,
  `store_id` INT UNSIGNED NOT NULL,
  PRIMARY KEY (`store_id`, `id`),
  INDEX `fk_outlet_store1_idx` (`store_id` ASC),
  CONSTRAINT `fk_outlet_store1`
    FOREIGN KEY (`store_id`)
    REFERENCES `store` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);
-- +mig Down
DROP TABLE `outlet`;
