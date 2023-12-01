import * as React from "react";
import Typography from "@mui/material/Typography";
import "../App.css";
import Grid from "@mui/material/Grid";
import { Tile } from "../Util/TileTS";
import { Flex } from "../Util/Flex.jsx";
import { Link } from "react-router-dom";
import { useMediaQuery } from "@mui/material";
import teamPagePhoto from "./teampage.png";
import radarPhoto from "./radar.png";
import goalsPhoto from "./goals.png";
import erdPhoto from "./erd.png";
import DisplayImage from "../Util/DisplayImage";

function About() {
  const isSmallerThanLG = useMediaQuery("(max-width: 1260px)");
  return (
    <Grid container spacing={2}>
      <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
        <Tile
          sx={{
            justifyContent: "space-between",
            flexDirection: isSmallerThanLG ? "column" : "row",
            gap: "24px",
          }}
        >
          <Typography
            sx={{
              fontSize: { xs: "32px", sm: "32px" },
              marginLeft: "10px",
              fontWeight: "bold",
            }}
          >
            About Us
          </Typography>
          <Flex
            style={{
              justifyContent: "right",
              flexDirection: isSmallerThanLG ? "column" : "row",
              gap: isSmallerThanLG ? "12px" : "32px",
              fontSize: isSmallerThanLG ? "28px" : "32px",
              marginRight: "10px",
            }}
          >
            <Link to={"/"}>
              <Typography
                sx={{
                  fontSize: isSmallerThanLG ? "20px" : "32px",
                  marginLeft: "10px",
                  "&:hover": {
                    textDecoration: "underline", // Add underline on hover
                  },
                }}
              >
                Home
              </Typography>
            </Link>
            <Link to={"/tutorial"}>
              <Typography
                sx={{
                  fontSize: isSmallerThanLG ? "20px" : "32px",
                  marginLeft: "10px",
                  "&:hover": {
                    textDecoration: "underline", // Add underline on hover
                  },
                }}
              >
                Tutorial
              </Typography>
            </Link>
            <Link to={"/members"}>
              <Typography
                sx={{
                  fontSize: isSmallerThanLG ? "20px" : "32px",
                  marginLeft: "10px",
                  "&:hover": {
                    textDecoration: "underline", // Add underline on hover
                  },
                }}
              >
                Our Team
              </Typography>
            </Link>
          </Flex>
        </Tile>
      </Grid>
      <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
        <Tile
          sx={{ justifyContent: "left", alignItems: "left", textAlign: "left" }}
        >
          <Typography
            sx={{ fontSize: { sm: "20px", lg: "20px" }, padding: "10px" }}
          >
            MapEleven aims to provide soccer fans with tools to better
            understand the game by visualizing complex data. Through data
            visualization we can provide users with objective tools to assist
            them in interpreting matches, players and teams. Users currently
            enhance their understanding of the game through subjective resources
            such as television punditry and journalism. Current visualization
            tools for objective data only scratch the surface of the data being
            collected. They rely on simple numerical inferences such as showing
            the number of shots, passes and fouls for each team in a match. The
            subjective view provided by pundits only offer a partial picture of
            the game. We incorporate various objective data visualizations
            giving users a comprehensive understanding of the game
          </Typography>
        </Tile>
      </Grid>
      <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
        <Tile
          sx={{ justifyContent: "left", alignItems: "left", textAlign: "left" }}
        >
          <DisplayImage
            location={""}
            src={teamPagePhoto}
            alt={"Team Page"}
            width={"100%"}
            height={"auto"}
          />
        </Tile>
      </Grid>
      <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
        <Tile
          sx={{ justifyContent: "left", alignItems: "left", textAlign: "left" }}
        >
          <Typography
            sx={{ fontSize: { sm: "20px", lg: "20px" }, padding: "10px" }}
          >
            On the front-end, we utilized the power of React to create a
            seamless and dynamic user experience. Leveraging TypeScript, we
            ensured a robust and type-safe codebase that enhances both
            reliability and maintainability. Vite, our chosen build tool, not
            only optimizes development processes but also significantly improves
            overall application performance. For captivating and insightful
            charts, we rely on the Recharts React graph library — a powerful
            tool that adds depth and clarity to the rich data we provide. Built
            on the sturdy foundation of D3, Recharts enhances our ability to
            create stunning visuals and brings versatility to our charting
            capabilities. To ensure a cohesive and polished user interface, we
            turned to Material-UI (MUI), a cutting-edge React component library.
            MUI has been instrumental in achieving an overall uniform and
            professional look, offering a rich set of components that harmonize
            with our design principles. With this well-rounded front-end tech
            stack, we empower you with an immersive soccer data experience that
            not only captivates but also maintains a sleek and professional
            aesthetic.
          </Typography>
        </Tile>
      </Grid>
      <Grid item xs={12} sm={12} md={12} lg={3} width="100%">
        <Tile>
          <DisplayImage
            location={""}
            src={radarPhoto}
            alt={"Team Page"}
            width={"100%"}
            height={"auto"}
          />
        </Tile>
      </Grid>
      <Grid item xs={12} sm={12} md={12} lg={9} width="100%">
        <Tile>
          <DisplayImage
            location={""}
            src={goalsPhoto}
            alt={"Team Page"}
            width={"90%"}
            height={"auto"}
          />
        </Tile>
      </Grid>
      <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
        <Tile
          sx={{ justifyContent: "left", alignItems: "left", textAlign: "left" }}
        >
          <Typography
            sx={{ fontSize: { sm: "20px", lg: "20px" }, padding: "10px" }}
          >
            The backend of the MapEleven project, architected entirely in Go,
            interacts with API-FOOTBALL to fetch external data. Dedicated
            controllers per endpoint (such as team, standings, league, player)
            facilitate data procurement, after which the data is passed to their
            respective models. Employing Ent - an Object-Relational Mapping
            library for Go - these models orchestrate data insertion or updates
            within the PostgreSQL database, ensuring efficient data handling.
          </Typography>
        </Tile>
      </Grid>
      <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
        <Tile>
          <DisplayImage
            location={""}
            src={erdPhoto}
            alt={"Team Page"}
            width={"100%"}
            height={"auto"}
          />
        </Tile>
      </Grid>
      <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
        <Tile
          sx={{ justifyContent: "left", alignItems: "left", textAlign: "left" }}
        >
          <Typography
            sx={{ fontSize: { sm: "20px", lg: "20px" }, padding: "10px" }}
          >
            On the server-side operations, the backend utilizes the Go Fiber web
            framework to construct routes (endpoints), effectively responding to
            client-side requests for specific data (fixtures, team, player).
            These routes call serializers, which leverage Ent in managing
            uniform and optimized data delivery back to the frontend. These
            routes invoke their corresponding serializers, which utilize Ent to
            efficiently and consistently retrieve and package data from the
            database for delivery to the frontend. In doing so, Ent's
            capabilities are harnessed to perform complex data pulls from
            multiple tables in a single query. This operation showcases the
            power and flexibility of MapEleven, exemplifying its robust,
            streamlined, and dataintensive backend process.
          </Typography>
        </Tile>
      </Grid>
    </Grid>
  );
}

export default About;
