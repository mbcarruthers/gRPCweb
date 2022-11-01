Have protoc-gen-grpc and protoc-gen-grpc-go or whatever the latter is called installed. Whatever gives you an error trying to use `make build_protoc`.

Project consists of essentially 4(ish) parts: 1) React front-end 
						*including Web/* in src which containts javascript protocol buffers build from make. You can find the command with make or make help.
						 Just trying a vanilla make in the directory containing the makefile will automatically call the about function.
					      2) Gin backend serving running on port 30000 serving ReactJS front-end. Does not contain any proto-stuff
					      3) gRPCserver/server(the client in the gRPCserver directory was just there to test the gRPCserver/server
					      4) The Envoy docker container information
					      	Envoy dir with Dockerfile and envoy.yaml file.

1)open up terminal. open 3 tabs within in. 1 tab in gRPCserver/envoy, one tab in $ROOT/server, one tab in gRPCserver directory

2) Open up WebStorm in app directory, and open up GoLand in gRPCserver directory, or $ROOT/server directory, Whatever you'd be working on, the backend or gRPCserver.

3) Run make commands to build go and web-grpc, so `make build_protoc` and `make build_protoc_grpc_web`.

4) Build gRPCserver in tab#3, `make build_server`

3) in terminal #1(gRPCserver/envoy) run `sudo docker build . -t envoy_proxy`
	get container_id and run `sudo docker run -p 8000:8000 --network=host ${CONTAINER_ID}`
	That'll get the envoy proxy running in Docker ready to accept grpc-web requests to localhost:8000 to 127.0.0.0:55051
4) Open terminal #2 and `go run ./` to start main and get it hosting the ReactJS code.

5) Open terminal #3 in gRPCserver and run the server with `./bin/gRPCserver/server`


Open up a browser and navigate to localhost:3000, and click the home button. Check the terminal for that sweet sweet console.log

NOTE: MUST run docker container with --network=host, will not route correctly without that.

This needs to be cleaned up but it works and I am saving it for now.

Dear god, as if just this did not take long enough.
