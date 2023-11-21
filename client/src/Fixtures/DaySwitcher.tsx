import React, { useState } from "react";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";
import { useMediaQuery } from "@mui/material";
import ChevronLeftIcon from "@mui/icons-material/ChevronLeft";
import ChevronRightIcon from "@mui/icons-material/ChevronRight";
import Calendar from 'react-calendar';
import './Calendar.css';

import Dialog from "@mui/material/Dialog";
import DialogTitle from "@mui/material/DialogTitle";
import DialogContent from "@mui/material/DialogContent";
import DialogActions from "@mui/material/DialogActions";


type DaySwitcherProps = {
    onDateChange: (newDate: Date) => void;
};

const styles = {
    root: {
        display: "flex",
        alignItems: "center",
        justifyContent: "space-between",
        width: "100%",
        marginBottom: "16px",
        border: "1.5px solid var(--dark2)",
        borderRadius: "12px",
    },
    dateText: {
        margin: "0 16px",
        cursor: "pointer",
    },
    arrowButton: {
        color: "white",
        borderRadius: "10px",
        width: "20x",
        height: "40px",
        fontSize: "20px",
        transition: "background-color 0.3s",
        "&:hover": {
            backgroundColor: "var(--focus)",
        },
    },
};

const DaySwitcher: React.FC<DaySwitcherProps> = ({ onDateChange }) => {
    const [currentDate, setCurrentDate] = useState(new Date());
    const [isCalendarOpen, setCalendarOpen] = useState(false);
    const [isDatePickerOpen, setDatePickerOpen] = useState(false);
    const isSmallerThanLG = useMediaQuery("(max-width: 1260px)");

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

    const handleDateTextClick = () => {
        setDatePickerOpen(!isDatePickerOpen);
    };

    const handleCalendarClick = () => {
        setDatePickerOpen(false);
    };

    const options: Intl.DateTimeFormatOptions = {
        weekday: "long",
        year: "numeric",
        month: "long",
        day: "numeric",
    };

    const minDate = new Date();
    const maxDate = new Date();
    maxDate.setDate(maxDate.getDate() + 21);
    minDate.setDate(minDate.getDate() - 21);

    const isDateOutOfRange = (date) => date < minDate || date > maxDate;

    const tileClassName = ({ date, view }) => {
        if (view === 'month' && isDateOutOfRange(date)) {
            return 'out-of-range';
        }
        return null;
    };

    return (
        <Box sx={styles.root}>
            <Button sx={styles.arrowButton} onClick={handleBackwardClick}>
                <ChevronLeftIcon />
            </Button>
            <Typography
                variant="h6"
                component="h2"
                gutterBottom
                sx={{
                    ...styles.dateText,
                    fontSize: !isSmallerThanLG ? "20px" : "16px",
                }}
                onClick={handleDateTextClick}
            >
                {currentDate.toDateString() === new Date().toDateString()
                    ? "Today"
                    : currentDate.toLocaleDateString("en-US", options)}
            </Typography>
            <Button sx={styles.arrowButton} onClick={handleForwardClick}>
                <ChevronRightIcon />
            </Button>
            <Dialog
                open={isDatePickerOpen}
                onClose={() => setDatePickerOpen(false)}
                maxWidth="sm"
                fullWidth>
                <DialogTitle>Select a Date</DialogTitle>
                <DialogContent>
                    <Calendar
                        onChange={(date) => {
                            setCurrentDate(date as Date);
                            onDateChange(date as Date);
                            setCalendarOpen(false);
                        }}
                        value={currentDate}
                        minDate={minDate}
                        maxDate={maxDate}
                        tileClassName={tileClassName}
                        onClickDay={handleCalendarClick}
                    />
                </DialogContent>
                <DialogActions>
                    <Button onClick={() => setDatePickerOpen(false)}>Cancel</Button>
                </DialogActions>
            </Dialog>
        </Box>
    );
};

export default DaySwitcher;