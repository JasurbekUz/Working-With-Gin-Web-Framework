create database livechat;

create extension pgcrypto;

drop table if exists users cascade;

create table users (
	user_id serial not null primary key,
	user_name character varying(32) not null,
	email character varying(64) not null,
	pass_word character varying(60) not null
);

-- MOCK DATA:

insert into users (user_name, email, pass_word) values ('johne', 'doee@gmail.com', crypt('2112', gen_salt('bf')));

insert into users (user_name, email, pass_word) 
values ('Jones', 'jones@gmail.com', crypt('jone0012', gen_salt('bf')))
returning *;