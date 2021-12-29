CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    uuid uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    login VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    role INTEGER NOT NULL,
);

CREATE TABLE books (
    uuid uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    status INTEGER NOT NULL,
    holder_uuid uuid REFERENCES users(uuid)
);

INSERT INTO books (name, description, status) VALUES('Война и мир', 'Книга о том, как Лев Толстой пролил кувшин с водой на свои труды', 1);
INSERT INTO books (name, description, status) VALUES('Война и мир', 'Книга о том, как Лев Толстой пролил кувшин с водой на свои труды', 1);
INSERT INTO books (name, description, status) VALUES('Преступление и наказание', 'Инструкция о том, как не надо пользоваться топором', 1);
INSERT INTO books (name, description, status) VALUES('Преступление и наказание', 'Инструкция о том, как не надо пользоваться топором', 1);
INSERT INTO books (name, description, status) VALUES('Преступление и наказание', 'Инструкция о том, как не надо пользоваться топором', 1);