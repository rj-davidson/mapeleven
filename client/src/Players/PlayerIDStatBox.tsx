import * as React from "react";
import Typography from "@mui/material/Typography";
import { Tile } from "../Util/TileTS";
import Box from "@mui/material/Box";

interface PlayerIDStatBoxProps {
  stat: string;
  value: number | string | undefined;
}

const PlayerIDStatBox: React.FC<PlayerIDStatBoxProps> = ({ stat, value }) => {
  return (
    <Tile
      sx={{
        justifyContent: "space-between",
        alignItems: "center",
        background: "var(--dark1)",
        border: "none",
      }}
    >
      <Typography sx={{ fontSize: "18px" }}>{stat}</Typography>
      <Box
        sx={{
          background: "var(--dark0)",
          borderRadius: "10px",
          padding: "2px 10px 2px 10px",
        }}
      >
        <Typography sx={{ fontSize: "18px", color: "var(--light0)" }}>
          {value}
        </Typography>
      </Box>
    </Tile>
  );
};

export default PlayerIDStatBox;
