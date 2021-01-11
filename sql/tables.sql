CREATE EXTENSION IF NOT EXISTS citext;

create table if not exists users
(
    id       serial primary key,
    nickname citext unique not null,
    about    text,
    email    citext unique,
    fullname text
);
create index users_covering_index on users (lower(nickname), nickname, email, about, fullname);
create unique index users_nickname_index on users (nickname);
cluster users using users_nickname_index;


create table if not exists forum
(
    id      serial primary key,
    posts   integer       not null default 0,
    slug    citext unique not null,
    threads integer       not null default 0,
    title   text          not null,
    "user"  citext        not null
);
create unique index forum_slug_index on forum (slug);
create unique index forum_slug_id_index on forum (id, slug);
create index forum_covering_index on forum (slug, id, title, "user", threads, posts);
cluster forum using forum_slug_id_index;


create table if not exists forum_users
(
    forum    citext not null,
    nickname citext not null,
    primary key (nickname, forum)
);

create table if not exists thread
(
    id      serial primary key,
    author  citext      not null,
    created timestamptz not null,
    forum   citext      not null,
    message text        not null,
    slug    citext,
    title   text,
    votes   integer     not null default 0
);
create unique index thread_slug_index on thread (slug) where slug <> '';
create index thread_slug_id_index on thread (id, slug);
create index thread_forum_created_index on thread (forum, created);
cluster thread using thread_forum_created_index;


create table if not exists post
(
    id       serial primary key,
    author   citext      not null,
    created  timestamptz not null,
    isEdited bool                 default false,
    message  text        not null,
    parent   integer     not null default 0,
    forum    citext      not null,
    thread   integer     not null,
    tree     integer[]
);
create index on post (id, created);
create index posts_id_thread_index on post (thread, id);
create index posts_thread_tree_index on post (thread, tree);
create index parent_tree_1 on post ((tree[1]), tree, id);


create table votes
(
    id       serial primary key,
    nickname citext  not null,
    thread   integer not null references thread (id),
    voice    integer,
    prev     integer not null default 0,
    constraint unique_uservoice_for_thread unique (nickname, thread)
);
create unique index forum_users_nickname ON users (lower(nickname));



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



create or replace function set_post_parent_tree() returns trigger as
$set_post_parent_tree$
begin
    update post set tree=array_append(new.tree, new.id) where id = new.id;
    return new;
end;
$set_post_parent_tree$ language plpgsql;

create trigger on_create_post_set_parents
    after insert
    on post
    for each row
execute procedure set_post_parent_tree();


GRANT ALL PRIVILEGES ON DATABASE dbms_db TO dbms_user;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO dbms_user;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO dbms_user;
