CREATE TABLE persons (
    person_id bigserial not null primary key,
    name text not null,
    age integer not null,
    work text,
    address text not null
);