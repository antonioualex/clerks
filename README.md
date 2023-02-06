# Clerks

### Introduction
This project is an implementation of a backend server, which is written in Go and uses ArangoDB as its database.
The architecture of the project follows the Domain-Driven Design (DDD) approach.
The server provides two endpoints for populating the database with 5000 random user entries and for retrieving
the list of clerks based on a few parameters.


### Endpoints
The server provides the following endpoints:

- A **POST** Method with endpoint: **/populate**, e.g., `localhost:8000/populate`. <br>
  This endpoint adds 5000 random user entries from the Randomuser.com API to the database.
  Each user's information such as name, email, phone number, picture, and registration date is stored.
- A **GET** Method with endpoint **/clerks** , e.g., `localhost:8000/clerks`.
  This endpoint returns a list of users sorted by registration date, with the most recent users
  appearing first.
  You can specify optional parameters to filter and paginate the
  results. The parameters are:
    - **limit**: Specifies the maximum number of users returned to the response.
      The value of limit must be a number between 1 and 100, e.g., `localhost:8000/clerks?limit=5`.
    - **offset**: Indicates the starting point for pagination by including the starting_after
      or ending_before value from a previous response. The starting_after value indicates the next page,
      while the ending_before value indicates the previous page. Pages are counted starting from 0, e.g.,
      `localhost:8000/clerks?offset=7`.
    - **email**: Filters the users based on their email address, e.g., `localhost:8000/clerks?email=cl@erk.com`.

You can use multiple parameters by adding the `&` symbol between the parameters, e.g.,
`localhost:8000/clerks?limit=5&offset=10`.



### HowTo

#### Setup Server
`Ensure that you have Docker and Docker Compose installed on your system.`

Docker Compose is used to setup the server and the database for easy reproduction and testing of the project.

- Navigate to the **infra/docker** directory of the project.
- Run **docker-compose up** to build and start the services defined in the Docker Compose file.
- The backend server will be available at `http://localhost:8572`.

To stop the services, press Ctrl + C and run **docker-compose down** to remove the containers and networks
created by the services.

#### Run Clerk Service

`Ensure that you have Go installed on your system. `
- Navigate to the root directory of the project in your terminal.
- Run `go build ./cmd/main.go`  to build the binary executable for the service.
- Start the service by running the following command: ```./cmd/main```

###### Optional enviroment variables

```
SERVE_ON_PORT= Set service port, default value is 8000
ARANGO_URL= Database server address, by default is http://localhost:8572, it is recommended not be changed due to its connection with the docker-compose.yaml. 
```