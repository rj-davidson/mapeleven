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
            Stats
          </Typography>
        </Grid>
        <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
          <PlayerIDStatBox
            stat="Lineups"
            value={playerData?.statistics?.[0].games.lineups}
          />
        </Grid>
        <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
          <PlayerIDStatBox
            stat="Minutes"
            value={playerData?.statistics?.[0].games.minutes}
          />
        </Grid>
      </Grid>
    </Tile>
  );
};

export default PlayerStats;
