# courses-ada

## Creating data
1. You will need [Go](https://golang.org/doc/) and [PostgreSQL](https://www.postgresql.org/).
2. After installing PostgreSQL do the following:
    ```
    sudo -u postgres -i // or psql postgres on macOS
    create user YOURUSERNAME with password 'YOURPASSWORD';
    create database  YOURDATBASENAME owner YOURUSERNAME;
    ```
3. Inside [database.json](https://github.com/kdjohnson/courses-ada/blob/master/database.json) you need to alter 
    ```
    {"username":"YOUR_DATABASE_NAME","password":"YOUR_PASSWORD","url":"localhost","dbname":"YOUR_DATABASE_NAME"}
    ```
    to match the according values that you used to create your user and database 
4. Run the command `go run main.go`
    * This will create tables with the names meeting, courses, instructors, and grades.
    * It will also start a server listening out on port `8082`. You can see the json that will be returned for terms [here](http://localhost:8082/api/terms) 
5. Inside the [sql.sql](https://github.com/kdjohnson/courses-ada/blob/master/sql.sql) file is where you'll find the data that will be populated into the database. So alter it as you please.
6. To insert your new data from the SQL file, run the following command `psql -d DATABASE_NAME -U USERNAME -W < sql.sql` You should see something similar to: 
    ```
    INSERT 0 1
    INSERT 0 1
    INSERT 0 1
    INSERT 0 1
    ```

    Also note that it will not check for duplicates so if you do not want that you will need to comment out what you do not want. 
7. To view your data in a colomn layout do the following `psql -d DATABASE_NAME -U USER`. Then `select * from terms;` and you should see data.

## Dropping data
Run the following command to drop all tables created: `psql -d DATABASE_NAME -U USERNAME -W < drop.sql`

## Modifying data
Note if you ever want to change data in you table. Do the following `update TABLE set COLUMN = VALUE where id = ID_NUMBER_OF_ROW;`

## Running React code
1. Install [yarn](https://yarnpkg.com/)
2. Then, simply run `yarn start` from the base directory and it should open in your default browser 

## Running Go code
1. Do a `go get` to install the golang postgresql driver
2. Then `go run main.go` and the server will be started on port 8082

## Credits
Thanks to [Brandon Powell](https://github.com/bpowell) for a lot of help with the Go code. 