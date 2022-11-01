import React from "react";

const Admin = () => {
    return(
        <div className="container text-center">
            <h3 className="text-center">Admin Page</h3>
            <button className="btn btn-outline-dark" onClick={() => {
                alert("Admin Button Clicked");
            }}>
                Admin Button
            </button>
        </div>
    )
}
export default Admin;