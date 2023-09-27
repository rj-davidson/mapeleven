import * as React from 'react';
import {Box} from '@mui/material';

type PageBodyProps = {
    children: React.ReactNode;
};

const PageBody: React.FC<PageBodyProps> = ({children}) => {
    return (
        <Box
            sx={{
                display: 'flex',
                flexDirection: 'column',
                alignItems: 'center',
                padding: '1rem',
                margin: '0 auto',
                width: '100%',
                maxWidth: 1440,
                gap: 24,
                paddingTop: '108px',
            }}
        >
            {children}
        </Box>
    );
};

export default PageBody;
