create table `group_n_group` (
 `created_at` datetime not null default current_timestamp
,`created_by` int not null
,`organization_id` int not null
,`child_user_group_id` int not null
,`parent_user_group_id` int not null
,primary key(`organization_id`, `child_user_group_id`, `parent_user_group_id`)
,foreign key(`created_by`) references `app_user`(`id`) on delete cascade
,foreign key(`organization_id`) references `organization`(`id`) on delete cascade
,foreign key(`child_user_group_id`) references `user_group`(`id`) on delete cascade
,foreign key(`parent_user_group_id`) references `user_group`(`id`) on delete cascade
);
