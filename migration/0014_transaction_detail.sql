-- +mig Up
CREATE TABLE IF NOT EXISTS `transaction_detail` (
  `id` INT UNSIGNED NOT NULL,
  `transaction_id` INT UNSIGNED NOT NULL,
  `variant_id` INT UNSIGNED NULL,
  `item_id` INT UNSIGNED NULL,
  `quantity` INT,
  `comment` VARCHAR(255),
  `price` INT,
  `option` VARCHAR(255),
  PRIMARY KEY (`id`, `transaction_id`),
  INDEX `fk_transaction1_idx` (`transaction_id` ASC),
  INDEX `fk_variant1_idx` (`variant_id` ASC, `item_id` ASC),
  CONSTRAINT `fk_variant_transaction1`
    FOREIGN KEY (`variant_id` , `item_id`)
    REFERENCES `variant` (`id` , `item_id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_transaction1`
    FOREIGN KEY (`transaction_id`)
    REFERENCES `transaction` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);
-- +mig Down
DROP TABLE `transaction_detail`;
