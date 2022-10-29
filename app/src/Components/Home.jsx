import React from "react";

const Home = () => {
    return(
        <div className="container text-center">
            <div className="row">
                <h3 className="text-center">Home Page</h3>
            </div>
            <button className="btn btn-outline-dark text-center" onClick={() => {
                alert("Button is Working")
            }}>
              Home Button
            </button>
        </div>
    )
}
export default Home;