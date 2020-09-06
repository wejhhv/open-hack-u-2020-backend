CREATE table users (
    id SERIAL NOT NULL,
    name varchar(20) NOT NULL,
    point integer,
    PRIMARY KEY (id), 
    UNIQUE(name)
);