# graphql-sample

## Development Guide

### Installation

- Install Git

  See [Git Installation](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

- Install Go (Golang)

  See [Golang Installation](https://golang.org/doc/install)

- Clone this repo in `$GOPATH/src/github.com/bickyeric`

- Install Dep

  See [Dep Installation](https://golang.github.io/dep/docs/installation.html)

- Install dependencies

  ```sh
  cd $GOPATH/src/github.com/bickyeric/graphql-sample
  dep ensure
  ```

- Copy env.sample and if necessary, modify the env value(s)

  ```sh
  cp .env.sample .env
  ```

- Prepare MySQL database

  ```sh
  make db-up
  make db-seed
  ```

- Run Locally

  ```sh
  make run
  ```