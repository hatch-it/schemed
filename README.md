# Schemed

 > Hatch a plan with your friends, effortlessly.

[![Go Report Card](https://goreportcard.com/badge/github.com/hatch-it/schemed)](https://goreportcard.com/report/github.com/hatch-it/schemed) [![GoDoc](https://godoc.org/github.com/hatch-it/schemed?status.svg)](https://godoc.org/github.com/hatch-it/schemed)

Schemed is an online service that aims to make it easier for people to hangout without repetitive communication.

## API Overview

The Schemed API is a *RESTful* web service, which operates upon the idea of **resources** and **CRUD**.

### Resources

Resources are representations of the data that you'll be using from our API.
Here are the resource types that will be available to you:
  - Users
  - Events
  - Venues

### Operations

There are four basic operations which make up the acronym CRUD.
Each operation is accessible through HTTP by sending requests to `https://schemed.io/api`.

Operation | Request method and URL | Response (in JSON)
--------- | ---------------------- | ------------------
Create    | `POST /resource`       | Created resource with ID
Read      | `GET /resource`        | Requested resources
Update    | `PATCH /resource/id`   | Updated resource
Delete    | `DELETE /resource/id`  | Deleted resource

A more detailed description of REST operations can be found [here](http://www.restapitutorial.com/lessons/httpmethods.html).

### Examples

#### JavaScript

Create an Event
```js
const response = await fetch('https://schemed.io/api/events', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
  },
  body: JSON.stringify({
    title: 'End of the world party',
    place: 'Heaven',
    host: 'Jesus',
  }),
})
const event = await response.json()
const { id, title, place, host } = event
```

To retrieve an Event by ID
```js
const response = await fetch('https://schemed.io/api/events/123')
const event = await response.json()
const { id, title, place, host } = event
```

Update an Event by ID
```js
const response = await fetch('https://schemed.io/api/events/123', {
  method: 'PATCH',
  headers: {
    'Content-Type': 'application/json',
  },
  body: JSON.stringify({
    title: 'End of the world party',
    place: 'Hell',
    host: 'Satan',
  }),
})
const event = await response.json()
const { id, title, place, host } = event
```

Delete an Event by ID
```js
const response = await fetch('https://schemed.io/api/events/123', {
  method: 'DELETE',
})
const event = await response.json()
const { id, title, place, host } = event
```

## How to contribute

Clone this repo then use [Docker](https://www.docker.com/) to get it running.
 > Note: Make sure you have your `GOPATH` all set up.

```shell
go get github.com/puradox/schemed
cd $GOPATH/src/github.com/puradox/schemed
docker-compose up
```

Be sure to use [conventional commit messages](https://conventionalcommits.org/) when contributing back to this repo.

Happy hacking!
