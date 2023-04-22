import { useState } from 'react';
import { Box, Button, Typography } from '@mui/material';

type DaySwitcherProps = {
    onDateChange: (newDate: Date) => void;
};

const styles = {
    root: {
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'space-between',
        width: '100%',
        marginBottom: '16px',
        marginTop: '16px',
        backgroundColor: 'var(--dark1)',
        borderRadius: '12px',
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

const DaySwitcher: React.FC<DaySwitcherProps> = ({ onDateChange }) => {
    const [currentDate, setCurrentDate] = useState(new Date());

    const handleBackwardClick = () => {
        const newDate = new Date(currentDate.getTime() - 86400000);
        setCurrentDate(newDate);
        onDateChange(newDate);
    };

    const handleForwardClick = () => {
        const newDate = new Date(currentDate.getTime() + 86400000);
        setCurrentDate(newDate);
        onDateChange(newDate);
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
