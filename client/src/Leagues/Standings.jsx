import React from 'react';
import { Table, TableHead, TableBody, TableRow, TableCell, Paper, Typography } from '@mui/material';
import { Flex } from '../Util/Flex.jsx';
import DisplayImage from '../Util/DisplayImage';
import { Link } from 'react-router-dom';

const StandingsTable = ({ data }) => {
    const sortedData = [...data].sort((a, b) => a.rank - b.rank);

    //rgba(59, 66, 82, 1)
    const getBackgroundColor = (size, rank) => {
        if (rank === 1) {
            return 'rgba(59, 80, 82, 1)';
        } else if (rank > 1 && rank < 5) {
            return 'rgba(59, 66, 100, 1)';
        } else if (rank > size - 3) {
            return 'rgba(80, 66, 82, 1)';
        } else {
            return 'var(--dark1)';
        }
    };

    return (
        <Paper elevation={3} sx={{ width: '100%', backgroundColor: 'var(--dark0)', borderRadius: '12px' }}>
            <Table sx={{ borderRadius: '12px' }}>
                <TableHead>
                    <TableRow>
                        <TableCell sx={{ fontWeight: 'bold' }}>Team</TableCell>
                        <TableCell sx={{ fontWeight: 'bold' }} align='right'>
                            Played
                        </TableCell>
                        <TableCell sx={{ fontWeight: 'bold' }} align='right'>
                            Won
                        </TableCell>
                        <TableCell sx={{ fontWeight: 'bold' }} align='right'>
                            Draw
                        </TableCell>
                        <TableCell sx={{ fontWeight: 'bold' }} align='right'>
                            Lost
                        </TableCell>
                        <TableCell sx={{ fontWeight: 'bold' }} align='right'>
                            Points
                        </TableCell>
                        <TableCell sx={{ fontWeight: 'bold' }} align='right'>
                            Form
                        </TableCell>
                    </TableRow>
                </TableHead>
                <TableBody sx={{ backgroundColor: 'var(--dark1)' }}>
                    {sortedData.map((row, index) => (
                        <TableRow key={index} sx={{ backgroundColor: getBackgroundColor(sortedData.length, row.rank) }}>
                            <TableCell>
                                <Flex style={{ justifyContent: 'left', alignItems: 'center', gap: '8px' }}>
                                    <Typography sx={{ fontWeight: 'bold' }}>{row.rank}</Typography>
                                    <DisplayImage src={row.logo} alt={row.name} sx={{ marginLeft: '16px' }} />
                                    <Link to={`/teams/${row.slug}`}>
                                        <Typography
                                            sx={{
                                                '&:hover': {
                                                    textDecoration: 'underline', // Add underline on hover
                                                },
                                            }}
                                        >
                                            {row.name}
                                        </Typography>
                                    </Link>
                                </Flex>
                            </TableCell>
                            <TableCell align='right'>{row.played}</TableCell>
                            <TableCell align='right'>{row.wins}</TableCell>
                            <TableCell align='right'>{row.draw}</TableCell>
                            <TableCell align='right'>{row.losses}</TableCell>
                            <TableCell align='right'>{row.points}</TableCell>
                            <TableCell align='right'>{row.form}</TableCell>
                        </TableRow>
                    ))}
                </TableBody>
            </Table>
        </Paper>
    );
};

export default StandingsTable;
