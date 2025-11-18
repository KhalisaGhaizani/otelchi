# otelchi

This fork adds "status code" and "outcome" attributes for mterics.
the "outcome" attribute allows to easily filter by "success" or "failure".

[![compatibility-test](https://github.com/KhalisaGhaizani/otelchi/actions/workflows/compatibility-test.yaml/badge.svg)](https://github.com/KhalisaGhaizani/otelchi/actions/workflows/compatibility-test.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/KhalisaGhaizani/otelchi)](https://goreportcard.com/report/github.com/KhalisaGhaizani/otelchi)
[![Documentation](https://godoc.org/github.com/KhalisaGhaizani/otelchi?status.svg)](https://pkg.go.dev/mod/github.com/KhalisaGhaizani/otelchi)

OpenTelemetry instrumentation for [go-chi/chi](https://github.com/go-chi/chi).

Essentially this is an adaptation from [otelmux](https://github.com/open-telemetry/opentelemetry-go-contrib/tree/main/instrumentation/github.com/gorilla/mux/otelmux) but instead of using `gorilla/mux`, we use `go-chi/chi`.

Currently, this library can only instrument traces and metrics.

Contributions are welcomed!

## Install

```bash
$ go get github.com/KhalisaGhaizani/otelchi
```

## Examples

See [examples](./examples) for details.

## Why Port This?

I was planning to make this project as part of the Open Telemetry Go instrumentation project. However, based on [this comment](https://github.com/open-telemetry/opentelemetry-go-contrib/pull/986#issuecomment-941280855) they no longer accept new instrumentation. This is why I maintain this project here.
