create extension if not exists "uuid-ossp";

create table if not exists events
(
    id          uuid        not null default uuid_generate_v4() primary key,
    title       varchar(64) not null,
    time        timestamp   without time zone
                            not null,
    duration    bigint      not null,
    user_id     uuid        not null,
    day_key     timestamp   not null,
    week_key    timestamp   not null,
    month_key   timestamp   not null
);

create index if not exists events_day_key_index on events (day_key);
create index if not exists events_week_key_index on events (week_key);
create index if not exists events_month_key_index on events (month_key);
