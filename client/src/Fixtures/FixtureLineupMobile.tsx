import * as React from "react";
import { Tile } from "../Util/TileTS";
import { Flex } from "../Util/Flex";
import { useEffect } from "react";
import { Link } from "react-router-dom";
import fieldMobile from "./fieldMobile.svg";
import LineupPlayer from "./LineupPlayer";

const url = import.meta.env.VITE_API_URL;

interface FixtureLineupProps {
  homePlayers: any[];
  awayPlayers: any[];
}

const FixtureLineup: React.FC<FixtureLineupProps> = ({
  homePlayers,
  awayPlayers,
}) => {
  const [homePhotos, setHomePhotos] = React.useState<{
    [slug: string]: string;
  }>({});
  const [awayPhotos, setAwayPhotos] = React.useState<{
    [slug: string]: string;
  }>({});

  useEffect(() => {
    homePlayers.forEach((player) => {
      const { slug } = player.player;
      fetch(`${url}/players/${slug}`)
        .then((response) => response.json())
        .then((jsonData) => {
          setHomePhotos((prevState) => ({
            ...prevState,
            [slug]: jsonData.photo,
          }));
        })
        .catch((error) => console.error(error));
    });
    awayPlayers.forEach((player) => {
      const { slug } = player.player;
      fetch(`${url}/players/${slug}`)
        .then((response) => response.json())
        .then((jsonData) => {
          setAwayPhotos((prevState) => ({
            ...prevState,
            [slug]: jsonData.photo,
          }));
        })
        .catch((error) => console.error(error));
    });
  }, []);

  const homeColumns = {};

  homePlayers.forEach((player) => {
    const { x, y, slug, name, number } = player.player;

    if (x === 0 || y === 0) {
      // Skip (0,0) positions (subs)
      return;
    }

    if (!homeColumns[x]) {
      homeColumns[x] = [];
    }

    homeColumns[x].push({ slug, name, number, y });
  });

  const awayColumns = {};
  awayPlayers.forEach((player) => {
    const { x, y, slug, name, number } = player.player;

    if (x === 0 || y === 0) {
      // Skip (0,0) positions (subs)
      return;
    }

    if (!awayColumns[x]) {
      awayColumns[x] = [];
    }

    awayColumns[x].push({ slug, name, number, y });
  });

  return (
    <Tile
      sx={{
        alignItems: "center",
        backgroundImage: `url(${fieldMobile})`,
        backgroundSize: "cover",
        backgroundPosition: "center",
        flexDirection: "column",
      }}
    >
      <Flex
        style={{
          flexDirection: "column",
          width: "50%",
          alignItems: "center",
          justifyContent: "space-around",
          padding: "20px",
          gap: "20px",
        }}
      >
        {Object.keys(homeColumns).map((x) => (
          <Flex
            key={x}
            style={{
              flexDirection: "row-reverse",
              justifyContent: "space-around",
              alignItems: "flex-start",
              gap: "8%",
            }}
          >
            {homeColumns[x].map((player, index) => (
              <Link to={`/players/${player.slug}`} key={player.slug}>
                <LineupPlayer
                  player={player}
                  index={index}
                  photo={homePhotos[player.slug]}
                  home={true}
                  mobile={true}
                  number={player.number}
                />
              </Link>
            ))}
          </Flex>
        ))}
      </Flex>
      <Flex
        style={{
          flexDirection: "column-reverse",
          width: "50%",
          alignItems: "center",
          justifyContent: "space-around",
          padding: "20px",
          gap: "20px",
        }}
      >
        {Object.keys(awayColumns).map((x) => (
          <Flex
            key={x}
            style={{
              flexDirection: "row",
              justifyContent: "space-around",
              alignItems: "flex-start",
              gap: "8%",
            }}
          >
            {awayColumns[x].map((player, index) => (
              <Link to={`/players/${player.slug}`} key={player.slug}>
                <LineupPlayer
                  player={player}
                  index={index}
                  photo={awayPhotos[player.slug]}
                  home={false}
                  mobile={true}
                  number={player.number}
                />
              </Link>
            ))}
          </Flex>
        ))}
      </Flex>
    </Tile>
  );
};

export default FixtureLineup;
