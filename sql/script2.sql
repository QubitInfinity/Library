-- Create and populate tables

DROP TABLE IF EXISTS loans;
DROP TABLE IF EXISTS members;
DROP TABLE IF EXISTS books;


CREATE TABLE IF NOT EXISTS members (
                                       id bigserial primary key,
                                       full_name varchar(50) NOT NULL
    )

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS members
    OWNER to postgres;

INSERT INTO members (full_name) VALUES('John Smith');
INSERT INTO members (full_name) VALUES('Jules Verne');


CREATE TABLE IF NOT EXISTS books (
                                     id bigserial primary key,
                                     title varchar(50) NOT NULL,
    author varchar(50) NOT NULL
    )

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS books
    OWNER to postgres;

INSERT INTO books (author, title) VALUES('Jules Verne', 'Around the World in 80 Days');
INSERT INTO books (author, title) VALUES('Jules Verne', 'Twenty Thousand Leagues Under the Sea');
INSERT INTO books (author, title) VALUES('Carl Sagan', 'Cosmos');
INSERT INTO books (author, title) VALUES('Bertrand Russell', 'The Principia Mathematica');
INSERT INTO books (author, title) VALUES('Me Mine', 'The Greatest Fiction');
INSERT INTO books (author, title) VALUES('Me Mine', 'The Unavailable Book');




CREATE TABLE IF NOT EXISTS loans (
                                     id bigserial primary key,
                                     book_id bigserial NOT NULL REFERENCES books ,
                                     member_id bigserial REFERENCES members,
                                     borrowed DATE,
                                     returned DATE
)

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS loans
    OWNER to postgres;

-- borrow
INSERT INTO loans (book_id, member_id, borrowed) VALUES('6', '1', '2023/11/01');


-- return
UPDATE loans SET returned = '2023/11/15'
WHERE loans.book_id = '6'
  AND loans.member_id = '1'
  AND loans.returned IS NULL;

-- borrow
INSERT INTO loans (book_id, member_id, borrowed) VALUES('6', '1', '2024/11/29');