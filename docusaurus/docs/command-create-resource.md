---
id: command-create-resource
title: create resource
sidebar_label: create resource
---

## Command
`csctl create resource <name> -f ./manifest`

Creates a `resource` you can then deploy to your Containership cluster or environments. A `resource` is simply one or more Kubernetes resources associated with a `name` to be referenced in deployments.

## Parameters
* name - name of the resource to be created

## Arguments
* file (-f) - Path to manifest file(s) to be associated with resource

## Examples
```
> csctl create resource redis -f ./redis.yaml

Succesfully created resource `redis`
```
