---
id: overview
title: Overview
sidebar_label: Overview
---

## Overview

The Containership Control Client allows you to easily label clusters, define `resources` and deploy your resources to one or more Containership clusters.

## Getting Started

Login to your Containership Cloud account with your authentication method:

```
# csctl login <method>
> csctl login github

> Please Enter Your Username: containershipbot
> Please Enter Your Password: **********
> Please Enter your 2FA Code: ******

You have been successfully logged in!@
```

> See [command-login](command-login) for more information on the available options.

Once you are logged-in you can retrieve inforation from your Containership Cloud account such as the `attached` clusters.

```
> csctl get clusters

ID                                      NAME            ORGANIZATION    LABELS
26eb6caf-cc1d-42e6-80c1-5b256e94648b    cluster1        sample-org      team=ui,env=dev
2149a670-e184-4bfe-b200-60a94035512c    cluster2        sample-org      team=api,env=dev
```

In order to deploy resources to different clusters you need to first create the `resource` and also create an `environment`

```
> csctl create resource myredisapp -f ./redis.yaml

Successfully created resource `myredisapp`

> csctl create environment dev

Successfully created environment `dev`
```

You cannot promote that app to the environment quite yet, however. If you try you will see that the environment does not know about any target clusters to deploy to!

```
> csctl promote myredisapp --to dev

Failed to promote `myredisapp` to environment `dev`.
There were no associated clusters linked with the environment.
```

You can easily associate clusters with and `environment` by creating one or more `links`. The following command creates a link to all clusters with the `env` label set to `dev`.

```
> csctl create environment-link dev -l env=dev

Successfully created environment link on the `dev` environment.
```

Now when you promote resources to that environment, all associated clusters will be deployed with the `resource`.

```
> csctl promote myredisapp --to dev

myredisapp has been successfully promoted to the `dev` environment.
It has been deployed to the following clusters:

cluster1
cluster2
```

> All of our commands support the `--dry-run` flag so you can see what modfiications will take place before you run them.

Take a look at each of the individual command documentation pages for more advanced information!

## Coming Soon

* Integration with Helm for your defined `resources` and `environments`
* You will soon be able to automate the provisioning of your infrastructure through the Containership Cloud Control
* Support for the upcoming `cluster federation` standard being worked on in the Kubernetes community
