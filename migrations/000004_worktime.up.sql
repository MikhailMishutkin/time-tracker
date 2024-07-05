CREATE TABLE IF NOT EXISTS worktime (
    id serial not null PRIMARY KEY,
    start timestamptz,
    finish timestamptz,
    task_id integer,
    user_uuid integer,
    FOREIGN KEY (task_id) REFERENCES task (id),
    FOREIGN KEY (user_uuid) REFERENCES users (uuid) ON DELETE CASCADE
)
