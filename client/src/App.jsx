import React from 'react';
import './App.css';
import dark from './Theme/MapTheme.tsx';
import Layout from './Layout/Layout.tsx';
import Players from './Players/Players.jsx';
import TeamPageAll from './Teams/TeamPageAll.jsx';
import TeamPageSingle from './Teams/TeamPageSingle';
import Leagues from './Leagues/Leagues.jsx';
import LeaguePage from './Leagues/LeaguePage.jsx';
import About from './About/About.tsx';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import { ThemeProvider } from '@mui/material';
import HomePage from './Home/HomePage';

function App() {
    return (
        <ThemeProvider theme={dark}>
            <div className='App'>
                <Router>
                    <Layout>
                        <Routes>
                            <Route path='/fixtures/:slug' element='' />
                            <Route path='/players/:slug' element={<Players />} />
                            <Route path='/players' element={<Players />} />
                            <Route path='/teams/:slug' element={<TeamPageSingle />} />
                            <Route path='/teams' element={<TeamPageAll />} />
                            <Route path='/leagues/:slug' element={<LeaguePage />} />
                            <Route path='/leagues' element={<Leagues />} />
                            <Route path='/about' element={<About />} />
                            <Route path='/' element={<HomePage />} />
                        </Routes>
                    </Layout>
                </Router>
            </div>
        </ThemeProvider>
    );
}

export default App;
