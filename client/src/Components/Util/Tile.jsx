import React from 'react';

export const Tile = ({children, style}) => {
    const tileStyle = {
        display: 'flex',
        overflow: 'hidden',
        justifyContent: 'center',
        backgroundColor: 'var(--dark2)',
        borderRadius: '12px',
        padding: '12px',
        ...style,
    };

    return (
        <div style={tileStyle}>
            {children}
        </div>
    );
};