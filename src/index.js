import React from "react";
import {render} from "react-dom";
import App from "./components/app.jsx";
import Survey from "./components/survey.jsx"

//import our stylesheet so webpack puts it into the bundle
import "./css/main.css";
import {Router, Route, IndexRoute, hashHistory} from "react-router";

//TODO: replace the JSX here with a Router configuration
//from the react router module (already a dependency in package.json)
var router = (
    <Router history= {hashHistory}>
        <Route path="/" component={Survey}>
        </Route>
    </Router>
)
render(router, document.getElementById("app"));
