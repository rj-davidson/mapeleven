import * as React from "react";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import Grid from "@mui/material/Grid";
import Skeleton from "@mui/material/Skeleton";
import PlayerPageSingleColumn from "./PlayerPageSingleColumn";
import PlayerPageSingleMain from "./PlayerPageSingleMain";
import { APIPlayer } from "../Models/api_types";

const url = import.meta.env.VITE_API_URL;

function PlayerPageSingle() {
  const { slug } = useParams();
  const [playerData, setPlayerData] = useState<APIPlayer>();
  const [loading, setLoading] = useState<boolean>(true);

  useEffect(() => {
    fetch(`${url}/players/${slug}`)
      .then((response) => response.json())
      .then((jsonData) => {
        setPlayerData(jsonData as APIPlayer);
        setLoading(false);
      })
      .catch((error) => console.error(error));
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
      <Grid item xs={12} sm={12} md={12} lg={3} width="100%">
        <PlayerPageSingleColumn
          slug={slug && slug !== "" ? slug : ""}
          playerData={playerData}
        />
      </Grid>
      <Grid item xs={12} sm={12} md={12} lg={9} width="100%">
        <PlayerPageSingleMain playerData={playerData} />
      </Grid>
    </Grid>
  );
}

export default PlayerPageSingle;
