import { useState } from 'react';
import './DaySwitcher.css';

type DaySwitcherProps = {};

const DaySwitcher: React.FC<DaySwitcherProps> = () => {
    const [currentDate, setCurrentDate] = useState(new Date());

    const handleBackwardClick = () => {
        const newDate = new Date();
        newDate.setDate(currentDate.getDate() - 1);
        setCurrentDate(newDate);
    };

    const handleForwardClick = () => {
        const newDate = new Date();
        newDate.setDate(currentDate.getDate() + 1);
        setCurrentDate(newDate);
    };

    return (
        <div>
            <div>
                <button onClick={handleBackwardClick}>{'<'}</button>
                <span id={"current-date"}>{currentDate.toDateString()}</span>
                <button onClick={handleForwardClick}>{'>'}</button>
            </div>
        </div>
    );
};

export default DaySwitcher;
