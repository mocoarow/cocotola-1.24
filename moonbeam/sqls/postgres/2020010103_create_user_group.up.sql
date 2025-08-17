create table user_group (
 id serial not null
,version int not null default 1
,created_at timestamp not null default current_timestamp
,updated_at timestamp not null default current_timestamp
,created_by int not null
,updated_by int not null
,organization_id int not null
,key_name varchar(20) not null
,name varchar(40) not null
,description text
,removed bool not null
,primary key(id)
,unique(organization_id, key_name)
,foreign key(created_by) references app_user(id) on delete cascade
,foreign key(updated_by) references app_user(id) on delete cascade
,foreign key(organization_id) references organization(id) on delete cascade
);
