-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

INSERT INTO `frozen_food_stock` (`name`, `quantity`) VALUES
    ('Frozen bananas', '132'),
    ('Frozen pineapple', '56'),
    ('Frozen strawberries', '80'),
    ('Frozen peas', '98'),
    ('Frozen edamame', '230'),
    ('Frozen salmon', '514'),
    ('Frozen chicken', '981'),
    ('Frozen beef', '898'),
    ('Frozen spinach', '432'),
    ('Chicken nugget', '674');

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back

TRUNCATE TABLE `frozen_food_stock`;
