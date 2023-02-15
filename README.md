# The order receiver

## Project goal
### In the database:
* Deploy postgresql locally
* Create your own database
* Set up your user
* Create tables to store the received data
### In the service:
* Connect and subscribe to a channel in nats-streaming
* Write received data in Postgres
* Save received data in memory in the service (Cache)
* In case of service crash restore cache from Postgres
* Set up a http server and show data by id from cache
* Build simple interface to display received data, to query it by id
### Additional info:
* The data is static, based on that think about the storage model in Cache and in pg. The model is in the file model.json
* Anything can be sent into your channel, so think about how to avoid problems with it.
* To check if your subscription works online, make a separate script to publish data in the channel
* Think about how not to lose data if there are errors or problems with the service
* Deploy Nats-streaming locally

## Tech stack
[![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)](https://go.dev)
[![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)](https://www.postgresql.org)
[![STAN](https://svgshare.com/i/qMu.svg)](https://nats.io)

## Structure
* `bin/` - build files
* `cmd/` - main application
* `cmd/configs` - configs files
* `internal/` - internal application files
* `internal/streaming` - streaming files (nats-streaming)
* `internal/testing` - some test data
* `internal/web` - http-server
* `pkg/` - public files

## How to build
**Use the Makefile for build and run the application**
* `make build` - builds the application and saves the output binary file to bin/app.out
* `make run` - builds the application and runs the bin-file
* `make clean` - cleans object files from package source directories (for example: app.out)