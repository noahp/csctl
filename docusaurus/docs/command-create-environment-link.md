---
id: command-create-environment-link
title: create environment-link
sidebar_label: create environment-link
---

## Command
`csctl create environment-link <env_name> <...cluster_ids>`

`csctl create environment-link <env_name> --cluster-label key=value`

Links an `environment` to Containership Clusters in one of two ways.

1. Explicitly pass one or more `cluster_ids`
2. Specify one or more `--cluster-label` flags that match `labels` you have previously set on your Containership Clusters.

You can also optionally specify a namespace that the given link is associated with.

> Explicit `cluster_ids` take precedence. If you create multiple links for an environment, it will first process all the `cluster_ids` before moving on to `cluster-labels`.

## Parameters
* name - name of the environment to create a link on
* cluster_ids - one of more clusters to associate the link with

## Arguments
* cluster-label (-l) - Cluster label to be used in the linking
* namespace (-n) - Default Namespace the environment link applies to

## Examples
```
> csctl create environment-link dev 2149a670-e184-4bfe-b200-60a94035512c -n development

Succesfully created new link in the `dev` environment

> csctl create environment-link dev -l team=ui -l env=development

Succesfully created new link in the `dev` environment
```
