
-- source this file for the first time.


-- user table create query
CREATE TABLE users (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    role VARCHAR(255) NOT NULL,
    is_active TINYINT(1) NOT NULL,
    phone BIGINT NOT NULL,
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
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    modified_at DATETIME
);


-- cart table create query
CREATE TABLE cart (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT NOT NULL,
    product_id BIGINT NOT NULL,
    qty INT NOT NULL,
    total_amount DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);


-- order table create table
CREATE TABLE orders (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT NOT NULL,
    products JSON NOT NULL,
    order_status ENUM('placed', 'dispatched', 'completed', 'cancelled') NOT NULL,
    dispatched DATETIME,
    order_value DECIMAL(10, 2) NOT NULL,
    is_active TINYINT(1) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    modified_at DATETIME,
    FOREIGN KEY (user_id) REFERENCES users(id)
);