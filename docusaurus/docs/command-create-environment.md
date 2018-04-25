---
id: command-create-environment
title: create environment
sidebar_label: create environment
---

## Command
`csctl create environment <name>`

Creates an `environment` you can associate with Containership clusters and labels. An `environment` is simply a target you can promote your resources between. You can also optionally specify a default `namespace` that an environment is associated with.

> Creating an `environment` is not all that useful unless you link containership `clusters` or `labels` with it. See [command-environment-link](command-environment-link) for more information.

## Parameters
* name - name of the environment to be created

## Arguments
* namespace (-n) - Default Namespace the environment applies to

## Examples
```
> csctl create environment dev -n development

Succesfully created environment `dev`
```
