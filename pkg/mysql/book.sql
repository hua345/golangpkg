-- auto-generated definition
create table book
(
    id        bigint                      not null,
    book_name varchar(32)                 not null,
    price     decimal(10, 2) default 0.00 not null,
    book_desc varchar(256)                null,
    constraint book_id_uindex
        unique (id)
)
    charset = utf8;

alter table book
    add primary key (id);

