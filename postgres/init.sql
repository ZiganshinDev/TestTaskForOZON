CREATE TABLE IF NOT EXISTS urls
(
    id      serial        not null,
    original_url varchar(1024) not null,
    short_url   varchar(10)   not null,
    UNIQUE (short_url)
);