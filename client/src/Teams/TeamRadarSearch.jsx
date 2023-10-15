import * as React from 'react';
import TextField from '@mui/material/TextField';
import Autocomplete from '@mui/material/Autocomplete';
import { useEffect, useState } from 'react';
import { createFilterOptions, Paper } from '@mui/material';
import SearchIcon from '@mui/icons-material/Search';

const url = import.meta.env.VITE_API_URL;

export default function TeamRadarSearch({ handleSearch }) {
    const [data, setData] = useState([]);
    const maxFilter = 6;

    useEffect(() => {
        // Send a GET request to the API.
        fetch(`${url}/teams`)
            .then(response => response.json())
            .then(jsonData => {
                setData(jsonData);
            })
            .catch(error => console.error(error));
    }, []);

    const customRenderInput = params => (
        <TextField
            {...params}
            fullWidth
            variant='outlined'
            label={
                <div style={{ display: 'flex', alignItems: 'center' }}>
                    <SearchIcon sx={{ marginRight: '8px' }} />
                    <label>Search team to compare</label>
                </div>
            }
            InputProps={{
                ...params.InputProps,
                onKeyDown: e => {
                    if (e.key === 'Enter') {
                        e.stopPropagation();
                    }
                },
                style: {
                    /* your custom CSS styles for the input field */
                },
            }}
        />
    );

    const customRenderOption = (props, option) => (
        <li {...props} style={{ marginBottom: 8 }}>
            <img
                src={`${url}/` + option.badge}
                alt={option.name.short}
                style={{ marginRight: '16px', width: '32px', height: '32px', overflow: 'hidden' }}
            />
            {option.name.long}
        </li>
    );

    const CustomPaperComponent = ({ children }) => {
        return (
            <Paper style={{ padding: '8px', backgroundColor: 'var(--dark0)', color: 'var(--light2)' }}>
                {children}
            </Paper>
        );
    };

    const defaultFilterOptions = createFilterOptions();
    const filterOptions = (options, state) => {
        return defaultFilterOptions(options, state).slice(0, maxFilter);
    };

    return (
        <Autocomplete
            disablePortal
            disableClearable
            options={data}
            getOptionLabel={option => option.name.long}
            sx={{ width: '%100' }}
            onChange={handleSearch}
            noOptionsText='No results'
            popupIcon={''}
            filterOptions={filterOptions}
            PaperComponent={CustomPaperComponent}
            renderInput={customRenderInput}
            renderOption={customRenderOption}
            ListboxProps={{
                style: {
                    maxHeight: '80vh',
                },
            }}
        />
    );
}
