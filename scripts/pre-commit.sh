#!/bin/sh
set -e
golangci-lint run --enable-all --disable gci,maligned
courtney -e
go tool cover -func=coverage.out
