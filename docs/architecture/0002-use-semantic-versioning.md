# 2. Use Semantic Versioning

Date: 2021-05-08

## Status

Accepted

## Context

Terraform providers are expected to follow [Semantic Versioning](http://semver.org/), according to the [Terraform Best Practices](https://www.terraform.io/docs/extend/best-practices/versioning.html#versioning-specification).

> Observing that Terraform plugins are in many ways analogous to shared libraries in a programming language, we adopted a version numbering scheme that follows the guidelines of Semantic Versioning. In summary, this means that with a version number of the form `MAJOR`.`MINOR`.`PATCH`, the following meanings apply:
>
> * Increasing only the patch number suggests that the release includes only bug fixes, and is intended to be functionally equivalent.
> * Increasing the minor number suggests that new features have been added but that existing functionality remains broadly compatible.
> * Increasing the major number indicates that significant breaking changes have been made, and thus extra care or attention is required during an upgrade.
>
> Version numbers above `1.0.0` signify stronger compatibility guarantees, based on the rules above. Each increasing level can also contain changes of the lower level (e.g., `MINOR` can contain `PATCH` changes).

## Decision

We will follow this advice.

## Consequences

To be clarified later.
