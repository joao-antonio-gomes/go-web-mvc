create table public.products
(
    id          serial
        constraint products_pk_id
            primary key,
    name        varchar,
    description varchar,
    price       decimal,
    quantity    integer
);

