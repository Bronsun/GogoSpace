# GoGoApps NASA

### How to run a project
 
 ```
 make build
 ```
 or

 ```
 docker-compose up --build
 ```

### Requests

```
[GET] /pictures?from=<start_date>&to=<end_date>
```

Query parameters:

Both parameters `from` and `to` are required

Validation:
- `From` cannot be placed in the future
- `To` cannot be before `from`
- Both are required


### Tests
Test check ValidateDate() function in ../models/

Runing tests
```
cd models
go test
```