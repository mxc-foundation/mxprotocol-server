-- +migrate Up
CREATE TABLE IF NOT EXISTS topup (
     id SERIAL PRIMARY KEY,
     fk_ext_account_sender INT REFERENCES  ext_account(id) NOT NULL,
     fk_ext_account_receiver INT REFERENCES  ext_account(id) NOT NULL,
     fk_ext_currency INT REFERENCES ext_currency(id) NOT NULL,
     value NUMERIC(28,18) NOT NULL,
     tx_approved_time TIMESTAMP,
     tx_hash varchar (128) NOT NULL UNIQUE
);

CREATE OR REPLACE FUNCTION topup_req_apply (
    v_fk_ext_account_sender INT,
    v_fk_ext_account_receiver INT,
    v_fk_ext_currency INT,
    v_value NUMERIC(28,18),
    v_tx_approved_time TIMESTAMP,
    v_tx_hash VARCHAR(128),
    v_fk_wallet_sender INT,
    v_fk_wallet_receiver INT,
    v_payment_cat PAYMENT_CATEGORY
) RETURNS INT
    LANGUAGE plpgsql
AS $$

declare topup_id INT;

BEGIN

    INSERT INTO topup (
        fk_ext_account_sender,
        fk_ext_account_receiver,
        fk_ext_currency,
        value,
        tx_approved_time,
        tx_hash )
    VALUES (
               v_fk_ext_account_sender ,
               v_fk_ext_account_receiver,
               v_fk_ext_currency,
               v_value,
               v_tx_approved_time,
               v_tx_hash
           )RETURNING id INTO topup_id;


    INSERT INTO internal_tx (
        fk_wallet_sender,
        fk_wallet_receiver,
        payment_cat,
        tx_internal_ref,
        value,
        time_tx )
    VALUES (
               v_fk_wallet_sender,
               v_fk_wallet_receiver,
               v_payment_cat,
               topup_id,
               v_value,
               v_tx_approved_time)
    ;


    UPDATE
        wallet
    SET
        balance = balance + v_value
    WHERE
            id = v_fk_wallet_receiver
    ;

    RETURN topup_id;

END;
$$;

-- +migrate Down