# newrelic-demo

[![Build Status](https://travis-ci.org/Huang-Wei/newrelic-demo.svg?branch=master)](https://travis-ci.org/Huang-Wei/newrelic-demo)

This is a demo trying to showcase how to use expose prometheus-like metric/event using newrelic API.

## Running locally

Ensure to set following env variables:

- NEW_RELIC_LICENSE_KEY - `export NEW_RELIC_LICENSE_KEY=979axxxxxxxxxxxxxxxxxxxxxxxxx`
- NAMESPACE - `export NAMESPACE=default`
- CLUSTER - `export CLUSTER=$(hostname)`)