DROP TABLE IF EXISTS checkouts CASCADE;
DROP INDEX IF EXISTS idx_id_phone_number_name_created_at;

DROP TABLE IF EXISTS checkout_items;
DROP INDEX IF EXISTS idx_checkout_id;

DROP TABLE IF EXISTS products CASCADE;
DROP INDEX IF EXISTS idx_id_sku_name_category_location_is_available_price_stock_created_at;

DROP TABLE IF EXISTS staffs CASCADE;
DROP INDEX IF EXISTS idx_id_phone_number;

DROP TABLE IF EXISTS customers CASCADE;
DROP INDEX IF EXISTS idx_id_phone_number_name_created_at;

DROP EXTENSION IF EXISTS "uuid-ossp";
