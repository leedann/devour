import React from "react";
import {render} from "react-dom";
import App from "./components/app.jsx";

//import our stylesheet so webpack puts it into the bundle
import "./css/main.css";

//TODO: replace the JSX here with a Router configuration
//from the react router module (already a dependency in package.json)
render((
    <App/>
    ), document.getElementById("app"));
