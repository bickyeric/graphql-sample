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
  `website` VARCHAR(150) NULL,
  `twitter` VARCHAR(50) NULL,
  `facebook` VARCHAR(50) NULL,
  `instagram` VARCHAR(50) NULL,
  `image` VARCHAR(255) NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_store_type_category1_idx` (`category_id` ASC, `type_id` ASC),
  CONSTRAINT `fk_store_type_category1`
    FOREIGN KEY (`type_id`, `category_id`)
    REFERENCES `store_category` (`type_id`, `id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);
-- +mig Down
DROP TABLE `store`;
