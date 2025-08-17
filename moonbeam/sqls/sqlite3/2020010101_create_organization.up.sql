create table `organization` (
 `id`integer primary key
,`version` int not null default 1
,`created_at` datetime not null default current_timestamp
,`updated_at` datetime not null default current_timestamp
,`created_by` int not null
,`updated_by` int not null
,`name` varchar(20) not null
,unique(`name`)
);
