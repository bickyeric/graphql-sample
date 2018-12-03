-- +mig Up
CREATE TABLE IF NOT EXISTS `employee` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `outlet_id` INT UNSIGNED NOT NULL,
  `store_id` INT UNSIGNED NOT NULL,
  `first_name` VARCHAR(50) NOT NULL,
  `last_name` VARCHAR(50) NOT NULL,
  `phone_number` VARCHAR(15) NOT NULL,
  `email` VARCHAR(150) NOT NULL,
  `password` VARCHAR(50) NOT NULL,
  `confirmed` BOOLEAN NOT NULL DEFAULT 0,
  `active` BOOLEAN NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  INDEX `fk_employee_outlet1_idx` (`outlet_id` ASC, `store_id` ASC),
  CONSTRAINT `fk_employee_outlet1`
    FOREIGN KEY (`store_id`, `outlet_id`)
    REFERENCES `outlet` (`store_id`, `id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);
-- +mig Down
DROP TABLE `employee`;
