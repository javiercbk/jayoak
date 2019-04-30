# Jayoak

Jayoak is a web application that allows a user to explore an audio analysis database. It is the main tool for a musical research.


## How does it work

A user logs in into application and can search sounds, most likely instrument notes, and compare the frequencies between different sounds. Some frequencies will alter other frequencies in a particular way, in other words some notes on an instrument will affect other notes in other instrument, the goal of this project is to know how.

By uploading a wav file, the user will get the frequency spectrum of the sound wave by performing a [fourier transform](https://en.wikipedia.org/wiki/Fourier_transform) using the [FFT](https://en.wikipedia.org/wiki/Fast_Fourier_transform) algorithm. Such frequency spectrum is stored in the database for later comparison between other sounds.

## Dependencies

* Go
* Docker
* [golangci-lint](github.com/golangci/golangci-lint/cmd/golangci-lint) go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

## How to build

```sh
make server
```

## How to run

At the moment the server doesn't do much, in fact I don't think you can do anything useful, but you can run the test though.


```sh
go test ./...
```

## TODO

- [x] (Backend) Create the main application skeleton.
- [x] (Backend) Create the database schema.
- [x] (Backend) Given a wav file, calculate the sound spectrum.
- [x] (Backend) Benchmark if a Postgres array type stores the frequencies faster that storing a single record for each frequency.
- [x] (Backend) Start a docker container when running the application's test.
- [x] (Architecture) Define if redis is worth the trouble or not (probably not).
- [ ] (Architecture) Define if the array approach pays off when searching and reading frequencies.
- [ ] (Backend) Authenticate a user.
- [ ] (Backend) Implement access control.
- [ ] (Frontend) Create the main frontend application skeleton with Vue.js
- [ ] (Frontend) Create the login page and logout logic.
- [ ] (Frontend) Create the upload sound page.
- [ ] (Frontend) Create the sound analysis and edition page.
- [ ] (Backend) Implement the sound searching handlers.
- [ ] (Frontend) Create the sound search page.


## Software related Q&A

#### What is the status of this project?

The project is in its very early stages. At the moment it is a web application that mostly run unit testing but anyone experienced enough should foresee how will the application's components are going to interact. This project is not rocket science, but not a trivial hello world either.

#### Why are you using Go?

Go is really suitable for web applications that performs heavy processing such as audio analysis.

Plus I love Go, it is a fantastic language and I have plenty of experience with it. My second option was [Rust](https://www.rust-lang.org).

#### Why are you using Gin? Don't you know that Go doesn't need a library to route requests?

Usually I use [Gorilla Mux](https://github.com/gorilla/mux) because I really like having variables within urls and Go standard router does not support them.
It's been a while since I last used [Gin](github.com/gin-gonic/gin) and I wanted to see where the project was going. I'm not in love but I don't hate it either, I still prefer Gorilla Mux though.

Regardless of which library I use, the general pattern I apply to initialize handlers is the same in every project and this project is no exception.

#### Why are you using Postgres?

The application's data is highly structured, so a relational database was the best choice and Postgres is arguably one of the best relational databases.

Also the project is storing a lot of frequencies related data. If you stick the the fully relational approach you will end up creating +1000 records per sound analyzed, would and array improve performance on read or write?

At the moment of writing, my benchmarks says that the array method is faster to store. But what about reading it? or comparing it? I want to know that as well. Postgres gives me the possibility of denormalizing data to test both approaches.

#### Why are you spinning a docker container before running tests?

I do not like mocking database queries, I believe that there is value on testing sql queries even though test might be a little more complex.

Also I do think that you don't have to be a sheep and think for yourself, and I think that testing works best for me if I have a database running my queries. I write test to deliver my best work, THEN I try to see how can I adapt my tests to provide value for the organization. In this case the organization is me and I approve myself.

Even though this approach is working just fine, I'm currently researching if this is the best approach for testing and trying to figure if there is a better way.


#### You have a package called "api", don't you know that packages should have a meaningful name and not some generic word?

Actually, the "api" directory is just a directory. It contains other packages (which are the application's API and route handlers) that used to be in the root directory, but I didn't want to bloat the root directory so I added the "api" directory. I'm open to suggestions here.

#### Why are you using Redis?

At the moment I want to do some benchmarking, but most likely I'll leave sessions stored inside Postgres.


#### DUDE, you left all your database secrets out in the open in a public repo?

Those are my local development environment keys.


## Jayoak uses the following open source libraries

* [go-audio](https://github.com/go-audio)
* [gin](github.com/gin-gonic/gin)
* A slightly modified version of [filetype](https://github.com/h2non/filetype)
* [sqlboiler](github.com/volatiletech/sqlboiler) and the postgres' driver [pq](github.com/lib/pq)
* [go.uuid](github.com/satori/go.uuid)
* [go-validator](gopkg.in/go-playground/validator.v9)