---
id: command-get-resource-version
title: get resource-version
sidebar_label: get resource-version
---

## Command
`csctl get resource-version [resource_name]`

Gets one or more `resource-version`.

## Parameters
* resource_name - name of the resource to retrieve resource-versions from

## Arguments
* --version (-v) - specific version of the resource to retrieve
* --json (-j) - output the full data contents in json

## Examples
```
> csctl get resource-version myredisresource

ID                                      VERSION     KUBERNETES_RESOURCES
26eb6caf-cc1d-42e6-80c1-5b256e94648b    1           3
2149a670-e184-4bfe-b200-60a94035512c    2           2

> csctl get resource-version myredisresource --version 2

{
    id: 26eb6caf-cc1d-42e6-80c1-5b256e94648b,
    version: 2,
    kubernetes_manifests: [
        ...
    ]
}
```
