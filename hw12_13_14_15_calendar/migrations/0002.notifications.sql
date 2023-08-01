create table if not exists notifications
(
    notification_id uuid    not null,
    user_id         uuid    not null,
    event           jsonb   not null,
    created_at      timestamp without time zone
);
