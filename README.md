# mypoints-rest-api
[![CI](https://github.com/yossdev/mypoints-rest-api/actions/workflows/CI.yml/badge.svg?branch=main)](https://github.com/yossdev/mypoints-rest-api/actions/workflows/CI.yml)

Backend application for MyPoints WebApp - Final Project Alterra Academy (MBKM)

## PTT Project

[MyPointsPPT](https://docs.google.com/presentation/d/1CknmZ4wCH17ST1ViROIu_vEyyxy4VZTl32VclzEOhjM/edit?usp=sharing)

## APIs Documentation
[MyPoints REST APIs](https://server.mypoints.site/api/v1/swagger/index.html)

## API Server tech-stack:

- Server code: **go1.17**
- REST Server: [**fiber v2**](https://docs.gofiber.io/)
- Database: **PostgreSQL**, **MongoDB**
- ORM: [**gorm**](https://gorm.io/docs/)
- Nginx & Cloudflare

## Third-party API used:
- [xendit](https://www.xendit.co/en-id/) - Payment gateway

## CI/CD

- [github actions](https://github.com/features/actions)
- [docker](https://www.docker.com/)
- [amazon EC2](https://aws.amazon.com/ec2/?ec2-whats-new.sort-by=item.additionalFields.postDateTime&ec2-whats-new.sort-order=desc)

## Other technology

- [air](https://github.com/cosmtrek/air) - ☁️ Live reload for Go apps.
- [xendit](https://github.com/xendit/xendit-go) - Go library from xendit
- [uuid](https://github.com/google/uuid) - Go package for UUIDs based on RFC 4122 and DCE 1.1: Authentication and Security Services.
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) - Package bcrypt implements Provos and Mazières's bcrypt adaptive hashing algorithm.
- [validator v10](https://github.com/go-playground/validator) - Go Struct and Field validation.
- [swaggo](https://github.com/swaggo/swag) - Automatically generate RESTful API documentation with Swagger 2.0 for Go.
- [fiber-swagger](https://github.com/arsmn/fiber-swagger) - Fiber middleware to automatically generate RESTful API documentation with Swagger 2.0.
- [viper](https://github.com/spf13/viper) - Go configuration with fangs.
- [mockery](https://github.com/vektra/mockery) - A mock code autogenerator for Golang.
- [testify](https://github.com/stretchr/testify) - A toolkit with common assertions and mocks that plays nicely with the standard library.

## Architecture
Clean Architecture - Uncle Bob
```
|-- configs
|-- infrastructures
|   |-- db
|-- internal
|   |-- middleware
|   |-- routes
|   |-- utils
|   |-- web
|-- src
|   |-- [module_name]
|       |-- dto
|       |-- entities
|       |-- handlers
|       |-- repositories
|       |-- router
|       |-- services
```

#### The diagram:
![golang clean architecture](https://github.com/yossdev/mypoints-rest-api/blob/main/docs/img/clean-arch.png)

## ERD
![ERD](https://github.com/yossdev/mypoints-rest-api/blob/main/docs/img/erd.jpg)