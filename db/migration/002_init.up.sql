-- +migrate Up
CREATE TABLE products
(
    product_id   SERIAL PRIMARY KEY,
    product_name VARCHAR(255),
    price        INT,
    quantity     INT
);