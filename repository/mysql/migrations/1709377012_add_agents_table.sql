-- +migrate Up
create table agents
(
    id           int primary key AUTO_INCREMENT,
    firstname         varchar(191) not null,
    lastname         varchar(191) not null,
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +migrate Down
DROP TABLE agents;