create table `user_group_details` (
 `id` integer primary key
,`version` int not null default 1
,`created_at` datetime not null default current_timestamp
,`updated_at` datetime not null default current_timestamp
,`created_by` int not null
,`updated_by` int not null
,`organization_id` int not null
,`user_group_id` int not null
,`details` json not null
,unique(`organization_id`, `user_group_id`)
,foreign key(`created_by`) references `app_user`(`id`) on delete cascade
,foreign key(`updated_by`) references `app_user`(`id`) on delete cascade
,foreign key(`organization_id`) references `organization`(`id`) on delete cascade
,foreign key(`user_group_id`) references `user_group`(`id`) on delete cascade
);
