-- +mig Up
CREATE TABLE IF NOT EXISTS `store` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `category_id` INT UNSIGNED NOT NULL,
  `type_id` INT UNSIGNED NOT NULL,
  `name` VARCHAR(255) NOT NULL,
  `address` VARCHAR(255) NOT NULL,
  `city` VARCHAR(255) NOT NULL,
  `state` VARCHAR(255) NOT NULL,
  `zip` VARCHAR(255) NOT NULL,
  `email` VARCHAR(50) NOT NULL,
  `description` VARCHAR(255) NOT NULL,
  `website` VARCHAR(150) NOT NULL,
  `twitter` VARCHAR(50) NOT NULL,
  `facebook` VARCHAR(50) NOT NULL,
  `instagram` VARCHAR(50) NOT NULL,
  `image` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_store_type_category1_idx` (`category_id` ASC, `type_id` ASC),
  CONSTRAINT `fk_store_type_category1`
    FOREIGN KEY (`category_id` , `type_id`)
    REFERENCES `store_category` (`id` , `type_id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);
-- +mig Down
DROP TABLE `store`;
