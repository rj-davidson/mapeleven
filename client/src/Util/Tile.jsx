import React from 'react';

export const Tile = ({children, style}) => {
    const tileStyle = {
        display: 'flex',
        overflow: 'hidden',
        justifyContent: 'center',
        backgroundColor: 'var(--dark2)',
        borderRadius: '12px',
        padding: '12px',
        paddingTop: '10px',
        width: '100%',
        border: '5px solid var(--dark1)',
        color: 'var(--light2)',
        ...style,
    };

    return (
        <div style={tileStyle}>
            {children}
        </div>
    );
};