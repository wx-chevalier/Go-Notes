#!/usr/bin/env bash
mockgen -destination mocks_test.go -package dbinterface github.com/agtorre/go-cookbook/chapter5/dbinterface DB,Transaction
