CREATE TABLE IF NOT EXISTS task (
    id serial not null PRIMARY KEY,
    title varchar UNIQUE,
    active boolean
)