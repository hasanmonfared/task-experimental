-- +migrate Up
create table delay_reports
(
    id            int primary key AUTO_INCREMENT,
    vendor_id     int not null,
    order_id      int not null,
    agent_id      int,
    delay_check   bool      default false,
    delivery_time timestamp,
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (order_id) REFERENCES orders (id),
    FOREIGN KEY (agent_id) REFERENCES agents (id),
    FOREIGN KEY (vendor_id) REFERENCES vendors (id)

);
-- +migrate Down
DROP TABLE delay_reports;