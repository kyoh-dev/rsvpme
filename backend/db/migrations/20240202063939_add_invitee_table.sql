-- migrate:up
CREATE TABLE invitee (
    id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT,
    rsvp BOOLEAN NOT NULL,
    event_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (event_id) REFERENCES event(id) ON DELETE CASCADE
);

CREATE FUNCTION invitee_log_updated_at() RETURNS trigger AS $invitee_log_updated_at$
    BEGIN
        NEW.updated_at = NOW();
        RETURN NEW;
    END;
$invitee_log_updated_at$ LANGUAGE plpgsql;

CREATE TRIGGER invitee_log_updated_at BEFORE UPDATE ON invitee
    FOR EACH ROW EXECUTE FUNCTION invitee_log_updated_at();

-- migrate:down
DROP TABLE invitee;
