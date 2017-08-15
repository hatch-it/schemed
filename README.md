# Schemed

 > Hatch a plan with your friends, effortlessly.

Schemed is an online service that aims to make it easier for people to hangout without repetitive communication.

## Getting started

Clone this repo then use [Docker](https://www.docker.com/) to get it running.
 > Note: `docker-compose` is automatically installed with [_Docker for Mac_](https://www.docker.com/docker-mac)
 > and [_Docker for Windows_](https://www.docker.com/docker-windows)

```shell
go get github.com/puradox/schemed
cd $GOPATH/src/github.com/puradox/schemed
docker-compose up
```

## API

The Schemed API is a RESTful web service, which you can learn more about at [REST API Tutorial](http://www.restapitutorial.com/).

For example,
```shell
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"username":"xyz","password":"xyz"}' \
  http://localhost:3000/api/login

```

### /api/_resource_
 > Where _resource_ can be: **users**, **events**, **venues**

 - GET /_resource_ - Fetches all the matching instances of _resource_
 - GET /_resource_/:id - Get the instance of _resource_ with the specified ID
 - POST /_resource_ - Create an instance of _resource_
 - POST /_resource_/:id - Update the instance of _resource_ with the specified ID
 - DELETE /_resource_/:id - Soft delete the instance of _resource_ with the specified ID

