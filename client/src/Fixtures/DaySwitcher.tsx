import { useState } from 'react';
import { Box, Button, Typography, useMediaQuery } from '@mui/material';
import ChevronLeftIcon from '@mui/icons-material/ChevronLeft';
import ChevronRightIcon from '@mui/icons-material/ChevronRight';

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
        border: '1.5px solid var(--dark2)',
        borderRadius: '12px',
    },
    dateText: {
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
    const isSmallerThanLG = useMediaQuery('(max-width: 1260px)');

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

    const options: Intl.DateTimeFormatOptions = {
        weekday: 'long',
        year: 'numeric',
        month: 'long',
        day: 'numeric',
    };

    return (
        <Box sx={styles.root}>
            <Button sx={styles.arrowButton} onClick={handleBackwardClick}>
                <ChevronLeftIcon />
            </Button>
            <Typography
                variant='h6'
                component='h2'
                gutterBottom
                sx={{ ...styles.dateText, fontSize: !isSmallerThanLG ? '20px' : '16px' }}
            >
                {currentDate.toDateString() === new Date().toDateString()
                    ? 'Today'
                    : currentDate.toLocaleDateString('en-US', options)}
            </Typography>
            <Button sx={styles.arrowButton} onClick={handleForwardClick}>
                <ChevronRightIcon />
            </Button>
        </Box>
    );
};

export default DaySwitcher;
