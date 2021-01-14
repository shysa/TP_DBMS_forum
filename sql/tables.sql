CREATE EXTENSION IF NOT EXISTS CITEXT;

drop table if exists users;
drop table if exists forum;
drop table if exists forum_users;
drop table if exists thread;
drop table if exists post;
drop table if exists votes;

create table if not exists users
(
    id       serial primary key,
    nickname citext unique not null,
    about    text,
    email    citext unique,
    fullname text
);

create table if not exists forum
(
    id      serial primary key,
    posts   integer       not null default 0,
    slug    citext unique not null,
    threads integer       not null default 0,
    title   text          not null,
    "user"  citext        not null
);

create table if not exists forum_users
(
    forum    citext not null,
    nickname citext not null,
    primary key (nickname, forum)
);

create table if not exists thread
(
    id      serial primary key,
    created timestamptz not null,
    message text        not null,
    slug    citext,
    title   text,
    votes   integer     not null default 0,

    author  citext      not null,
    forum   citext      not null
);
create unique index thread_slug_index on thread (slug) where slug <> '';
create index thread_forum_created_index on thread (forum, created);

create table if not exists post
(
    id       serial primary key,
    created  timestamptz not null,
    isEdited bool                 default false,
    message  text        not null,
    parent   integer     not null default 0,
    tree     bigint[]             default array []::bigint[],

    thread   integer     not null,
    author   citext      not null,
    forum    citext      not null
);
create index posts_id_thread_index on post (thread, id);
create index posts_thread_tree_index on post (thread, tree);
create index parent_tree_1 on post (tree, (tree[1]));

create table votes
(
    id       serial primary key,
    thread   integer not null references thread (id),
    voice    integer,
    prev     integer not null default 0,

    nickname citext  not null,
    constraint unique_uservoice_for_thread unique (nickname, thread)
);


create or replace function add_forum_thread() returns trigger as
$add_forum_thread$
begin
    update forum set threads=forum.threads + 1 where slug = new.forum;
    return new;
end;
$add_forum_thread$ language plpgsql;

create trigger on_add_forum_thread
    after insert
    on thread
    for each row
execute procedure add_forum_thread();

GRANT ALL PRIVILEGES ON DATABASE dbms_db TO dbms_user;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO dbms_user;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO dbms_user;

