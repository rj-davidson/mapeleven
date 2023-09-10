import React from 'react';
import './App.css';
import FixturesPage from './Components/Pages/Fixtures/FixturesPage.tsx';
import Players from './Components/Pages/Players/Players.jsx';
import Clubs from './Components/Pages/Clubs/Clubs.jsx';
import Leagues from './Components/Pages/Leagues/Leagues.jsx';
import About from './Components/Pages/About/About.tsx';
import {BrowserRouter as Router, Routes, Route} from 'react-router-dom';
import {ThemeProvider} from '@mui/material';
import dark from './Components/Theme/MapTheme.tsx';
import Layout from './Components/Layout/Layout.tsx';

function App() {
    return (
        <ThemeProvider theme={dark}>
            <div className='App'>
                <Router>
                    <Layout>
                        <Routes>
                            <Route path='/fixtures/:slug' element={<FixturesPage/>}/>
                            <Route path='/players/:slug' element={<Players/>}/>
                            <Route path='/clubs/:slug' element={<Clubs/>}/>
                            <Route path='/leagues/:slug' element={<Leagues/>}/>
                            <Route path='/about' element={<About/>}/>
                            <Route path='/' element={<FixturesPage/>}/>
                        </Routes>
                    </Layout>
                </Router>
            </div>
        </ThemeProvider>
    );
}

export default App;
