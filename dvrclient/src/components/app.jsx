import React from "react";
import Login from "./Login.jsx";
import Register from "./Register.jsx";
import Gatherings from "./gatherings.jsx";
import Home from "./home.jsx";
import Welcome from "./welcome.jsx";
import Survey from "./survey.jsx";
import CreateEvent from "./createEvent.jsx";
import GatheringRequests from "./gatheringRequests.jsx";
import SurveyEnd from "./surveyConf.jsx";
import SurveyGoals from "./surveyGoals.jsx";
import SurveyAllergies from "./surveyAllergies.jsx";
import GatheringPage from "./gatheringEvent.jsx";
import { createBrowserHistory } from "history";
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
const HomePage = () => <Home />
const App = () => (
  <Router>
    <Switch>
      <Route exact path="/" component={Login} />
      <Route path="/register" component={Register} />
      <Route path="/home" component={HomePage} />
      <Route exact path="/social" component={Gatherings} />
      <Route path="/welcome" component={Welcome} />
      <Route path="/survey" component={Survey} />
      <Route path="/create" component={CreateEvent} />
      <Route path="/info" component={GatheringRequests} />
      <Route path="/surveyend" component={SurveyEnd} />
      <Route path="/selectGoals" component={SurveyGoals} />
      <Route path="/allergyinfo" component={SurveyAllergies} />
      <Route path="/social/event" component={GatheringPage} />
    </Switch>
  </Router>
);
export const history = createBrowserHistory({
    forceRefresh: true,
});
export default App;