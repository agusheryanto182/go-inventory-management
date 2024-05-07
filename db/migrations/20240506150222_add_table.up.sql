CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS staffs (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(50) NOT NULL, 
    password VARCHAR(255) NOT NULL,
    phone_number VARCHAR(16) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
) WITH (fillfactor=70);

CREATE INDEX IF NOT EXISTS idx_id_phone_number ON staffs(id, phone_number) WITH (fillfactor=100);

CREATE TABLE IF NOT EXISTS products(
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(30) NOT NULL,
    sku VARCHAR(30) NOT NULL,
    category VARCHAR(30) NOT NULL,
    image_url TEXT NOT NULL, 
    notes VARCHAR(200) NOT NULL,
    price FLOAT NOT NULL,
    stock INT NOT NULL,
    location VARCHAR(200) NOT NULL,
    is_available BOOLEAN NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
) WITH (fillfactor=60);

CREATE INDEX IF NOT EXISTS idx_id_sku_name_category_location_is_available_price_stock_created_at ON products(id, sku, name, category, location, is_available, price, stock, created_at) WITH (fillfactor=100);

CREATE TABLE IF NOT EXISTS customers (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(50) NOT NULL,
    phone_number VARCHAR(16) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
) WITH (fillfactor=70);

CREATE INDEX IF NOT EXISTS idx_id_phone_number_name_created_at ON customers(id, phone_number, name, created_at) WITH (fillfactor=100);

CREATE TABLE IF NOT EXISTS checkouts(
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    customer_id UUID NOT NULL,
    product_id UUID NOT NULL,
    quantity INT NOT NULL,
    paid FLOAT NOT NULL,
    change FLOAT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
) WITH (fillfactor=60);
