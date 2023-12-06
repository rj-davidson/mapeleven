import * as React from "react";
import Card from "@mui/material/Card";
import CardContent from "@mui/material/CardContent";
import { Flex } from "../Util/Flex.jsx";
import { Link } from "react-router-dom";
import ListFixtureTeam from "./ListFixtureTeam";
import ListFixtureScore from "./ListFixtureScore";
import ListFixtureTime from "./ListFixtureTime";

interface ListFixtureCardProps {
  slug: string;
  status: string | undefined;
  elapsed: number | undefined;
  homeName: string;
  homeSlug: string;
  homeLogo: string;
  homeScore: number | undefined;
  awayName: string;
  awaySlug: string;
  awayLogo: string;
  awayScore: number | undefined;
  time: Date | undefined;
}

const ListFixtureCard: React.FC<ListFixtureCardProps> = ({
  slug,
  status,
  elapsed,
  homeName,
  homeSlug,
  homeLogo,
  homeScore,
  awayName,
  awaySlug,
  awayLogo,
  awayScore,
  time,
}) => {
  return (
    <Link to={`/fixtures/${slug}`}>
      <Card>
        <CardContent sx={{ padding: "8px" }}>
          <Flex
            style={{
              flexDirection: "row",
              gap: "4px",
              justifyContent: "center",
              alignItems: "center",
            }}
          >
            <ListFixtureTime status={status} elapsed={elapsed} />
            <ListFixtureTeam
              slug={homeSlug}
              name={homeName}
              logo={homeLogo}
              home={true}
            />
            <ListFixtureScore
              homeScore={homeScore}
              awayScore={awayScore}
              status={status}
              time={time}
            />
            <ListFixtureTeam
              slug={awaySlug}
              name={awayName}
              logo={awayLogo}
              home={false}
            />
          </Flex>
        </CardContent>
      </Card>
    </Link>
  );
};

export default ListFixtureCard;
