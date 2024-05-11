CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS staffs (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(50) NOT NULL, 
    password VARCHAR(255) NOT NULL,
    phone_number VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
) WITH (fillfactor=70);

CREATE INDEX IF NOT EXISTS idx_id_phone_number ON staffs(id, phone_number) WITH (fillfactor=100);

CREATE TABLE IF NOT EXISTS products(
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(30) NOT NULL,
    sku VARCHAR(30) NOT NULL,
    category VARCHAR(30) NOT NULL,
    image_url TEXT NOT NULL, 
    stock INT NOT NULL,
    notes VARCHAR(200) NOT NULL,
    price INT NOT NULL,
    location VARCHAR(200) NOT NULL,
    is_available BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
) WITH (fillfactor=60);

CREATE INDEX IF NOT EXISTS idx_id_sku_name_category_location_is_available_price_stock_created_at ON products(id, sku, name, category, location, is_available, price, stock, created_at) WITH (fillfactor=100);

CREATE TABLE IF NOT EXISTS customers (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    phone_number VARCHAR(50) NOT NULL,
    name VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
) WITH (fillfactor=70);

CREATE INDEX IF NOT EXISTS idx_id_phone_number_name_created_at ON customers(id, phone_number, name, created_at) WITH (fillfactor=100);

CREATE TABLE IF NOT EXISTS checkouts (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    customer_id UUID NOT NULL REFERENCES customers(id) ON DELETE CASCADE,
    paid INT NOT NULL,
    change INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
) WITH (fillfactor=60);


CREATE TABLE IF NOT EXISTS checkout_items (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    checkout_id UUID NOT NULL REFERENCES checkouts(id) ON DELETE CASCADE,
    product_id UUID NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    quantity INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
) WITH (fillfactor=60);

CREATE INDEX IF NOT EXISTS idx_checkout_id ON checkout_items(checkout_id) WITH (fillfactor=100)