CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS relationships (
    id SERIAL PRIMARY KEY,
    first_email_id INT,
    second_email_id INT,
    status INT NOT NULL
);

