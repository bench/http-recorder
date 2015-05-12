<img src="http://www.ucd.ie/building/resource/buttons/beta.gif" alt="beta" style="width: 150px;align:center;"/>

# Http-recorder

Http-recorder can be used for transparently recording all requests sent by a system via [HTTP](http://tools.ietf.org/html/).

It provides :
* an `HTTP endpoint` that accepts any kind of requests on any path, any method, any headers and any payload
* an `API` to retrieved stored requests using a specific quey syntax (see above).

Http requests are stored in a FIFO implemented in an LRU way.

## Usage

Clone project

    git clone https://github.com/BenC-/http-recorder.git

Build project

    cd http-recorder
    make install

Run http-recorder

     bin/http-recorder -monitorePort 12345 -retrieverPort 23456

Start testing !

## Query syntax

*host:23456?containspath=youyou : return the 

TO COMPLETE

## Contributing

The project is developed in [Go](http://golang.org/).

See the [Golang documentation](https://golang.org/doc/) for more information about the language.

If you would like to submit pull requests, please feel free to apply.

## Dependencies

* Golang
* Make 
