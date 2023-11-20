import * as React from "react";
import Grid from "@mui/material/Grid";
import { Tile } from "../Util/TileTS";
import TeamIDStatBox from "./TeamIDStatBox";
import { APITeam } from "../Models/api_types";

interface TeamStatsProps {
  teamData: APITeam | undefined;
}

const TeamStats: React.FC<TeamStatsProps> = ({ teamData }) => {
  return (
    <Tile
      sx={{
        flexDirection: "column",
        gap: "24px",
      }}
    >
      <Grid container spacing={1}>
        <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
          <TeamIDStatBox
            stat="Wins"
            value={
              (teamData?.competitions?.[0].stats?.fixtures.wins.home || 0) +
              (teamData?.competitions?.[0].stats?.fixtures.wins.away || 0)
            }
          />
        </Grid>
        <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
          <TeamIDStatBox
            stat="Draws"
            value={
              (teamData?.competitions?.[0].stats?.fixtures.draws.home || 0) +
              (teamData?.competitions?.[0].stats?.fixtures.draws.away || 0)
            }
          />
        </Grid>
        <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
          <TeamIDStatBox
            stat="Losses"
            value={
              (teamData?.competitions?.[0].stats?.fixtures.loses.home || 0) +
              (teamData?.competitions?.[0].stats?.fixtures.loses.away || 0)
            }
          />
        </Grid>
        <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
          <TeamIDStatBox
            stat="Goals"
            value={
              teamData?.competitions?.[0].stats?.goals.for.total.total || 0
            }
          />
        </Grid>
        <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
          <TeamIDStatBox
            stat="Average Goals"
            value={
              teamData?.competitions?.[0].stats?.goals.for.average.total || 0
            }
          />
        </Grid>
        <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
          <TeamIDStatBox
            stat="Matches Scored"
            value={
              (teamData?.competitions?.[0].stats?.fixtures.played.home || 0) +
              (teamData?.competitions?.[0].stats?.fixtures.played.away || 0) -
              (teamData?.competitions?.[0].stats?.failed_to_score.total || 0)
            }
          />
        </Grid>
        <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
          <TeamIDStatBox
            stat="Failed To Score"
            value={
              teamData?.competitions?.[0].stats?.failed_to_score.total || 0
            }
          />
        </Grid>
        <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
          <TeamIDStatBox
            stat="Clean Sheets"
            value={teamData?.competitions?.[0].stats?.clean_sheet.total || 0}
          />
        </Grid>
      </Grid>
    </Tile>
  );
};

export default TeamStats;
