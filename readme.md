# Project
go project with grapqh
```
cmd
   ---> server: main program
internal
  |---> model: model of service examlple users, course, etc
   ---> users: internal service of users (endpont, service, repository)
pkg
   ---> graphql: graphql packages and configurations
```

# Example

```sh
cd pkg/graphql
go run server.go
```

```graphql
GraphQL in Insomnia
POST method

# create User
mutation createUser {
  createUser(input: { name: "Nahuel", email: "nlcostamagna@gmail.com", phone: "12212312", addressId: "123" }) {
		id
    name
		email
  }
}

# get User
query findUsers {
  users {
    id
    name
		email
  }
}

```

# Update Schema

```sh
cd pkg/graphql
go run github.com/99designs/gqlgen generate
```