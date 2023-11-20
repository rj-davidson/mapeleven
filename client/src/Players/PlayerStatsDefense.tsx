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
            Defense
          </Typography>
        </Grid>
        <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
          <PlayerIDStatBox
            stat="Duels Wons"
            value={playerData?.statistics?.[0].defense.duels_won}
          />
        </Grid>
        <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
          <PlayerIDStatBox
            stat="Blocks"
            value={playerData?.statistics?.[0].defense.blocks}
          />
        </Grid>
        <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
          <PlayerIDStatBox
            stat="Interceptions"
            value={playerData?.statistics?.[0].defense.interceptions}
          />
        </Grid>
      </Grid>
    </Tile>
  );
};

export default PlayerStats;
