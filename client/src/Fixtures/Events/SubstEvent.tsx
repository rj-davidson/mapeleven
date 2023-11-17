import { APIFixtureEvent } from "../../Models/api_types";
import React from "react";
import Typography from "@mui/material/Typography";
import { Link } from "react-router-dom";
import Box from "@mui/material/Box";
import ArrowUpward from "@mui/icons-material/ArrowUpward";
import ArrowDownward from "@mui/icons-material/ArrowDownward";
import { EventTime } from "./EventTime";

type SubstEventProps = {
  event: APIFixtureEvent;
  away: boolean;
};

export const SubstEvent: React.FC<SubstEventProps> = ({ event, away }) => {
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
                {event.assist && event.assist.name}{" "}
                <ArrowUpward style={{ fill: "var(--green)" }} />
              </>
            ) : (
              <>
                <ArrowUpward style={{ fill: "var(--green)" }} />{" "}
                {event.assist && event.assist.name}
              </>
            )}
          </Box>
        </Typography>
      </Link>
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
                  {event.player.name}{" "}
                  <ArrowDownward
                    style={{ fill: "var(--red)", fontSize: "16px" }}
                  />{" "}
                </>
              ) : (
                <>
                  <ArrowDownward
                    style={{ fill: "var(--red)", fontSize: "16px" }}
                  />{" "}
                  {event.player.name}
                </>
              )}
            </Box>
          </Typography>
        </Link>
      </Box>
    </div>
  );
};
