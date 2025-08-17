create table `user_group` (
 `id` integer primary key
,`version` int not null default 1
,`created_at` datetime not null default current_timestamp
,`updated_at` datetime not null default current_timestamp
,`created_by` int not null
,`updated_by` int not null
,`organization_id` int not null
,`key_name` varchar(20) not null
,`name` varchar(40) not null
,`description` text
,`removed` tinyint(1) not null
,unique(`organization_id`, `key_name`)
,foreign key(`created_by`) references `app_user`(`id`) on delete cascade
,foreign key(`updated_by`) references `app_user`(`id`) on delete cascade
,foreign key(`organization_id`) references `organization`(`id`) on delete cascade
);
