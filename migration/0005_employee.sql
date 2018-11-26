-- +mig Up
CREATE TABLE IF NOT EXISTS `employee` (
  `id` INT UNSIGNED NOT NULL,
  `outlet_id` INT UNSIGNED NOT NULL,
  `store_id` INT UNSIGNED NOT NULL,
  `first_name` VARCHAR(50),
  `last_name` VARCHAR(50),
  `phone_number` VARCHAR(15),
  `email` VARCHAR(150),
  `password` VARCHAR(50),
  `confirmed` BOOLEAN,
  `active` BOOLEAN,
  PRIMARY KEY (`id`),
  INDEX `fk_employee_outlet1_idx` (`outlet_id` ASC, `store_id` ASC),
  CONSTRAINT `fk_employee_outlet1`
    FOREIGN KEY (`outlet_id` , `store_id`)
    REFERENCES `outlet` (`id` , `store_id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);
-- +mig Down
DROP TABLE `employee`;
