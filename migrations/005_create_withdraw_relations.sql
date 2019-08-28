-- +migrate Up

CREATE TYPE TX_STATUS AS ENUM (
    'NOT_SENT_TO_PS',
    'PENDING',
    'SUCCESSFUL'
);

CREATE TABLE IF NOT EXISTS withdraw (
    id SERIAL PRIMARY KEY,
    fk_ext_account_sender INT REFERENCES  ext_account(id) NOT NULL,
    fk_ext_account_receiver INT REFERENCES  ext_account(id) NOT NULL,
    fk_ext_currency INT REFERENCES  ext_currency(id) NOT NULL,
    value NUMERIC(28,18) NOT NULL,
    fk_withdraw_fee INT REFERENCES  withdraw(id) NOT NULL,
    tx_sent_time TIMESTAMP NOT NULL,
    tx_stat tx_status NOT NULL,
    tx_approved_time TIMESTAMP,
    fk_query_id_payment_service INT ,
    tx_hash varchar (128)
);

CREATE OR REPLACE FUNCTION withdraw_success (withdrawId INT, txHash varchar(128), txAprvdTime TIMESTAMP) RETURNS void
    LANGUAGE plpgsql
AS $$
BEGIN
    UPDATE withdraw
    SET	tx_stat = 'SUCCESSFUL',
           tx_approved_time = txAprvdTime,
           tx_hash = txHash
    WHERE
            id = withdrawId ;
END;
$$;


CREATE OR REPLACE FUNCTION withdraw_req_init (
    v_fk_ext_account_sender INT,
    v_fk_ext_account_receiver INT,
    v_fk_ext_currency INT,
    v_value NUMERIC(28,18),
    v_fk_withdraw_fee INT,
    v_tx_sent_time TIMESTAMP,
    v_tx_stat tx_status,
    v_fk_wallet_sender INT,
    v_fk_wallet_receiver INT,
    v_payment_cat PAYMENT_CATEGORY,
    v_value_fee_included NUMERIC(28,18)
) RETURNS INT
    LANGUAGE plpgsql
AS $$

declare wdr_id INT;

BEGIN
    INSERT INTO withdraw (
        fk_ext_account_sender,
        fk_ext_account_receiver,
        fk_ext_currency,
        value,
        fk_withdraw_fee,
        tx_sent_time,
        tx_stat)
    VALUES (
               v_fk_ext_account_sender ,
               v_fk_ext_account_receiver,
               v_fk_ext_currency,
               v_value,
               v_fk_withdraw_fee,
               v_tx_sent_time,
               v_tx_stat
           )RETURNING id INTO wdr_id;


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
               wdr_id,
               v_value_fee_included,
               v_tx_sent_time)
    ;


    UPDATE
        wallet
    SET
        balance = balance - v_value_fee_included
    WHERE
            id = v_fk_wallet_sender
    ;

    RETURN wdr_id;

END;
$$;

-- +migrate Down