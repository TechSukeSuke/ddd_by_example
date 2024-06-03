-- +goose Up
CREATE TABLE IF NOT EXISTS companies
(
    id          BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    name        VARCHAR(255)  NOT NULL DEFAULT '',
    description VARCHAR(1000) NOT NULL DEFAULT '',
    area        VARCHAR(255)  NOT NULL DEFAULT '',
    genre       VARCHAR(255)  NOT NULL DEFAULT '',
    closing_day VARCHAR(255)  NOT NULL DEFAULT '',
    thumnail    VARCHAR(1000) NOT NULL DEFAULT ''
);

-- +goose Down
DROP TABLE companies;
