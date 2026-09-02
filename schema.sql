create table if not exists expenses (
    id integer primary key autoincrement,
    name string not null,
    amount real not null,
    category string not null,
    date text not null default current_timestamp
);