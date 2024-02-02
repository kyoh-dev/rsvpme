CREATE TABLE IF NOT EXISTS "schema_migrations" (version varchar(128) primary key);
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
CREATE TRIGGER event_update_updated_at
AFTER UPDATE ON event
BEGIN
    UPDATE event
    SET updated_at = CURRENT_TIMESTAMP
    WHERE id = NEW.id;
END;
CREATE TABLE invitee (
    id INTEGER NOT NULL PRIMARY KEY,
    uuid TEXT NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT,
    rsvp BOOLEAN NOT NULL,
    event_id INTEGER NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (event_id) REFERENCES event(id) ON DELETE CASCADE
);
CREATE UNIQUE INDEX invitee_uuid_unique_idx ON invitee(uuid);
CREATE TRIGGER invitee_update_updated_at
AFTER UPDATE ON invitee
BEGIN
    UPDATE invitee
    SET updated_at = CURRENT_TIMESTAMP
    WHERE id = NEW.id;
END;
-- Dbmate schema migrations
INSERT INTO "schema_migrations" (version) VALUES
  ('20240202061936'),
  ('20240202063939');
