import React from "react";
import "./App.css";
import dark from "./Theme/MapTheme.tsx";
import Layout from "./Layout/Layout.tsx";
import Players from "./Players/Players.jsx";
import TeamPageAll from "./Teams/TeamPageAll.jsx";
import TeamPageSingle from "./Teams/TeamPageSingle";
import LeaguesPageAll from "./Leagues/LeaguePageAll.jsx";
import LeaguePage from "./Leagues/LeaguePage.jsx";
import About from "./About/About.tsx";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { ThemeProvider } from "@mui/material";
import HomePage from "./Home/HomePage";
import PlayerPageSingle from "./Players/PlayerPageSingle";
import FixturePageSingle from "./Fixtures/FixturePageSingle";
import FixturePageSingleDemo from "./Fixtures/FixturePageSingleDemo";

function App() {
  return (
    <ThemeProvider theme={dark}>
      <div className="App">
        <Router>
          <Layout>
            <Routes>
              <Route path="/fixtures/:slug" element={<FixturePageSingle />} />
              <Route
                path="/fixtures-demo/:slug"
                element={<FixturePageSingleDemo />}
              />
              <Route path="/players/:slug" element={<PlayerPageSingle />} />
              <Route path="/players" element={<Players />} />
              <Route path="/teams/:slug" element={<TeamPageSingle />} />
              <Route path="/teams" element={<TeamPageAll />} />
              <Route path="/leagues/:slug" element={<LeaguePage />} />
              <Route path="/leagues" element={<LeaguesPageAll />} />
              <Route path="/about" element={<About />} />
              <Route path="/" element={<HomePage />} />
            </Routes>
          </Layout>
        </Router>
      </div>
    </ThemeProvider>
  );
}

export default App;
