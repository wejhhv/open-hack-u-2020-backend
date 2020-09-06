CREATE table responses (
    id SERIAL NOT NULL,
    user_id integer NOT NULL,
    comment_id integer NOT NULL,
    comment TEXT,
    date_time timestamp NOT NULL,
    PRIMARY KEY (id), 
    foreign key (user_id) references users(id)
        on delete cascade
        on update no action,
    foreign key (comment_id) references comments(id)
        on delete cascade
        on update no action
);