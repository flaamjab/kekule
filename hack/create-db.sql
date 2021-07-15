create table Item (
  id integer primary key,
  sku text generated always as ("c" || category || "n" || id),
  name string,
  price double,
  category integer,
  foreign key(category) references Category(id)
);

create table Category (
  id integer primary key,
  name string
);

insert into Category (name) values ("Groceries");
insert into Category (name) values ("Currency");
insert into Category (name) values ("Clothing");
