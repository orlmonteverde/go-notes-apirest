CREATE TABLE notes (
	id smallserial NOT NULL,
	title varchar(150) NOT NULL,
	description text NOT NULL,
	created_at timestamp DEFAULT now(),
	updated_at timestamp NOT NULL,
	CONSTRAINT pk_notes PRIMARY KEY(id)
);
