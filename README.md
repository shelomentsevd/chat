# Chat

# Build

`
make build
`

# Routes

## Registration

POST /api/v1/registration

`
 curl -v -X POST -F 'name=john@doe.com' -F 'password=password' http://localhost:3000/api/v1/registration
`

## Authentication

Basic

## Create chat
`
curl -H "Authorization: Basic am9obkBkb2UuY29tOnBhc3N3b3Jk" -H "Content-Type: application/vnd.api+json" -X POST -d '{"data":{"attributes":{"name":"Уютный чатик"},"id":"10","relationships":{"users":{"data":[{"id":"5","type":"users"},{"id":"4","type":"users"}]}},"type":"chats"}}' http://localhost:3000/api/v1/chats/
`
## Get chat list
`
curl -v -GET -H "Authorization: Basic am9obkBkb2UuY29tOnBhc3N3b3Jk" http://localhost:3000/api/v1/chats/
`
## Get user list in chat
`
curl -v -GET -H "Authorization: Basic am9obkBkb2UuY29tOnBhc3N3b3Jk" http://localhost:3000/api/v1/chats/1/users
`
## Get message list in chat
`
curl -v -GET -H "Authorization: Basic am9obkBkb2UuY29tOnBhc3N3b3Jk" http://localhost:3000/api/v1/chats/1/messages
`
## Send message to chat
`
curl -v -H "Authorization: Basic am9obkBkb2UuY29tOnBhc3N3b3Jk" -H "Content-Type: application/vnd.api+json" -X POST -d '{"data":{"attributes":{"content":"Heldlo everyone!"},"type":"messages"}}' http://localhost:3000/api/v1/chats/1/messages
`
## Join chat
`
curl -X POST -H "Authorization: Basic am9obkBkb2UuY29tOnBhc3N3b3Jk" http://localhost:3000/api/v1/chats/1/join
`
## Leave chat
`
curl -X POST -H "Authorization: Basic am9obkBkb2UuY29tOnBhc3N3b3Jk" http://localhost:3000/api/v1/chats/1/leave`