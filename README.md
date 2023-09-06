# Location API

Provides GraphQL APIs to manage locations

## Development and Contributing

- [Development Guide](docs/development.md)
- [Contributing](https://infratographer.com/community/contributing/)
- [NATS](https://nats.io/)

### Getting Started

To begin start by opening the devcontainer as outlined in the [Development Guide](docs/development.md)

**To initialize the database:**
1. `go build`
1. `./location-api migrate up`

**To run the api**
```sh
make go-run
```

**Interacting with the GraphQL queries:**
> Go to localhost:XXXX/playground

## Example GraphQL Queries

### Create Location Mutation

Input:
```graphql
mutation {
    locationCreate(input:{
        name: "DA",
        ownerID: "testtnt-123456",
        description: "DA Metro"
    }) {
        location{
            id
        }
    }
}
```

Output:
```json
{
    "data": {
        "locationCreate": {
            "location": {
                "id": "lctnloc-NW_5CWPx3xeOuh4Ak2WgS"
            }
        }
    }
}
```