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
import { Tile } from "../Util/Tile";
import FixtureLineup from "./FixtureLineup";
import FixtureLineupMobile from "./FixtureLineupMobile";

const url = import.meta.env.VITE_API_URL;

function FixturePageSingle() {
  const { slug } = useParams();
  const isSmallerThanLG = useMediaQuery("(max-width: 1260px)");
  const [fixtureData, setFixtureData] = useState<APIFixtureSet>();
  const [loading, setLoading] = useState<boolean>(true);

  useEffect(() => {
    fetch(`${url}/fixtures/${slug}`)
      .then((response) => response.json())
      .then((jsonData) => {
        setFixtureData(jsonData as APIFixtureSet);
        setLoading(false);
      })
      .catch();
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
          </Grid>
          <Grid item xs={12} sm={6}>
            <Tile style={{ flexDirection: "column" }}>
              {fixtureData?.events &&
                fixtureData.events.map((event, index) => (
                  <Grid item xs={12} key={index}>
                    {fixtureData.teams && (
                      <FixtureMatchEvents
                        event={event}
                        awayTeamSlug={fixtureData.teams.away.slug}
                      />
                    )}
                  </Grid>
                ))}
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
          </Grid>
        </Grid>
      </Grid>
    </Grid>
  );
}

export default FixturePageSingle;
