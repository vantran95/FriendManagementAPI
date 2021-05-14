CREATE TABLE IF NOT EXISTS users (
    id SERIAL       PRIMARY KEY,
    email           TEXT                     NOT NULL CHECK (email <> ''::text) UNIQUE,
    created_at      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS relationships (
    id SERIAL       PRIMARY KEY,
    first_email_id  INTEGER                  NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    second_email_id INTEGER                  NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    status          TEXT                     NOT NULL,
    created_at      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

