create table if not exists users
(
    id          int unsigned auto_increment primary key,
    username    varchar(255) not null,
    password    varchar(100) not null,
    created_at  timestamp default CURRENT_TIMESTAMP null,
    updated_at  timestamp default CURRENT_TIMESTAMP null,
    role        varchar(100) default 'user' not null,
    constraint username
        unique (username)
);