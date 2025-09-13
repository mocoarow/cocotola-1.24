create table `core_template` (
 `id`integer primary key
,`name` varchar(40) not null
,`deleted` boolean not null default 0
,unique(`name`)
);
