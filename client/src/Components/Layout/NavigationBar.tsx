import { useState, useEffect } from 'react';
import { Link, useLocation } from 'react-router-dom';
import './NavigationBar.css';

function NavigationBar(): JSX.Element {
    const [click, setClick] = useState(false);
    const [button, setButton] = useState(true);
    const location = useLocation();
    const [activePage, setActivePage] = useState(location.pathname.split('/')[1]);

    const handleClick = () => setClick(!click);
    const closeMenu = () => setClick(false);

    const showButton = () => {
        if (window.innerWidth <= 960) {
            setButton(false);
        } else {
            setButton(true);
        }
    };

    useEffect(() => {
        window.addEventListener('resize', showButton);
        return () => {
            window.removeEventListener('resize', showButton);
        };
    }, []);

    useEffect(() => {
        setActivePage(location.pathname.split('/')[1]);
    }, [location]);

    return (
        <nav className='navigationbar'>
            <div className='navigationbar-container'>
                <div className='menu-icon' onClick={handleClick}>
                    <i className={click ? 'fas fa-times' : 'fas fa-bars'} />
                </div>
                <div className={'title'}>
                    <Link to='/' className='nav-links' onClick={closeMenu}>
                        {' '}
                        mapeleven{' '}
                    </Link>
                </div>
                <ul className={click ? 'nav-menu active' : 'nav-menu'}>
                    <li className='nav-item'>
                        <Link
                            to='/fixtures'
                            className={activePage === 'fixtures' ? 'nav-links active' : 'nav-links'}
                            onClick={closeMenu}
                        >
                            {' '}
                            Fixtures{' '}
                        </Link>
                    </li>
                    <li className='nav-item'>
                        <Link
                            to='/players'
                            className={activePage === 'players' ? 'nav-links active' : 'nav-links'}
                            onClick={closeMenu}
                        >
                            {' '}
                            Players{' '}
                        </Link>
                    </li>
                    <li className='nav-item'>
                        <Link
                            to='/clubs'
                            className={activePage === 'clubs' ? 'nav-links active' : 'nav-links'}
                            onClick={closeMenu}
                        >
                            {' '}
                            Clubs{' '}
                        </Link>
                    </li>
                    <li className='nav-item'>
                        <Link
                            to='/leagues'
                            className={activePage === 'leagues' ? 'nav-links active' : 'nav-links'}
                            onClick={closeMenu}
                        >
                            {' '}
                            Leagues{' '}
                        </Link>
                    </li>
                    <li className='nav-item'>
                        <Link
                            to='/about'
                            className={activePage === 'about' ? 'nav-links active' : 'nav-links'}
                            onClick={closeMenu}
                        >
                            {' '}
                            About{' '}
                        </Link>
                    </li>
                </ul>
            </div>
        </nav>
    );
}

export default NavigationBar;
