import Footer from './Footer';
import { ReactNode } from 'react';
import NavigationBar from './NavigationBar';
import PageBody from './PageBody';

type LayoutProps = {
    children: ReactNode;
};

const Layout = ({ children }: LayoutProps) => {
    return (
        <>
            <NavigationBar />

            <PageBody>{children}</PageBody>

            <Footer companyName={'mapeleven'} year={new Date().getFullYear()} />
        </>
    );
};

export default Layout;
