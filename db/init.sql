CREATE TABLE IF NOT EXISTS products (
                                        id SERIAL PRIMARY KEY,
                                        name VARCHAR(255),
    price NUMERIC
    );

INSERT INTO products (name, price) VALUES
                                       ('Go Book', 29.99),
                                       ('Docker Stickers', 9.90);

CREATE TABLE IF NOT EXISTS orders (
                                      id SERIAL PRIMARY KEY,
                                      product_id INT REFERENCES products(id),
    quantity INT
    );

INSERT INTO orders (product_id, quantity) VALUES
                                              (1, 3),
                                              (2, 5);
