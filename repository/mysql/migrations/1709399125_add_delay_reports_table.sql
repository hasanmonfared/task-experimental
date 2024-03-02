-- +migrate Up
create table delay_reports
(
    id         int primary key AUTO_INCREMENT,
    order_id   int not null,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (order_id) REFERENCES orders (id)

);
-- +migrate Down
DROP TABLE delay_reports;