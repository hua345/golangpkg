-- auto-generated definition
create table book
(
    id        bigint                      not null
        constraint book_pkey
        primary key
        constraint book_id_uindex
        unique,
    book_name integer                     not null,
    price     numeric(10, 2) default 0.00 not null,
    book_desc varchar(256)
);

alter table book
    owner to fang;

