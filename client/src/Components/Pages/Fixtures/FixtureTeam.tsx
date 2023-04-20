import * as React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Box, Typography } from '@material-ui/core';

type FixtureTeamProps = {
    logo: string;
    name: string;
    score: string;
};

const useStyles = makeStyles({
    teamWrapper: {
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'space-between',
        height: 44,
    },
    teamLogo: {
        width: 44,
        height: 44,
    },
    teamName: {
        flexGrow: 1,
        paddingLeft: 10,
    },
    teamScore: {
        fontSize: 20,
        fontWeight: 'bold',
    },
});

const FixtureTeam: React.FC<FixtureTeamProps> = ({ logo, name, score }) => {
    const classes = useStyles();

    return (
        <Box className={classes.teamWrapper}>
            <img src={logo} alt={`${name} logo`} className={classes.teamLogo} />
            <Typography variant='body1' className={classes.teamName}>
                {name}
            </Typography>
            <Typography variant='body1' className={classes.teamScore}>
                {score}
            </Typography>
        </Box>
    );
};

export default FixtureTeam;
