import * as React from "react";
import { useEffect, useState } from "react";
import { APIScoreboardDate } from "../Models/api_types";
import Box from "@mui/material/Box";
import Grid from "@mui/material/Grid";
import Skeleton from "@mui/material/Skeleton";
import Typography from "@mui/material/Typography";
import { Tile } from "../Util/TileTS";
import DaySwitcher from "../Fixtures/DaySwitcher";
import ListFixtureTable from "./ListFixtureTable";

const url = import.meta.env.VITE_API_URL;

interface Fixtures {
  [key: string]: APIScoreboardDate;
}

export default function ListFixtures(): JSX.Element {
  const [fixtures, setFixtures] = useState<Fixtures>();
  const [loading, setLoading] = useState<boolean>(true);
  const [selectedDate, setSelectedDate] = useState<string>();

  useEffect(() => {
    const timeZone = Intl.DateTimeFormat().resolvedOptions().timeZone;
    fetch(`${url}/scoreboard?timezone=${encodeURIComponent(timeZone)}`)
      .then((response) => response.json())
      .then((jsonData) => {
        const newFixtures: Fixtures = {};

        jsonData.date.forEach((dateObj: APIScoreboardDate) => {
          const convertedDate = new Date(dateObj.date);
          newFixtures[convertedDate.toISOString().slice(0, 10)] = dateObj;
        });

        setFixtures(newFixtures);
        const date = new Date();
        const localDate = new Date(
          date.getTime() - date.getTimezoneOffset() * 60000
        )
          .toISOString()
          .slice(0, 10);

        setSelectedDate(localDate);

        setLoading(false);
      })
      .catch();
  }, []);

  const skeletonCards: JSX.Element[] = [];

  for (let i = 0; i < 10; i++) {
    skeletonCards.push(
      <Grid item xs={12} sm={12} md={12} lg={12} key={i}>
        <Skeleton
          variant="rectangular"
          height="52px"
          width="100%"
          sx={{ background: "var(--dark1)", borderRadius: "12px" }}
        />
      </Grid>
    );
  }

  // Function to retrieve the selected fixture
  const getSelectedFixture = () => {
    if (fixtures && selectedDate && fixtures[selectedDate]) {
      return fixtures[selectedDate];
    }
    return null;
  };

  const selectedFixture = getSelectedFixture();

  const handleDateChange = (newDate: Date) => {
    const localDate = new Date(
      newDate.getTime() - newDate.getTimezoneOffset() * 60000
    )
      .toISOString()
      .slice(0, 10);
    setSelectedDate(localDate);
  };

  return (
    <Tile sx={{ flexDirection: "column" }}>
      <DaySwitcher onDateChange={handleDateChange} />
      {loading ? (
        <Grid container spacing={1}>
          {skeletonCards}
        </Grid>
      ) : (
        <Box>
          {selectedFixture ? (
            <ListFixtureTable date={selectedFixture} />
          ) : (
            <Typography
              variant="h6"
              component="h2"
              sx={{ fontSize: "16px", textAlign: "center" }}
              paddingY={4}
            >
              No Fixtures
            </Typography>
          )}
        </Box>
      )}
    </Tile>
  );
}
