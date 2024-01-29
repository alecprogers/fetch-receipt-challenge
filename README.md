# Alec Rogers | Fetch Receipt Challenge

## Objective

This project is my solution to the Fetch Receipt Processor Challenge. I elected to use Go for this project, since that was stated to be Fetch's preferred language. I do not have prior experience with Go, but found it to be straightforward to get up to speed with.

I also decided to design it such that it can be run locally using only Go, or also run using the Serverless framework, which I have used extensively in the past. This required a bit more complexity in terms of project structure, but ended up working smoothly. This is the reason that I have separate `process-receipt` and `process-receipt-handler` packages. `process-receipt` contains the actual business logic to process the receipt, while `process-receipt-handler` just serves as an entrypoint for Serverless. The entrypoint to run locally is the `src/main.go` package.

## Assumptions

I found the requirements to be clear for this project and therefore did not factor any major assumptions into its design.

## Usage

Please begin by cloning this repo and navigating to it in your CLI.

### Running Locally with Go

_Requirements: Go_

To run this project locally using just Go, run the following command:

```bash
go run src/main.go
```

This will serve the application locally on port 8080.

### Running Locally with Serverless

_Requirements: Go, [NPM](https://docs.npmjs.com/downloading-and-installing-node-js-and-npm), [Serverless](https://www.serverless.com/framework/docs/getting-started), Docker_

Serverless offers a couple of methods for running the project locally. To invoke the lambda locally, run:

```bash
npm install
make build
serverless invoke local --function process-receipt --path test-receipt.json
```

Substitute the payload of your choice into `test-receipt.json`.

To serve the lambda as a local API, run:

```bash
npm install
make build
serverless offline --useDocker
```

I found this method to be quite slow when using Go, so I would not recommend it. You will likely have to increase the function timeout defined in `serverless.yml` in order for it to finish.

### Calling in AWS

_Requirements: None_

The primary benefit of using Serverless is the ease with which it allows setting up a basic REST API such as this one using AWS API Gateway and Lambda. It is currently deployed and available at the following endpoint:

```
https://e4l9b5t9z9.execute-api.us-east-2.amazonaws.com/receipts/process
```

Note that I only configured the `receipts/process` endpoint with Serverless due to the lack of persistent data storage. As a result, when calling this API locally with Serverless or directly in AWS, it will return both `id` and `points`. When running locally with Go, `receipts/process` will only return `id`, per the API spec provided. It stores the data in memory and will return `points` from the `receipts/{id}/points` endpoint.
