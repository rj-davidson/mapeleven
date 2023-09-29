import React from 'react';

export const Tile = ({children, style}) => {
    const tileStyle = {
        display: 'flex',
        overflow: 'hidden',
        justifyContent: 'center',
        backgroundColor: 'var(--dark2)',
        borderRadius: '12px',
        padding: '24px',
        width: '100%',
        border: '5px solid var(--dark1)',
        ...style,
    };

    return (
        <div style={tileStyle}>
            {children}
        </div>
    );
};