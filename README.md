# PenDown-Backend

## Before Start

0. Notice
   If you don't want to set up environment, you can jump to [set up](#set-up) directly and then [use docker](#use-docker)
    Although you download both golang and gcc, you may still start fail due to package dependency. So docker is the recommended method.
    <br />

1. Install golang
    You should follow the official [install instruction](https://go.dev/doc/install).
    <br />

2. Check golang version
    You can check in your terminal.

    ```shell
    go version

    // expected output, at least 1.17.6
    // go version go1.17.6
    ```

3. Install gcc
   Check you have gcc or type `gcc -v ` in terminal.
   

## Set up

1. Clone project

    ```shell
    git clone https://github.com/AOPLab/PenDown-be.git
    ```

2. Change directory

   ```shell
   cd github.com/AOPLab/PenDown-be
   ```

3. Copy configuration files

    ```shell
    cp .env.example .env
    ```

4. Put firebase file
   Put your firebase secret file (json format) in root.
   <br />

5. Edit `.env` file

    ```txt
    # PostgreSQL
    PG_HOST=
    PG_PORT=5432
    PG_USERNAME=postgres
    PG_PASSWORD=
    PG_DBNAME=

    # Firebase
    SA_PATH=YOUR_FIREBASE_FILE_PATH
    BUCKET_NAME=

    # jwt
    jwt_token=
    ```

6. Start backend service or [use docker](#use-docker)

    ```shell
    go mod download
    go run .
    ```

## Use docker

1. build

   ```shell
    docker build -t 'pendown-be' .
   ```

2. run

   ```shell
    docker run -d -p 8080:8080 pendown-be
   ```

## Unit test

* In root directory, execute `go test ./src/service -v`