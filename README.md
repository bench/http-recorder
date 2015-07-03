<img src="http://www.ucd.ie/building/resource/buttons/beta.gif" alt="beta" style="width: 150px;align:center;"/>

# Http-recorder

Http-recorder can be used for transparently recording all requests sent by any system via [HTTP](http://tools.ietf.org/html/).

It is made of :
* An `HTTP endpoint` that accepts any kind of HTTP requests
* A `API` to retrieve stored requests using a specific query syntax (see above).

## Usage

Start recorder

     ./http-recorder -recorderPort 80 -retrieverPort 8080

Send any kinf of HTTP request on recorder port

![Alt text](https://cloud.githubusercontent.com/assets/3688186/7613711/a88451b4-f992-11e4-8043-f58fa74c4c73.png "any request")

Pop stored request(s) by querying them on retriever port, so simple !

![Alt text](https://cloud.githubusercontent.com/assets/3688186/7613728/bc818812-f992-11e4-9e57-5190d38dc2a6.png "query request")


## Query syntax


#### by pathcontains
Return the oldest request whose path contains 'banana'

    http://hostname:23456?pathcontains=banana

#### by bodycontains
Return the oldest request whose body contains 'store'

    http://hostname:23456?pathcontains=store

#### by contenttype
Return the oldest request whose content type header is 'text/plain'

    http://hostname:23456?contenttype=text%2Fplain

#### by method
Return the oldest request whose method is 'PUT'

    http://hostname:23456?method=put


NB : Query syntax ignores case.	

## Build and install

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
