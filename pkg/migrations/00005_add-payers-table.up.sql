CREATE TABLE payers(
	id SERIAL PRIMARY KEY,
	address TEXT NOT NULL UNIQUE
);

