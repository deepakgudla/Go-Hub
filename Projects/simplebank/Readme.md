# Backend master class in GO

- [course](https://www.udemy.com/course/backend-master-class-golang-postgresql-kubernetes/)

- Design Develop and Deploy

- A simple bank with the following features
    - create and manage account
        - contains owner's name, balance and currency
    - record all balance changes to each of the account
        - create an account entry for each change
        - every time money is debited/credited, an account entry record will be created
    - Money Transfer transaction
        - perform money transfer between 2 accounts consistently within a transaction so that either both the accounts balance are updated successfully

- Docker
    - a container is one instance of the application contained in the image which is started by the docker run command
    - we can start multiple containers from one single image

- docker container runs in a seperate virtual container which is different from the host network that we are on 

- `docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:17-alpine`

- docker exec -it postgres17 psql -U root
    - -U flag to tell psql that we want to connect with the root user
- `it` flag is to run the command as an interactive TTY session
- **Postgresql_password**
    - The postgreql image sets up `trust` authentication locally so you may notice a password is not required when connecting from localhost when it is inside the same container. However, a password will be required if connecting from a different host/container...

- port mapping
    - basically a docker container runs in a seperate virtual network which is different from the host network, we cannot simply connect to the postgres server running on port 5432 of the container network, unless we tell docker to create one kind of `bridge` between localhost network and the container's network, this can be done using `-p` flag
- postgres container sets up a trust authentication locally, password is not required when connecting from localhost

- `tableplus` is a gui toll that can talk to many different kinds of database engines and it is very easy to use, and helps us a lot in speeding up development



- go through database migration in detail
    - `up/down migration` is a best practice when writing database migration
    - `up script` is run to make a forward change to the schema
    - `down script` is run if we want to revert the change made by the up-script
    - when we run the `migrate up` command, the `up script` files inside `db/migration` folder will be run sequentially by the order of their prefix version
    - when we run the `migrate down` command, the `down script` files inside `db/migration` folder will be run sequentially by the reverse order of their prefix version
    - command for migration
        - `migrate create -ext sql -dir db/migration -seq init_schema`
        - seq file is to represent a sequential version number for the migration file

- docker exec -it postgres17 /bin/sh

- createdb --username=root --owner=root simple_bank

- ```migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up```


- the schema migration table stores the latest applied migration version
- in this case, it is 1, because we have to run only 1 single migration file
- the dirty column tells us if the last migration has failed or not...
- If it fails, we must manually fix the issues to make the database state clean before trying to run any other migration versions



###### CRUD


- create
    - insert new records to the database
- READ
    - select or search records in the database

- UPDATE
    - change some fields of the record in the database

- DELETE
    - remove records from the database

- there are several ways to implement crud operations in golang
    - database/sql (low level standard library)
        - runs very fast and straightforward (upside)
        - manual mapping sql fields to variables (downside)
        - easy to make mistakes, not caught until runtime
        - if somehow the order of variables doesn't match, or if we forget to pass some arguments to the function call, the errors will only show up at runtime
    - gorm
        - high level object relational mapping library for golang
        - `CRUD` functions are already implemented
        - we only need to declare the models, and call the functions that gorm provided
        - must learn to write queries using gorm's provided functions
        - we have to learn how to declare association tags to make gorm understand the relationship between the tables, so that it can generate the correct SQL query
        - runs slowly on high load, run slowly when the traffic is high
        - gorm can run 3-5 times slower than the standard library
    - sqlx
        - runs nearly as fast as the standard library
        - field mappings are done via query text and struct tags
        - it provides some functions like `Select()` or `StructScan()` which will automatically scan the result into struct fields so we don't have to do the mapping manually like the `database/sql` package 
        - any errors in the query will only be caught at runtime
    - sqlc
        - runs very fast like `database/sql`
        - we just need to write queries, golang crud codes will be automatically generated
        - sqlc parses sql queries
        - any errors will be caught and reports rigth away - catches sql query errors before generating codes
        - checks the SQL query syntax before generating the codes


- `RETURNING *` clause is used to tell postgresql to return the value of all columns after inserting the record into accounts table
- after a table is created, we will always want to return its ID to the client

### Database Transaction

- it is a single unit of work
- often made up of multiple db operations
- 
- we need to perform transaction that combines some operations from several tables
- to transfer 10 usd from account1 to account 2, it contains 5 txns
    - create a transfer record with amount = 10
    - create an account entry for account1 with amount = -10, since money is moving out of this account
    - create an account entry for account 2 with amount = +10, since money is moving in to this account
    - subtract 10 from the balance of account1
    - add 10 to the balance of account2
- why do we need db transaction
    - to provide a reliable and consistent unit of work, even in case of system failure
    - to provide isolation between programs that access the database concurrently
    - in order to achieve the above two goals, a database txn must satisfy the ACID properties
- ACID
    - atomicity
        - either all operations(of the txn) complete successfully or the txn fails and (everything is rolled back) the db is unchanged
    - consistency
        - all constraints must be satisfied
        - the database should remain valid after the txn is executed
        - all data written to the database must be valid according to the predefined rules, including constraints, cascades and triggers
    - isolation
        - concurrent transactions must not affect each other
        - there are several levels of isolation that defines when the changes made by 1 transaction can be visible to others
    - Durability
        - data written by a successful txn must be recorded in persistent storage
        - basically means that all data written by a successful txn must stay in persistent storage and cannot be lost, even in the case of system failure
- How to run a SQL database transaction
    - start a txn with `BEGIN;` statement and end with `COMMIT;`
    - then we write a series of normal queries or operations, if all of them were succesful, we COMMIT the txn to make it permanent, the database will be changed to a new state, otherwise if any query fails we will ROLLBACK the txn
    - all changes made by the previous queries of the txn will be gone and the database stays the same as it was before the txn


- embedding inside the struct is called **composition**, it is a preferred way to extend struct functionality in golang instead of inheritance

- The best way to make sure that the db txn works well is to run it with several concurrent go routines

- channel is designed to connect concurrent Goroutines and allow them to safely share data with each other without explicit locking


































## Deployment of Application using Docker, kubernetes and AWS

- 