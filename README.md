# Golang Webscraper using Colly

This project is an example of a web scraper using Golang, Go Colly.

This web scraper gets the company information from the website and stores the information into PostgreSQL database.

### Running this example
#### Prerequisites
* Install Docker -  https://docs.docker.com/install/

#### Steps to run solution in docker container

* Checkout github repo - 
`git clone https://github.com/jeroldleslie/golang-web-scraper`
* Go to the repo path - `cd <path>/golang-web-scraper`
* Run command `docker-compose up` to build and bring up postgres, ingestor, api services in docker containers.

Note:
* Please wait for a while to inject datas into database, then check the datas using API

#### API Details


#### Get Single Company
*Url Path*: `<domain>/company/:cvrid`

*Method*: GET

*URL Query Params*: -

*URL Path Params:* **:cvrid** - mandatory

#### Get Companies
*Url Path*: `company/_search?limit=*&offset=*`

*Method*: GET

*URL Query Params*: **limit**(mandatory), **offset**(mandatory)

*URL Path Params:* -
 
**Note**:
This exersice is done as per the requirement
1. Used **Go Modules** for dependencies management
2. Used **docker-compose**
3. Ingestor and API service **running in different docker container**
4. API is build using **echo framework**
5. Used **go-pg library** for database transactions

**Bonus work done**
1. Used **goroutines**
2. Used **channels**
3. Usage of git commits/branches/PRs


## Actual Task

Use http://www.cvr.dk to: Parse companies information, insert the information to the database, expose a HTTP API to retrieve this information.

## Requirements:

1. Should use Go Modules for dependencies management
2. All components should run from docker-compose
3. The ingestor and API should be in different containers
4. The API should use Echo web framework
5. Database transactions should be handled through go-pg library
 
### Bonus:

- Use Redis
- Showcase usage of goroutines
- Showcase usage of channels
- Handle cases when ingestion would crash on the source response changes / unavailability
- Proper usage of git commits/branches/PRs
