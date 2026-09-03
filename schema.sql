create table if not exists expenses (
    id integer primary key autoincrement,
    name text not null,
    amount real not null,
    category text not null,
    date text not null default current_timestamp
);