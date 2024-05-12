IMAGE=CloudSilk/curd:v1.0.0
run:
	CURD_DISABLE_AUTH=false DUBBO_GO_CONFIG_PATH="./dubbogo.yaml" go run main.go
run-lift:
	CURD_DISABLE_AUTH=true DUBBO_GO_CONFIG_PATH="./dubbogo-lift.yaml" go run main.go
build-image:
	CGO_ENABLED=0  GOOS=linux  GOARCH=amd64 go build -o curd main.go
	sudo docker build -f local.Dockerfile -t ${IMAGE} .
	rm curd
test-image:
	docker run -v `pwd`:/workspace/code -p 48081:48081 --env DUBBO_GO_CONFIG_PATH="./code/dubbogo.yaml" --rm  ${IMAGE}
push-image:
	sudo docker push ${IMAGE}
gen-doc:
	swag init --parseDependency --parseInternal --parseDepth 2