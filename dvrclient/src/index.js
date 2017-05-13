import React from "react";
import {render} from "react-dom";
import App from "./components/app.jsx";

//import our stylesheet so webpack puts it into the bundle
import "./css/main.css";

render(<App />, document.getElementById("app"));
