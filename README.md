# Polling system

## Description

Service in Go that enables "The Decisives" to create polls, vote in real-time, and view poll results as they come in, keeping the debate lively and engaging

## API

Example request
```shell
$ curl --location 'localhost:8080/ping'
```

## Local run

### Installation

```shell
$ go mod download
```

### ENV
Locally fill local.env file with your configuration

#### Launch via IDE GoLand

- Download plugin [EnvFile](https://plugins.jetbrains.com/plugin/7861-envfile)
- Add running config: Edit Configurations -> Add New Configuration -> Go Build
- In configuration specify:
    - Run Kind - Package
    - Package Path - polling-system/cmd/app
    - Check the box "Run After Build"
    - Working Directory
    - Go to the tub "EnvFile"
    - Enable `EnvFile`, click the plus sign below and use keyboard shortcut `Command`+`Shift`+`Dot`
    - Choose local.env file
    - Save

- Start configuration (control+R)
