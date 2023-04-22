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
    },
    dateText: {
        fontSize: '24px',
        margin: '0 16px',
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
            <Button onClick={handleBackwardClick}>{'<'}</Button>
            <Typography sx={styles.dateText}>{currentDate.toDateString()}</Typography>
            <Button onClick={handleForwardClick}>{'>'}</Button>
        </Box>
    );
};

export default DaySwitcher;
