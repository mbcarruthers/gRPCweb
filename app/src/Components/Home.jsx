import React from "react";
import {EchoServiceClient} from "../Web/gRPCserver_grpc_web_pb.js";
import {EchoRequest} from "../Web/gRPCserver_pb.js";




const Home = () => {
    function EchoTest() {
        const echoServiceClient = new EchoServiceClient("localhost:8080",null,null);
        let req = new EchoRequest();
        req.setRequest("PiNg");

        echoServiceClient.echo(req,null, (err,response) => {
            if(err) {
                console.log(`grpc error! ${err}`);
                return
            } else {
                console.log(response.getError());
                console.log(response.getResponse());
                console.log("if that is PiNg-PONG then SUCCESS!");
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