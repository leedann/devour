import React from "react";
import {render} from "react-dom";
import App from "./components/app.jsx";
import { createBrowserHistory } from "history";
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';

//import our stylesheet so webpack puts it into the bundle
import "./css/main.css";

render(<App />, document.getElementById("app"));
