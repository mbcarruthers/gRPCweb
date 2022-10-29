import React from "react";
import './App.css';
import {BrowserRouter, Switch, Route, Link} from "react-router-dom";
import Home from "./Components/Home";
import Admin from "./Components/Admin";
const App = () => {
  return(
      <BrowserRouter>
              <div className="row d-flex m-1 position-sticky">
                  <div className="col-12">
                      <h2 className="gradient-text text-right">
                          Web gRPC
                      </h2>
                  </div>
              </div>
              <div className="row">
                  <div className="col-2 outline align-items-center">
                      <nav>
                          <ul className="list-unstyled list-group">
                              <li className="list-group-item m-2 text-center rounded">
                                  <Link to="/" className="btn btn-outline-dark text-white">Home</Link>
                              </li>
                              <li className="list-group-item m-2 text-center rounded">
                                  <Link to="/admin" className="btn btn-outline-dark text-white">Admin</Link>
                              </li>
                          </ul>
                      </nav>
                  </div>
                  <div className="col-10 outline">
                      <Switch>
                          <Route exact path="/">
                              <Home/>
                          </Route>
                          <Route exact path="/admin">
                              <Admin/>
                          </Route>
                      </Switch>
                  </div>
              </div>
      </BrowserRouter>
  )

}
export default App;