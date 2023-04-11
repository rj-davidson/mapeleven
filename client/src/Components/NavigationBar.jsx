import React, {useState} from "react";
import Link from "react-router-dom";


function NavigationBar(){
const [click, setClick] = useState(false);

const handleClick = () => setClick(!click);
const closeMenu = () => setClick(false)

    return(
        <nav className="navigationbar">
            <div className="navigationbar-container">
                <Link>Navigation Bar</Link>
                <div className='menu-icon' onClick={handleClick}>
                    <i className={click ? 'fas fa-time': 'fas fa-bars'}/>
                </div>
                <ul className={click ? 'nav-menu active' : 'nav-menu'}>
                    <li className='nav-item'>
                        <Link to='/' className = 'nav-links' onClick = {closeMenu}>
                            Home
                        </Link>
                    </li>
                </ul>
            </div>
        </nav>
    )
}