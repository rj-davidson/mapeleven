import * as React from 'react';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import IconButton from '@mui/material/IconButton';
import Typography from '@mui/material/Typography';
import MenuIcon from '@mui/icons-material/Menu';
import { Link } from 'react-router-dom';
import SearchBar from './SearchBar.jsx';
import { useState } from 'react';
import { Drawer, List, ListItem, ListItemText } from '@mui/material';
import { FormatListNumbered, Groups, Home, Person } from '@mui/icons-material';

export default function NavBar() {
    const [drawerOpen, setDrawerOpen] = useState(false);

    const handleDrawerOpen = () => {
        setDrawerOpen(true);
    };

    const handleDrawerClose = () => {
        setDrawerOpen(false);
    };

    return (
        <Box sx={{ flexGrow: 1 }}>
            <AppBar
                position='fixed'
                sx={{
                    backgroundColor: 'rgba(0,0,0,0.3)',
                    color: 'var(--light0)',
                    width: '100%',
                    backdropFilter: 'blur(10px)',
                }}
            >
                <Toolbar>
                    <IconButton
                        size='large'
                        edge='start'
                        color='inherit'
                        aria-label='open drawer'
                        onClick={handleDrawerOpen}
                        sx={{ mr: 2 }}
                    >
                        <MenuIcon />
                    </IconButton>
                    <Drawer anchor='left' open={drawerOpen} onClose={handleDrawerClose}>
                        <List sx={{ width: '200px' }}>
                            <Link to='/'>
                                <ListItem onClick={handleDrawerClose}>
                                    <Home sx={{ marginRight: '8px' }} />
                                    <ListItemText primary='Home' primaryTypographyProps={{ fontSize: '24px' }} />
                                </ListItem>
                            </Link>
                            <Link to='/players'>
                                <ListItem onClick={handleDrawerClose}>
                                    <Person sx={{ marginRight: '8px' }} />
                                    <ListItemText primary='Players' primaryTypographyProps={{ fontSize: '24px' }} />
                                </ListItem>
                            </Link>
                            <Link to='/teams'>
                                <ListItem onClick={handleDrawerClose}>
                                    <Groups sx={{ marginRight: '8px' }} />
                                    <ListItemText primary='Teams' primaryTypographyProps={{ fontSize: '24px' }} />
                                </ListItem>
                            </Link>
                            <Link to='/leagues'>
                                <ListItem onClick={handleDrawerClose}>
                                    <FormatListNumbered sx={{ marginRight: '8px' }} />
                                    <ListItemText primary='Leagues' primaryTypographyProps={{ fontSize: '24px' }} />
                                </ListItem>
                            </Link>
                        </List>
                    </Drawer>
                    <Typography
                        variant='body'
                        noWrap
                        component='div'
                        sx={{ flexGrow: 1, display: { xs: 'none', md: 'block' }, fontSize: '3rem', fontWeight: '500' }}
                    >
                        <Link to='/'>
                            <div style={{ display: 'inline', color: 'var(--light0)', fontWeight: '200' }}>map</div>
                            <div style={{ display: 'inline', color: 'var(--light2)' }}>eleven</div>
                        </Link>
                    </Typography>
                    <Box sx={{ width: '100%', maxWidth: { xs: 'none', md: 500 } }}>
                        <SearchBar />
                    </Box>
                </Toolbar>
            </AppBar>
        </Box>
    );
}
