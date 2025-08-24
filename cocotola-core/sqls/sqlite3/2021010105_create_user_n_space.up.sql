create table `core_user_n_space` (
 `created_at` datetime not null default current_timestamp
,`created_by` int not null
,`organization_id` int not null
,`app_user_id` int not null
,`space_id` int not null
,primary key(`organization_id`, `app_user_id`, `space_id`)
,foreign key(`created_by`) references `mb_app_user`(`id`) on delete cascade
,foreign key(`organization_id`) references `mb_organization`(`id`) on delete cascade
,foreign key(`app_user_id`) references `mb_app_user`(`id`) on delete cascade
,foreign key(`space_id`) references `core_space`(`id`) on delete cascade
);
