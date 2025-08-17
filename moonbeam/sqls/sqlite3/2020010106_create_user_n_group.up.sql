create table `mb_user_n_group` (
 `created_at` datetime not null default current_timestamp
,`created_by` int not null
,`organization_id` int not null
,`app_user_id` int not null
,`user_group_id` int not null
,primary key(`organization_id`, `app_user_id`, `user_group_id`)
,foreign key(`created_by`) references `mb_app_user`(`id`) on delete cascade
,foreign key(`organization_id`) references `mb_organization`(`id`) on delete cascade
,foreign key(`app_user_id`) references `mb_app_user`(`id`) on delete cascade
,foreign key(`user_group_id`) references `mb_user_group`(`id`) on delete cascade
);
