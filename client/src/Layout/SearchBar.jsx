import * as React from 'react';
import TextField from '@mui/material/TextField';
import Autocomplete from '@mui/material/Autocomplete';
import { useEffect, useState } from 'react';
import { createFilterOptions, Paper } from '@mui/material';
import { styled } from '@mui/material/styles';
import { useNavigate } from 'react-router-dom';
import SearchIcon from '@mui/icons-material/Search';

const url = import.meta.env.VITE_API_URL;

export default function SearchBar() {
    const [data, setData] = useState([]);
    const maxFilter = 6;

    useEffect(() => {
        // Send a GET request to the API.
        fetch(`${url}/search`)
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
            placeholder="Search players, teams, or leagues"
            InputProps={{
                ...params.InputProps,
                onFocus: (e) => {
                    if (!e.target.value) {
                        e.target.setAttribute(
                            "placeholder",
                            "Search players, teams, or leagues"
                        );
                    }
                },
                onBlur: (e) => {
                    if (!e.target.value) {
                        e.target.setAttribute(
                            "placeholder",
                            "Search players, teams, or leagues"
                        );
                    }
                },
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
                src={`${url}/` + option.image}
                alt={option.name}
                style={{ marginRight: '16px', width: '32px', height: '32px', overflow: 'hidden' }}
            />
            {option.name}
        </li>
    );

    const CustomPaperComponent = ({ children }) => {
        return (
            <Paper style={{ padding: '8px', backgroundColor: 'var(--dark0)', color: 'var(--light2)' }}>
                {children}
            </Paper>
        );
    };

    const GroupHeader = styled('div')(() => ({
        position: 'sticky',
        top: '-8px',
        padding: '12px 10px',
        fontWeight: 'bold',
        color: 'var(--light2)',
        backgroundColor: 'var(--dark0)',
    }));

    const GroupItems = styled('ul')({
        padding: 10,
    });

    const navigate = useNavigate();
    const handleSearch = (event, newValue) => {
        if (newValue) {
            // Assuming your data contains players, teams, and leagues with unique IDs.
            const selectedItem = data.find(item => item.slug === newValue.slug);

            if (selectedItem) {
                if (selectedItem.type === 'player') {
                    navigate(`/players/${selectedItem.slug}`);
                }
                if (selectedItem.type === 'team') {
                    navigate(`/teams/${selectedItem.slug}`);
                }
                if (selectedItem.type === 'league') {
                    navigate(`/leagues/${selectedItem.slug}`);
                }
                if (selectedItem.type === 'fixture') {
                    navigate(`/fixtures/${selectedItem.slug}`);
                }
            }
        }
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
            getOptionLabel={option => option.name}
            groupBy={option => option.type.toUpperCase()}
            sx={{ width: '%100', my: 2 }}
            onChange={handleSearch}
            noOptionsText='No results'
            popupIcon={''}
            filterOptions={filterOptions}
            PaperComponent={CustomPaperComponent}
            renderInput={customRenderInput}
            renderOption={customRenderOption}
            renderGroup={params => (
                <li key={params.key}>
                    <GroupHeader>{params.group}</GroupHeader>
                    <GroupItems>{params.children}</GroupItems>
                </li>
            )}
            ListboxProps={{
                style: {
                    maxHeight: '80vh',
                },
            }}
        />
    );
}
