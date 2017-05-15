import React from "react";
import DietTile from "./diet_tiles.jsx"
import Layout from "./Layout.jsx";

//the survey goals -- why the user is using 
export default class SurveyGoals extends React.Component {
    constructor(props) {
        super(props)
        this.state = {};
    }

    render() {
        //probally not pertinent to the functioning of the app
        var goals = ["All of these Reasons", "Plan group meals", "Learn to cook", "Eat Healthy", "Find Recipes", "Plan Meals", "Entertain", "Stick to a Diet", "Budget"]
        return(
            <Layout title="taste profile">
                    <header className="mdl-layout__header mdl-layout__header-row center_title mdl-color--white textAccent ">select your goals</header>
                    <div className="mdl-grid dietGrid">
                        {goals.map((d) => <DietTile title={d}/>)}
                    </div>
            </Layout>
        );
    }
}