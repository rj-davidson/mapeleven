import React, {useState} from 'react';
import {Link} from 'react-router-dom';
import {Button} from "./Button.jsx";
import './NavigationBar.css';
//import fetch from 'fetch';


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
                <div className={"title"}>
                    <Link to='/' className = 'nav-links' onClick = {closeMenu}> mapeleven </Link>
                </div>
                <ul className={click ? 'nav-menu active' : 'nav-menu'}>
                    <li className='nav-item'>
                        <Link to='/fixtures' className = 'nav-links' onClick = {closeMenu}> Fixtures </Link>
                    </li>
                    <li className='nav-item'>
                        <Link to='/players' className = 'nav-links' onClick = {closeMenu}> Players </Link>
                    </li>
                    <li className='nav-item'>
                        <Link to='/clubs' className = 'nav-links' onClick = {closeMenu}> Clubs </Link>
                    </li>
                    <li className='nav-item'>
                        <Link to='/leagues' className = 'nav-links' onClick = {closeMenu}> Leagues </Link>
                    </li>
                    <li className='nav-item'>
                        <Link to='/about' className = 'nav-links' onClick = {closeMenu}> About </Link>
                    </li>
                </ul>
            </div>
        </nav>
    );
}
/*<ul className={click ? 'nav-menu active' : 'nav-menu'}>
    <li className='nav-item'>
        <Link to='/Fixtures' className = 'nav-links' onClick = {closeMenu}>
            Fixtures
        </Link>
                <Button onClick={() => fetch('http://localhost:8080/apitest')}>Click me</Button>
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