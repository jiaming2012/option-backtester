
Goal is to create a golang backtester.

Data source will be theta data.

Pros:
- API is fast
- reasonably prices compared to competetitors
- extensive data

Con:
- Requires running a java client locally to interact with the server

We will dockerize the client to avoid having to run java on our host machine.

Build the ThetaData client:
``` bash
cd path/to/backtester/thetadata
docker build -t thetadata .
```

# Creating the application
To get started, we create a golang application that will fetch a OHLC of a past AAPL options contract.

## Conventions
- in models folder, only one struct per file. The name of the struct in camel case, should match the name of the file, with underscores

- model names should be as descriptive as possible

- pass a request struct, instead of a list of parameters, it makes it much clearer to read (show an example)

- services folder hold all business logic

See my other post about the benefits of using DTOs.

## Services
We create a function that fetches our data

## Models
Using the thetadata documentation (https://http-docs.thetadata.us/docs/theta-data-rest-api-v2/8mgtkq5bepz0n-ohlc), lets make some models to capture or Request and Response.

I advocate creating new types instead of using strings, as it makes the function signatures and struct definitions so much more descriptive

For example,
type ThetaDataHistOptionOHLCRequest struct {
	EndDate    time.Time   `json:"end_date"`
	Expiration time.Time   `json:"exp"`
	Interval   int         `json:"ivl"`
	Right      OptionType  `json:"right"`
	Root       StockSymbol `json:"root"`
}

is much more descriptive than
type ThetaDataHistOptionOHLCRequest struct {
	EndDate    string   `json:"end_date"`
	Expiration string   `json:"exp"`
	Interval   int         `json:"ivl"`
	Right      string  `json:"right"`
	Root       string `json:"root"`
}

In the example above it is not clear that the Theta data api only expects "C" for call, insteall of "call," "Call" or some variation. Root is immediately known to be a stock symbol. Also the api expects, EndDate and Expiration to be in format, e.g. 20231103


``` bash

docker build -t my_go_app .
docker run -d --name my_go_app_instance my_go_app
```

Create a new app:
``` bash
go mod init github.com/jiaming2012/option-backtester
go mod tidy
```

To run the docker compose file:
``` bash
THETA_DATA_USERNAME="myusername" THETA_DATA_PASSWORD="mypassword" docker-compose up
```