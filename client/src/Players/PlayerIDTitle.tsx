import * as React from "react";
import Typography from "@mui/material/Typography";
import { Flex } from "../Util/Flex.jsx";
import { Link } from "react-router-dom";
import HealingIcon from "@mui/icons-material/Healing";
import DisplayImage from "../Util/DisplayImage";
import { APIPlayer } from "../Models/api_types";
import Box from "@mui/material/Box";

interface PlayerIDTitleProps {
  slug: string;
  playerData: APIPlayer | undefined;
}

const PlayerIDTitle: React.FC<PlayerIDTitleProps> = ({ slug, playerData }) => {
  const getRatingColor = (rating) => {
    if (rating >= 7.5) {
      return "rgba(30, 180, 30, 1)";
    } else if (rating > 7.0) {
      return "var(--green)";
    } else if (rating > 6.0) {
      return "var(--yellow)";
    } else if (rating > 5.0) {
      return "var(--orange)";
    }
    return "var(--red)";
  };

  return (
    <Flex style={{ justifyContent: "space-between" }}>
      <Flex style={{ flexDirection: "column", alignItems: "left" }}>
        <Link to={`/players/${slug}`}>
          <Flex
            style={{ flexDirection: "row", alignItems: "center", gap: "8px" }}
          >
            <Typography
              sx={{
                fontSize: "24px",
                fontWeight: "bold",
                "&:hover": {
                  textDecoration: "underline", // Add underline on hover
                },
              }}
            >
              {playerData?.name}
            </Typography>
            <Flex>
              {playerData?.injured ?? (
                <HealingIcon style={{ fill: "var(--red)" }} />
              )}
            </Flex>
          </Flex>
        </Link>
        <Link to={`/teams/${playerData?.statistics?.[0].team.slug}`}>
          <Flex
            style={{ flexDirection: "row", gap: "6px", alignItems: "center" }}
          >
            <DisplayImage
              src={playerData?.statistics?.[0].team.logo}
              alt={playerData?.statistics?.[0].team.name}
              width={"24px"}
              height={"24px"}
            />
            <Typography
              sx={{
                fontSize: "14px",
                color: "var(--light0)",
                "&:hover": {
                  textDecoration: "underline", // Add underline on hover
                },
              }}
            >
              {playerData?.statistics?.[0].team.name}
            </Typography>
          </Flex>
        </Link>
      </Flex>

      <Box
        sx={{
          display: "flex",
          background: "var(--dark0)",
          borderRadius: "10px",
          padding: "0px 16px 0px 16px",
          maxHeight: "60px",
          justifyContent: "center",
          alignItems: "center",
        }}
      >
        <Typography
          sx={{
            fontSize: "22px",
            color: getRatingColor(
              parseFloat(
                playerData?.statistics?.[0].games.rating.substring(0, 3) ??
                  "0.0"
              )
            ),
          }}
        >
          {playerData?.statistics?.[0].games.rating.substring(0, 3)}
        </Typography>
      </Box>
    </Flex>
  );
};

export default PlayerIDTitle;
