import Footer from "./Footer";
import {ReactNode} from "react";

type LayoutProps = {
    children: ReactNode;
};

const Layout = ({ children }: LayoutProps) => {
    return (
        <>
            <main>{children}</main>
            <Footer companyName={'mapeleven'} year={new Date().getFullYear()} />
        </>
    );
};

export default Layout;
