## grpc-web example w/ Envoy sidecar Docker container 

---

Prequisites:Docker installed, protoc installed , protoc-gen-go installed, and protoc-gen-grpc-web as well.

---

### Contents 


Project consists of essentially 4 parts: 

1) React front-end, ( Contents are within the app directory )


2) Gin backend serving running on port 3000 serving ReactJS front-end. Does not contain any protocol buffer code. (Contents are within the server directory)


3) gRPCserver/server. ( Contents are within the gRPCserver/server directory)(Also,Note:the client dir will not be used.It was initially just there for testing)


4) The Envoy proxy.Will run in a Docker container.( Contents are within container/envoy directory)

---

### Running the code

Considering the code is essentially made up of essentially four directories it's best to have a few terminal windows open. I preffered 1 terminal with 3 tabs and 2 IDE's,

First things first, you'll need to have protoc, protoc-gen-go, protoc-gen-go-grpc installed. The scope of installing those is beyond this and up to you. Those will need
to be installed and added to you $PATH bash variable. You'll also need gin-gonic, and gin-gonic/contrib installed on the gin server code(located in the server directory).


The protocol buffer code needs to be generated before any compilation is done. Within the gRPCserver directory there is a Makefile with make recipes to facilitate all of that.
There is `make build_protoc` which will build the protoc files for go and also `make build_protoc_grpc_web` for the building of the javascript protocol buffer code.
The go code will output into the gRPCserver/proto directory and the Javascript built code will output into the app/src/Web directory to be built with NPM when the
react code is built.


1) In the gRPCserver directory run `make build_protoc` and `make build_protoc_grpc_web`. That will build the protocol buffer code. Then to build the go gRPC server code
run `make build_server`. The output of that will be put into the gRPCserver/bin/gRPCserver/server. You can run it with `./bin/gRPCserver/server` from within the 
highest gRPCserver directory, However, that is not necessary yet.


2) Now that the protocol buffer code is generated and the gRPCserver binary created a small tweek will need to be made with the generated javascript code.

In the app/src/Web directory there will have been a gRPCserver_pb.js file generated. Paste the following code above the GENERATED CODE DO NOT EDIT line near the top.


```/* eslint-disable */``` <br/>
```// @ts-nocheck```


3) Make sure to run `npm install` and then run `npm run build` which will build the contents of the ReactJS app into a directory called client, which will be served by the
 
 gin back-end server code.

4) In Docker container will be build in the container/envoy directory. To do so, from within the container/envoy directoy run `sudo docker build . -t envoy_proxy`.

Copy the Id of the image you just built and run it with `sudo docker run -p 8000:8000 --network=host ${ID}`. That envoy proxy will receive request at :8000 and redirect them to

127.0.0.1:55051, where the gRPCserver will be running.

5) Now that The envoy proxy is running you can start the web-server code, from within the server/ directory run `go run ./`, then run the gRPC server from within the 
    gRPCserver directory with `./bin/gRPCserver/server`. Navigate in your web browser to `localhost:3000` and the button one Home is to test that it is working properly.


Note:Just an example code to get it running. Took me way longer than it should have. This only contains one gRPC service and that is to ping the service. The Envoy proxy still needs some work as it will not stop within 10 seconds after 
a stop is initiated so docker just kills it. May as well just `sudo docker kill ${CONTAINER_ID}` instead of `sudo docker stop ${CONTAINER_ID}`. I'll need to learn more about Envoy to fix that,
which I intend to do.

Note: This is just boilerplate material.
