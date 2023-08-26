# syntax=docker/dockerfile:1
FROM golang:1.21-buster

WORKDIR /go/src/github.com/dinnerdonebetter/backend
ENV SKIP_PASETO_TESTS=TRUE
COPY . .

# to debug a specific test:
# ENTRYPOINT go test -parallel 1 -v -failfast github.com/dinnerdonebetter/backend/tests/integration -run TestIntegration/TestHouseholds_UsersHaveBackupHouseholdCreatedForThemWhenRemovedFromLastHousehold

ENTRYPOINT go test -v github.com/dinnerdonebetter/backend/tests/integration
