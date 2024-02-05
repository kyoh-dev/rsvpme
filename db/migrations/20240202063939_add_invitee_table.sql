-- migrate:up
PRAGMA foreign_keys = ON;

CREATE TABLE invitee (
    id INTEGER NOT NULL PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT,
    rsvp BOOLEAN NOT NULL,
    event_id INTEGER NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (event_id) REFERENCES event(id) ON DELETE CASCADE
);

CREATE TRIGGER IF NOT EXISTS invitee_update_updated_at
AFTER UPDATE ON invitee
BEGIN
    UPDATE invitee
    SET updated_at = CURRENT_TIMESTAMP
    WHERE id = NEW.id;
END;

-- migrate:down
DROP TABLE invitee;
