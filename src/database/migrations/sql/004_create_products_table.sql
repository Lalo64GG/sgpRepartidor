CREATE TABLE Product (
    id BIGINT(20) PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(6,2) NOT NULL,
    supplier_id BIGINT(20),
    FOREIGN KEY (supplier_id) REFERENCES Supplier(Id) ON DELETE CASCADE ON UPDATE CASCADE
);
