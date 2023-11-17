import { APIFixtureEvent } from "../../Models/api_types";
import React from "react";
import Typography from "@mui/material/Typography";

type EventTimeProps = {
  event: APIFixtureEvent;
};

export const EventTime: React.FC<EventTimeProps> = ({ event }) => {
  return (
    <Typography variant="subtitle1">
      {event.time.elapsed}
      {""}
      {event.time.extra !== undefined && event.time.extra > 0
        ? `+${event.time.extra}'`
        : `'`}
    </Typography>
  );
};
