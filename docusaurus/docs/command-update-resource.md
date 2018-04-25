---
id: command-update-resource
title: update resource
sidebar_label: update resource
---

## Command
`csctl update resource <name> -f ./manifest`

Updates a `resource` based on the given manifest(s). The new configuration will become the `latest` version by default and can then be `promoted` to new environments.

## Parameters
* name - name of the resource to be updated

## Arguments
* file (-f) - Path to manifest file(s) to be associated with resource

## Examples
```
> csctl update resource myredis -f ./redis.yaml

Succesfully updated resource `myredis`
```
