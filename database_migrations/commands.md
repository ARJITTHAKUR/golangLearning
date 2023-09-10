    # UP

 migrate -database="postgres://postgres:password@localhost:5432/golang_migration_test?sslmode=disable" -path="./migrations" up
    # Down
 migrate -database="postgres://postgres:password@localhost:5432/golang_migration_test?sslmode=disable" -path="./migrations" down 1