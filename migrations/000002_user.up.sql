CREATE TABLE IF NOT EXISTS users (
    uuid            serial not null PRIMARY KEY,
    surname         varchar,
    name            varchar,
    patronymic      varchar,
    address         varchar,
    passportSerie  integer,
    passportNumber integer
)