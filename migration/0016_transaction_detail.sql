-- +mig Up
CREATE TABLE IF NOT EXISTS `transaction_payment` (
  `transaction_id` INT UNSIGNED NOT NULL,
  `type` VARCHAR(45),
  `cash` INT,
  `comment` VARCHAR(255),
  PRIMARY KEY (`transaction_id`),
  INDEX `fk_transaction1_idx` (`transaction_id` ASC),
  CONSTRAINT `fk_transaction_payment1`
    FOREIGN KEY (`transaction_id`)
    REFERENCES `transaction` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);
-- +mig Down
DROP TABLE `transaction_payment`;
