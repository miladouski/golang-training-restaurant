create table orders
(
    id           serial  not null
        constraint orders_pk
            primary key,
    date         date    not null,
    table_number integer not null,
    waiter_id    integer not null,
    price        integer not null,
    payment      boolean not null
);

create table waiters
(
    waiter_id serial not null
        constraint waiters_pk
            primary key,
    full_name varchar(50)   not null
);

create table order_content
(
    order_id integer not null,
    dish_id  integer not null,
    quantity integer not null
);

create table menu
(
    dish_id          serial  not null
        constraint menu_pk
            primary key,
    name             varchar(50)    not null,
    dish_category_id integer not null,
    price            integer
);

create table ingredients
(
    ingredient_id serial  not null
        constraint ingredients_pk
            primary key,
    name          varchar(50)    not null,
    price         integer not null
);

create table dish_composition
(
    dish_id       integer not null,
    ingredient_id integer not null,
    quantity      integer not null
);

create table dish_category
(
    dish_category_id serial not null
        constraint dish_category_pk
            primary key,
    category_name    varchar(50)   not null
);

INSERT INTO public.orders (id, date, table_number, waiter_id, price, payment) VALUES (1, '2021-04-02', 1, 1, 300, true);
INSERT INTO public.orders (id, date, table_number, waiter_id, price, payment) VALUES (2, '2021-04-03', 2, 2, 225, true);
INSERT INTO public.orders (id, date, table_number, waiter_id, price, payment) VALUES (3, '2021-04-03', 3, 2, 213, true);
INSERT INTO public.orders (id, date, table_number, waiter_id, price, payment) VALUES (4, '2021-04-03', 2, 1, 225, true);
INSERT INTO public.orders (id, date, table_number, waiter_id, price, payment) VALUES (5, '2021-04-03', 1, 2, 123, false);

INSERT INTO public.waiters (waiter_id, full_name) VALUES (1, 'Mark');
INSERT INTO public.waiters (waiter_id, full_name) VALUES (2, 'Harry');
INSERT INTO public.waiters (waiter_id, full_name) VALUES (3, 'Thomas');
INSERT INTO public.waiters (waiter_id, full_name) VALUES (4, 'Oliver');
INSERT INTO public.waiters (waiter_id, full_name) VALUES (5, 'Jack');

INSERT INTO public.order_content (order_id, dish_id, quantity) VALUES (1, 1, 1);
INSERT INTO public.order_content (order_id, dish_id, quantity) VALUES (2, 2, 2);
INSERT INTO public.order_content (order_id, dish_id, quantity) VALUES (3, 3, 1);
INSERT INTO public.order_content (order_id, dish_id, quantity) VALUES (4, 4, 3);
INSERT INTO public.order_content (order_id, dish_id, quantity) VALUES (5, 5, 1);

INSERT INTO public.menu (dish_id, name, dish_category_id, price) VALUES (1, 'soup', 1, 23);
INSERT INTO public.menu (dish_id, name, dish_category_id, price) VALUES (2, 'pizza', 2, 123);
INSERT INTO public.menu (dish_id, name, dish_category_id, price) VALUES (3, 'stew', 2, 45);
INSERT INTO public.menu (dish_id, name, dish_category_id, price) VALUES (4, 'potatoes', 3, 45);
INSERT INTO public.menu (dish_id, name, dish_category_id, price) VALUES (5, 'salad', 1, 65);

INSERT INTO public.ingredients (ingredient_id, name, price) VALUES (1, 'potatoes', 12);
INSERT INTO public.ingredients (ingredient_id, name, price) VALUES (2, 'greens', 24);
INSERT INTO public.ingredients (ingredient_id, name, price) VALUES (3, 'meate', 31);
INSERT INTO public.ingredients (ingredient_id, name, price) VALUES (4, 'cheese', 12);
INSERT INTO public.ingredients (ingredient_id, name, price) VALUES (5, 'salad', 14);

INSERT INTO public.dish_composition (dish_id, ingredient_id, quantity) VALUES (1, 1, 1);
INSERT INTO public.dish_composition (dish_id, ingredient_id, quantity) VALUES (2, 2, 3);
INSERT INTO public.dish_composition (dish_id, ingredient_id, quantity) VALUES (3, 3, 2);
INSERT INTO public.dish_composition (dish_id, ingredient_id, quantity) VALUES (4, 4, 2);
INSERT INTO public.dish_composition (dish_id, ingredient_id, quantity) VALUES (5, 5, 2);

INSERT INTO public.dish_category (dish_category_id, category_name) VALUES (1, 'first course');
INSERT INTO public.dish_category (dish_category_id, category_name) VALUES (2, 'main course');
INSERT INTO public.dish_category (dish_category_id, category_name) VALUES (3, 'snack');
INSERT INTO public.dish_category (dish_category_id, category_name) VALUES (4, 'beverages');
INSERT INTO public.dish_category (dish_category_id, category_name) VALUES (5, 'garnish');
