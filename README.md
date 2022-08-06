# News CRUD API

This repository is CRUD API for news management applying the clean code with three layer of handler, service, and repository. Created using GO, using redis for caching, pure sql library (no ORM) and docker for containerizing.

## Relation
![image](https://user-images.githubusercontent.com/63847012/183245053-c9f03f0e-d27a-4c52-8c13-8f7c0c5d38df.png)


## Features

This service come into two endpoint

- `/api/v1/articles/` for articles related
  - GET
    - query
    - author
    - page
  - POST
  - DELETE
  - PATCH
- `/api/v1/authors/` for authors related
  - GET
    - page
  - POST
  - DELETE
  - PUT

some of them containing extra query parameters and all of the documentation could be found in the `doc` folder. Documentation created using postman collection so import it to your local postman and start using it~

## Technologies Used

1. Go
2. Docker
3. Redis
4. Postgresql
5. go-migrate
6. air
7. cobra-cmd

## How to run

1. Install all dependencies listed above
2. Create .env from .env.example and fill it up
3. Run the docker container using `docker-compose up -d` to run on the background or `docker-compose up --build` to run in the CLI
4. run
5. to run the server
   1. `air server` for hot reloading after update
   2. `go run main.go server`
