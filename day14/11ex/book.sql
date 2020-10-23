create table book
(
    id int unsigned auto_increment,
    name varchar(50) not null,
    price int unsigned default 0 not null,
    constraint book_pk
        primary key (id)
);

