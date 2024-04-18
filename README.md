# go-web-setup
Rapid development with Golang setup template, fortified with Echo, sqlx, sql-migrate, and zap⚡️ promises an extraordinary journey for your next project.

# Includes 
- User Schema
- Default APIs ( Register, Login, CRUD )
- JWT Authentication 


# Run

    git clone go-web-setup
    go get
    
    // setup .env refer .env.example
    source .env

    docker-compose up -d
    go run cmd/prod/main.go

# layer

    ├── app
    │   ├── dependencies.go
    │   ├── logger
    │   ├── model
    │   ├── repository
    │   ├── routes
    │   └── services
    ├── cmd
    │   ├── migrate
    │   ├── prod
    │   └── staging
    ├── config
    ├── logs
    └── migrations

# TODO
- Update readme
- Test case
- Queue management ( horizon )
- CSRF support