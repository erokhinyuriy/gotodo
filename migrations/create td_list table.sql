create table tb_list(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name character varying(50) not null,
    date timestamp without time zone
);