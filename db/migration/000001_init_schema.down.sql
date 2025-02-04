DROP TABLE IF EXISTS order_products, orders, products, customers CASCADE;
DROP TYPE IF EXISTS order_status, user_role;
DROP FUNCTION IF EXISTS update_modified_column();
DROP FUNCTION IF EXISTS prevent_order_deletion();