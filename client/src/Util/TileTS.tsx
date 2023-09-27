import { ReactNode, CSSProperties } from 'react';

interface TileProps {
    children: ReactNode;
    sx?: CSSProperties;
}

export const Tile: React.FC<TileProps> = ({ children, sx }) => {
    const tileStyle: CSSProperties = {
        display: 'flex',
        overflow: 'hidden',
        justifyContent: 'center',
        backgroundColor: 'var(--dark2)',
        borderRadius: '12px',
        padding: '24px',
        width: '100%',
        ...sx,
    };

    return <div style={tileStyle}>{children}</div>;
};
