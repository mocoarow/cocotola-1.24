create table `core_card` (
 `id`integer primary key
,`version` int not null default 1
,`created_at` datetime not null default current_timestamp
,`updated_at` datetime not null default current_timestamp
,`created_by` int not null
,`updated_by` int not null
,`organization_id` int not null
,`deck_id` int not null
,`template_id` int not null
,`content` json not null
,`owner_id` int not null
-- ,foreign key(`created_by`) references `mb_app_user`(`id`) on delete cascade
-- ,foreign key(`updated_by`) references `mb_app_user`(`id`) on delete cascade
-- ,foreign key(`organization_id`) references `mb_organization`(`id`) on delete cascade
,foreign key(`deck_id`) references `core_deck`(`id`) on delete cascade
,foreign key(`template_id`) references `core_template`(`id`) on delete cascade
);
