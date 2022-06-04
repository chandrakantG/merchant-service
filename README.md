# merchant-service

# Follow below steps to setup merchant service
1. open DB client and execute queries from `schema.sql` file
2. run below commands to start service
    * `go mod tidy -go=1.16 && go mod tidy -go=1.17`
    * `build .`
    * `./merchant-service`
3. kindly check `ApiDoc` for api requst and response
4. run below command to perform unit test
    * `go test server/*`