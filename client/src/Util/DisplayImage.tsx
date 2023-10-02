import * as React from 'react';
import { CSSProperties } from 'react';

const url = import.meta.env.VITE_API_URL;

interface DisplayImageProps {
    src: string;
    alt: string;
    location?: string;
    width?: string;
    height?: string;
    sx?: CSSProperties;
}

const DisplayImage: React.FC<DisplayImageProps> = ({
    src,
    alt,
    location = `${url}/`,
    width = '32px',
    height = '32px',
    sx,
}) => {
    return <img src={location + src} alt={alt} width={width} height={height} style={sx} />;
};

export default DisplayImage;
