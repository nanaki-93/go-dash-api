create table if not exists "user"
(
    id       uuid default gen_random_uuid() not null constraint user_pk primary key,
    name     text                           not null,
    mail     text                           not null,
    password text                           not null
    );

create table if not exists "group"
(
    id     uuid default gen_random_uuid() not null constraint group_pk primary key,
    name   text                           not null,
    colour text                           not null,
    icon   text
    );


create table if not exists threshold_type
(
    id          uuid default gen_random_uuid() not null constraint threshold_type_pk primary key,
    name        text,
    description text
    );


create table if not exists frequency
(
    id          uuid default gen_random_uuid() not null constraint frequency_pk primary key,
    name        text,
    description text
    );


create table if not exists recurrence
(
    id             uuid default gen_random_uuid() not null constraint recurrence_pk primary key,
    frequency_uid  uuid,
    time           timestamp with time zone,
    interval       numeric,
    days_of_week   text,
    days_of_months text,
    months_of_year text,
    start_date     timestamp with time zone,
    end_date       timestamp with time zone,
                                 occurrences    numeric,
                                 is_infinite    boolean,
                                 is_threshold   boolean,
                                 is_event       boolean
                                 );


create table if not exists reminder
(
    id                uuid default gen_random_uuid() not null constraint reminder_pk primary key,
    name              text,
    noe               text,
    priority          numeric,
    threshold         integer,
    alarm             timestamp with time zone,
                                    icon              text,
                                    group_id          uuid,
                                    threshold_type_id uuid,
                                    recurrence_id     uuid
                                    );

create table if not exists notification
(
    id          uuid default gen_random_uuid() not null constraint notification_pk primary key,
    notify_at   timestamp with time zone       not null,
                              reminder_id uuid                           not null
                              );


create table if not exists user_reminder
(
    user_id     uuid not null,
    reminder_id uuid not null,
    constraint user_reminder_pk primary key (reminder_id, user_id)
    );

alter table recurrence add constraint recurrence_frequency_uid_fk foreign key (frequency_uid) references frequency;
alter table reminder add constraint reminder_group_uid_fk foreign key (group_id) references "group";
alter table reminder add constraint reminder_threshold_type_uuid_fk foreign key (threshold_type_id) references threshold_type;
alter table reminder add constraint reminder_recurrence_uid_fk foreign key (recurrence_id) references recurrence;
alter table notification add constraint notification_reminder_uid_fk foreign key (reminder_id) references reminder;
alter table user_reminder add constraint user_reminder_user_uid_fk foreign key (user_id) references "user";
alter table user_reminder add constraint user_reminder_reminder_uid_fk foreign key (reminder_id) references reminder;



