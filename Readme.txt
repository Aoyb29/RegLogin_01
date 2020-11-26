-- auto-generated definition
create table users
(
    id   serial   not null
        constraint users_pkey
            primary key,
    name char(20) not null,
    psw  char(20) not null,
    sex  char(2)  not null,
    age  integer  not null,
    tel  char(20)
);

alter table users
    owner to postgres;

INSERT INTO public.users (id, name, psw, sex, age, tel) VALUES (2, 'qwe                 ', '111111              ', '男 ', 12, 'qwe111              ');
INSERT INTO public.users (id, name, psw, sex, age, tel) VALUES (3, 'asd                 ', '333333              ', '女 ', 18, 'asd333              ');
INSERT INTO public.users (id, name, psw, sex, age, tel) VALUES (1, 'ayb                 ', '123123              ', '男 ', 23, 'ayb123              ');