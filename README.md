# card-rest 
### This project implements the REST API for "Create a new Deck", "Open a deck" and "Draw Cards"

### Requirements:
[Docker](https://www.docker.com/products/docker-desktop)
Go 1.16 or higher
[Golang-migrage](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

### Usage
Clone the repository with:
`git clone https://github.com/haiqicun/card-rest.git`

Start the server and Postgresql database in docker with:
`make start-server`

Migrate up the database with:
`make migrate-up`

Run the tests with:
`make test`

Stop the server in docker with:
`make stop-server`

### API Endpoints

#### Create a new Deck:
1.
`curl -X POST http://localhost:8080/deck -H "Content-type: application/json" -d '{ "shuffled": true}'`

The response is: 
```
{
  "deck_id":"8505959b-e834-44f1-a89f-578d0d608040",
  "shuffled":true,
  "remaining":52
}
```
2.
`curl -X POST http://localhost:8080/deck -H "Content-type: application/json" -d '{ "shuffled": false, "codes":"8S,6D,KH,10C"}'`

The response is:
```
{
  "deck_id":"f926bb84-375a-4021-b0bb-09a1706a44d3",
  "shuffled":false,
  "remaining":4
}
```

#### Open a Deck

`curl http://localhost:8080/deck/"f926bb84-375a-4021-b0bb-09a1706a44d3"`

The response is:
```
{
  "deck_id":"f926bb84-375a-4021-b0bb-09a1706a44d3",
  "shuffled":false,"remaining":4,
  "cards":[
          {"value":"8","suit":"SPADES","code":"8S"},
          {"value":"6","suit":"DIAMONDS","code":"6D"},
          {"value":"KING","suit":"HEARTS","code":"KH"},
          {"value":"10","suit":"CLUBS","code":"10C"}
          ]
}
```

### Draw cards of a Deck

`curl -X PUT http://localhost:8080/deck/"f926bb84-375a-4021-b0bb-09a1706a44d3"/draw/"2"`

The response is:
```
{
   "deck_id":"f926bb84-375a-4021-b0bb-09a1706a44d3",
   "shuffled":false,"remaining":2,
   "drawncards":[
                {"value":"8","suit":"SPADES","code":"8S"},
                {"value":"6","suit":"DIAMONDS","code":"6D"}
                ]
 }
```
The "drawncards" are the cards removed from the deck.

## TODO:
Add more unit tests and integration tests to test the router.
