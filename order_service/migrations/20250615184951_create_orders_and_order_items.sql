-- +goose Up
CREATE TABLE orders (
                        id UUID PRIMARY KEY,
                        user_id UUID NOT NULL,
                        total_price DECIMAL NOT NULL,
                        delivery_price DECIMAL NOT NULL,
                        currency TEXT NOT NULL,
                        status TEXT NOT NULL,
                        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE order_items (
                             id UUID PRIMARY KEY,
                             order_id UUID NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
                             product_id UUID NOT NULL,
                             quantity INT NOT NULL,
                             unit_price DECIMAL NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS order_items;
DROP TABLE IF EXISTS orders;
