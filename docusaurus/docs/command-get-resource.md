---
id: command-get-resource
title: get resource
sidebar_label: get resource
---

## Command
`csctl get resource [resource_name]`

Gets one or more `resources`.

## Parameters
* resource_name - name of the resource to retrieve

## Examples
```
> csctl get resource

RESOURCE            VERSIONS    ENVIRONMENTS
myredisresource     3           qa
myotherresource     2           dev,qa,...
```
