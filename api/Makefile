check_install:
	@if not exist swagger (echo Swagger not found, installing... & go get -u github.com/go-swagger/go-swagger/cmd/swagger)

swagger: check_install
	@echo Generating Swagger specification...
	@swagger generate spec -o swagger.yaml --scan-models
