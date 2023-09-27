import { ReactNode } from 'react';
import Footer from './Footer';
import NavBar from './NavBar';
import PageBody from './PageBody';
import { Grid } from '@mui/material';

type LayoutProps = {
    children: ReactNode;
};

const Layout = ({ children }: LayoutProps) => {
    return (
        <Grid container direction='column'>
            <Grid item>
                <NavBar />
            </Grid>

            <Grid item>
                <PageBody>{children}</PageBody>
            </Grid>

            <Grid item>
                <Footer companyName={'mapeleven'} year={new Date().getFullYear()} />
            </Grid>
        </Grid>
    );
};

export default Layout;
