<a href="https://www.velocityworks.io/home">Velocity Works Coding Demo</a>

# Golang-Demo

This Golang application consumes a JSON payload from https://www.data.gov/, populates a database and displays the database contents on a web page.

Frameworks used:

- Echo to build the website: https://github.com/labstack/echo
- Resty to retrieve data from Data.gov: https://github.com/go-resty/resty
- Jsonparser to process the data: https://github.com/buger/jsonparser
- GORM to populate the database: https://github.com/jinzhu/gorm

## Setup

Create a `.env` file with the following variables:

```bash
POSTGRES_USER=user
POSTGRES_PASSWORD=admin
POSTGRES_DB=renewable_energy
DB_PORT=5432
DB_HOST=db
```

## Run

The app is setup with Docker. Run it with `docker-compose`:

```bash
docker-compose up;
```

The app will be available in `http://localhost:8080/renewables`
