import { APIFixtureEvent } from "../../Models/api_types";
import RectangleRounded from "@mui/icons-material/RectangleRounded";
import React from "react";
import Typography from "@mui/material/Typography";
import { Link } from "react-router-dom";
import Box from "@mui/material/Box";
import { EventTime } from "./EventTime";

type CardEventProps = {
  event: APIFixtureEvent;
  away: boolean;
};

export const CardEvent: React.FC<CardEventProps> = ({ event, away }) => {
  const getIcon = (event: APIFixtureEvent) => {
    if (event.type === "Card" && event.detail === "Yellow Card") {
      return (
        <RectangleRounded
          style={{ fill: "var(--yellow)", transform: "rotate(90deg)" }}
        />
      );
    }

    if (event.type === "Card" && event.detail === "Red Card") {
      return (
        <RectangleRounded
          style={{ fill: "var(--red)", transform: "rotate(90deg)" }}
        />
      );
    }
  };

  return (
    getIcon(event) && (
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
                  {event.player.name} {getIcon(event)}
                </>
              ) : (
                <>
                  {getIcon(event)} {event.player.name}
                </>
              )}
            </Box>
          </Typography>
        </Link>
        {event.comments && (
          <Box>
            <Typography color={"text.secondary"} fontSize={"16px"}>
              <Box gap={1} display="inline-flex" alignItems="center">
                {event.comments}
              </Box>
            </Typography>
          </Box>
        )}
      </div>
    )
  );
};
