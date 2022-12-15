USE
stat_db;

SET
FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS endpoint;
DROP TABLE IF EXISTS stat;
DROP TABLE IF EXISTS account;
SET
FOREIGN_KEY_CHECKS=1;

CREATE TABLE endpoint
(
    id_endpoint INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name        VARCHAR(255) NOT NULL
);

CREATE TABLE stat
(
    id_stat     INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    visited     DATETIME DEFAULT (DATE_FORMAT(NOW(), "%Y-%m-%d %H:%i")),
    fk_endpoint INT UNSIGNED NOT NULL
);

CREATE TABLE account
(
    id_account VARCHAR(255) NOT NULL PRIMARY KEY,
    id_user    VARCHAR(255) NOT NULL,
    acc_limit  INT UNSIGNED NOT NULL,
    openDate   DATETIME DEFAULT (DATE_FORMAT(NOW(), "%Y-%m-%d %H:%i")),
    acc_type   VARCHAR(255) NOT NULL
);

ALTER TABLE stat
    ADD CONSTRAINT fkc_endpoint_stat
        FOREIGN KEY (fk_endpoint)
            REFERENCES endpoint (id_endpoint)
            ON UPDATE CASCADE
            ON DELETE CASCADE;

INSERT INTO endpoint (name)
VALUES ("api/v1/account"),
       ("api/v1/accounts/:type"),
       ("api/v1//accounts/:type/transactions"),
       ("api/v1/account/:accountID"),
       ("api/v1/account/:accountID/deposit"),
       ("api/v1/account/:accountID/withdraw"),
       ("api/v1/account/:accountID/close"),
       ("api/v1/account/:accountID");
