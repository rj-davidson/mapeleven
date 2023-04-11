import React, {useState} from 'react';
import {Link} from 'react-router-dom';



function NavigationBar(){
const [click, setClick] = useState(false);

const handleClick = () => setClick(!click);
const closeMenu = () => setClick(false);
// q: How to make a navigation bar that has 4 different tabs in the top right corner?




    return(
        <nav className="navigationbar">
            <div className="navigationbar-container">
                <div className='menu-icon' onClick={handleClick}>
                    <i className={click ? 'fas fa-time': 'fas fa-bars'}/>
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