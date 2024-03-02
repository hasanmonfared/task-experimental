-- +migrate Up
create table vendors
(
    id           int primary key AUTO_INCREMENT,
    name         varchar(191) not null,
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +migrate Down
DROP TABLE vendors;