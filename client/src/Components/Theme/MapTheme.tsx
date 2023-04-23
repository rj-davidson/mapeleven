import { createTheme } from '@mui/material/styles';

const dark = createTheme({
    palette: {
        primary: {
            main: '#D08770',
        },
        secondary: {
            main: '#3B4252',
        },
        text: {
            primary: '#D8DEE9',
            secondary: '#4C566A',
        },
    },
    components: {
        MuiCard: {
            styleOverrides: {
                root: {
                    backgroundColor: 'var(--dark0)',
                    border: '1px solid var(--dark1)',
                    borderRadius: '10px',
                    boxShadow: 'none',
                },
            },
        },
    }
});

export default dark;
