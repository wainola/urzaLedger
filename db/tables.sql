create table Expenses (
    id text primary key not null,
    expense text not null,
    amount int not null,
    source text not null,
    createdAt text not null,
    updatedAt text default null,
    deletedAt text default null
);

create table Incomes (
    id text primary key not null,
    income text not null,
    amount int not null,
    source text not null,
    createdAt text not null,
    updatedAt text default null,
    deletedAt text default null
);

create table Users (
    id text primary key not null,
    name text not null,
    lastname text not null,
    email text not null,
    age int not null,
    createdAt text not null,
    updatedAt text default null,
    deletedAt text default null
);

alter table Expenses add column user_id text references Users(id);
alter table Incomes add column user_id text references Users(id);

