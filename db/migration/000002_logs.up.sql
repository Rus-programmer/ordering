CREATE TABLE logs
(
    id           BIGSERIAL PRIMARY KEY,
    method       VARCHAR(20) NOT NULL,
    path         TEXT        NOT NULL,
    status_code  INT         NOT NULL,
    elapsed_time VARCHAR(50) NOT NULL,
    time         TIMESTAMPTZ NOT NULL
);