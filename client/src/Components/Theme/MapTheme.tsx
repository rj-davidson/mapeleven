import {CommonColors, createTheme} from '@mui/material/styles';

interface CustomColors extends CommonColors {
    dark0: string;
    dark1: string;
    dark2: string;
    dark3: string;
    light0: string;
    light1: string;
    light2: string;
    focus: string;
    red: string;
    orange: string;
    yellow: string;
    green: string;
    blue: string;
    indigo: string;
}

const customTheme = createTheme({
    palette: {
        primary: {
            main: '#3854FC',
        },
        secondary: {
            main: '#3B4252',
        },
        text: {
            primary: '#D8DEE9',
            secondary: '#E5E9F0',
        },
        common: {
            dark0: '#2E3440',
            dark1: '#3B4252',
            dark2: '#434C5E',
            dark3: '#4C566A',
            light0: '#D8DEE9',
            light1: '#E5E9F0',
            light2: '#ECEFF4',
            focus: '#3854FC',
            red: '#BF616A',
            orange: '#D08770',
            yellow: '#EBCB8B',
            green: '#A3BE8C',
            blue: '#5E81AC',
            indigo: '#B48EAD',
        } as CustomColors,
    },
    typography: {
        fontFamily: 'Poppins'
    },
    components: {
        MuiCard: {
            styleOverrides: {
                root: {
                    borderRadius: '10px',
                    boxShadow: 'none',
                    width: '100%',
                    backgroundColor: 'var(--dark1)',
                    border: '3px solid var(--dark1)',
                    margin: '0',
                    padding: '0',
                    '&:hover': {
                        border: '3px solid var(--focus)',
                        backgroundColor: 'var(--dark2)',
                    },
                },
            },
        },
        MuiAutocomplete: {
            styleOverrides: {
                root: {
                    '& .MuiOutlinedInput-root': {
                        backgroundColor: 'var(--dark2)',
                        '& fieldset': {
                            border: 'none', // Remove the border
                        },
                    },
                    '& .MuiInputLabel-root.Mui-focused': {
                        color: 'var(--light2)', // Change to your desired label color
                    },
                },
            },
        },
    },
});

export default customTheme;
