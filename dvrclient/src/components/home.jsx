import React from "react"
import Layout from "./Layout.jsx";
import { Link } from 'react-router-dom';

export default class Home extends React.Component {
    render() {
        return(
            <Layout title="devour">
                <div className="mdl-grid">
                    <div className="homeWrapper mdl-grid mdl-cell--4-col mdl-cell--4-col-phone mdl-cell--4-col-tablet">
                        <Link to="/recipes">
                            <div className="homeCards mdl-card mdl-shadow--4dp" >
                                <div className="mdl-card__title mdl-card--expand">
                                    <img src="/img/homeIco/svg/spatula.svg" className="homeIco" alt="exploreIcon"/>
                                </div>
                                <div className="mdl-card__actions menuItem">
                                    Explore
                                </div>
                            </div>  
                        </Link>
                    </div>
                    <div className="homeWrapper mdl-grid mdl-cell--4-col mdl-cell--4-col-phone mdl-cell--4-col-tablet">
                        <Link to="/social">
                            <div className="homeCards mdl-card mdl-shadow--4dp" >
                                <div className="mdl-card__title mdl-card--expand">
                                    <img src="/img/homeIco/svg/food.svg" className="homeIco" alt="exploreIcon"/>
                                </div>
                                <div className="mdl-card__actions menuItem">
                                    Socialize
                                </div>
                            </div>
                        </Link>
                    </div>
                    <div className="homeWrapper mdl-grid mdl-cell--4-col mdl-cell--4-col-phone mdl-cell--4-col-tablet">
                        <Link to="/budget">
                            <div className="homeCards mdl-card mdl-shadow--4dp" >
                                <div className="mdl-card__title mdl-card--expand">
                                    <img src="/img/homeIco/svg/restaurant-2.svg" className="homeIco" alt="exploreIcon"/>
                                </div>
                                <div className="mdl-card__actions menuItem">
                                    Budget
                                </div>
                            </div>  
                        </Link>
                    </div>
                    <div className="homeWrapper mdl-grid mdl-cell--4-col mdl-cell--4-col-phone mdl-cell--4-col-tablet">
                        <Link to="/plan">
                            <div className="homeCards mdl-card mdl-shadow--4dp" >
                                <div className="mdl-card__title mdl-card--expand">
                                    <img src="/img/homeIco/svg/restaurant-5.svg" className="homeIco" alt="exploreIcon"/>
                                </div>
                                <div className="mdl-card__actions menuItem">
                                    Plan
                                </div>
                            </div>  
                        </Link>
                    </div>
                    <div className="homeWrapper mdl-grid mdl-cell--4-col mdl-cell--4-col-phone mdl-cell--4-col-tablet">
                        <Link to="/recipebook">
                            <div className="homeCards mdl-card mdl-shadow--4dp" >
                                <div className="mdl-card__title mdl-card--expand">
                                    <img src="/img/homeIco/svg/app.svg" className="homeIco" alt="exploreIcon"/>
                                </div>
                                <div className="mdl-card__actions menuItem">
                                    My Recipe Book
                                </div>
                            </div>  
                        </Link>
                    </div>
                    <div className="homeWrapper mdl-grid mdl-cell--4-col mdl-cell--4-col-phone mdl-cell--4-col-tablet">
                        <Link to="/friends">
                            <div className="homeCards mdl-card mdl-shadow--4dp" >
                                <div className="mdl-card__title mdl-card--expand">
                                    <img src="/img/homeIco/svg/table.svg" className="homeIco" alt="exploreIcon"/>
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

