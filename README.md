# PenDown-Backend

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

4. Edit `.env` file

    ```txt
    PG_HOST=localhost
    PG_PORT=5432
    PG_USERNAME=
    PG_PASSWORD=
    PG_DBNAME=
    PORT=8080
    ```

5. Start backend service

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
