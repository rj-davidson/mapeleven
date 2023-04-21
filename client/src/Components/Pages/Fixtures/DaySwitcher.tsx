import { useState } from 'react';
import { Box, Button, Typography } from '@mui/material';

type DaySwitcherProps = {};

const styles = {
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
    arrowButton: {
        color: 'white',
        borderRadius: '10px',
        width: '20x',
        height: '40px',
        fontSize: '20px',
        transition: 'background-color 0.3s',
        '&:hover': {
            backgroundColor: 'var(--focus)',
        },
    },
};

const DaySwitcher: React.FC<DaySwitcherProps> = () => {
    const [currentDate, setCurrentDate] = useState(new Date());

    const handleBackwardClick = () => {
        setCurrentDate(prevDate => new Date(prevDate.getTime() - 86400000)); // 86400000 = 24 hours * 60 minutes * 60 seconds * 1000 milliseconds
    };

    const handleForwardClick = () => {
        setCurrentDate(prevDate => new Date(prevDate.getTime() + 86400000));
    };

    return (
        <Box sx={styles.root}>
            <Button sx={styles.arrowButton} onClick={handleBackwardClick}>{'<'}</Button>
            <Typography sx={styles.dateText}>{currentDate.toDateString()}</Typography>
            <Button sx={styles.arrowButton} onClick={handleForwardClick}>{'>'}</Button>
        </Box>
    );
};

export default DaySwitcher;
