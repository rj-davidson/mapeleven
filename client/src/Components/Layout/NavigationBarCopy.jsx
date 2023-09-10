import * as React from 'react';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import IconButton from '@mui/material/IconButton';
import Typography from '@mui/material/Typography';
import MenuIcon from '@mui/icons-material/Menu';
import {Link} from "react-router-dom";
import SearchBar from "./SearchBar.jsx"

export default function NavBar() {
    return (
        <Box sx={{flexGrow: 1}}>
            <AppBar position="fixed" sx={{backgroundColor: 'var(--dark1)', color: 'var(--light0)', width: '100%'}}>
                <Toolbar>
                    <IconButton
                        size="large"
                        edge="start"
                        color="inherit"
                        aria-label="open drawer"
                        sx={{mr: 2}}
                    >
                        <MenuIcon/>
                    </IconButton>
                    <Typography
                        variant="body"
                        noWrap
                        component="div"
                        sx={{flexGrow: 1, display: {xs: 'none', md: 'block'}, fontSize: '3rem', fontWeight: '500'}}
                    >
                        <Link to='/'>
                            <div style={{display: "inline", color: "var(--light0)", fontWeight: '200'}}>
                                map
                            </div>
                            <div style={{display: "inline", color: "var(--light2)"}}>
                                eleven
                            </div>
                        </Link>
                    </Typography>
                    <Box sx={{width: '100%', maxWidth: {xs: 'none', md: 500}}}>
                        <SearchBar/>
                    </Box>
                </Toolbar>
            </AppBar>
        </Box>
    );
}
