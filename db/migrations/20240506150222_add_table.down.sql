DROP TABLE IF EXISTS checkouts;
DROP INDEX IF EXISTS idx_id_phone_number_name_created_at;

DROP TABLE IF EXISTS customers;

DROP INDEX IF EXISTS idx_id_sku_name_category_location_is_available_price_stock_created_at;
DROP TABLE IF EXISTS products;

DROP INDEX IF EXISTS idx_id_phone_number;
DROP TABLE IF EXISTS staffs;

DROP EXTENSION IF EXISTS "uuid-ossp";