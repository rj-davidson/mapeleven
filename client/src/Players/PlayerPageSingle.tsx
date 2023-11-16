import * as React from "react";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import Box from "@mui/material/Box";
import Grid from "@mui/material/Grid";
import Skeleton from "@mui/material/Skeleton";
import PlayerPageSingleColumn from "./PlayerPageSingleColumn";
import { Tile } from "../Util/Tile";
import PlayerPageSingleMain from "./PlayerPageSingleMain";

const url = import.meta.env.VITE_API_URL;

function PlayerPageSingle() {
  const { slug } = useParams();

  const [name, setName] = useState<string>("");
  const [firstName, setFirstName] = useState<string>("");
  const [lastName, setLastName] = useState<string>("");
  const [photo, setPhoto] = useState<string>("");
  const [age, setAge] = useState<number>(0);
  const [height, setHeight] = useState<string>("");
  const [weight, setWeight] = useState<string>("");
  const [injured, setInjured] = useState<boolean>(false);
  const [loading, setLoading] = useState<boolean>(true);

  useEffect(() => {
    fetch(`${url}/players/${slug}`)
      .then((response) => response.json())
      .then((jsonData) => {
        setName(jsonData.name);
        setFirstName(jsonData.firstname);
        setLastName(jsonData.lastname);
        setPhoto(jsonData.photo);
        setAge(jsonData.age);
        setHeight(jsonData.height);
        setWeight(jsonData.weight);
        setInjured(jsonData.injured);
        setLoading(false);
      })
      .catch((error) => console.error(error));
  }, [slug]);

  return loading ? (
    <Grid container spacing={2} direction="column">
      <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
        <Skeleton
          variant="rectangular"
          height="150px"
          width="100%"
          sx={{ background: "gray", borderRadius: "12px" }}
        />
      </Grid>
      <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
        <Skeleton
          variant="rectangular"
          height="600px"
          width="100%"
          sx={{ background: "gray", borderRadius: "12px" }}
        />
      </Grid>
    </Grid>
  ) : (
    <Grid container spacing={2}>
      <Grid item xs={12} sm={12} md={12} lg={3} width="100%">
        <PlayerPageSingleColumn
          slug={slug && slug !== "" ? slug : ""}
          name={name}
          firstName={firstName}
          lastName={lastName}
          photo={photo}
          age={age}
          height={height}
          weight={weight}
          injured={injured}
        />
      </Grid>
      <Grid item xs={12} sm={12} md={12} lg={9} width="100%">
        <PlayerPageSingleMain name={name} />
      </Grid>
    </Grid>
  );
}

export default PlayerPageSingle;
