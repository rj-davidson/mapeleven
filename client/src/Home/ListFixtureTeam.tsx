import * as React from "react";
import Typography from "@mui/material/Typography";
import { useMediaQuery } from "@mui/material";
import DisplayImage from "../Util/DisplayImage";
import { Flex } from "../Util/Flex.jsx";
import { Link } from "react-router-dom";

interface ListFixtureTeamProps {
  slug: string;
  name: string;
  logo: string;
  home: boolean;
}

const ListFixtureTeam: React.FC<ListFixtureTeamProps> = ({
  slug,
  name,
  logo,
  home,
}) => {
  const isSmallerThanLG = useMediaQuery("(max-width: 1260px)");

  return home ? (
    <Flex
      style={{
        justifyContent: "right",
        alignItems: "center",
        width: "100%",
        gap: "8px",
      }}
    >
      <Link to={`/teams/${slug}`}>
        <Typography
          variant="h6"
          component="h2"
          sx={{
            fontSize: !isSmallerThanLG ? "16px" : "12px",
            textAlign: "right",
            "&:hover": {
              textDecoration: "underline",
            },
          }}
        >
          {name}
        </Typography>
      </Link>
      <DisplayImage src={logo} alt={name} />
    </Flex>
  ) : (
    <Flex
      style={{
        justifyContent: "left",
        alignItems: "center",
        width: "100%",
        gap: "8px",
      }}
    >
      <DisplayImage src={logo} alt={name} />
      <Link to={`/teams/${slug}`}>
        <Typography
          variant="h6"
          component="h2"
          sx={{
            fontSize: !isSmallerThanLG ? "16px" : "12px",
            textAlign: "left",
            "&:hover": {
              textDecoration: "underline",
            },
          }}
        >
          {name}
        </Typography>
      </Link>
    </Flex>
  );
};

export default ListFixtureTeam;
