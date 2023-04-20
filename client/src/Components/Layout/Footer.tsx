import * as React from 'react';

type FooterProps = {
    companyName: string;
    year: number;
};

const Footer: React.FC<FooterProps> = ({ companyName, year }) => {
    return (
        <footer>
            <div>{companyName} &copy; {year}</div>
        </footer>
    );
};

export default Footer;
