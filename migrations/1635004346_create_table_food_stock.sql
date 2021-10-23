-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE `frozen_food_stock` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(100) NOT NULL,
  `count` INT NOT NULL,
  PRIMARY KEY (`id`));

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE `frozen_food_stock`;
