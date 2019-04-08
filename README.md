# Real Estate Bidding 

### Go lang application to:
    - Create/Read/Update/Delete real estate 
    - Place a bid on existing estates
    - Fetch a winning bid for a second price auction on an estate
    
### Requirements:
    - Go lang
    - dep ( for dependency management )
    - mongo 3.6
    
### To run tests:
 ##### At project root level run ```go test```
 
### To run locally:
    1. Install mongo 3.6
    2. Modify confg.yml for mongo db uri, db and collection name
    3. run ```go build start.go```
    4. [http://localhost:8080/](http://localhost:8080/) on browser will land you at graphql playground with schema, query and mutation info
    
### To run container: 
    1. At project root level run command ```docker-compose up```
    2. [http://localhost:80/](http://localhost:80/) on browser will land you at graphql playground with schema, query and mutation info