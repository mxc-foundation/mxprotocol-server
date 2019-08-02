-- +migrate Up
INSERT INTO wallet (
    fk_org_la ,
    type,
    balance
) VALUES (0, "SUPER_ADMIN", 8.888), (1, "USER", 8.888);

-- +migrate Down
DELETE FROM wallet WHERE fk_org_la = 0;
DELETE FROM wallet WHERE fk_org_la = 1;