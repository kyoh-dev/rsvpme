-- migrate:up
CREATE TABLE event (
    id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    uuid UUID NOT NULL DEFAULT gen_random_uuid(),
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    start_datetime TIMESTAMP NOT NULL,
    finish_datetime TIMESTAMP,
    address TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX event_uuid_unique_idx ON event(uuid);

CREATE FUNCTION event_log_updated_at() RETURNS trigger AS $event_log_updated_at$
    BEGIN
        NEW.updated_at = NOW();
        RETURN NEW;
    END;
$event_log_updated_at$ LANGUAGE plpgsql;

CREATE TRIGGER event_log_updated_at BEFORE UPDATE ON event
    FOR EACH ROW EXECUTE FUNCTION event_log_updated_at();

-- migrate:down
DROP TABLE event;
