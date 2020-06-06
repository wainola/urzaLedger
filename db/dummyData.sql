insert into Users (id, name, lastname, email, age, createdAt) values (
    "69d894e0-a40e-4391-ad44-ef7a43777edf",
    "John",
    "Doe",
    "john@mail.com",
    33,
    date("now")
);

insert into Users (id, name, lastname, email, age, createdAt) values (
    "bcab95a9-412d-463d-ad22-8994b076bcfc",
    "Cena",
    "John",
    "cena@mail.com",
    45,
    date("now")
);

insert into Incomes (id, income, amount, source, createdAt, user_id) values (
    "a59aa9e1-4fd9-4ad3-ba0a-3b8a98c42537",
    "salary",
    2250000,
    "company",
    date("now"),
    "69d894e0-a40e-4391-ad44-ef7a43777edf"
);

insert into Incomes (id, income, amount, source, createdAt, user_id) values (
    "a85e3fb2-81c9-496b-ad22-3f6ce6c1e6f5",
    "salary",
    3200000,
    "company",
    date("now"),
    "bcab95a9-412d-463d-ad22-8994b076bcfc"
);

insert into Expenses (id, expense, amount, source, createdAt, user_id) values (
    "9c159c6e-b799-4649-a6c7-7edcfd441a3c",
    "Credit card payment",
    300000,
    "account",
    date("now"),
    "bcab95a9-412d-463d-ad22-8994b076bcfc"
);


insert into Expenses (id, expense, amount, source, createdAt, user_id) values (
    "db8e34a7-d0bf-4200-8b12-aa8e87f42448",
    "Motorcycle quote",
    580000,
    "account",
    date("now"),
    "69d894e0-a40e-4391-ad44-ef7a43777edf"
);