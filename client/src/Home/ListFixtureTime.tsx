import * as React from "react";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";

interface ListFixtureTimeProps {
  status: string | undefined;
  elapsed: number | undefined;
}

const ListFixtureTime: React.FC<ListFixtureTimeProps> = ({
  status,
  elapsed,
}) => {
  let content = "";

  if (status === "Match Finished") {
    content = "FT";
  } else if (elapsed) {
    content = `${elapsed}'`;
  }

  return (
    <Box sx={{ width: "5%" }}>
      <Typography
        variant="h6"
        component="h2"
        sx={{
          fontSize: "16px",
          color: status !== "Match Finished" ? "var(--green)" : "inherit",
        }}
      >
        {content}
      </Typography>
    </Box>
  );
};

export default ListFixtureTime;
