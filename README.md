# mapeleven

## Project Description
mapeleven is a comprehensive, soccer-focused web application that provides users with access to live scores, club, league, player and fixture info. The data is presented through a variety of visual techniques, enhancing the overall user experience. The user-friendly interface, complete with a search bar and a date selection feature, allows users to effortlessly search for players, teams, and leagues, or check out upcoming and past fixtures.

## Technical Details

### Frontend

The frontend is built with [ReactJS](https://reactjs.org/), complemented by [TypeScript](https://www.typescriptlang.org/) and Javascript, along with [Vite](https://vitejs.dev/) as a build tool, and [Yarn](https://yarnpkg.com) as a package manager. For the interface, the [MUI](https://mui.com/) library is used for creating a robust and responsive UI.

### Backend

The backend is crafted in [Go](https://golang.org/), using [Fiber](https://gofiber.io/) as a web framework, and [Ent](https://entgo.io/) as an ORM. Data is stored in a PostgreSQL database.

## Build Instructions

The project is containerized using Docker, so ensure that [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/) are installed. Also make sure that make has been installed as well. This will be utilized in the command to deploy the project.

The primary command to deploy the project:

```bash
make docker-up
```

To run the client individually ensure that yarn is installed.

The primary command to deploy the client:
```bash
yarn run start
```

This command builds and starts the Docker containers defined in the Docker Compose files of the project.

## External API Usage

This project relies on [API-FOOTBALL](https://www.api-football.com/) as its primary data source. To run Mapeleven, you must have an API key from API-FOOTBALL. Head over to their website, register, and obtain your API key.

## Environment Configuration

To configure the application, you need to set up environment files for both the frontend and backend components of the project.:
-  For the backend, edit the `api/.env` file with your corresponding database details, host details, and API-FOOTBALL key like so:
    ```
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=<name>
    DB_PASS=<pass>
    DB_NAME=<db name>

    ENV_PATH=
    IP_HOST=localhost
    DOMAIN=<when necessary>

    API_KEY=<API-FOOTBALL API Key>
    API_HOST=api-football-v1.p.rapidapi.com
    ```

-  For the frontend, modify the `client/.env.local` file to set the URL of the API server:
    ```
    VITE_API_URL=http://localhost:8080 //this will change if the user deploys to a server
    ```

## Server Deployment

To deploy this project to a Debian server with Nginx, follow these steps:

1. Clone the project onto your server.
2. Install Docker and Docker Compose if not already installed.
3. Configure your environment files (`api/.env` and `client/.env.local`) as described above.
4. Once Docker and Docker Compose are installed and your environment files are configured, you can use the `make docker-up` command to build and start the Docker containers.

Nginx configuration is beyond the scope of this README. Please consult the official [Nginx documentation](https://nginx.org/en/docs/) for instructions on setting up Nginx as a reverse proxy and for serving static files.


## External Libraries
Our project makes use of the following key external libraries:

#### Frontend
* [D3](https://d3js.org/)
* [ESLint](https://eslint.org/)
* [React](https://reactjs.org/)
* [MUI](https://mui.com/)
* [Prettier](https://prettier.io/)
* [TypeScript](https://www.typescriptlang.org/)
* [Vite](https://vitejs.dev/)
* [Yarn](https://yarnpkg.com)

#### Backend
* [Ent](https://entgo.io/ent)
* [GoCron](https://github.com/go-co-op/gocron)
* [Fiber](https://github.com/gofiber/fiber/v2)
* [Viper](https://github.com/spf13/viper)

## User Usage

The mapeleven web application is designed to offer a seamless user experience on a variety of devices, from desktops to tablets and smartphones. Simply navigate to the site in your web browser of choice, and use the search bar to find players, teams, or leagues, or use the date selector on the homepage to view upcoming and past fixtures.

### Technical Requirements

The mapeleven web app is built to run efficiently on any device with a modern web browser, such as Chrome, Firefox, Safari, or Edge. It is mobile-friendly and responsive, ensuring a smooth experience on smartphones and tablets.