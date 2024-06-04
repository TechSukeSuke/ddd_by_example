## At First
This repository is the ddd's example based on golang.
In the DB is using MySQL, and it is possible for us to perform the program with golang run methods. in addtion to those some method, there is a lot of runnnig options, such as delve, using the go test etc.

Actually, if you would rather those command run, use make command that is placed the root directory.

Don't worry about the runnable environment! I have been preparated it on docker compose.

Let's start with your understanding!

I appreciate that you visit this repository. I'm glad that I have meet you through this repository.

I'm thinking that this repository become the resource like OSS.
you can not merge for master branch but do some actions.
For example, if you want to improve a part of this repository's code, you can create either the issue or the PR, and which is able to include the improved code and state your thinking about performance.

So, I expect to your actions anytime. Let's try with us!

## Concept
This repository is assuming the situation like getting the company information when calls API.

The type of API is 3 pattern.
The one of those is api to check whether server is going on.
Other APIs, which are some apis to get the company information.

Those API's endpoint is below.
```
# pattern 1
http://localhost:8080/api/ping
# pattern 2
http://localhost:8080/api/v1/companies
# pattern 3
http://localhost:8080/api/v1/companies/:id
```

if you are the engineer that have see a API schema every day, you feel this API's endpoint very easy to understand.
As focus on pattern3, all most of someone can be understand but I explain it just in case.
:id keyword is the changeable values. The allowed values are number 1, 2, 3.
This numbers is the value that is containing DB called company details table.
So, if you were to add on another number into company details table, you can be able to try to call api on condition the another number.

Please try it!

## Table of Contents
- Setup
  - First, when you use the Make command like the below, you must add the specific argument.
    - Chose either `dev` or `prd` mode.
      - This repository is assuming to be able to use which environments. So, if when you use only the local environment to confirm, please chose `dev` mode.
  - You should run `set-up` command using Make command. The detail of this command is the below.
- Useable Make commands (when you construct a few containers.)
  - `set-up`
    - The command that should run before stand up the development environment.
  - `build-backend`
    - The command in order to build on backend container using docker compose.
    - In addition this command, `build-mysql` command will perform at the background. So you dont' have to run `build-mysql` command in this case.
  - `build-mysql`
    - The command in order to build on mysql container using docker compose.
  - `up-backend`
    - The command in order to start up the builded container.
    - In addition this command, `up-mysql` command will perform at the background. So you dont' have to run `up-mysql` command in this case.
  - `up-mysql`
    - The command in order to start up the builded container.
    - If you run this command, the mysql's container will stand up.
    - And after container is created, some migration files is run into the db server.
  - `down-backend`
    - The command in order to shout down the running contaienr.
    - It is depending on docker compose down command.
  - `down-mysql`
    - The command in order to shout donw the running contaienr.
    - It is depending on docker compose down command.
  - `wire`
    - The command in order to inject dependence each directories through DI tool.
  - `go-tidy`
    - The command in order to tidy to get the need packages in backend container.

- Useable Make commands (when you migrate the database on mysql container.)
  - `build-migrate`
    - The command in order to build the container that can only be runnable after the mysql container start on.
  - `up-migrate`
    - The only command to be able to run the migrations files after the mysql container run on.
    - It is depending on goose package.
  - `down-migrate`
    - The command in order to destory the existing table on mysql.
    - It is depending on goose package.
  - `seed`
    - The command in order to insert the dummy data into mysql's tables.
