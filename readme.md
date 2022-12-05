# Example

For the time being

```sh
cd pkg/graphql
go run go run server.go
```

```graphql
GraphQL in Insomnia
POST method

# create User
mutation createUser {
  createUser(input: { name: "Nahuel", email: "nlcostamagna@gmail.com" }) {
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

# create Todo
mutation createTodo {
  createTodo(input: { text: "todo", userId: "T45" }) {
    user {
      id
			name
    }
    text
    done
  }
}
```


todo:
 - add gokit in the project