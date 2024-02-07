-- migrate:up
CREATE VIEW event_detail AS (
    SELECT 
        uuid, 
        title, 
        description, 
        start_datetime, 
        finish_datetime, 
        address,
        (
            SELECT 
                json_agg(
                json_build_object(
                'first_name', invitee.first_name,
                'last_name', invitee.last_name,
                'email', invitee.email,
                'rsvp', invitee.rsvp
                )
                ) AS invitees_array
            FROM invitee
            WHERE event.id = invitee.event_id
        ) AS invitees
    FROM event
);

-- migrate:down
DROP VIEW event_detail;
