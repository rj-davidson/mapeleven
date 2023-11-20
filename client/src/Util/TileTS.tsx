import { ReactNode, CSSProperties } from "react";

interface TileProps {
  children: ReactNode;
  sx?: CSSProperties;
}

// eslint-disable-next-line react/prop-types
export const Tile: React.FC<TileProps> = ({ children, sx }) => {
  const tileStyle: CSSProperties = {
    display: "flex",
    overflow: "hidden",
    justifyContent: "center",
    backgroundColor: "var(--dark2)",
    borderRadius: "12px",
    padding: "12px",
    paddingTop: "10px",
    width: "100%",
    color: "var(--light2)",
    ...sx,
  };

  return <div style={tileStyle}>{children}</div>;
};
