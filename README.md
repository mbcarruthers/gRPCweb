## grpc-web example w/ Envoy sidecar Docker container 

---

Prequisites:Docker installed, protoc installed , protoc-gen-go installed, and protoc-gen-grpc-web as well.

---

### Contents 


Project consists of essentially 4(ish) parts: 

1) React front-end 


2) Gin backend serving running on port 30000 serving ReactJS front-end. Does not contain any protocol buffer code.


3) gRPCserver/server(the client in the gRPCserver directory was just there to test the gRPCserver/server


4) The Envoy directory(in gRPCserver) with the Dockerfile && envoy.yaml configuration file.

---

### Running the code

1)open up terminal. open 3 tabs within in. 1 tab in gRPCserver/envoy, one tab in $ROOT/server, one tab in gRPCserver directory


 Run make commands to build go and web-grpc, so `make build_protoc` and `make build_protoc_grpc_web`.


Build gRPCserver in terminal #3, `make build_server`


In terminal #1(gRPCserver/envoy) run `sudo docker build . -t envoy_proxy`

	get container_id and run `sudo docker run -p 8000:8000 --network=host ${CONTAINER_ID}`

	That'll get the envoy proxy running in Docker ready to accept grpc-web requests to localhost:8000 to 127.0.0.0:55051


Open terminal #2 and `go run ./` to start main and get it hosting the ReactJS code.


Open terminal #3 in gRPCserver and run the server with `./bin/gRPCserver/server`


Open up a browser and navigate to localhost:3000, and click the home button. Check the terminal for that sweet sweet console.log

NOTE: MUST run docker container with --network=host, will not route correctly without that.

This needs to be cleaned up but it works and I am saving it for now.

This is just a Boilerplate application to say, hey it works and this is how
