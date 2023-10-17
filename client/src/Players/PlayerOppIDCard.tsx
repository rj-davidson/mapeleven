import * as React from 'react';
import { Tile } from '../Util/Tile.jsx';
import { Flex } from '../Util/Flex.jsx';
import DisplayImage from '../Util/DisplayImage';
import PlayerIDTitle from './PlayerIDTitle';

interface TeamOppIDCardProps {
    slug: string;
    name: string;
    photo: string;
    injured: boolean;
}

const TeamOppIDCard: React.FC<TeamOppIDCardProps> = ({ slug = '', name = 'Team Name', injured, photo }) => {
    return (
        <Tile
            style={{
                flexDirection: 'column',
                justifyContent: 'space-between',
                height: '100%',
                background: 'var(--dark1)',
                border: '5px solid var(--blue)',
            }}
        >
            <PlayerIDTitle slug={slug} name={name} country={'Country'} flag={''} injured={injured} />

            <Flex style={{ justifyContent: 'center', alignItems: 'center', height: '100%' }}>
                <DisplayImage
                    src={photo}
                    alt={name}
                    width={'108px'}
                    height={'108'}
                    sx={{
                        borderRadius: '10px',
                        boxShadow: 'rgba(0,0,0,0.2) 0 0 10px',
                    }}
                />
            </Flex>
        </Tile>
    );
};

export default TeamOppIDCard;
