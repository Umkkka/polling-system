# Polling system

## Description

Service in Go that enables "The Decisives" to create polls, vote in real-time, and view poll results as they come in, keeping the debate lively and engaging

## Questions

When developing this API, I made the following **key assumptions**:

- Users will create polls with the ability to select only one answer option from those provided.
- When attempting to vote in a non-existent poll or for a non-existent answer option in a poll, the API should notify the user of this with the corresponding error.
- Websocket connection will be used for displaying poll results, as it allows real-time results display.

In the process of developing the API, I encountered a number of **compromises**:

- Data storage. Poll data and results are stored in maps, but they can easily be replaced with tables in a traditional relational database.
- Access to results and information about polls is organized in a way that only one process can work with this data at any given time. This slows down the API's performance, but ensures consistency and data integrity.

If this project were to evolve into a **full-scale application for the real world**, here are some enhancements or next steps that could be prioritized to improve its functionality, usability, and technical reliability:

- Database Integration: Transition from using in-memory maps to a robust database system for storing poll data and results. This will ensure scalability, data persistence, and efficient data management.

- User Authentication and Authorization: Implement a secure user authentication system to allow only authorized users to create, vote on, and view polls. This will enhance security and protect user data.

- Real-Time Notifications: Implement real-time notifications to inform users of new polls, updates on polls they have participated in, and live result updates without the need to manually refresh the page.

- Performance Optimization: Conduct performance optimizations such as caching mechanisms, load balancing, and scalability measures to handle increasing traffic and maintain fast response times.

- User Feedback and Analytics: Collect user feedback to continuously improve the application based on user preferences and behaviors, and utilize analytics to track user engagement and identify areas for enhancement.

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
