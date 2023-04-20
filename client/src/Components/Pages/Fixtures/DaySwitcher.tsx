import { useState } from 'react';
import { Box, Button, Typography } from '@material-ui/core';
import { makeStyles } from '@material-ui/core/styles';

type DaySwitcherProps = {};

const useStyles = makeStyles({
    root: {
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'space-between',
        width: '100%',
        marginBottom: '16px',
    },
    dateText: {
        fontSize: '24px',
        margin: '0 16px',
    },
});

const DaySwitcher: React.FC<DaySwitcherProps> = () => {
    const classes = useStyles();
    const [currentDate, setCurrentDate] = useState(new Date());

    const handleBackwardClick = () => {
        const newDate = new Date();
        newDate.setDate(currentDate.getDate() - 1);
        setCurrentDate(newDate);
    };

    const handleForwardClick = () => {
        const newDate = new Date();
        newDate.setDate(currentDate.getDate() + 1);
        setCurrentDate(newDate);
    };

    return (
        <Box className={classes.root}>
            <Button onClick={handleBackwardClick}>&lt;</Button>
            <Typography className={classes.dateText}>{currentDate.toDateString()}</Typography>
            <Button onClick={handleForwardClick}>&gt;</Button>
        </Box>
    );
};

export default DaySwitcher;
