import Footer from "./Footer";
import {ReactNode} from "react";
import NavigationBar from "./NavigationBar";

type LayoutProps = {
    children: ReactNode;
};

const Layout = ({ children }: LayoutProps) => {
    return (
        <>
            <NavigationBar />
            <main>{children}</main>
            <Footer companyName={'mapeleven'} year={new Date().getFullYear()} />
        </>
    );
};

export default Layout;
