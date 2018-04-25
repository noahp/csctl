---
id: command-promote
title: promote
sidebar_label: promote
---

## Command
`csctl promote [...resources] --from=<from_env> --to=<to_env>`

Promotes one or more resources to an environment. You can specify a `from` environment in order to `promote` the given `versioned` resources in the environment to a new environment.

You can also promote arbitrary `resources` and `versions` to a given `environment` by ommitting the `--from` argument.

## Parameters
* resources - One or more resources to be promoted. Resource version defaults to `latest` if not specified. You can specify a given resource version by adding the version suffix (e.g. myredisresource:4)

> To see a list of your existing resources version history, take a look at the [get-resource-history](get-resource-history) command.

## Arguments
* --from (-f) - Environment to select existing resources from
* --to (-t) - Environment to promote the specified resources to
* --strategy (-s) - Strategy for operations in the case of failure. Options are `rollback-on-failure` or `continue-on-failure`

## Examples
```
> csctl promote myredisresource:5 --to dev

Succesfully promoted myredisresource to the `dev` environment.

> csctl promote myotherresource --to dev

Succesfully promoted myotherresource:2 to the `dev` environment.

> csctl promote --from dev --to qa

The following applications have been promoted from `dev` to `qa`:

myredisresource:5
myotherresource:2

> csctl promote --from dev --to qa --strategy continue-on-failure

The following applications were succesfully promoted:
myredisresource

The following applications failed to deploy:
badresource
```
