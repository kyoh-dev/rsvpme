import type { JSX } from "preact";
import type { Event } from "~/server/services/event/schema";

interface ViewEventProps extends JSX.HTMLAttributes<HTMLDivElement> {
  event: Event;
}

export default function ViewEvent({ event, ...props }: ViewEventProps) {
  return (
    <div {...props}>
      <h1 className="text-4xl font-bold text-neutral mb-2">{event.title}</h1>
      {event.description && (
        <p className="text-lg text-base-content/80 mb-6">{event.description}</p>
      )}
      <div className="space-y-4">
        <div className="flex items-center gap-2">
          <span className="font-semibold">Date:</span>
          {event.date ? (
            <span>{new Date(event.date).toLocaleDateString()}</span>
          ) : (
            <span className="text-base-content/60 flex items-center gap-2">
              <span>Date not set</span>
            </span>
          )}
        </div>
        <div className="flex items-center gap-2">
          <span className="font-semibold">Address:</span>
          {event.address ? (
            <span>{event.address}</span>
          ) : (
            <span className="text-base-content/60 flex items-center gap-2">
              <span>Location not specified</span>
            </span>
          )}
        </div>
        <div className="flex items-center gap-2">
          <span className="font-semibold">RSVP by:</span>
          {event.rsvpByDate ? (
            <span>{new Date(event.rsvpByDate).toLocaleDateString()}</span>
          ) : (
            <span className="text-base-content/60 flex items-center gap-2">
              <span>No RSVP deadline set</span>
            </span>
          )}
        </div>
      </div>
    </div>
  );
}
