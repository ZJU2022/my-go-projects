-- 创建表结构
CREATE TABLE users (
    id INT PRIMARY KEY,
    username VARCHAR(50),
    email VARCHAR(50)
);

CREATE TABLE products (
    id INT PRIMARY KEY,
    name VARCHAR(100),
    price DECIMAL(10, 2)
);

CREATE TABLE orders (
    id INT PRIMARY KEY,
    user_id INT,
    product_id INT,
    quantity INT,
    order_date DATE,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);

-- 插入数据到 users 表
INSERT INTO users (id, username, email) VALUES
(1, 'john_doe', 'john@example.com'),
(2, 'jane_doe', 'jane@example.com'),
(3, 'alice', 'alice@example.com');

-- 插入数据到 products 表
INSERT INTO products (id, name, price) VALUES
(1, 'Laptop', 899.99),
(2, 'Smartphone', 499.99),
(3, 'Tablet', 299.99);

-- 插入数据到 orders 表
INSERT INTO orders (id, user_id, product_id, quantity, order_date) VALUES
(1, 1, 1, 1, '2023-07-01'),
(2, 2, 2, 2, '2023-07-02'),
(3, 3, 3, 1, '2023-07-03');

