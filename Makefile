check_install: 
	which swagger || go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger:check_install
	swagger generate spec -o swaggerui/swagger.json --scan-models