CREATE TABLE IF NOT EXISTS albums
(
    title character varying COLLATE pg_catalog."default",
    artist character varying COLLATE pg_catalog."default",
    price numeric,
    id integer NOT NULL DEFAULT nextval('"albums_Id_seq"'::regclass),
    CONSTRAINT albums_pkey PRIMARY KEY (id)
)