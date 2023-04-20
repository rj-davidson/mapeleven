import { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import './NavigationBar.css';

function NavigationBar(): JSX.Element {
    const [click, setClick] = useState(false);
    const [button, setButton] = useState(true);

    const handleClick = () => setClick(!click);
    const closeMenu = () => setClick(false);

    const showButton = () => {
        if (window.innerWidth <= 960) {
            setButton(false);
        } else {
            setButton(true);
        }
    }

    useEffect(() => {
        window.addEventListener('resize', showButton);
        return () => {
            window.removeEventListener('resize', showButton);
        };
    }, []);

    return (
        <nav className="navigationbar">
            <div className="navigationbar-container">

                <div className='menu-icon' onClick={handleClick}>
                    <i className={click ? 'fas fa-times' : 'fas fa-bars'} />
                </div>
                <div className={"title"}>
                    <Link to='/' className='nav-links' onClick={closeMenu}> mapeleven </Link>
                </div>
                <ul className={click ? 'nav-menu active' : 'nav-menu'}>
                    <li className='nav-item'>
                        <Link to='/fixtures' className='nav-links' onClick={closeMenu}> Fixtures </Link>
                    </li>
                    <li className='nav-item'>
                        <Link to='/players' className='nav-links' onClick={closeMenu}> Players </Link>
                    </li>
                    <li className='nav-item'>
                        <Link to='/clubs' className='nav-links' onClick={closeMenu}> Clubs </Link>
                    </li>
                    <li className='nav-item'>
                        <Link to='/leagues' className='nav-links' onClick={closeMenu}> Leagues </Link>
                    </li>
                    <li className='nav-item'>
                        <Link to='/about' className='nav-links' onClick={closeMenu}> About </Link>
                    </li>
                </ul>
            </div>
        </nav>
    );
}

export default NavigationBar;
