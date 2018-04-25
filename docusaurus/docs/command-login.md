---
id: command-login
title: login
sidebar_label: login
---

## Command
`csctl login <method>`

Login to Containership Cloud based on your account's authentication `method`. Based on the `method` chosen, you will then be prompted for more information in order to authenticate your user.

## Parameters
* method - type of authentication method (email, github, bitbucket, auth0)

## Arguments
* username (-u) - Username for authentication method
* password (-p) - Password for authentication method

## Examples
```
> csctl login github

> Please Enter Your Username: containershipbot
> Please Enter Your Password: **********
> Please Enter your 2FA Code: ******

You have been successfully logged in!
```
