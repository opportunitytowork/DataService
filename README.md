# DataService

Service between API and DB. (Local DB for development)

## How does it work

Acts as a handler between the API and the DB. A config instructs the data service which DB to interact with (local, remote). The data service then acts as a middle man between the API and DB

## Testing

We should be able to spin up a local MySQL Container and use that to interact with the DB
