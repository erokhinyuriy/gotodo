create table td_task(
    id uuid NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    list_id uuid NOT NULL,
    name character varying(50) not null COLLATE pg_catalog."ru-RU-x-icu",
    description character varying(500) COLLATE pg_catalog."ru-RU-x-icu",
    date timestamp without time zone DEFAULT now(),
    constraint td_task_fk_td_list foreign key (list_id) references td_list (id) match simple on update no action on delete no action
);

create index td_task_list_id_idx on td_task using btree(list_id asc nulls last);