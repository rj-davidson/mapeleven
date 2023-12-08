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
import { useMediaQuery } from "@mui/material";
import LinearProgress from "@mui/material/LinearProgress";

interface FixturePageSingleMainProps {
  data: APIFixtureSet | undefined;
}

const FixturePageSingleMain: React.FC<FixturePageSingleMainProps> = ({
  data,
}) => {
  const isMobile = useMediaQuery("(max-width: 1260px)");

  return (
    data &&
    data.teams && (
      <Grid item xs={12} sm={12}>
        <Tile>
          <Grid container spacing={2} justifyContent="center" padding={"16px"}>
            <Grid
              container
              item
              xs={12}
              justifyContent="center"
              alignItems="center"
              spacing={isMobile ? 0 : 3}
            >
              <Grid item xs={4}>
                {isMobile ? (
                  <Box
                    style={{
                      display: "flex",
                      flexDirection: "column",
                      alignItems: "center",
                      justifyContent: "flex-end",
                    }}
                  >
                    <DisplayImage
                      src={data.teams.home.logo}
                      alt={data.teams.home.name}
                      width="64px"
                      height="64px"
                    />
                    <Link to={`/teams/${data.teams.home.slug}`}>
                      <Typography
                        variant="subtitle2"
                        sx={{ "&:hover": { textDecoration: "underline" } }}
                      >
                        {data.teams.home.name}
                      </Typography>
                    </Link>
                  </Box>
                ) : (
                  <Box
                    style={{
                      display: "flex",
                      justifyContent: "flex-end",
                      alignItems: "center",
                    }}
                    gap={1}
                  >
                    <Link to={`/teams/${data.teams.home.slug}`}>
                      <Typography
                        variant="h4"
                        sx={{ "&:hover": { textDecoration: "underline" } }}
                        paddingRight="8px"
                      >
                        {data.teams.home.name}
                      </Typography>
                    </Link>
                    <DisplayImage
                      src={data.teams.home.logo}
                      alt={data.teams.home.name}
                      width="96px"
                      height="96px"
                    />
                  </Box>
                )}
              </Grid>

              <Grid item xs={4}>
                {data.fixture?.status !== "Not Started" && (
                  <>
                    <Typography variant={isMobile ? "h5" : "h4"} align="center">
                      {data.fixture?.homeTeamScore} -{" "}
                      {data.fixture?.awayTeamScore}
                    </Typography>

                    <Typography
                      variant={isMobile ? "subtitle2" : "h6"}
                      align="center"
                      paddingTop={isMobile ? 0.5 : 1}
                      fontSize={isMobile ? "14px" : undefined}
                    >
                      {[
                        "In Progress",
                        "First Half",
                        "Second Half",
                        "Extra Time",
                      ].indexOf(data.fixture?.status ?? "") !== -1
                        ? `${data.fixture?.elapsed}'`
                        : data.fixture?.status ?? `${data.fixture?.elapsed}'`}
                    </Typography>

                    {[
                      "In Progress",
                      "First Half",
                      "Second Half",
                      "Extra Time",
                    ].indexOf(data.fixture?.status ?? "") !== -1 && (
                      <Box
                        sx={{
                          width: "100%",
                          color: "var(--green)",
                          display: "flex",
                          justifyContent: "center",
                        }}
                      >
                        <Box
                          width={isMobile ? "32px" : "44px"}
                          paddingRight={"2px"}
                        >
                          <LinearProgress color={"inherit"} />
                        </Box>
                      </Box>
                    )}
                  </>
                )}
              </Grid>

              <Grid item xs={4}>
                {isMobile ? (
                  <Box
                    style={{
                      display: "flex",
                      flexDirection: "column",
                      alignItems: "center",
                      justifyContent: "flex-start",
                    }}
                  >
                    <DisplayImage
                      src={data.teams.away.logo}
                      alt={data.teams.away.name}
                      width="64px"
                      height="64px"
                    />
                    <Link to={`/teams/${data.teams.away.slug}`}>
                      <Typography
                        variant="subtitle2"
                        sx={{ "&:hover": { textDecoration: "underline" } }}
                      >
                        {data.teams.away.name}
                      </Typography>
                    </Link>
                  </Box>
                ) : (
                  <Box
                    style={{
                      display: "flex",
                      justifyContent: "flex-start",
                      alignItems: "center",
                    }}
                    gap={1}
                  >
                    <DisplayImage
                      src={data.teams.away.logo}
                      alt={data.teams.away.name}
                      width="96px"
                      height="96px"
                    />
                    <Link to={`/teams/${data.teams.away.slug}`}>
                      <Typography
                        variant="h4"
                        sx={{ "&:hover": { textDecoration: "underline" } }}
                        paddingLeft="8px"
                      >
                        {data.teams.away.name}
                      </Typography>
                    </Link>
                  </Box>
                )}
              </Grid>
            </Grid>
            <Grid
              container
              item
              xs={12}
              justifyContent="space-between"
              alignItems="flex-start"
            >
              <Grid item xs={12} sm={4} paddingRight={16}>
                {data.events && !isMobile
                  ? data.events
                      .filter(
                        (event) =>
                          data.teams !== undefined &&
                          event.team.name === data.teams.home.name &&
                          event.type === "Goal" &&
                          event.detail !== "Missed Penalty"
                      )
                      .map((event, index) => (
                        <Typography
                          variant="subtitle1"
                          align="right"
                          key={index}
                        >
                          <Box
                            display="flex"
                            justifyContent="flex-end"
                            alignItems="center"
                          >
                            {`${event.player.name} `}{" "}
                            <Box ml={1}>
                              <EventTime event={event} />
                            </Box>
                          </Box>
                        </Typography>
                      ))
                  : undefined}
              </Grid>

              <Grid item xs={12} sm={4}>
                <Typography
                  variant={isMobile ? "subtitle2" : "h6"}
                  align="center"
                >
                  {data.fixture &&
                    new Date(data.fixture.date.toString()).toDateString()}
                  {data.fixture &&
                    data.fixture.status === "Not Started" &&
                    ` - ${new Date(
                      data.fixture.date.toString()
                    ).toLocaleTimeString([], {
                      hour: "2-digit",
                      minute: "2-digit",
                    })}`}
                </Typography>
                <Typography
                  variant={isMobile ? "body2" : "subtitle1"}
                  align="center"
                >
                  <Box gap={1} display="inline-flex" alignItems="center">
                    <Sports sx={{ fontSize: isMobile ? "16px" : "24px" }} />{" "}
                    {data.fixture?.referee}
                  </Box>
                </Typography>
              </Grid>

              <Grid item xs={12} sm={4} paddingLeft={16}>
                {data.events && !isMobile
                  ? data.events
                      .filter(
                        (event) =>
                          event.type === "Goal" &&
                          data.teams !== undefined &&
                          event.team.name === data.teams.away.name &&
                          event.detail !== "Missed Penalty"
                      )
                      .map((event, index) => (
                        <Typography
                          variant="subtitle1"
                          align="left"
                          key={index}
                        >
                          <Box
                            display="flex"
                            justifyContent="flex-start"
                            alignItems="center"
                          >
                            {`${event.player.name} `}{" "}
                            <Box ml={1}>
                              <EventTime event={event} />
                            </Box>
                          </Box>
                        </Typography>
                      ))
                  : undefined}
              </Grid>
            </Grid>
          </Grid>
        </Tile>
      </Grid>
    )
  );
};

export default FixturePageSingleMain;
