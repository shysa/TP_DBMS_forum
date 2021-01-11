CREATE USER dbms_user WITH password 'qwerty123456';

create database dbms_db
    with owner dbms_user
    encoding 'utf8'
    TABLESPACE = pg_default
;
