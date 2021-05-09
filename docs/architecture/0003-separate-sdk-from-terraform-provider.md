# 3. Separate SDK from Terraform Provider

Date: 2021-05-08

## Status

Accepted

## Context

Terraform recommends that providers not directly interact with the target API, but go through a layer of abstraction via a Go SDK. While the prototypes of a Go SDK [do exist](https://github.com/moltin/go-epcc-client), it is not under active development and additionally the work to implement resources would introduce a fair bit of complexity to the build and development process. 

## Decision

We will structure the code along the lines of an SDK and provider separator, and keep a clean separation internally, however everything will be done in one repository.

## Consequences

Building and developing the project in the short term will be easier, it will be harder to keep the layers and dependencies between the projects isolated without additional automation.