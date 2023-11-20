import * as React from "react";
import Typography from "@mui/material/Typography";
import { Flex } from "../Util/Flex.jsx";
import DisplayImage from "../Util/DisplayImage";
import { Link } from "react-router-dom";
import Box from "@mui/material/Box";

interface TeamIDTitleProps {
  slug: string | undefined;
  name: string | undefined;
  country: string | undefined;
  flag: string | undefined | null;
  leagueSlug?: string | undefined;
  leagueName?: string | undefined;
  leagueLogo?: string | undefined;
}

const TeamIDTitle: React.FC<TeamIDTitleProps> = ({
  name,
  country,
  flag,
  slug,
  leagueSlug,
  leagueName,
  leagueLogo,
}) => {
  return (
    <Flex style={{ flexDirection: "row", justifyContent: "space-between" }}>
      <Flex style={{ flexDirection: "column" }}>
        <Link to={`/teams/${slug}`}>
          <Typography
            sx={{
              fontSize: "24px",
              fontWeight: "bold",
              "&:hover": {
                textDecoration: "underline", // Add underline on hover
              },
            }}
          >
            {name}
          </Typography>
        </Link>

        <Flex
          style={{ flexDirection: "row", gap: "6px", alignItems: "center" }}
        >
          <DisplayImage
            src={flag}
            alt={country}
            width={"20px"}
            height={"16px"}
          />
          <Typography sx={{ fontSize: "14px", color: "var(--light0)" }}>
            {country}
          </Typography>
        </Flex>
      </Flex>
      {leagueLogo && leagueSlug && (
        <Flex style={{ flexDirection: "column" }}>
          <Flex
            style={{ flexDirection: "row", gap: "6px", alignItems: "center" }}
          >
            <Link to={`/leagues/${leagueSlug}`}>
              <Box
                sx={{
                  display: "flex",
                  background: "var(--dark1)",
                  borderRadius: "10px",
                  padding: "12px 12px 12px 12px",
                  maxHeight: "80px",
                  justifyContent: "center",
                  alignItems: "center",
                }}
              >
                <DisplayImage
                  src={leagueLogo}
                  alt={leagueName}
                  width={"48px"}
                  height={"48px"}
                  sx={{ margin: "4px" }}
                />
              </Box>
            </Link>
          </Flex>
        </Flex>
      )}
    </Flex>
  );
};

export default TeamIDTitle;
