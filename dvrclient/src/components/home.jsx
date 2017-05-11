import React from "react"
import Gatherings from "./gatherings.jsx"
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import Layout from "./Layout.jsx";
import { Link } from 'react-router-dom';

export default class Home extends React.Component {
    constructor(props) {
        super(props)
    }
    render() {
        return(
            <Layout title="devour">
                <div className="homePage mdl-layout mdl-js-layout mdl-grid">
                    <div className="homeWrapper mdl-grid mdl-cell--4-col mdl-cell--2-col-phone mdl-cell--4-col-tablet">
                        <Link to="/recipes">
                            <div className="homeCards mdl-card mdl-shadow--4dp" >
                                <div className="mdl-card__title mdl-card--expand">
                                </div>
                                <div className="mdl-card__actions menuItem">
                                    Explore
                                </div>
                            </div>  
                        </Link>
                    </div>
                    <div className="homeWrapper mdl-grid mdl-cell--4-col mdl-cell--2-col-phone mdl-cell--4-col-tablet">
                        <Link to= "/social">
                            <div className="homeCards mdl-card mdl-shadow--4dp" >
                                <div className="mdl-card__title mdl-card--expand">
                                </div>
                                <div className="mdl-card__actions menuItem">
                                    Socialize
                                </div>
                            </div>
                        </Link>
                    </div>
                    <div className="homeWrapper mdl-grid mdl-cell--4-col mdl-cell--2-col-phone mdl-cell--4-col-tablet">
                        <Link to="/budget">
                    <       div className="homeCards mdl-card mdl-shadow--4dp" >
                                <div className="mdl-card__title mdl-card--expand">
                                </div>
                                <div className="mdl-card__actions menuItem">
                                    Budget
                                </div>
                            </div>  
                        </Link>
                    </div>
                    <div className="homeWrapper mdl-grid mdl-cell--4-col mdl-cell--2-col-phone mdl-cell--4-col-tablet">
                        <Link to="/plan">
                            <div className="homeCards mdl-card mdl-shadow--4dp" >
                                <div className="mdl-card__title mdl-card--expand">
                                </div>
                                <div className="mdl-card__actions menuItem">
                                    Plan
                                </div>
                            </div>  
                        </Link>
                    </div>
                    <div className="homeWrapper mdl-grid mdl-cell--4-col mdl-cell--2-col-phone mdl-cell--4-col-tablet">
                        <Link to="/recipebook">
                            <div className="homeCards mdl-card mdl-shadow--4dp" >
                                <div className="mdl-card__title mdl-card--expand">
                                </div>
                                <div className="mdl-card__actions menuItem">
                                    Recipes
                                </div>
                            </div>  
                        </Link>
                    </div>
                    <div className="homeWrapper mdl-grid mdl-cell--4-col mdl-cell--2-col-phone mdl-cell--4-col-tablet">
                        <Link to="/friends">
                            <div className="homeCards mdl-card mdl-shadow--4dp" >
                                <div className="mdl-card__title mdl-card--expand">
                                </div>
                                <div className="mdl-card__actions menuItem">
                                    Friends
                                </div>
                            </div>  
                        </Link>
                    </div>
                </div>
        </Layout>
        );
    }
}

