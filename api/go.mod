module FriendApi

go 1.16

replace InternalUserManagement => ../internal

require (
	InternalUserManagement v0.0.0-00010101000000-000000000000
	github.com/go-chi/chi/v5 v5.0.3
	github.com/lib/pq v1.10.1
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/subosito/gotenv v1.2.0
)
