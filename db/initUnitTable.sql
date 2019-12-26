create table units
(
	id serial not null,
	name varchar(60) not null
);

create unique index units_id_uindex
	on units (id);

alter table units
	add constraint units_pk
		primary key (id);