create table workbook (
 id serial not null
,version int not null default 1
,created_at timestamp not null default current_timestamp
,updated_at timestamp not null default current_timestamp
,created_by int not null
,updated_by int not null
,organization_id int not null
,name varchar(40) not null
,problem_type varchar(20) not null
,lang2 varchar(2) not null
,description text
,content jsonb not null
,primary key(id)
);
