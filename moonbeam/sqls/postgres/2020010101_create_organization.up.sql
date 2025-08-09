create table organization (
 id serial not null
,version int not null default 1
,created_at timestamp not null default current_timestamp
,updated_at timestamp not null default current_timestamp
,created_by int not null
,updated_by int not null
,name varchar(20) not null
,primary key(id)
,unique(name)
);
