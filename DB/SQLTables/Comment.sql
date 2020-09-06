CREATE table comments (
    id SERIAL NOT NULL,
    user_id integer NOT NULL,
    emotion_id integer NOT NULL,
    prefecture varchar(20) NOT NULL,
    latitude double precision NOT NULL,
    longtitude double precision NOT NULL,
    comment TEXT,
    date_time timestamp NOT NULL,
    PRIMARY KEY (id), 
    foreign key (user_id) references users(id)
        on delete cascade
        on update no action
);