import React from 'react';

export const Flex = ({children, style}) => {
    const flexStyle = {
        display: 'flex',
        ...style,
    };

    return (
        <div style={flexStyle}>
            {children}
        </div>
    );
};