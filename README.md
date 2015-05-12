# Http-recorder

Http-recorder can be used for transparently recording all requests sent by a system via [HTTP](http://tools.ietf.org/html/).

It provides :
* an `HTTP endpoint` that accepts any kind of requests on any path, any method, any headers and any payload
* an `API` to retrieved stored requests using a specific quey syntax (see above).

![beta](http://leitorcabuloso.com.br/wp-content/uploads/2013/01/beta.jpg 200px)

## Usage

Clone project

    git clone https://github.com/BenC-/http-recorder.git

Build project

    make install

Run http-recorder

     cd bin
     ./http-recorder -monitoredPort 12345 -retrieverPort 23456

Start testing !

## Contributing

The project is developed in [Go](http://golang.org/).

See the [Golang documentation](https://golang.org/doc/) for more information about the language.

If you would like to submit pull requests, please feel free to apply.

## Dependencies

* Golang
* Make 
