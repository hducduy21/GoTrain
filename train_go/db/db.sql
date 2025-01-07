-- Mock PostgreSQL SQL Script

-- Create a mock table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert mock data into the users table
INSERT INTO users (username, email) VALUES
('john_doe', 'john_doe@example.com'),
('jane_smith', 'jane_smith@example.com'),
('alice_wonder', 'alice_wonder@example.com'),
('bob_builder', 'bob_builder@example.com');

-- Create another mock table
CREATE TABLE orders (
    order_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    product_name VARCHAR(100) NOT NULL,
    order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    amount DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

-- Insert mock data into the orders table
INSERT INTO orders (user_id, product_name, amount) VALUES
(1, 'Laptop', 999.99),
(2, 'Smartphone', 699.99),
(3, 'Headphones', 199.99),
(1, 'Monitor', 299.99),
(4, 'Tablet', 399.99);

-- Query to verify data
SELECT * FROM users;
SELECT * FROM orders;
