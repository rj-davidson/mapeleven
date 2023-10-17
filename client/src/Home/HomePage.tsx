import * as React from 'react';
import { Box, Grid, useMediaQuery } from '@mui/material';
import DisplayPlayers from '../Players/DisplayPlayersCopy';
import { Tile } from '../Util/TileTS';
import ListTeams from '../Teams/ListTeams';
import ListFixtures from './ListFixtures';

export default function FixturesPage(): JSX.Element {
    const isSmallerThanLG = useMediaQuery('(max-width: 1260px)');

    return (
        <Box>
            {!isSmallerThanLG ? (
                <Grid container spacing={2} direction='row'>
                    <Grid item xs={12} sm={12} md={12} lg={2.5} width='100%'>
                        <Tile>
                            <DisplayPlayers limit={10} skeleton={10} />
                        </Tile>
                    </Grid>

                    <Grid item xs={12} sm={12} md={12} lg={7} width='100%'>
                        <ListFixtures />
                    </Grid>

                    <Grid item xs={12} sm={12} md={12} lg={2.5} width='100%'>
                        <Tile>
                            <ListTeams limit={10} skeleton={10} />
                        </Tile>
                    </Grid>
                </Grid>
            ) : (
                <Grid container spacing={2} direction='row'>
                    <Grid item xs={12} sm={12} md={12} lg={7} width='100%'>
                        <ListFixtures />
                    </Grid>
                    <Grid item xs={12} sm={12} md={12} lg={2.5} width='100%'>
                        <Tile>
                            <DisplayPlayers limit={10} skeleton={10} />
                        </Tile>
                    </Grid>
                    <Grid item xs={12} sm={12} md={12} lg={2.5} width='100%'>
                        <Tile>
                            <ListTeams limit={10} skeleton={10} />
                        </Tile>
                    </Grid>
                </Grid>
            )}
        </Box>
    );
}
