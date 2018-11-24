-- +mig Up
CREATE TABLE IF NOT EXISTS `sales` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `total_price` INT NULL DEFAULT 0,
  `status` VARCHAR(45) NULL DEFAULT "",
  `comment` TEXT NULL,
  `created_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
  `customer_id` INT UNSIGNED NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_sales_customers1_idx` (`customer_id` ASC),
  CONSTRAINT `fk_sales_customers1`
    FOREIGN KEY (`customer_id`)
    REFERENCES `customers` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE);
-- +mig Down
DROP TABLE `sales`;
