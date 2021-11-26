# Prixfixe

## dev dependencies

The following tools are prerequisites for development work:

- [go](https://golang.org/) 1.17+
- [docker](https://docs.docker.com/get-docker/)
- [docker-compose](https://docs.docker.com/compose/install/)
- [wire](https://github.com/google/wire) for dependency management
- [make](https://www.gnu.org/software/make/) for task running
- [terraform](https://learn.hashicorp.com/tutorials/terraform/install-cli) for deploying/formatting

## dev setup

It's a good idea to run `mage make lintegration_tests` before commits. You won't catch every error, but you'll catch the simplest ones that waste CI (and consequently your) time.

## running the server

1. clone this repository
2. run `make run`
3. [http://localhost:8888/](http://localhost:8888/)

## working on the frontend

2. run `mage run`
2. in a different terminal, run `mage frontendAutoBuild`
3. edit and have fun