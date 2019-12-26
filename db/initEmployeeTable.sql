create table employees
(
	id serial not null,
	name varchar(60) not null,
	age int,
	unit_id int not null
);

create unique index employees_id_uindex
	on employees (id);

alter table employees
	add constraint employees_pk
		primary key (id);
