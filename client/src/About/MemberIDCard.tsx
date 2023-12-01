import * as React from "react";
import Box from "@mui/material/Box";
import Grid from "@mui/material/Grid";
import { Tile } from "../Util/TileTS";
import { Flex } from "../Util/Flex.jsx";
import DisplayImage from "../Util/DisplayImage";
import Typography from "@mui/material/Typography";
import LinkedIn from "@mui/icons-material/LinkedIn";

interface MemberIDCardProps {
  name: string;
  picture: string;
  age: number;
  contact: string;
  linkedin: string;
}

const MemberIDCard: React.FC<MemberIDCardProps> = ({
  name,
  picture,
  contact,
  linkedin,
}) => {
  return (
    <Box>
      <Tile
        sx={{
          flexDirection: "column",
          gap: "32px",
        }}
      >
        <Flex
          style={{
            justifyContent: "space-between",
            alignItems: "center",
          }}
        >
          <Flex style={{ flexDirection: "column", alignItems: "left" }}>
            <Flex
              style={{
                flexDirection: "row",
                alignItems: "center",
                gap: "8px",
              }}
            >
              <Typography
                sx={{
                  fontSize: "24px",
                  fontWeight: "bold",
                }}
              >
                {name}
              </Typography>
            </Flex>
          </Flex>
          <a target="_blank" rel="noopener noreferrer" href={linkedin}>
            <LinkedIn sx={{ fontSize: "32px" }} />
          </a>
        </Flex>

        <Flex
          style={{
            justifyContent: "center",
            alignItems: "center",
            marginTop: "-20px",
          }}
        >
          <DisplayImage
            location={""}
            src={picture}
            alt={name}
            width={"216px"}
            height={"216px"}
            sx={{
              borderRadius: "10px",
              border: "5px solid var(--light0)",
            }}
          />
        </Flex>
        <Grid container spacing={1}>
          <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
            <Tile
              sx={{
                justifyContent: "space-between",
                alignItems: "center",
                background: "var(--dark1)",
                border: "none",
              }}
            >
              {contact}
            </Tile>
          </Grid>
        </Grid>
      </Tile>
    </Box>
  );
};

export default MemberIDCard;
