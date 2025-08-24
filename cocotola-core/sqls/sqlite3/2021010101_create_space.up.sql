create table `core_space` (
 `id`integer primary key
,`version` int not null default 1
,`created_at` datetime not null default current_timestamp
,`updated_at` datetime not null default current_timestamp
,`created_by` int not null
,`updated_by` int not null
,`organization_id` int not null
,`owner_id` int not null
,`key` varchar(20) not null
,`name` varchar(20) not null
,unique(`organization_id`, `owner_id`, `key`)
,foreign key(`created_by`) references `mb_app_user`(`id`) on delete cascade
,foreign key(`updated_by`) references `mb_app_user`(`id`) on delete cascade
,foreign key(`organization_id`) references `mb_organization`(`id`) on delete cascade
);
