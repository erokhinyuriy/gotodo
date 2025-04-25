create table td_list(
    id uuid NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    name character varying(50) not null COLLATE pg_catalog."ru-RU-x-icu",
    date timestamp without time zone
);