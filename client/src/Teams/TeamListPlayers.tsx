import * as React from "react";
import Box from "@mui/material/Box";
import Grid from "@mui/material/Grid";
import Typography from "@mui/material/Typography";
import { Flex } from "../Util/Flex.jsx";
import PlayerCard from "../Players/PlayerCardCopy";
import { APIPlayer } from "../Models/api_types";

interface DisplayPlayersProps {
  playerList: APIPlayer[] | undefined;
}

const DisplayPlayers: React.FC<DisplayPlayersProps> = ({ playerList }) => {
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
        <Typography
          sx={{ fontSize: { xs: "20px", sm: "20px" }, textAlign: "left" }}
        >
          Squad
        </Typography>
      </Flex>
      <Grid container spacing={1}>
        {playerList?.map((player, playerIndex) => (
          <Grid item xs={12} sm={12} md={12} lg={3} key={playerIndex}>
            <PlayerCard
              slug={player.slug}
              name={player.name}
              photo={player.photo}
            />
          </Grid>
        ))}
      </Grid>
    </Box>
  );
};

export default DisplayPlayers;
