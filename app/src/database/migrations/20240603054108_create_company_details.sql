-- +goose Up
CREATE TABLE IF NOT EXISTS company_details
(
    id            BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    company_id    BIGINT UNSIGNED NOT NULL,
    access        VARCHAR(1000)   NOT NULL DEFAULT '',
    opening_hours VARCHAR(255)    NOT NULL DEFAULT '',
    contact       VARCHAR(500)    NOT NULL DEFAULT '',
    thumnails     JSON            NOT NULL DEFAULT (JSON_ARRAY())
);;

-- +goose Down
DROP TABLE company_details;
