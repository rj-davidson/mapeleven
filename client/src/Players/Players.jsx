import React from "react";
import "../App.css";
import { Tile } from "../Util/TileTS";
import Grid from "@mui/material/Grid";
import DisplayPlayers from "./DisplayPlayersCopy.tsx";

function Players() {
  return (
    <Grid container spacing={2}>
      <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
        <Tile>
          <DisplayPlayers skeleton={15} limit={20} />
        </Tile>
      </Grid>
    </Grid>
  );
}

export default Players;
