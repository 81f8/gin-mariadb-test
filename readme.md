# API Documentation

## Get All Countries
`GET /countries`
Returns a list of all countries.

**Example Request:**

GET /countries


**Example Response:**

HTTP/1.1 200 OK
Content-Type: application/json

[
  {
    "countryID": 1,
    "countryName": "Country 1",
    "area": 1000.0,
    "language": "English",
    "continentName": "Continent 1"
  },
  {
    "countryID": 2,
    "countryName": "Country 2",
    "area": 2000.0,
    "language": "French",
    "continentName": "Continent 2"
  },
  ...
]


## Get a Specific Country
`GET /countries/{id}`
Returns the details of a specific country identified by its `id`.

**Example Request:**

GET /countries/1


**Example Response:**

HTTP/1.1 200 OK
Content-Type: application/json

{
  "countryID": 1,
  "countryName": "Country 1",
  "area": 1000.0,
  "language": "English",
  "continentName": "Continent 1"
}


## Create a New Country
`POST /countries`
Creates a new country with the provided data.

**Example Request:**

POST /countries
Content-Type: application/json

{
  "countryName": "New Country",
  "area": 1500.0
}


**Example Response:**

HTTP/1.1 201 Created
Content-Type: application/json

{
  "countryID": 3,
  "countryName": "New Country",
  "area": 1500.0
}


## Update an Existing Country
`PUT /countries/{id}`
Updates the details of an existing country identified by its `id`.

**Example Request:**

PUT /countries/1
Content-Type: application/json

{
  "countryName": "Updated Country",
  "area": 1200.0
}


**Example Response:**

HTTP/1.1 200 OK
Content-Type: application/json

{
  "message": "Country updated successfully"
}


## Delete a Country
`DELETE /countries/{id}`
Deletes a country identified by its `id`.

**Example Request:**

DELETE /countries/1


**Example Response:**

.....

