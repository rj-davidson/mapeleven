import * as React from 'react';
import { Box, Grid } from '@mui/material';

type FooterProps = {
    companyName: string;
    year: number;
};

const Footer: React.FC<FooterProps> = ({ companyName, year }) => {
    return (
        <footer style={{ textAlign: 'center', padding: '12px' }}>
            {companyName} &copy; {year}
        </footer>
    );
};

export default Footer;
