import React, {useState} from 'react';
import {Link} from 'react-router-dom';
import {Button} from "./Button.jsx";
import './NavigationBar.css';
import fetch from 'fetch';


function NavigationBar(){
const [click, setClick] = useState(false);
const [button, setButton] = useState(true);

const handleClick = () => setClick(!click);
const closeMenu = () => setClick(false);

const showButton = () => {
    if(window.innerWidth <= 960){
        setButton(false);
    } else {
        setButton(true);
    }
}

window.addEventListener('resize', showButton);


    return(
        <nav className="navigationbar">
            <div className="navigationbar-container">
                <div className='menu-icon' onClick={handleClick}>
                    <i className={click ? 'fas fa-times': 'fas fa-bars'}/>
                </div>
                <ul className={click ? 'nav-menu active' : 'nav-menu'}>
                    <li className='nav-item'>
                        <Link to='/Fixtures' className = 'nav-links' onClick = {closeMenu}> Fixtures </Link>
                    </li>
                    <li className='nav-item'>
                        <Link to='/Players' className = 'nav-links' onClick = {closeMenu}> Players </Link>
                    </li>
                    <li className='nav-item'>
                        <Link to='/Clubs' className = 'nav-links' onClick = {closeMenu}> Clubs </Link>
                    </li>
                    <li className='nav-item'>
                        <Link to='/Competitions' className = 'nav-links' onClick = {closeMenu}> Competitions </Link>
                    </li>
                    <li className='nav-item'>
                        <Link to='/About' className = 'nav-links' onClick = {closeMenu}> About </Link>
                    </li>
                </ul>
                {button && <Button buttonStyle='button--primary'>Good luck Ethan!</Button>}
                //q: How to make a button that sends a request to the server?

                <Button onClick={() => fetch('http://localhost:8080/apitest')}>Click me</Button>
            </div>
        </nav>
    );
}
/*<ul className={click ? 'nav-menu active' : 'nav-menu'}>
    <li className='nav-item'>
        <Link to='/Fixtures' className = 'nav-links' onClick = {closeMenu}>
            Fixtures
        </Link>
    </li>
    <li className='nav-item'>
        <Link to='/Players' className = 'nav-links' onClick = {closeMenu}>
            Players
        </Link>
    </li>
    <li className='nav-item'>
        <Link to='/Clubs' className = 'nav-links' onClick = {closeMenu}>
            Clubs
        </Link>
    </li>
    <li className='nav-item'>
        <Link to='/Competitions' className = 'nav-links' onClick = {closeMenu}>
            Competitions
        </Link>
    </li>
    <li className='nav-item'>
        <Link to='/About' className = 'nav-links' onClick = {closeMenu}>
            About
        </Link>
    </li>
</ul> */
export default NavigationBar;