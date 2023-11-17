import { APIFixtureEvent } from "../../Models/api_types";
import React from "react";
import Typography from "@mui/material/Typography";
import { Link } from "react-router-dom";
import SportsSoccer from "@mui/icons-material/SportsSoccer";
import Box from "@mui/material/Box";
import AirlineStops from "@mui/icons-material/AirlineStops";
import { EventTime } from "./EventTime";

type GoalEventProps = {
  event: APIFixtureEvent;
  away: boolean;
};

export const GoalEvent: React.FC<GoalEventProps> = ({ event, away }) => {
  return (
    <div style={{ textAlign: away ? "right" : "left" }}>
      <EventTime event={event} />
      <Link to={`/players/${event.player.slug}`}>
        <Typography
          sx={{
            "&:hover": {
              textDecoration: "underline",
            },
            fontSize: "20px",
          }}
        >
          <Box gap={1} display="inline-flex" alignItems="center">
            {away ? (
              <>
                {event.player.name} <SportsSoccer />
              </>
            ) : (
              <>
                <SportsSoccer /> {event.player.name}
              </>
            )}
          </Box>
        </Typography>
      </Link>
      {event.assist && (
        <Box>
          <Link to={`/players/${event.player.slug}`}>
            <Typography
              color={"text.secondary"}
              fontSize={"16px"}
              sx={{
                "&:hover": {
                  textDecoration: "underline",
                },
              }}
            >
              <Box gap={1} display="inline-flex" alignItems="center">
                {away ? (
                  <>
                    {event.assist.name}{" "}
                    <AirlineStops sx={{ fontSize: "16px" }} />
                  </>
                ) : (
                  <>
                    <AirlineStops sx={{ fontSize: "16px" }} />{" "}
                    {event.assist.name}
                  </>
                )}
              </Box>
            </Typography>
          </Link>
        </Box>
      )}
    </div>
  );
};
