-- migrate:up
CREATE TABLE event (
    id INTEGER NOT NULL PRIMARY KEY,
    uuid TEXT NOT NULL,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    start_datetime TEXT NOT NULL,
    finish_datetime TEXT,
    address TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX event_uuid_unique_idx ON event(uuid);

CREATE TRIGGER IF NOT EXISTS event_update_updated_at
AFTER UPDATE ON event
BEGIN
    UPDATE event
    SET updated_at = CURRENT_TIMESTAMP
    WHERE id = NEW.id;
END;


-- migrate:down
DROP TABLE event;
