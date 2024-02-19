create table if not exists produce (
    id integer primary key,
    code text,
    name text,
    price float
);
create index produce_name_idx on produce (name);

insert into produce (code, name, price) values ('A12T-4GH7-QPL9-3N4M', 'lettuce', 3.46);
insert into produce (code, name, price) values ('E5T6-9UI3-TH15-QR88', 'peach', 2.99);
insert into produce (code, name, price) values ('YRT6-72AS-K736-L4AR', 'green pepper', 0.79);
insert into produce (code, name, price) values ('TQ4C-VV6T-75ZX-1RMR', 'gala apple', 3.59);