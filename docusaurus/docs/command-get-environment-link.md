---
id: command-get-environment-link
title: get environment-link
sidebar_label: get environment-link
---

## Command
`csctl get environment-link [env_name]`

Gets one or more `environment-links`.

## Parameters
* env_name - name of the environment to retrieve environment-links from

## Examples
```
> csctl get environment-link

ID                                      ENVIRONMENT     APPS    LINKS
26eb6caf-cc1d-42e6-80c1-5b256e94648b    dev             3       1
2149a670-e184-4bfe-b200-60a94035512c    qa              2       1

> csctl get environment-link qa

ID                                      ENVIRONMENT     APPS    LINKS
2149a670-e184-4bfe-b200-60a94035512c    qa              2       1
```
