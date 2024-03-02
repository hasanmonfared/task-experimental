-- +migrate Up
create table trips
(
    id         int primary key AUTO_INCREMENT,
    order_id   int not null,
    status     varchar(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (order_id) REFERENCES orders (id)

);
-- +migrate Down
DROP TABLE trips;