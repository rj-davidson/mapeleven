import * as React from "react";
import Grid from "@mui/material/Grid";
import Typography from "@mui/material/Typography";
import { APIFixtureSet } from "../Models/api_types";
import DisplayImage from "../Util/DisplayImage";
import { Tile } from "../Util/TileTS";
import { EventTime } from "./Events/EventTime";
import Box from "@mui/material/Box";
import Sports from "@mui/icons-material/Sports";
import { Link } from "react-router-dom";

interface FixturePageSingleMainProps {
  data: APIFixtureSet | undefined;
}

const FixturePageSingleMain: React.FC<FixturePageSingleMainProps> = ({
  data,
}) => {
  return (
    data &&
    data.teams && (
      <Grid item xs={12} sm={12}>
        <Tile>
          <Grid container spacing={2} justifyContent="center" padding={"16px"}>
            <Grid item xs={12} sm={4}>
              <Box
                style={{
                  display: "flex",
                  alignItems: "center",
                  justifyContent: "center",
                }}
                gap={1}
              >
                <DisplayImage
                  src={data.teams.home.logo}
                  alt={data.teams.home.name}
                  width={"100px"}
                  height={"100px"}
                />
                <Link to={`/teams/${data.teams.home.slug}`}>
                  <Typography
                    variant="h6"
                    align="center"
                    sx={{
                      fontSize: "42px",
                      "&:hover": {
                        textDecoration: "underline", // Add underline on hover
                      },
                    }}
                  >
                    {data.teams.home.name}
                  </Typography>
                </Link>
              </Box>

              {data.events ? (
                data.events
                  .filter(
                    (event) =>
                      data.teams !== undefined &&
                      event.team.name === data.teams.home.name &&
                      event.type === "Goal"
                  )
                  .map((event, index) => (
                    <Typography variant="subtitle1" align="center" key={index}>
                      <Box
                        display={"flex-center"}
                        justifyContent={"center"}
                        alignItems={"center"}
                      >
                        {`${event.player.name} `}{" "}
                        <Box ml={1}>
                          {" "}
                          <EventTime event={event} />
                        </Box>
                      </Box>
                    </Typography>
                  ))
              ) : (
                <Typography variant="body2">No events available.</Typography>
              )}
            </Grid>
            <Grid item xs={12} sm={4}>
              <Typography
                variant="h4"
                align="center"
                sx={{
                  fontSize: "42px",
                }}
              >
                {data.fixture?.homeTeamScore} - {data.fixture?.awayTeamScore}
              </Typography>
              <Typography
                variant="h6"
                align="center"
                sx={{
                  fontSize: "20px",
                }}
              >
                {data.fixture?.elapsed}&#39;
              </Typography>
            </Grid>
            <Grid item xs={12} sm={4}>
              <Box
                style={{
                  display: "flex",
                  alignItems: "center",
                  justifyContent: "center",
                }}
                gap={1}
              >
                <DisplayImage
                  src={data.teams.away.logo}
                  alt={data.teams.away.name}
                  width={"100px"}
                  height={"100px"}
                ></DisplayImage>
                <Link to={`/teams/${data.teams.away.slug}`}>
                  <Typography
                    variant="h6"
                    align="center"
                    sx={{
                      fontSize: "42px",
                      "&:hover": {
                        textDecoration: "underline", // Add underline on hover
                      },
                    }}
                  >
                    {data.teams.away.name}
                  </Typography>
                </Link>
              </Box>
              {data.events ? (
                data.events
                  .filter(
                    (event) =>
                      event.type === "Goal" &&
                      data.teams !== undefined &&
                      event.team.name === data.teams.away.name
                  )
                  .map((event, index) => (
                    <Typography variant="subtitle1" align="center" key={index}>
                      <Box
                        display={"flex-center"}
                        justifyContent={"center"}
                        alignItems={"center"}
                      >
                        {`${event.player.name} `}{" "}
                        <Box ml={1}>
                          {" "}
                          <EventTime event={event} />
                        </Box>
                      </Box>
                    </Typography>
                  ))
              ) : (
                <Typography variant="body2">No events available.</Typography>
              )}
            </Grid>
            <Grid item xs={12} sm={12} sx={{ textAlign: "center" }}>
              <Typography
                variant="body1"
                align="center"
                sx={{
                  fontSize: "20px",
                }}
              >
                {data.fixture &&
                  new Date(data.fixture.date.toString()).toDateString()}{" "}
                {data.fixture?.timezone}
              </Typography>
            </Grid>
            <Grid item xs={12} sm={12}>
              <Typography
                variant="subtitle1"
                align="center"
                sx={{
                  fontSize: "16px",
                }}
              >
                <Box gap={1} display="inline-flex" alignItems="center">
                  <Sports sx={{ fontSize: "24px" }} /> {data.fixture?.referee}
                </Box>
              </Typography>
            </Grid>
          </Grid>
        </Tile>
      </Grid>
    )
  );
};

export default FixturePageSingleMain;
