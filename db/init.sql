-- CREATE TABLE users (
--                        id UUID PRIMARY KEY,
--                        email TEXT UNIQUE NOT NULL,
--                        password_hash TEXT NOT NULL,
--                        name TEXT NOT NULL,
--                        role_id INT REFERENCES roles(id),
--                        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
-- );

CREATE TABLE categories (
                            id UUID PRIMARY KEY,
                            name TEXT NOT NULL,
                            created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE orders (
                        id UUID PRIMARY KEY,
                        user_id UUID NOT NULL,
                        total_price DECIMAL NOT NULL,
                        delivery_price DECIMAL NOT NULL,
                        currency TEXT NOT NULL,
                        status TEXT NOT NULL,
                        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE products (
                          id SERIAL PRIMARY KEY,
                          name VARCHAR(100) NOT NULL,
                          description TEXT,
                          price NUMERIC(10, 2) NOT NULL,
                          category_id INTEGER,
                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO products (name, price) VALUES
                                       ('Go Book', 29.99),
                                       ('Docker Stickers', 9.90);

-- INSERT INTO orders (product_id, quantity) VALUES
--                                               (1, 3),
--                                               (2, 5);
