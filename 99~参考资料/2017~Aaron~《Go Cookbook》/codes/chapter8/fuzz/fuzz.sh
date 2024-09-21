#!/usr/bin/env bash
go-fuzz-build github.com/agtorre/go-cookbook/chapter8/fuzz
go-fuzz -bin=./fuzz-fuzz.zip -workdir=output
