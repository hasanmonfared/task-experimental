-- +migrate Up
create table orders
(
    id            int primary key AUTO_INCREMENT,
    user_id       int          not null,
    vendor_id     int          not null,
    order_id      int          not null,
    name          varchar(191) not null,
    delivery_time date,
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (vendor_id) REFERENCES vendors (id),
    FOREIGN KEY (order_id) REFERENCES orders (id)

);
-- +migrate Down
DROP TABLE orders;