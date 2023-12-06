import * as React from "react";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import { useMediaQuery } from "@mui/material";

interface ListFixtureScoreProps {
  homeScore: number | undefined;
  awayScore: number | undefined;
  status: string | undefined;
  time: Date | undefined;
}

const ListFixtureScore: React.FC<ListFixtureScoreProps> = ({
  homeScore,
  awayScore,
  status,
  time,
}) => {
  const isMobile = useMediaQuery("(max-width: 1260px)");
  const fontSize = isMobile ? "13px" : "16px";
  const timeString = time
    ? time
        .toLocaleTimeString([], { hour: "2-digit", minute: "2-digit" })
        .replace(/^0+/, "")
    : "";
  const displayScore =
    status === "Not Started" ? timeString : `${homeScore} - ${awayScore}`;

  return (
    <Box
      sx={{
        paddingX: isMobile ? 1 : 2,
        width: isMobile ? "20%" : "25%",
        display: "flex",
        flexDirection: "column",
        justifyContent: "center",
        alignItems: "center",
      }}
    >
      <Typography
        variant="subtitle2"
        sx={{
          fontSize: fontSize,
          textAlign: "center",
          whiteSpace: "nowrap",
        }}
      >
        {displayScore}
      </Typography>
    </Box>
  );
};

export default ListFixtureScore;
