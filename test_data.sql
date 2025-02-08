-------------------------------------------
-- Customers
-------------------------------------------
INSERT INTO customers (username, role, hashed_password)
VALUES ('john_doe', 'user', '$2a$10$ABC123'),
       ('alice_smith', 'admin', '$2a$10$DEF456'),
       ('robert_johnson', 'user', '$2a$10$GHI789'),
       ('emily_wilson', 'user', '$2a$10$JKL012'),
       ('michael_brown', 'user', '$2a$10$MNO345'),
       ('sarah_davis', 'admin', '$2a$10$PQR678'),
       ('david_miller', 'user', '$2a$10$STU901'),
       ('linda_anderson', 'user', '$2a$10$VWX234'),
       ('james_taylor', 'user', '$2a$10$YZA567'),
       ('patricia_thomas', 'admin', '$2a$10$BCD890'),
       ('william_moore', 'user', '$2a$10$EFG123'),
       ('jennifer_jackson', 'user', '$2a$10$HIJ456'),
       ('richard_white', 'user', '$2a$10$KLM789'),
       ('mary_harris', 'user', '$2a$10$NOP012'),
       ('charles_martin', 'admin', '$2a$10$QRS345'),
       ('susan_thompson', 'user', '$2a$10$TUV678'),
       ('joseph_garcia', 'user', '$2a$10$WXY901'),
       ('margaret_martinez', 'user', '$2a$10$ZAB234'),
       ('thomas_robinson', 'user', '$2a$10$CDE567'),
       ('dorothy_clark', 'admin', '$2a$10$FGH890'),
       ('daniel_rodriguez', 'user', '$2a$10$IJK123'),
       ('lisa_lewis', 'user', '$2a$10$LMN456'),
       ('paul_lee', 'user', '$2a$10$OPQ789'),
       ('nancy_walker', 'user', '$2a$10$RST012'),
       ('mark_hall', 'admin', '$2a$10$UVW345'),
       ('betty_allen', 'user', '$2a$10$XYZ678'),
       ('donald_young', 'user', '$2a$10$ABC901'),
       ('sandra_king', 'user', '$2a$10$DEF234'),
       ('george_wright', 'user', '$2a$10$GHI567'),
       ('karen_scott', 'admin', '$2a$10$JKL890'),
       ('matthew_green', 'user', '$2a$10$MNO123'),
       ('ashley_adams', 'user', '$2a$10$PQR456'),
       ('kevin_nelson', 'user', '$2a$10$STU789'),
       ('kimberly_hill', 'user', '$2a$10$VWX012'),
       ('brian_ramirez', 'admin', '$2a$10$YZA345'),
       ('emma_campbell', 'user', '$2a$10$BCD678'),
       ('ronald_mitchell', 'user', '$2a$10$EFG901'),
       ('amanda_roberts', 'user', '$2a$10$HIJ234'),
       ('jason_carter', 'user', '$2a$10$KLM567'),
       ('stephanie_phillips', 'admin', '$2a$10$NOP890'),
       ('jeffrey_evans', 'user', '$2a$10$QRS123'),
       ('rebecca_turner', 'user', '$2a$10$TUV456'),
       ('ryan_torres', 'user', '$2a$10$WXY789'),
       ('sharon_parker', 'user', '$2a$10$ZAB012'),
       ('gary_collins', 'admin', '$2a$10$CDE345'),
       ('cynthia_edwards', 'user', '$2a$10$FGH678'),
       ('timothy_stewart', 'user', '$2a$10$IJK901'),
       ('kathryn_flores', 'user', '$2a$10$LMN234'),
       ('eric_morris', 'user', '$2a$10$OPQ567'),
       ('christine_nguyen', 'admin', '$2a$10$RST890');

-------------------------------------------
-- Products
-------------------------------------------
INSERT INTO products (name, price, quantity)
VALUES ('UltraBook Pro', 1299.99, 15),
       ('Wireless Keyboard', 59.99, 85),
       ('Noise-Canceling Headphones', 299.00, 40),
       ('Smartphone X', 899.00, 25),
       ('4K Monitor', 450.00, 30),
       ('Ergonomic Chair', 349.99, 10),
       ('Gaming Mouse', 79.99, 60),
       ('Bluetooth Speaker', 129.50, 75),
       ('External SSD 1TB', 149.99, 50),
       ('E-Reader', 199.00, 20),
       ('Fitness Tracker', 129.99, 45),
       ('Coffee Maker', 89.95, 30),
       ('Electric Kettle', 45.00, 65),
       ('Air Fryer', 129.99, 25),
       ('Robot Vacuum', 399.00, 15),
       ('Smart Watch', 249.00, 35),
       ('Wireless Earbuds', 179.00, 50),
       ('Portable Charger', 39.99, 100),
       ('Backpack', 79.00, 40),
       ('Travel Mug', 24.99, 120),
       ('Desk Lamp', 55.00, 60),
       ('Monitor Stand', 89.00, 30),
       ('USB-C Hub', 65.00, 45),
       ('Mechanical Keyboard', 149.99, 25),
       ('Webcam HD', 99.00, 50),
       ('Graphics Tablet', 299.00, 20),
       ('Smart Bulb', 29.99, 80),
       ('Electric Toothbrush', 89.95, 40),
       ('Hair Dryer', 69.99, 30),
       ('Steam Iron', 54.99, 35),
       ('Yoga Mat', 34.99, 50),
       ('Resistance Bands', 29.00, 65),
       ('Dumbbell Set', 149.00, 15),
       ('Treadmill', 999.00, 5),
       ('Camping Tent', 199.00, 20),
       ('Sleeping Bag', 89.00, 25),
       ('Hiking Backpack', 159.00, 15),
       ('Water Filter', 79.99, 30),
       ('Solar Charger', 129.00, 20),
       ('Binoculars', 149.00, 10),
       ('Digital Camera', 599.00, 8),
       ('Action Camera', 299.00, 12),
       ('Drone Pro', 799.00, 5),
       ('VR Headset', 499.00, 7),
       ('Board Game', 49.99, 40),
       ('Puzzle 1000pc', 34.99, 30),
       ('Cookbook', 29.95, 50),
       ('Novel', 19.99, 75),
       ('Gardening Tools', 89.00, 25),
       ('Plant Fertilizer', 24.99, 60);

-------------------------------------------
-- Orders
-------------------------------------------
INSERT INTO orders (customer_id, status)
VALUES (1, 'pending'),
       (2, 'confirmed'),
       (3, 'cancelled'),
       (4, 'pending'),
       (5, 'confirmed'),
       (6, 'pending'),
       (7, 'confirmed'),
       (8, 'cancelled'),
       (9, 'pending'),
       (10, 'confirmed'),
       (11, 'pending'),
       (12, 'confirmed'),
       (13, 'cancelled'),
       (14, 'pending'),
       (15, 'confirmed'),
       (16, 'pending'),
       (17, 'confirmed'),
       (18, 'cancelled'),
       (19, 'pending'),
       (20, 'confirmed'),
       (21, 'pending'),
       (22, 'confirmed'),
       (23, 'cancelled'),
       (24, 'pending'),
       (25, 'confirmed'),
       (26, 'pending'),
       (27, 'confirmed'),
       (28, 'cancelled'),
       (29, 'pending'),
       (30, 'confirmed'),
       (31, 'pending'),
       (32, 'confirmed'),
       (33, 'cancelled'),
       (34, 'pending'),
       (35, 'confirmed'),
       (36, 'pending'),
       (37, 'confirmed'),
       (38, 'cancelled'),
       (39, 'pending'),
       (40, 'confirmed'),
       (41, 'pending'),
       (42, 'confirmed'),
       (43, 'cancelled'),
       (44, 'pending'),
       (45, 'confirmed'),
       (46, 'pending'),
       (47, 'confirmed'),
       (48, 'cancelled'),
       (49, 'pending'),
       (50, 'confirmed');

-------------------------------------------
-- Order Products
-------------------------------------------
INSERT INTO order_products (order_id, product_id, quantity)
VALUES    (1, 3, 2),
(1, 7, 1),
(1, 12, 3),
(1, 45, 1),
(2, 5, 1),
(2, 18, 2),
(3, 9, 4),
(4, 2, 1),
(4, 22, 2),
(5, 14, 3),
(5, 27, 2),
(6, 33, 1),
(6, 41, 1),
(7, 6, 2),
(7, 19, 1),
(8, 50, 3),
(9, 8, 2),
(10, 11, 1),
(10, 25, 2),
(11, 37, 1),
(12, 4, 1),
(12, 16, 1),
(13, 29, 2),
(14, 35, 1),
(15, 10, 3),
(16, 13, 1),
(17, 21, 2),
(18, 38, 1),
(19, 26, 1),
(20, 31, 3),
(21, 40, 2),
(22, 17, 1),
(23, 49, 2),
(24, 36, 1),
(25, 44, 1),
(27, 23, 1),
(28, 39, 1),
(29, 28, 3),
(30, 34, 1),
(31, 24, 2),
(32, 48, 1),
(33, 43, 1),
(34, 32, 2),
(35, 47, 1),
(36, 42, 3),
(37, 1, 1),
(38, 15, 2),
(39, 20, 1),
(40, 46, 1),
(41, 50, 2),
(42, 9, 1),
(43, 7, 3),
(44, 3, 1),
(45, 18, 2),
(46, 5, 1),
(47, 12, 2),
(48, 22, 1),
(49, 27, 3),
(50, 33, 1),
(1, 19, 2),
(2, 28, 1),
(3, 37, 3),
(4, 44, 1),
(5, 50, 2),
(6, 6, 1),
(7, 13, 1),
(8, 21, 2),
(9, 29, 1),
(10, 35, 3),
(11, 40, 2),
(12, 2, 1),
(13, 17, 1),
(14, 25, 2),
(15, 31, 1),
(16, 38, 1),
(17, 45, 3),
(18, 49, 2),
(19, 5, 1),
(20, 10, 2),
(21, 16, 1),
(22, 23, 2),
(23, 30, 1),
(25, 41, 3),
(26, 46, 2),
(27, 4, 1),
(28, 11, 1),
(29, 19, 2),
(30, 26, 1),
(31, 32, 3),
(32, 39, 1),
(33, 44, 2),
(34, 50, 1),
(35, 7, 1),
(36, 14, 2),
(37, 20, 3),
(38, 27, 1),
(39, 34, 2),
(40, 40, 1),
(41, 45, 1),
(42, 2, 2),
(43, 9, 1),
(44, 16, 3),
(45, 22, 1),
(46, 29, 2),
(47, 35, 1),
(48, 41, 1),
(49, 47, 2),
(50, 4, 1),
(1, 8, 2),
(2, 15, 1),
(3, 24, 3),
(4, 31, 1),
(5, 38, 2),
(6, 44, 1),
(7, 50, 3),
(8, 6, 1),
(9, 13, 2),
(10, 20, 1),
(11, 27, 1),
(12, 34, 2),
(13, 40, 3),
(14, 46, 1),
(15, 3, 2),
(16, 10, 1),
(17, 17, 2),
(18, 24, 1),
(19, 31, 3),
(20, 38, 1),
(21, 45, 2),
(22, 2, 1),
(23, 9, 1),
(24, 16, 2),
(25, 23, 1),
(26, 30, 3),
(27, 37, 1),
(28, 44, 2),
(29, 1, 1),
(30, 8, 2),
(31, 15, 1),
(32, 22, 3),
(33, 29, 1),
(34, 36, 2),
(35, 43, 1),
(36, 50, 1),
(37, 7, 2),
(38, 14, 1),
(39, 21, 3),
(40, 28, 1),
(41, 35, 2),
(42, 42, 1),
(43, 49, 3),
(44, 6, 1),
(45, 13, 2),
(46, 20, 1),
(47, 27, 2),
(48, 34, 1),
(49, 41, 3),
(50, 48, 1)