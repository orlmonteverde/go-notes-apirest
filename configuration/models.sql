CREATE TABLE users (
	id smallserial NOT NULL,
	name varchar(150) NOT NULL,
	password varchar(256) NOT NULL,
	role varchar(10) NOT NULL DEFAULT 'writer',
	created_at timestamp DEFAULT now(),
	updated_at timestamp NOT NULL,
	CONSTRAINT pk_users PRIMARY KEY(id)
);

CREATE TABLE notes (
	id smallserial NOT NULL,
	user_id smallint NOT NULL,
	title varchar(150) NOT NULL,
	description text NOT NULL,
	created_at timestamp DEFAULT now(),
	updated_at timestamp NOT NULL,
	CONSTRAINT pk_notes PRIMARY KEY(id)
);
