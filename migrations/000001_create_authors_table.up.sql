CREATE TABLE IF NOT EXISTS authors (
	id varchar(255) NOT NULL,
	name varchar(255) NOT NULL,
	CONSTRAINT authors_pkey PRIMARY KEY (id)
);

INSERT INTO authors
(id, "name")
VALUES('1', 'root');
INSERT INTO authors
(id, "name")
VALUES('2', 'test_user');
INSERT INTO authors
(id, "name")
VALUES('3', 'admin');