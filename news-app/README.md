# Tracing Implementation in a Simple News App

## Introduction

This repository contains a simple application that demonstrates how Tracing can be implemented. The app features a single endpoint, `/news`, which retrieves news items from a database and utilizes OpenTelemetry for tracing capabilities.

## Prerequisites

Before setting up and running the application, ensure you have the following prerequisites installed:

- [Docker](https://www.docker.com/) to run the database and Jaeger.
- [Go](https://golang.org/) to build and run the application.
- [cURL](https://curl.se/) for making HTTP requests.

## Installation and Setup

Follow the steps below to set up the application for tracing:


1. **Start the Database**: Navigate to the `deploy-db` directory and start the database using Docker Compose:
   ```shell
   cd deploy-db
   docker-compose up -d
   ```

2. **Start a Local Instance of Jaeger**: Start a local instance of Jaeger to capture traces:
    ```shell
    docker run --rm -it -p 6831:6831/udp -p 16686:16686 -p 14269:14269 --name jaeger jaegertracing/all-in-one:latest
    ```
    The jaegertracing/all-in-one image is a self-contained, all-in-one distribution of the Jaeger tracing system. It includes all the necessary components to set up a local Jaeger instance for tracing.

3. **Run the Go Application**: 
    ```shell
    go run .
    ```

4. **Query the News Endpoint**: To test the tracing capabilities, query the news endpoint using cURL or any HTTP client. For example:
    ```shell
    curl http://localhost:3333/news
    ```

5. **View Traces**: Traces for the query to the /news endpoint and interactions with the database should be visible on the Jaeger web interface. You can access the Jaeger interface at http://localhost:16686/.


## Implementation Details

For a more in-depth understanding of how tracing is implemented within the application, you can refer to the following key components:

- **TracedStore (store.go)**: This component contains the implementation of the traced data store. It provides an example of how to wrap a store with Tracing.

- **telemetry/telemetry.go**: This component is responsible for the setup of a tracing provider to export traces to Jaeger.
