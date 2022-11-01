import React from "react";
import {EchoServiceClient} from "../Web/gRPCserver_grpc_web_pb.js";
import {EchoRequest,EchoResponse} from "../Web/gRPCserver_pb.js";


// const grpc = require("grpc-web");



const Home = () => {
    function EchoTest() {
        const echoServiceClient = new EchoServiceClient("http://localhost:8000/",null,null);
        console.log("request created");
        const req = new EchoRequest();
        req.setRequest(`value from the font-end ${req}`);

        echoServiceClient.echo(req,{}, (err,response) => {
            if(err) {
                console.log(`grpc error! ${JSON.stringify(err)}`);
                return
            } else {
                console.log(response.getResponse());
                // console.log(`\n ${response}`)
                return
            }
        })
    }
    return(
        <div className="container text-center">
            <div className="row">
                <h3 className="text-center">Home Page</h3>
            </div>
            <button className="btn btn-outline-dark text-center" onClick={EchoTest}>
              Home Button
            </button>
        </div>
    )
}
export default Home;
