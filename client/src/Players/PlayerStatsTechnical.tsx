import * as React from "react";
import Grid from "@mui/material/Grid";
import { Tile } from "../Util/TileTS";
import { APIPlayer } from "../Models/api_types";
import PlayerIDStatBox from "./PlayerIDStatBox";
import Typography from "@mui/material/Typography";

interface PlayerStatsProps {
  playerData: APIPlayer | undefined;
}

const PlayerStats: React.FC<PlayerStatsProps> = ({ playerData }) => {
  return (
    <Tile
      sx={{
        flexDirection: "column",
        gap: "24px",
      }}
    >
      <Grid container spacing={1}>
        <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
          <Typography
            sx={{
              fontSize: "20px",
              textAlign: "left",
              whiteSpace: "nowrap",
            }}
          >
            Technical
          </Typography>
        </Grid>
        <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
          <PlayerIDStatBox
            stat="Dribble Attempts"
            value={playerData?.statistics?.[0].technical.dribble_attempts}
          />
        </Grid>
        <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
          <PlayerIDStatBox
            stat="Dribbles Success"
            value={playerData?.statistics?.[0].technical.dribbles_success}
          />
        </Grid>
        <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
          <PlayerIDStatBox
            stat="Assists"
            value={playerData?.statistics?.[0].shooting.assists}
          />
        </Grid>
        <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
          <PlayerIDStatBox
            stat="Passes"
            value={playerData?.statistics?.[0].technical.passes}
          />
        </Grid>
        <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
          <PlayerIDStatBox
            stat="Key Passes"
            value={playerData?.statistics?.[0].technical.key_passes}
          />
        </Grid>
      </Grid>
    </Tile>
  );
};

export default PlayerStats;
