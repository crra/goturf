# goTurf - TurfJS like library for golang

goTurf is a (partial) port of the popular [TurfJS][turfjs] library to golang.

It's build upon [go.geojson] (be sure to check out [orb]) and just offers convenience functions with the semantics of TurfJS.

## (partially) implemented packages

- circle
- bbox
- random
- meta

## Run tests

```
go test -coverprofile cover.out
go tool cover -html=cover.out -o testcoverage.html
```

Open the file `testcoverage.html` in your browser

[go.geojson]: http://github.com/paulmach/go.geojson
[orb]: https://github.com/paulmach/orb
[turfjs]: https://github.com/Turfjs/turf
