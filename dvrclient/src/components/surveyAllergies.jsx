import React from "react";
import DietTile from "./diet_tiles.jsx"
import Layout from "./Layout.jsx";

export default class SurveyAllergies extends React.Component {
    constructor(props) {
        super(props)
        this.state = {};
    }

    render() {
        //testing for allergies-- will have to get the most popular allergies from yummly
        var allergies = ["Nothing", "Nuts", "Gluten", "Mushrooms", "Shellfish", "Fish", "Dairy", "Corn", "Soy"]
        return(
            <Layout title="taste profile">
                <header className="mdl-layout__header mdl-layout__header-row center_title mdl-color--white textAccent">Allergies or Food Restrictions?</header>
                <div className="mdl-grid dietGrid">
                    {allergies.map((d) => <DietTile title={d}/>)}
                </div>
            </Layout>
        );
    }
}