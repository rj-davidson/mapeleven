import React from 'react';
import {Table, TableHead, TableBody, TableRow, TableCell, Paper, Box, Typography} from '@mui/material';
import {Flex} from "../Util/Flex.jsx";

const StandingsTable = ({data}) => {

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
            return 'var(--dark1)'
        }
    };

    return (
        <Paper elevation={3} sx={{width: '100%', backgroundColor: 'var(--dark0)', borderRadius: '12px'}}>
            <Table sx={{borderRadius: '12px'}}>
                <TableHead>
                    <TableRow>
                        <TableCell sx={{fontWeight: 'bold'}}>Team</TableCell>
                        <TableCell sx={{fontWeight: 'bold'}} align="right">Played</TableCell>
                        <TableCell sx={{fontWeight: 'bold'}} align="right">Won</TableCell>
                        <TableCell sx={{fontWeight: 'bold'}} align="right">Draw</TableCell>
                        <TableCell sx={{fontWeight: 'bold'}} align="right">Lost</TableCell>
                        <TableCell sx={{fontWeight: 'bold'}} align="right">Points</TableCell>
                    </TableRow>
                </TableHead>
                <TableBody sx={{backgroundColor: 'var(--dark1)'}}>
                    {sortedData.map((row, index) => (
                        <TableRow key={index} sx={{backgroundColor: getBackgroundColor(sortedData.length, row.rank)}}>
                            <TableCell>
                                <Flex style={{justifyContent: 'left', alignItems: 'center', gap: '8px'}}>
                                    <Typography sx={{fontWeight: 'bold'}}>
                                        {row.rank}
                                    </Typography>
                                    <Box
                                        component="img"
                                        sx={{
                                            maxHeight: {xs: 32, sm: 32},
                                            maxWidth: {xs: 32, sm: 32},
                                            marginLeft: '16px'
                                        }}
                                        alt={name}
                                        src={'http://localhost:8080/' + row.logo}
                                    />
                                    {row.name}
                                </Flex>
                            </TableCell>
                            <TableCell align="right">{row.played}</TableCell>
                            <TableCell align="right">{row.wins}</TableCell>
                            <TableCell align="right">{row.draw}</TableCell>
                            <TableCell align="right">{row.losses}</TableCell>
                            <TableCell align="right">{row.points}</TableCell>
                        </TableRow>
                    ))}
                </TableBody>
            </Table>
        </Paper>
    );
};

export default StandingsTable;
