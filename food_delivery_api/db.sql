drop table order_product;

drop table product;

drop table category;

drop table supplier;

drop table "order";

drop table customer;

CREATE TABLE IF NOT EXISTS category (
                                        id SERIAL PRIMARY KEY,
                                        name VARCHAR(128) NOT NULL,
                                        image BYTEA NOT NULL,
                                        description VARCHAR(258) NOT NULL
);

ALTER TABLE category OWNER TO fooddelivery;

CREATE TABLE IF NOT EXISTS supplier (
                                        id SERIAL PRIMARY KEY,
                                        name VARCHAR(128) NOT NULL,
                                        type VARCHAR(128) NOT NULL,
                                        image BYTEA NOT NULL,
                                        supplier_address VARCHAR(256) NOT NULL,
                                        open_time VARCHAR(5) NOT NULL,
                                        close_time VARCHAR(5) NOT NULL
);

ALTER TABLE supplier OWNER TO fooddelivery;

CREATE TABLE IF NOT EXISTS product (
                                       id SERIAL PRIMARY KEY,
                                       name VARCHAR(128) NOT NULL,
                                       supplier_id INTEGER NOT NULL REFERENCES supplier,
                                       category_id INTEGER NOT NULL REFERENCES category,
                                       price REAL NOT NULL,
                                       image BYTEA NOT NULL,
                                       description VARCHAR(258)
);

ALTER TABLE product OWNER TO fooddelivery;

CREATE TABLE IF NOT EXISTS customer (
                                        id SERIAL PRIMARY KEY,
                                        first_name VARCHAR(56) NOT NULL,
                                        last_name VARCHAR(56) NOT NULL,
                                        username VARCHAR(128) NOT NULL UNIQUE,
                                        age INTEGER NOT NULL,
                                        email VARCHAR(128) NOT NULL UNIQUE,
                                        phone VARCHAR(15) NOT NULL,
                                        password VARCHAR(128) NOT NULL,
                                        customer_address VARCHAR(256) NOT NULL,
                                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE customer OWNER TO fooddelivery;

CREATE TABLE IF NOT EXISTS "order" (
                                       id SERIAL PRIMARY KEY,
                                       customer_id INTEGER NOT NULL REFERENCES customer,
                                       total_price REAL NOT NULL,
                                       payment_method VARCHAR(56) NOT NULL,
                                       order_status VARCHAR(56) NOT NULL,
                                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                       customer_address VARCHAR(256) NOT NULL
);

ALTER TABLE "order" OWNER TO fooddelivery;

CREATE TABLE IF NOT EXISTS order_product (
                                             order_id INTEGER NOT NULL REFERENCES "order",
                                             product_id INTEGER NOT NULL REFERENCES product,
                                             quantity INTEGER NOT NULL DEFAULT 1,
                                             PRIMARY KEY (order_id, product_id)
);

-- Insert sample data into the category table
INSERT INTO category (name, image, description)
VALUES
    ('Pizza', 'categories/pizza.svg', 'Authentic Italian pizzas with a variety of toppings.'),
    ('Burgers', 'categories/burger.svg', 'Juicy and flavorful burger creations to satisfy your cravings.'),
    ('Drinks', 'categories/drinks.svg', 'Wide selection of refreshing and flavorful beverages.'),
    ('Sandwiches', 'categories/sandwich.svg', 'Delicious sandwiches made with high-quality meats and fresh ingredients.'),
    ('Desserts', 'categories/dessert.svg', 'Indulgent desserts and sweet treats for a delightful finish.'),
    ('SavorCombo', 'categories/savorcombo.svg', 'Irresistible combos that offer a mix of flavors to enjoy.');


-- Insert sample data into the supplier table
INSERT INTO supplier (name, image, type, supplier_address, open_time, close_time)
VALUES
    ('Burger King', 'suppliers/burger-king-icon.svg', 'description', 'Budapest, orczy ut 34, Hungary', '06:30', '21:00'),
    ('Starbucks', 'suppliers/starbucks-icon.svg', 'Cafe', 'Budapest, Vaci ut 1, Hungary', '07:00', '22:00'),
    ('Taco Bell', 'suppliers/taco-bell-icon.svg', 'Fast food', 'Budapest, Nagykorut 5, Hungary', '07:30', '20:30'),
    ('McDonald''s', 'suppliers/mcdonalds-icon.svg', 'Fast food', 'Budapest, Andrassy ut 123, Hungary', '08:00', '20:00'),
    ('Subway', 'suppliers/subway-icon.svg', 'Fast food', 'Budapest, Nagy kerut 7, Hungary', '00:00', '24:00'),
    ('Pizza Hut', 'suppliers/pizza-hut-icon.svg', 'Pizzeria', 'Budapest, Szent gellert ter 4, Hungary', '09:00', '21:00');

-- Insert sample data into product table for additional items
INSERT INTO product (name, supplier_id, category_id, price, image, description)
VALUES
    ('BBQ Bliss', 1, 6, 660, 'products/barbecuemenu.png', 'Delicious assortment of grilled favorites.'),
    ('Cheeseburger', 4, 2, 700, 'products/cheeseburger.png', 'Classic cheeseburger with juicy beef patty.'),
    ('Tuna Salad', 3, 6, 520, 'products/tunasalad.png', 'Fresh and tangy tuna salad with greens.'),
    ('Burrito', 5, 4, 800, 'products/burrito.png', 'A satisfying and flavorful burrito creation.'),
    ('Cola 0.5L', 4, 3, 220, 'products/cola05.png', 'Refreshing cola in a convenient 0.5L bottle.'),
    ('Fanta 0.5L', 6, 3, 220, 'products/fanta.png', 'Zesty Fanta in a 0.5L bottle for a burst of flavor.'),
    ('Small Bucket', 3, 6, 1100, 'products/kfcsmallbucket.png', 'A small bucket filled with KFC goodness.'),
    ('Cheesecake', 2, 5, 450, 'products/cheesecake.png', 'Rich and creamy cheesecake dessert.'),
    ('Extra Bucket', 4, 6, 1150, 'products/extrabucket.png', 'An extra-large bucket of fried delights.'),
    ('Ice Cream', 2, 5, 350, 'products/icecream.png', 'Chilled and tasty ice cream for your cravings.'),
    ('Margarita', 6, 1, 780, 'products/margarita.png', 'Savor the classic Margarita pizza.'),
    ('Orange Juice', 3, 3, 250, 'products/orangejuice.png', 'Refreshing and pure orange juice.'),
    ('Pepsi 0.33L', 1, 3, 180, 'products/pepsi033.png', 'Enjoy a fizzy Pepsi in a compact size.'),
    ('Mac Menu', 4, 2, 800, 'products/macmenu.png', 'A satisfying meal with a Mac twist.'),
    ('Frappuccino', 2, 3, 380, 'products/milkfrappuccino.png', 'Creamy milk frappuccino for a delightful treat.'),
    ('Pepperoni', 6, 1, 820, 'products/pepperoni.png', 'Savory pepperoni pizza loaded with flavor.'),
    ('Red Bull', 5, 3, 280, 'products/redbull.png', 'Energy-boosting Red Bull drink for your day.'),
    ('Tacos', 3, 4, 700, 'products/tacos.png', 'A delicious and filling taco meal.');



-- Insert sample data into the customer table
INSERT INTO customer (first_name, last_name, username, age, email, phone, password, customer_address, created_at)
VALUES
    ('Georgiy', 'Iakunenko', 'Goshan3097', 30, 'test@gmail.com', '+43534252535', '', 'Budapest, Main Street 1, Hungary', current_timestamp);


-- Insert sample data into the order table
INSERT INTO "order" (customer_id, total_price, order_status, payment_method, created_at, customer_address)
VALUES
    (1, 600, 'DONE', 'Cash', current_timestamp, 'Budapest, Main Street 1, Hungary');


-- Insert sample data into the order_product table
INSERT INTO order_product (order_id, product_id, quantity)
VALUES
    (1, 1, 2),
    (1, 3, 1);