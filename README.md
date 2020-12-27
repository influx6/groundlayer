# Groundlayer

[![GoDoc](https://img.shields.io/badge/api-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/influx6/groundlayer)

Opinionated web application library with a twist of disappearing templates.

Groundlayer combines the best ideas into a small but powerful library capable of building
performing server-rendered applications with ease. It's the culmination of all ideas running 
through my head for years, each learnt through playing with different technologies.


## Getting

```
go get -u github.com/influx6/groundlayer
```

## Using

Create a sample `hello` web application:

```bash
groundlayer new -package=github.com/user/hello
```

Then run the application:

```bash
cd hello && go run main.go start --env-file=.env.dev
```


## Contact

Ewetumo Alexander [@influx6](http://twitter.com/influx6)

## License

Source code is available under the MIT [License](/LICENSE).
