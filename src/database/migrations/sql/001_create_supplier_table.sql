CREATE TABLE IF NOT EXISTS Supplier (
    Id BIGINT(20) PRIMARY KEY AUTO_INCREMENT,
    Name VARCHAR(255) NOT NULL,
    ContactInfo VARCHAR(255),
    Address VARCHAR(255),
    Email VARCHAR(255) UNIQUE NOT NULL,
    Password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);