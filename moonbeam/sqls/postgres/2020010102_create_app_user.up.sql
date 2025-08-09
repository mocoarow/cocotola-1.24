create table app_user (
 id serial not null
,version int not null default 1
,created_at timestamp not null default current_timestamp
,updated_at timestamp not null default current_timestamp
,created_by int not null
,updated_by int not null
,organization_id int not null
,login_id varchar(200)
,hashed_password varchar(200)
,username varchar(40)
,provider varchar(40)
,provider_id varchar(40)
,provider_access_token text
,provider_refresh_token text
,removed bool not null
,primary key(id)
,unique(organization_id, login_id)
,foreign key(organization_id) references organization(id) on delete cascade
);
