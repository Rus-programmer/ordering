CREATE TYPE order_status AS ENUM ('pending', 'confirmed', 'cancelled');
CREATE TYPE user_role AS ENUM ('user', 'admin');

CREATE TABLE customers
(
    id              bigserial PRIMARY KEY,
    username        VARCHAR(100) NOT NULL UNIQUE,
    role            user_role    NOT NULL DEFAULT 'user',
    hashed_password VARCHAR(250) NOT NULL,
    created_at      TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE TABLE orders
(
    id          BIGSERIAL PRIMARY KEY,
    customer_id BIGINT       NOT NULL REFERENCES customers (id) ON DELETE SET NULL,
    status      order_status NOT NULL DEFAULT 'pending',
    is_deleted  BOOLEAN      NOT NULL DEFAULT FALSE,
    created_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE TABLE products
(
    id         BIGSERIAL PRIMARY KEY,
    name       VARCHAR(250)   NOT NULL,
    price      NUMERIC(10, 2) NOT NULL CHECK (price >= 0),
    quantity   int            NOT NULL CHECK (quantity >= 0),
    created_at TIMESTAMPTZ    NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ    NOT NULL DEFAULT NOW()
);

CREATE TABLE order_products
(
    order_id   BIGINT NOT NULL REFERENCES orders (id),
    product_id BIGINT NOT NULL REFERENCES products (id) ON DELETE CASCADE,
    quantity   INT    NOT NULL CHECK (quantity > 0),
    PRIMARY KEY (order_id, product_id)
);

CREATE FUNCTION update_modified_column()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$
    LANGUAGE plpgsql;

CREATE TRIGGER update_customers_updated_at
    BEFORE UPDATE
    ON customers
    FOR EACH ROW
EXECUTE FUNCTION update_modified_column();

CREATE TRIGGER update_products_updated_at
    BEFORE UPDATE
    ON products
    FOR EACH ROW
EXECUTE FUNCTION update_modified_column();

CREATE TRIGGER update_orders_updated_at
    BEFORE UPDATE
    ON orders
    FOR EACH ROW
EXECUTE FUNCTION update_modified_column();

CREATE OR REPLACE FUNCTION prevent_order_deletion()
    RETURNS TRIGGER AS
$$
BEGIN
    RAISE EXCEPTION 'Direct deletion of orders is prohibited. Use soft delete (UPDATE).';
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER prevent_order_deletion_trigger
    BEFORE DELETE
    ON orders
    FOR EACH ROW
EXECUTE FUNCTION prevent_order_deletion();

CREATE INDEX idx_products_name ON products (name);
CREATE INDEX idx_orders_status ON orders (status);