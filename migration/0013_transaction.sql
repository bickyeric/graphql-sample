-- +mig Up
CREATE TABLE IF NOT EXISTS `transaction` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `receipt_number` VARCHAR(45) NULL,
  `employee_id` INT UNSIGNED NOT NULL,
  `customer_id` INT UNSIGNED NULL,
  `comment` TEXT NULL,
  `status` VARCHAR(15) NOT NULL,
  `discount_id` INT UNSIGNED NULL,
  `tax_id` INT UNSIGNED NULL,
  `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX `fk_transaction_employee1_idx` (`employee_id` ASC),
  INDEX `fk_transaction_customer1_idx` (`customer_id` ASC),
  INDEX `fk_transaction_discount1_idx` (`discount_id` ASC),
  INDEX `fk_transaction_tax1_idx` (`tax_id` ASC),
  CONSTRAINT `fk_transaction_employee1`
    FOREIGN KEY (`employee_id`)
    REFERENCES `employee` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_transaction_customer1`
    FOREIGN KEY (`customer_id`)
    REFERENCES `customer` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_transaction_discount1`
    FOREIGN KEY (`discount_id`)
    REFERENCES `discount` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_transaction_tax1`
    FOREIGN KEY (`tax_id`)
    REFERENCES `tax` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);
-- +mig Down
DROP TABLE `transaction`;
