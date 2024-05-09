# Polling system

## Description

Service in Go that enables "The Decisives" to create polls, vote in real-time, and view poll results as they come in, keeping the debate lively and engaging

## Questions

In solving this problem, an API with minimal functionality was designed to solve the tasks set:
- poll creation
- voting
- live results

It is worth noting that the Go language map was chosen as the data warehouse, which can be easily replaced with a real database. To do this, a separate layer was implemented for working with the repository layer.

In addition, the mechanism for real-time notification of survey results by sending messages to websockets has not been fully implemented. However, the intended functionality is demonstrated in the API logs

The following can be highlighted as further steps to expand this project:

- Creation and implementation of a database in the project
- Creation web interface for the ability to display the results of polls in real time and vote
- Additional thinking over the functionality of working in parallel mode under high load

## API

API contains three methods for working with surveys and a websocket connection for real-time notification of survey results:

#### Create Poll

Method allows you to create a poll. The required options are the name of the poll and the list of voting options in the poll. If the poll is successfully created method return uuid of the poll.

```shell
$ curl --location 'localhost:8080/api/v1/poll' \
--header 'Content-Type: application/json' \
--data '{
    "title": "Poll",
    "options": [
        "yes",
        "no"
    ]
}'
```

#### Get Poll

Method allows you to get information about the survey by the uuid. If no polls are found for the specified identifier, the method returns an error.

```shell
$ curl --location --request GET 'localhost:8080/api/v1/poll?uuid=7a2b8e58-2694-4201-84b3-415b92dcbbe5' \
--header 'Content-Type: application/json' \
--data '{
    "uuid": "ef70b1c6-cc45-441d-9b39-187b34ea6225"
}'
```

#### Vote

Method allows you to vote in an existing poll. If the survey does not exist or the response is not present in the survey options, the method returns an error. If the voting result is successfully saved, updated information about the results of a particular survey is transmitted to the websockets 

```shell
$ curl --location --request POST 'localhost:8081/api/v1/vote?uuid=7a2b8e58-2694-4201-84b3-415b92dcbbe5&answer=yes'
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
