import Typography from "@mui/material/Typography";
import Grid from "@mui/material/Grid";
import Divider from "@mui/material/Divider";
import { type APIFixtureEvent } from "../../Models/api_types";
import RectangleRounded from "@mui/icons-material/RectangleRounded";
import SportsSoccer from "@mui/icons-material/SportsSoccer";
import Support from "@mui/icons-material/Support";
import ArrowUpward from "@mui/icons-material/ArrowUpward";
import ArrowDownward from "@mui/icons-material/ArrowDownward";
import { Link } from "react-router-dom";
import React from "react";

type MatchEventProps = {
  event: APIFixtureEvent;
  awayTeamSlug: string;
  homeTeamSlug: string;
};

const FixtureMatchEvent: React.FC<MatchEventProps> = ({
  event,
  awayTeamSlug,
  homeTeamSlug,
}) => {
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
    <Grid container alignItems="center" spacing={2}>
      <Grid item xs={12} sx={{ paddingLeft: "16px", paddingRight: "16px" }}>
        {event.type === "subst" && event.team.slug === homeTeamSlug ? (
          <div>
            <Typography
              variant="body1"
              sx={{
                fontSize: "20px",
              }}
            >
              {event.time.elapsed}`{" "}
              {event.time.extra &&
                event.time.extra > 0 &&
                `+${event.time.extra}'`}
            </Typography>

            <Link to={`/players/${event.player.slug}`}>
              <Typography
                sx={{
                  // eslint-disable-next-line @typescript-eslint/naming-convention
                  "&:hover": {
                    textDecoration: "underline", // Add underline on hover
                  },
                  fontSize: "20px",
                }}
              >
                <ArrowDownward style={{ fill: "var(--red)" }} />
                {event.player.name}
              </Typography>
            </Link>

            <br />
            <Link to={`/players/${event.assist && event.assist.slug}`}>
              <Typography
                sx={{
                  "&:hover": {
                    textDecoration: "underline", // Add underline on hover
                  },
                  fontSize: "20px", // Adjust the font size as needed
                }}
              >
                <ArrowUpward style={{ fill: "var(--green)" }} />
                {event.assist && event.assist.name}
              </Typography>
            </Link>
          </div>
        ) : null}

        {event.type === "subst" && event.team.slug === awayTeamSlug ? (
          <div style={{ textAlign: "right" }}>
            <Typography
              variant="body1"
              sx={{
                fontSize: "20px", // Adjust the font size as needed
              }}
            >
              {event.time.elapsed}`{" "}
              {event.time.extra &&
                event.time.extra > 0 &&
                `+${event.time.extra}'`}
            </Typography>
            <Link to={`/players/${event.player.slug}`}>
              <Typography
                sx={{
                  "&:hover": {
                    textDecoration: "underline", // Add underline on hover
                  },
                  fontSize: "20px", // Adjust the font size as needed
                }}
              >
                <ArrowDownward style={{ fill: "var(--red)" }} />
                {event.player.name}
              </Typography>
            </Link>

            <br />
            <Link to={`/players/${event.assist && event.assist.slug}`}>
              <Typography
                sx={{
                  "&:hover": {
                    textDecoration: "underline", // Add underline on hover
                  },
                  fontSize: "20px", // Adjust the font size as needed
                }}
              >
                <ArrowUpward style={{ fill: "var(--green)" }} />
                {event.assist && event.assist.name}
              </Typography>
            </Link>
          </div>
        ) : null}

        {event.type === "Goal" && event.team.slug === homeTeamSlug ? (
          <div>
            <Typography
              variant="body1"
              sx={{
                fontSize: "20px",
              }}
            >
              {event.time.elapsed}`{" "}
              {event.time.extra &&
                event.time.extra > 0 &&
                `+${event.time.extra}'`}
            </Typography>
            <Link to={`/players/${event.player.slug}`}>
              <Typography
                sx={{
                  "&:hover": {
                    textDecoration: "underline", // Add underline on hover
                  },
                  fontSize: "20px", // Adjust the font size as needed
                }}
              >
                <SportsSoccer />
                {event.player.name}
              </Typography>
            </Link>
            {event.assist && (
              <>
                <br />
                <Link to={`/players/${event.player.slug}`}>
                  <Typography
                    sx={{
                      "&:hover": {
                        textDecoration: "underline", // Add underline on hover
                      },
                      fontSize: "20px", // Adjust the font size as needed
                    }}
                  >
                    <Support />
                    {event.assist.name}
                  </Typography>
                </Link>
              </>
            )}
          </div>
        ) : null}

        {event.type === "Goal" && event.team.slug === awayTeamSlug ? (
          <div style={{ textAlign: "right" }}>
            <Typography
              variant="body1"
              sx={{
                fontSize: "20px", // Adjust the font size as needed
                fontFamily: "Arial, sans-serif", // Specify the desired font-family
              }}
            >
              {event.time.elapsed}`{" "}
              {event.time.extra &&
                event.time.extra > 0 &&
                `+${event.time.extra}'`}
            </Typography>
            <Link to={`/players/${event.player.slug}`}>
              <Typography
                sx={{
                  "&:hover": {
                    textDecoration: "underline", // Add underline on hover
                  },
                  fontSize: "20px", // Adjust the font size as needed
                }}
              >
                <SportsSoccer />
                {event.player.name}
              </Typography>
            </Link>
            {event.assist && (
              <>
                <br />
                <Link to={`/players/${event.player.slug}`}>
                  <Typography
                    sx={{
                      "&:hover": {
                        textDecoration: "underline", // Add underline on hover
                      },
                      fontSize: "20px", // Adjust the font size as needed
                      fontFamily: "Arial, sans-serif", // Specify the desired font-family
                    }}
                  >
                    <Support />
                    {event.assist.name}
                  </Typography>
                </Link>
              </>
            )}
          </div>
        ) : null}

        {getIcon(event) && (
          <Typography
            variant="body2"
            sx={{
              fontSize: "20px",
            }}
          >
            <Typography
              variant="body1"
              sx={{
                fontSize: "20px", // Adjust the font size as needed
              }}
            >
              {event.time.elapsed}`{" "}
              {event.time.extra &&
                event.time.extra > 0 &&
                `+${event.time.extra}'`}
            </Typography>
            <Link to={`/players/${event.player.slug}`}>
              <Typography
                sx={{
                  // eslint-disable-next-line @typescript-eslint/naming-convention
                  "&:hover": {
                    textDecoration: "underline", // Add underline on hover
                  },
                  fontSize: "20px", // Adjust the font size as needed
                }}
              >
                {getIcon(event)} {event.player.name}
              </Typography>
            </Link>
          </Typography>
        )}

        {event.comments && (
          <Typography
            variant="caption"
            sx={{
              fontSize: "20px", // Adjust the font size as needed
              fontFamily: "Arial, sans-serif", // Specify the desired font-family
            }}
          >
            {event.comments}
          </Typography>
        )}
      </Grid>

      {/* Add a horizontal divider for separation */}
      <Grid item xs={12}>
        <Divider />
      </Grid>
    </Grid>
  );
};

export default FixtureMatchEvent;
