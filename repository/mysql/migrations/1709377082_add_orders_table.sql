-- +migrate Up
create table orders
(
    id            int primary key AUTO_INCREMENT,
    user_id       int          not null,
    vendor_id     int          not null,
    delivery_time date,
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (vendor_id) REFERENCES vendors (id)

);
-- +migrate Down
DROP TABLE orders;