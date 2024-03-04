-- +migrate Up
create table orders
(
    id            int primary key AUTO_INCREMENT,
    user_id       int                                        not null,
    vendor_id     int                                        not null,
    delivery_time int,
    status        enum ('submit','in_doing','ready_to_send') not null,
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (vendor_id) REFERENCES vendors (id),
    INDEX (user_id, vendor_id)


);
-- +migrate Down
DROP TABLE orders;