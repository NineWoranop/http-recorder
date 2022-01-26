# http-recorder
This is a application for record http response as static files.


The current goals/features that are:

- export metrics as prometheus format to be pain text file (dat file)
- work as period (default - collect metrics every 10 seconds)

# Architecture
![Architecture diagram][architecture]

[architecture]: document/architecture.png "Architecture Diagram"

### Building from source

To build http-recorder from source code, You require:
* Go [version 1.17 or greater](https://golang.org/doc/install).
* Git

Build

    $ git clone https://github.com/NineWoranop/http-recorder.git
    $ go build main.go

Run

    $ http-recorder -url=http://localhost:9090/metrics

## Command Arguments

|Argument           |  Default                    | Description           |
|-------------------|:---------------------------:|:----------------------|
|cert               |                             |client certificate file|
|key                |                             |client certificate's key file|
|accept-invalid-cert|false                        |Accept any certificate during TLS handshake. Insecure, use only for testing|
|scrape-interval    |10s                          |Scrape interval to fetch metrics and write dat file|
|path               |                             |Path for write dat file|
|total-dat-file     |1                            |Number of dat files to write|
|url                |http://localhost:9090/metrics|Url endpoint to fetch metrics|

## Examples
#####1.Write/Replace single file at current folder (Expected to write 000001.dat file)

    $ http-recorder -autorepeat=false

#####2.Automatically write 3 data files at "/data/" folder

    $ http-recorder -path=/data/ -total-dat-file=3

#####3.Write 60 data files at current folder on every 2 seconds and run on port 8080 with any ip addresses

    $ http-recorder -scrape-interval=2s -total-dat-file=60 -url=http://localhost:9090/metrics

## License

Unlicense, see [LICENSE](https://github.com/NineWoranop/http-recorder/blob/master/LICENSE).
