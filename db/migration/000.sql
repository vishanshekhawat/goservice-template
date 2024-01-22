CREATE OR REPLACE FUNCTION trigger_set_timestamp()
    RETURNS TRIGGER AS $$
BEGIN
    IF NEW.* IS DISTINCT FROM OLD.* THEN
        NEW.modified_at = CURRENT_TIMESTAMP;
    END IF;

    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE SEQUENCE IF NOT EXISTS public.user_id_seq;

CREATE TABLE  IF NOT EXISTS public.users (
    id int NOT NULL DEFAULT nextval('user_id_seq'::regclass),

    name text NOT NULL,
    email text NOT NULL,
    is_active bool NOT NULL DEFAULT false,
    password text NOT NULL,
    salt text NOT NULL,

    created_at TIMESTAMP without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (id)
);

CREATE TRIGGER set_timestamp BEFORE UPDATE ON users FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();
