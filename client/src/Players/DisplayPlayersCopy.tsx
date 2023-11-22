import * as React from "react";
import { useEffect, useState } from "react";
import Box from "@mui/material/Box";
import Grid from "@mui/material/Grid";
import Skeleton from "@mui/material/Skeleton";
import Typography from "@mui/material/Typography";
import { Flex } from "../Util/Flex.jsx";
import Person from "@mui/icons-material/Person";
import PlayerCard from "./PlayerCardCopy";
import { APIPlayer } from "../Models/api_types";

const url = import.meta.env.VITE_API_URL;

interface DisplayPlayersProps {
  limit?: number;
  skeleton: number;
}

const DisplayPlayers: React.FC<DisplayPlayersProps> = ({ limit, skeleton }) => {
  const [playerData, setPlayerData] = useState<APIPlayer[]>();
  const [loading, setLoading] = useState<boolean>(true);

  useEffect(() => {
    fetch(`${url}/players`)
      .then((response) => response.json())
      .then((jsonData) => {
        const sortedData = jsonData.sort((a: APIPlayer, b: APIPlayer) => {
          const aPopularity = a.popularity !== undefined ? a.popularity : 0;
          const bPopularity = b.popularity !== undefined ? b.popularity : 0;
          return bPopularity - aPopularity;
        });

        if (limit) {
          setPlayerData(sortedData.slice(0, limit) as APIPlayer[]);
        } else {
          setPlayerData(sortedData as APIPlayer[]);
        }

        setLoading(false);
      })
      .catch((error) => console.error(error));
  }, []);

  const skeletonCards: JSX.Element[] = [];

  for (let i = 0; i < skeleton; i++) {
    skeletonCards.push(
      <Grid item xs={12} sm={12} md={12} lg={12} key={i}>
        <Skeleton
          variant="rectangular"
          height="52px"
          width="100%"
          sx={{ background: "var(--dark1)", borderRadius: "12px" }}
        />
      </Grid>
    );
  }

  return (
    <Box sx={{ width: "100%" }}>
      <Flex
        style={{
          flexDirection: "row",
          justifyContent: "space-between",
          alignItems: "center",
          margin: "0 4px 16px 2px",
        }}
      >
        <Typography variant="h6" component="h2" gutterBottom>
          Popular Players
        </Typography>
        <Person />
      </Flex>
      {loading ? (
        <Grid container spacing={1}>
          {skeletonCards}
        </Grid>
      ) : (
        <Grid container spacing={1}>
          {playerData &&
            playerData.map((player, playerIndex) => (
              <Grid item xs={12} sm={12} md={12} lg={12} key={playerIndex}>
                <PlayerCard
                  slug={player.slug}
                  name={player.name}
                  photo={player.photo}
                />
              </Grid>
            ))}
        </Grid>
      )}
    </Box>
  );
};

export default DisplayPlayers;
