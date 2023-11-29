CREATE TABLE questions (
    id SERIAL PRIMARY KEY,
    text VARCHAR(255),
    options VARCHAR ARRAY,
    correct_option_index INT
);