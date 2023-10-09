import * as React from 'react';
import { useEffect, useState } from 'react';
import { APIScoreboardDate } from '../Models/api_types';
import { Box, Grid, Skeleton, Typography } from '@mui/material';
import { Tile } from '../Util/TileTS';
import DaySwitcher from '../Fixtures/DaySwitcher';
import ListFixtureTable from './ListFixtureTable';

const url = import.meta.env.VITE_API_URL;

interface Fixtures {
    [key: string]: APIScoreboardDate;
}

export default function ListFixtures(): JSX.Element {
    const [fixtures, setFixtures] = useState<Fixtures>(null);
    const [loading, setLoading] = useState<boolean>(true);
    const [selectedDate, setSelectedDate] = useState(new Date().toISOString().slice(0, 10));

    function convertDate(dateString) {
        // Split the input string by "-"
        const dateParts = dateString.split('-');

        const day = parseInt(dateParts[1]);
        const month = parseInt(dateParts[0]) - 1; // Subtract 1 from the month (JavaScript months are 0-indexed)
        const year = parseInt(dateParts[2]);

        return new Date(year, month, day);
    }

    useEffect(() => {
        fetch(`${url}/scoreboard`)
            .then(response => response.json())
            .then(jsonData => {
                const newFixtures: Fixtures = {};

                jsonData.date.forEach((dateObj: APIScoreboardDate) => {
                    const convertedDate = convertDate(dateObj.date);
                    newFixtures[convertedDate.toISOString().slice(0, 10)] = dateObj;
                });

                setFixtures(newFixtures);
                setLoading(false);
            })
            .catch(error => console.error(error));
    }, []);

    const skeletonCards: JSX.Element[] = [];

    for (let i = 0; i < 10; i++) {
        skeletonCards.push(
            <Grid item xs={12} sm={12} md={12} lg={12} key={i}>
                <Skeleton
                    variant='rectangular'
                    height='52px'
                    width='100%'
                    sx={{ background: 'var(--dark1)', borderRadius: '12px' }}
                />
            </Grid>,
        );
    }

    // Function to retrieve the selected fixture
    const getSelectedFixture = () => {
        if (fixtures && fixtures[selectedDate]) {
            return fixtures[selectedDate];
        }
        return null; // Return null if no fixture is found for the selected date
    };

    const selectedFixture = getSelectedFixture();

    const handleDateChange = (newDate: Date) => {
        setSelectedDate(newDate.toISOString().slice(0, 10));
    };

    return (
        <Tile sx={{ flexDirection: 'column' }}>
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
                        <Typography variant='h6' component='h2' sx={{ fontSize: '16px', textAlign: 'center' }}>
                            No fixtures
                        </Typography>
                    )}
                </Box>
            )}
        </Tile>
    );
}
