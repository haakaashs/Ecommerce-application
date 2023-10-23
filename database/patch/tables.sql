
-- source this file for the first time.

-- user table create query
CREATE TABLE users (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    is_active TINYINT(1) NOT NULL,
    phone BIGINT NOT NULL UNIQUE,
    address VARCHAR(255),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);


-- product table create query
CREATE TABLE products (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    product_category ENUM('premium', 'regular', 'budget') NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    inventory_qty INT NOT NULL,
    is_active TINYINT(1) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);


-- cart table create query
CREATE TABLE cart (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT NOT NULL,
    cart_products JSON NOT NULL,
    qty INT NOT NULL,
    total_amount DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);


-- order table create table
CREATE TABLE orders (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT NOT NULL,
    order_products JSON NOT NULL,
    order_status ENUM('placed', 'dispatched', 'completed', 'cancelled') NOT NULL,
    dispatched DATETIME,
    order_value DECIMAL(10, 2) NOT NULL,
    is_active TINYINT(1) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);