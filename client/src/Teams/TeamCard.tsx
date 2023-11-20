import * as React from "react";
import Box from "@mui/material/Box";
import Card from "@mui/material/Card";
import CardContent from "@mui/material/CardContent";
import Typography from "@mui/material/Typography";
import DisplayImage from "../Util/DisplayImage";
import { Flex } from "../Util/Flex.jsx";
import { Link } from "react-router-dom";

interface TeamCardProps {
  slug: string;
  nameLong: string;
  nameShort: string;
  badge: string;
}

const TeamCard: React.FC<TeamCardProps> = ({
  slug,
  nameLong,
  nameShort,
  badge,
}) => {
  return (
    <Link to={`/teams/${slug}`}>
      <Card>
        <CardContent sx={{ padding: "8px" }}>
          <Flex
            style={{
              flexDirection: "row",
              gap: "16px",
              justifyContent: "space-between",
              alignItems: "center",
            }}
          >
            <Box
              sx={{
                display: "flex",
                justifyContent: "left",
                alignItems: "center",
              }}
            >
              <DisplayImage src={badge} alt={nameShort} />
            </Box>
            <Typography
              variant="h6"
              component="h2"
              sx={{ fontSize: "16px", flex: "1" }}
            >
              {nameLong}
            </Typography>
          </Flex>
        </CardContent>
      </Card>
    </Link>
  );
};

export default TeamCard;
