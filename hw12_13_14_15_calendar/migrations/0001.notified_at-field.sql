alter table events add column if not exists "notified_at" timestamp;
create index if not exists events_time_index on events (time);
