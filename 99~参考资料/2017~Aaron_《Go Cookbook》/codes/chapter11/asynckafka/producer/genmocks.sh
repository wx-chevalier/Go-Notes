#!/usr/bin/env bash
mockgen -destination mocks_test.go -package main gopkg.in/Shopify/sarama.v1 AsyncProducer
