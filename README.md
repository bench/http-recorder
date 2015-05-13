<img src="http://www.ucd.ie/building/resource/buttons/beta.gif" alt="beta" style="width: 150px;align:center;"/>

# Http-recorder

Http-recorder can be used for transparently recording all requests sent by any system via [HTTP](http://tools.ietf.org/html/).

It is made of :
* An `HTTP endpoint` that accepts any kind of HTTP requests
* A `API` to retrieve stored requests using a specific query syntax (see above).

## Usage

Send any HTTP request on Http-recorder

![Alt text](https://cloud.githubusercontent.com/assets/3688186/7613711/a88451b4-f992-11e4-8043-f58fa74c4c73.png "any request")

Query Http-recorder to retrieve stored request(s)

![Alt text](https://cloud.githubusercontent.com/assets/3688186/7613728/bc818812-f992-11e4-9e57-5190d38dc2a6.png "query request")


## Query syntax

#### pathcontains
Return a request in FIFO whose path contains value

    http://hostname:23456?pathcontains=banana

## Install

Clone project

    git clone https://github.com/BenC-/http-recorder.git

Build project

    cd http-recorder
    make install

Run http-recorder

     bin/http-recorder -recorderPort 12345 -retrieverPort 23456


![Alt text](https://cloud.githubusercontent.com/assets/3688186/7613417/e5d9c12c-f990-11e4-81ac-168327735bef.png "http-recorder")


Start testing !


## Under the hood

Http requests are stored in a FIFO implemented in an LRU way, to prevent uncontrolled memory increase.

REST API supports `long polling HTTP` pattern, as proposed by [IETF](https://tools.ietf.org/id/draft-thomson-hybi-http-timeout-00.xml), using Request-Timeout header.

### Roadmap
* Support more complete query syntax
* Provide HTML visualization of stored requests
* Adapt API to make it restful
* Support binary request and store base64 encoded body


## Contributing

The project is developed in [Golang](http://golang.org/).

See the [Golang documentation](https://golang.org/doc/) for more information about the language.

If you would like to submit pull requests, please feel free to apply.

## Dependencies

* Golang
* Make 
