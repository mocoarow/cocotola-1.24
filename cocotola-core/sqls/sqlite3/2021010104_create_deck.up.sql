create table `core_deck` (
 `id`integer primary key
,`version` int not null default 1
,`created_at` datetime not null default current_timestamp
,`updated_at` datetime not null default current_timestamp
,`created_by` int not null
,`updated_by` int not null
,`organization_id` int not null
,`space_id` int not null
,`owner_id` int not null
,`folder_id` int not null
,`template_id` int not null
,`name` varchar(40) not null
,`lang2` varchar(2) not null
,`description` text
,foreign key(`created_by`) references `mb_app_user`(`id`) on delete cascade
,foreign key(`updated_by`) references `mb_app_user`(`id`) on delete cascade
,foreign key(`organization_id`) references `mb_organization`(`id`) on delete cascade
,foreign key(`space_id`) references `core_space`(`id`) on delete cascade
,foreign key(`owner_id`) references `mb_app_user`(`id`) on delete cascade
,foreign key(`folder_id`) references `core_folder`(`id`) on delete cascade
,foreign key(`template_id`) references `core_template`(`id`) on delete cascade
);
