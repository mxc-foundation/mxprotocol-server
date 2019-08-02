-- +migrate Up
INSERT INTO withdraw_fee (
    fk_ext_currency,
    fee,
    insert_time,
    status
) VALUES (1, 0.2, now(), "ACTIVE");

-- +migrate Down
DELETE FROM withdraw_fee where fk_ext_currency = 1;