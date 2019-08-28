-- +migrate Up
INSERT INTO ext_account (
    fk_wallet,
    fk_ext_currency,
    account_adr,
    insert_time,
    status,
    latest_checked_block
) VALUES (1, 1, "0x9e5cDB74452D3C44324F9704694A46a926d32874", now(), "ACTIVE", 0),
         (2, 1, "0x2bD170F9DbCE6C78bC9a8ff5da5a4273258EaADD", now(), "ACTIVE", 0);

-- +migrate Down
DELETE FROM ext_account WHERE fk_wallet = 1;
DELETE FROM ext_account WHERE fk_wallet = 2;
