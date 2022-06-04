# PenDown-Backend

## Before You Start

**Notice** (for Windows users):
   
> If you don't want to install golang and gcc on your machine, you can jump to the [Set Up](#set-up) section and then go for [Option 2 - Using Docker](#using-docker)
>   
> Although you download both golang and gcc, you may still start fail due to package dependency. 
>    
> Therefore, **[Docker](https://www.docker.com/get-started/) is the recommended method**.

1. Install golang
   
   You should follow the official [install instruction](https://go.dev/doc/install).

2. Check golang version
   
   You can check in your terminal.

   ```shell
   go version
   
   // expected output, at least 1.17.6
   // go version go1.17.6
   ```

3. Install gcc
   
   Check you have gcc or type `gcc -v ` in terminal. (Version at least 8.1.0 or above)
   

## Set Up

1. Clone project

    ```shell
    git clone https://github.com/AOPLab/PenDown-be.git
    ```

2. Change directory

   ```shell
   cd PenDown-be
   ```

3. Copy configuration files

    ```shell
    cp .env.example .env
    ```

4. Put firebase file (`pendown-firebase.json` file)
   
   Put your firebase secret file (json format) in root. Follow [official document](https://firebase.google.com/) to get it.

   Or you can simply contact us to get the file, which is recommended.


5. Edit `.env` file

    ```txt
    # PostgreSQL
    PG_HOST=
    PG_PORT=5432
    PG_USERNAME=postgres
    PG_PASSWORD=
    PG_DBNAME=

    # Firebase
    SA_PATH=YOUR_FIREBASE_SECRET_FILE_PATH
    BUCKET_NAME=<YOUR_FIREBASE_BUCKET_NAME>.appspot.com

    # jwt
    jwt_token=ANY_STRING
    ```

6. Choose either [Option 1](#option-1---using-go-directly) or [Option 2](#option-2---using-docker) to start up the server.


### Option 1 - Using Go Directly

1. Start backend service 

    ```shell
    go mod download
    go mod tidy
    go run .
    ```

### Option 2 - Using Docker

Before using docker, you have to complete [set up](#set-up).

1. build

   ```shell
    docker build -t 'pendown-be' .
   ```

2. run

   ```shell
    docker run -p 8080:8080 pendown-be
   ```


After all this, you shall be able to visit http://localhost:8080/ in your browser, and it should display "404 page not found". 

## Unit Testing

> In root directory, execute `go test ./src/service -v`

### Unit Tests We Did

1. **註冊帳號**
   * Test_AddUser_Case_1 (Add normal user and success)
   * Test_AddUser_Case_2 (Add existing user and fail)
   * Test_AddGoogleUser
   
2. **登入帳號**
   * Test_FindUserByUsername
   * Test_FindUserByGoogleId
   
3. **上傳筆記**
   * Test_AddNote_Case_1 (Without course)
   * Test_AddNote_Case_2 (With course)
   * Test_AddNote_Case_3 (Without course)
   * Test_UpdatePdfFilename
   
4. **搜尋筆記**
   * Test_GetNoteByIdWithCourse
   * Test_SearchNoteAll (Search note)
   
5. **購買筆記**
   * Test_CheckUserBuyNote_Case_1 (User have bought the note)
   * Test_CheckUserBuyNote_Case_2 (User have not bought the note)
   * Test_BuyNote_Case_1 (User have not enough beans)
   * Test_BuyNote_Case_2 (User have enough beans)
   
6. **個人筆記瀏覽**
   * Test_GetUserNoteById
   
7. **其他**
   * Test_FindUserByAccountID

