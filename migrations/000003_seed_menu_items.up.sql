INSERT INTO menu_items (id, name, description, category, base_price, image_url, stock_count, customizations) VALUES
(1, 'Classic Espresso', 'Rich, bold shot of pure coffee perfection. Our signature blend roasted locally.', 'Coffee', 3.50, 'https://images.unsplash.com/photo-1510591509098-f4fdc6d0ff04', 50, '[1, 2]'),
(2, 'Cappuccino', 'Equal parts espresso, steamed milk, and velvety foam. A morning classic.', 'Coffee', 4.50, 'https://images.unsplash.com/photo-1572442388796-11668a67e53d', 50, '[1, 2, 3]'),
(3, 'Caramel Latte', 'Smooth espresso with steamed milk and house-made caramel syrup. Sweet indulgence.', 'Coffee', 5.25, 'https://images.unsplash.com/photo-1461023058943-07fcbe16d735', 50, '[1, 2, 3, 4]'),
(4, 'Iced Americano', 'Espresso shots over ice, topped with cold water. Simple, strong, refreshing.', 'Coffee', 4.00, 'https://images.unsplash.com/photo-1517487881594-2787fef5ebf7', 50, '[1, 2, 3]'),
(5, 'Flat White', 'Velvety microfoam poured over a double ristretto. Bold yet smooth.', 'Coffee', 4.75, 'https://images.unsplash.com/photo-1559496417-e7f25c6c740c', 50, '[2, 3]'),
(6, 'Mocha', 'Rich chocolate meets espresso with steamed milk. Topped with whipped cream.', 'Coffee', 5.50, 'https://images.unsplash.com/photo-1607260550778-aa9d29444ce1', 50, '[1, 2, 3, 4, 5]'),
(8, 'Vanilla Latte', 'Creamy latte sweetened with Madagascar vanilla. Comfort in a cup.', 'Coffee', 5.00, 'https://images.unsplash.com/photo-1579992357154-faf4bde95b3d', 50, '[1, 2, 3, 4]'),

(9, 'English Breakfast Tea', 'Full-bodied black tea blend. Perfect morning companion.', 'Tea', 3.00, 'https://images.unsplash.com/photo-1594631661960-247f9c2f5e50', 50, '[1, 2]'),
(10, 'Chamomile Tea', 'Calming herbal infusion with gentle floral notes. Naturally caffeine-free.', 'Tea', 3.25, 'https://images.unsplash.com/photo-1597318112274-6e94e4c2f908', 50, '[1]'),
(11, 'Matcha Latte', 'Premium Japanese matcha whisked with steamed milk. Earthy and energizing.', 'Tea', 5.25, 'https://images.unsplash.com/photo-1536013734839-4187e92e0ff4', 50, '[1, 2, 3, 4]'),

(14, 'Butter Croissant', 'Flaky, buttery layers baked fresh every morning. Light and golden.', 'Pastries', 3.75, 'https://images.unsplash.com/photo-1555507036-ab1f4038808a', 30, '[]'),
(17, 'Banana Bread', 'House-made with ripe bananas and walnuts. Moist and comforting.', 'Pastries', 3.75, 'https://images.unsplash.com/photo-1606890737304-57a1ca8a5b62', 20, '[]'),
(19, 'Cinnamon Roll', 'Soft, pillowy roll swirled with cinnamon sugar and topped with cream cheese frosting.', 'Pastries', 4.50, 'https://images.unsplash.com/photo-1626094309830-abbb0c99da4a', 25, '[]'),

(21, 'Breakfast Sandwich', 'Scrambled eggs, cheddar cheese, and your choice of bacon or sausage on a toasted English muffin.', 'Food', 7.50, 'https://images.unsplash.com/photo-1619096252214-ef06c45683e3', 30, '[7]'),
(23, 'Bagel with Cream Cheese', 'Fresh-baked bagel with house-whipped cream cheese. Simple perfection.', 'Food', 4.50, 'https://images.unsplash.com/photo-1600629080398-c49842021b8a', 40, '[8]'),

(24, 'Hot Chocolate', 'Rich Belgian chocolate melted into steamed milk. Topped with whipped cream.', 'Specialty', 4.50, 'https://images.unsplash.com/photo-1517578239113-b03992dcdd25', 50, '[2, 5]'),
(27, 'Affogato', 'Vanilla gelato drowned in a hot shot of espresso. Italian dessert-drink hybrid.', 'Specialty', 5.50, 'https://images.unsplash.com/photo-1514849302-984523450cf4', 20, '[1]');

SELECT setval('menu_items_id_seq', (SELECT MAX(id) FROM menu_items));