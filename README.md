# Advanced Golang

#### Testing
```
go test -v *.go
```

#### Auto reload (AIR)

https://github.com/air-verse/air

```
# for init config
air init

# fo start project server
air
```

## DATABASE
```
# postgres driver install

go get github.com/lib/pq
```



migrations cli: https://github.com/golang-migrate/migrate

bash
```
migrate create -seq -ext sql -dir ./cmd/migrate/migrations create_users
```