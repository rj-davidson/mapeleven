import Footer from './Footer';
import { ReactNode } from 'react';
import NavigationBar from './NavigationBar';
import PageBody from './PageBody';
import {Grid} from "@mui/material";

type LayoutProps = {
    children: ReactNode;
};

const Layout = ({ children }: LayoutProps) => {
    return (
        <Grid container direction="column">
            <NavigationBar />

            <PageBody>{children}</PageBody>

            <Footer companyName={'mapeleven'} year={new Date().getFullYear()} />
        </Grid>
    );
};

export default Layout;
