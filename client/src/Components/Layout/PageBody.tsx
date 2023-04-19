import * as React from 'react'
import './PageBody.css';

type PageBodyProps = {
    children: React.ReactNode;
};

const PageBody: React.FC<PageBodyProps> = ({ children }) => {
    return <div className="page-body">{children}</div>;
};

export default PageBody;
