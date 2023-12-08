import * as React from "react";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import Grid from "@mui/material/Grid";
import Skeleton from "@mui/material/Skeleton";
import { useMediaQuery } from "@mui/material";
import { APIFixtureSet } from "../Models/api_types";
import FixtureMatchEvents from "./Events/FixtureMatchEvents";
import FixturePageSingleMain from "./FixturePageSingleMain";
import TeamIDCardFixturePage from "../Teams/TeamIDCardFixturePage";
import { Tile } from "../Util/TileTS";
import FixtureLineup from "./FixtureLineup";
import FixtureLineupMobile from "./FixtureLineupMobile";
import FixturePlayerStepperAway from "./FixturePlayerStepperAway";
import FixturePlayerStepperHome from "./FixturePlayerStepperHome";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";

const url = import.meta.env.VITE_API_URL;

function FixturePageSingle() {
  const { slug } = useParams();
  const isSmallerThanLG = useMediaQuery("(max-width: 1260px)");
  const [fixtureData, setFixtureData] = useState<APIFixtureSet>();
  const [loading, setLoading] = useState<boolean>(true);

  useEffect(() => {
    const fetchData = () => {
      slug &&
        fetch(`${url}/fixtures/${slug}`)
          .then((response) => response.json())
          .then((jsonData) => {
            setFixtureData(jsonData as APIFixtureSet);
            setLoading(false);
          })
          .catch();
    };

    fetchData();
    const intervalId = setInterval(fetchData, 10000);
    return () => clearInterval(intervalId);
  }, [slug]);
  return loading ? (
    <Grid container spacing={2} direction="column">
      <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
        <Skeleton
          variant="rectangular"
          height="150px"
          width="100%"
          sx={{ background: "gray", borderRadius: "12px" }}
        />
      </Grid>
      <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
        <Skeleton
          variant="rectangular"
          height="600px"
          width="100%"
          sx={{ background: "gray", borderRadius: "12px" }}
        />
      </Grid>
    </Grid>
  ) : (
    <Grid container spacing={2}>
      <Grid item xs={12} sm={12}>
        <FixturePageSingleMain data={fixtureData} />
      </Grid>
      {fixtureData?.lineups &&
        fixtureData.lineups.home &&
        fixtureData.lineups.home.matchPlayers.length > 0 &&
        fixtureData.lineups.away &&
        fixtureData.lineups.away.matchPlayers.length > 0 && (
          <Grid item xs={12} sm={12}>
            {isSmallerThanLG ? (
              <FixtureLineupMobile
                homePlayers={fixtureData.lineups.home.matchPlayers}
                awayPlayers={fixtureData.lineups.away.matchPlayers}
              />
            ) : (
              <FixtureLineup
                homePlayers={fixtureData.lineups.home.matchPlayers}
                awayPlayers={fixtureData.lineups.away.matchPlayers}
              />
            )}
          </Grid>
        )}
      <Grid item xs={12} sm={12}>
        <Grid container spacing={2}>
          <Grid item xs={12} sm={3}>
            {fixtureData && fixtureData.teams && (
              <TeamIDCardFixturePage
                slug={fixtureData.teams.home.slug}
                name={fixtureData.teams.home.name}
                badge={fixtureData.teams.home.logo}
              />
            )}
            <div style={{ marginTop: "16px" }}>
              {fixtureData && fixtureData.teams && fixtureData.events && (
                <FixturePlayerStepperHome
                  events={fixtureData.events}
                  homeTeamName={fixtureData.teams.home.name}
                />
              )}
            </div>
          </Grid>
          <Grid item xs={12} sm={6}>
            <Tile sx={{ flexDirection: "column" }}>
              {fixtureData?.events ? (
                fixtureData.events
                  .sort((a, b) => a.time.elapsed - b.time.elapsed)
                  .map((event, index) => (
                    <Grid item xs={12} key={index}>
                      {fixtureData.teams && (
                        <FixtureMatchEvents
                          event={event}
                          awayTeamSlug={fixtureData.teams.away.slug}
                        />
                      )}
                    </Grid>
                  ))
              ) : (
                <Box
                  display="flex"
                  alignItems="center"
                  justifyContent="center"
                  paddingY={10}
                >
                  <Typography variant={"subtitle1"}>No Events</Typography>
                </Box>
              )}
            </Tile>
          </Grid>
          <Grid item xs={12} sm={3}>
            {fixtureData && fixtureData.teams && (
              <TeamIDCardFixturePage
                slug={fixtureData.teams.away.slug}
                name={fixtureData.teams.away.name}
                badge={fixtureData.teams.away.logo}
              />
            )}
            <div style={{ marginTop: "16px" }}>
              {fixtureData && fixtureData.teams && fixtureData.events && (
                <FixturePlayerStepperAway
                  events={fixtureData.events}
                  awayTeamName={fixtureData.teams.away.name}
                />
              )}
            </div>
          </Grid>
        </Grid>
      </Grid>
    </Grid>
  );
}

export default FixturePageSingle;
