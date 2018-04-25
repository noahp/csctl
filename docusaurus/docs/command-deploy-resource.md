---
id: command-deploy-resource
title: deploy resource
sidebar_label: deploy resource
---

## Command
`csctl deploy <resource_name> <...cluster_id>`

`csctl deploy <resource_name> --cluster-label key=val`

Deploys the given `resource_name` to one or more Containership clusters. The `deploy` command supports to options for targeting clusters.

1. Explicitly pass one or more `cluster_ids`
2. Specify one or more `--cluster-label` flags that match `labels` you have previously set on your Containership Clusters.

> All of the cluster labels are logically `ANDed` together

## Parameters
* resource_name - name of the resource to be deployed
* cluster_id - Containership Cluster ID of the cluster to deploy to

## Arguments
* cluster-label (-l) - Key-value cluster label to match against when determining whether or not to deploy to a cluster
* namespace (-n) - Namespace to deploy the resource to (overrides any existing namespace specified in the manifest of the resource)
* interactive (-i) - Runs the deployment in interactive mode, allowing you to specify which clusters to deploy to


## Examples
```
# example deploying to cluster ID
> csctl deploy redis 2149a670-e184-4bfe-b200-60a94035512c

Successfully deployed `redis` to cluster 2149a670-e184-4bfe-b200-60a94035512c

# example deploying with cluster-labels
> csctl deploy redis -l team=ui -l env=qa

Successfully deployed redis to the following clusters:
26eb6caf-cc1d-42e6-80c1-5b256e94648b
2149a670-e184-4bfe-b200-60a94035512c
```
