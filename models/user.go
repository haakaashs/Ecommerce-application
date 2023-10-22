package models

type User struct {
	Id uint64
}

// CREATE TABLE users (
//     id BIGINT AUTO_INCREMENT PRIMARY KEY,
//     name VARCHAR(255) NOT NULL,
//     password VARCHAR(255) NOT NULL,
//     email VARCHAR(255) NOT NULL,
//     role VARCHAR(255) NOT NULL,
//     is_active TINYINT(1) NOT NULL,
//     phone BIGINT NOT NULL,
//     address VARCHAR(255),
//     created_at DATETIME DEFAULT CURRENT_TIMESTAMP
// );
