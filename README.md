# Virtual-Minds-Go-Server
A lightweight server build in GoLang, using gin to handle api interactions and gorm to handle database interactions. Serves two endpoints:
1. ingests a request by a customer
2. serves aggregated statistics by date on a customer

## How It Works
- Virtual-Minds-Go-Server is a small server written in Go using Gin to handle requests, and GORM to handle object-relational mapping to the datastore
- Included is also a small database to test this API locally

### Architectural Choices Made
#### Controller 
- gin handles concurrency automatically as is the most lightweight routing implementation in go
- In addition, gin handles things like JSON validation and can be configured to attach Middleware easily should we want more security or flexibility

#### Database
- Though we can get more flexibility out of a NoSQL database, the current customer data is simple enough we can assume the schema will remain fixed and won't contain nested data
- PostGres, MySQL and MariaDB would all be valid options for such cases - with CustomerID being the partition key and Date being the clustering key
- In a high-throughput setting, it would be advisable to have a MASTER database that handles writes, and SLAVE replicas (because we have more reads than writes) that are load-balanced by least connections
- Alternatively, MASTER-MASTER nodes can be configured with proper conflict resolution code, or we can use an external tool like Kafka or Pulsar to replicate data to the different nodes

#### ORM
- gorm protects your code from SQL injections and ensures type validity to prevent badly interpreted insert data into SQL 
- gorm automatically migrates code and can create and conform to new data types without them existing in the target database
- For a NoSQL Database implementation, for example with Cassandra, goclqx would also work but isn't a fully functioning ORM like gorm or xorm

### Requirements 
- go >= 1.22
- docker >= 24.0.5
- docker-compose >= 2.20.3
- ports 8080 and 3306 for local runs

### Configuration 
- Project settings are found in config.yaml and are loaded in using *viper*
- Database connection settings are configured in database.yml
- Test database data can be inserted and configured under ./sql

### To run tests
Run `make test` in the terminal

### Running server locally 
1. Make sure you have Docker running in the background, as well as an internet connection to build the MySQL Image
2. Run `make start-db` to run a docker image with a test database in detached mode
3. Run `make` which will format, test, and build your code before running it
4. When you are done testing, exit the server with control-c (or command-c) and turn off the docker containers by running `make kill-db`

### Running server in a Docker container
1. Make sure you have Docker running in the background, as well as an internet connection to build the MYSQL and Server image
2. Run `make start-db` to run a docker image with a test database in detached mode
3. Run `build-docker` which will build your code into a local dockser image 
4. Run `run-docker` which will run the image you built into a container called `server`
5. When you are done testing, exit the server with control-c (or command-c) and turn off the docker containers by running `make kill-db`
6. Also, you can kill the running endpoint by running `make kill-db`

### To test endpoints (Swagger UI)
open `http://localhost:8080/docs/swagger/index.html`
If you expose different ports on your docker container, don't forget to adjust the path accordingly. 
**Note: Header user-agent is overwritten in Chromium based browsers**

### To test endpoints using curl 
For the customer endpoint (with your own values): `curl --location 'http://localhost:8080/api/v1/customer' --header 'User-Agent: test-agent' \--header 'Content-Type: text/plain' \--data ' {"customerId":1,"remoteIP": "121.11.55.2","timestamp":1709125373}'` 
For the statistics endpoint: `curl --location 'http://localhost:8080/api/v1/statistics?customer=1&date=20240228'` 
