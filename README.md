<a href="https://www.velocityworks.io/home">Velocity Works Coding Demo</a>
# Golang-Demo

This Golang application consumes a JSON payload from https://www.data.gov/, populates a database and displays the database contents on a web page.

Frameworks used:

- Echo to build the website: https://github.com/labstack/echo
- Resty to retrieve data from Data.gov: https://github.com/go-resty/resty
- Jsonparser to process the data: https://github.com/buger/jsonparser
- GORM to populate the database: https://github.com/jinzhu/gorm
