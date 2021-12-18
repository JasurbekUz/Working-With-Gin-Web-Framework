drop table if exists users cascade;

create table users (
	user_id serial not null primary key,
	full_name character varying(64) not null,
	user_name character varying(16) not null,
	user_role character varying(16) not null
);

-- MOCK DATA

insert into users (full_name, user_name, user_role)
values ('Axror Usmonov', 'axror', 'San''atkor'), ('Sherali Jo''orayev', 'sherali', 'San''atkor'), ('Botir', 'Qodirov', 'San''atkor');