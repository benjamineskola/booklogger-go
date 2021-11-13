#!/bin/sh
set -e
golangci-lint run
courtney -e
go tool cover -func=coverage.out
