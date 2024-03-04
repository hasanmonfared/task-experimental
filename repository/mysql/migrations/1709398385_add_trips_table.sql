-- +migrate Up
create table trips
(
    id         int primary key AUTO_INCREMENT,
    order_id   int not null unique,
    status        enum ('assigned','at_vendor','picked','delivered'),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (order_id) REFERENCES orders (id),
    INDEX(order_id)
);
-- +migrate Down
DROP TABLE trips;