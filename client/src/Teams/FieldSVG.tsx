import * as React from 'react';

const Field = () => (
    <svg width='100%' height='100%' viewBox='0 0 600 400'>
        {/* Grass Background */}
        <rect x='0' y='0' width='100%' height='100%' fill='transparent' />

        {/* Outer Boundary */}
        <rect x='50' y='50' width='500' height='300' fill='transparent' stroke='#fff' strokeWidth='4' />

        {/* Center Circle */}
        <circle cx='300' cy='200' r='50' fill='transparent' stroke='#fff' strokeWidth='4' />

        {/* Center Spot */}
        <circle cx='300' cy='200' r='5' fill='#fff' />

        {/* Large Penalty Area - Left */}
        <rect x='50' y='125' width='100' height='150' fill='transparent' stroke='#fff' strokeWidth='4' />

        {/* Large Penalty Area - Right */}
        <rect x='450' y='125' width='100' height='150' fill='transparent' stroke='#fff' strokeWidth='4' />

        {/* Large Penalty Spot - Left */}
        <circle cx='120' cy='200' r='5' fill='#fff' />

        {/* Large Penalty Spot - Right */}
        <circle cx='480' cy='200' r='5' fill='#fff' />

        {/* Small Penalty Area - Left */}
        <rect x='50' y='155' width='40' height='90' fill='transparent' stroke='#fff' strokeWidth='4' />

        {/* Small Penalty Area - Right */}
        <rect x='510' y='155' width='40' height='90' fill='transparent' stroke='#fff' strokeWidth='4' />

        {/* Small Penalty Spot - Left */}
        <circle cx='120' cy='200' r='3' fill='#fff' />

        {/* Small Penalty Spot - Right */}
        <circle cx='480' cy='200' r='3' fill='#fff' />

        {/* Halfway Line */}
        <line x1='300' y1='50' x2='300' y2='350' stroke='#fff' strokeWidth='4' />

        {/* Red Dots */}
        <circle cx='69' cy='200' r='10' fill='var(--red)' stroke='var(--dark0)' strokeWidth='3px' />
        <circle cx='190' cy='90' r='10' fill='var(--red)' stroke='var(--dark0)' strokeWidth='3px' />
        <circle cx='170' cy='160' r='10' fill='var(--red)' stroke='var(--dark0)' strokeWidth='3px' />
        <circle cx='170' cy='240' r='10' fill='var(--red)' stroke='var(--dark0)' strokeWidth='3px' />
        <circle cx='190' cy='315' r='10' fill='var(--red)' stroke='var(--dark0)' strokeWidth='3px' />
        <circle cx='280' cy='125' r='10' fill='var(--red)' stroke='var(--dark0)' strokeWidth='3px' />
        <circle cx='230' cy='200' r='10' fill='var(--red)' stroke='var(--dark0)' strokeWidth='3px' />
        <circle cx='280' cy='275' r='10' fill='var(--red)' stroke='var(--dark0)' strokeWidth='3px' />
        <circle cx='370' cy='110' r='10' fill='var(--red)' stroke='var(--dark0)' strokeWidth='3px' />
        <circle cx='400' cy='200' r='10' fill='var(--red)' stroke='var(--dark0)' strokeWidth='3px' />
        <circle cx='370' cy='290' r='10' fill='var(--red)' stroke='var(--dark0)' strokeWidth='3px' />
    </svg>
);

export default Field;
