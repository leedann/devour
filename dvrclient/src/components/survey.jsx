import React from "react";
import Layout from "./Layout.jsx";
import { Dropdown } from 'semantic-ui-react'
import { Button } from 'semantic-ui-react'
const myHeader = new Headers();
myHeader.append('Authorization', localStorage.getItem("Authorization"));
const baseurl = 'https://dvrapi.leedann.me/v1/'

const dietOptions = [
  { key: 'vegetarian', text: 'Vegetarian', value: 'Vegetarian' },
  { key: 'vegan', text: 'Vegan', value: 'Vegan' },
  { key: 'ovo vegetarian', text: 'Ovo vegetarian', value: 'Ovo vegetarian' },
  { key: 'lacto vegetarian', text: 'Lacto vegetarian', value: 'Lacto vegetarian' },
  { key: 'pescetarian', text: 'Pescetarian', value: 'Pescetarian' }
];

const allergyOptions = [
  { key: 'dairy', text: 'Dairy', value: 'Dairy' },
  { key: 'egg', text: 'Egg', value: 'Egg' },
  { key: 'gluten', text: 'Gluten', value: 'Gluten' },
  { key: 'peanut', text: 'Peanut', value: 'Peanut' },
  { key: 'seafood', text: 'Seafood', value: 'Seafood' },
  { key: 'sesame', text: 'Sesame', value: 'Sesame' },
  { key: 'soy', text: 'Soy', value: 'Soy' },
  { key: 'sulfite', text: 'Sulfite', value: 'Sulfite' },
  { key: 'tree nut', text: 'Tree Nut', value: 'Tree Nut' },
  { key: 'wheat', text: 'Wheat', value: 'Wheat' } 
];
const goalOptions = [
  { key: 'all of these reasons', text: 'All of these reasons', value: 'all' },
  { key: 'planning', text: 'Planning group meals', value: 'Planning group meals' },
  { key: 'learning', text: 'Learning how to cook', value: 'Learning how to cook' },
  { key: 'healthy', text: 'Eating Healthy', value: 'Eating Healthy' },
  { key: 'plan', text: 'Plan meals', value: 'Plan Meals' },
  { key: 'diet', text: 'Dieting', value: 'Dieting' },
  { key: 'budget', text: 'Budget', value: 'Budget' },
  { key: 'other', text: 'Other', value: 'Other' },
];
const DietDrop = () => (
  <Dropdown id="dietSurvey" placeholder='Diet' fluid multiple selection options={dietOptions} />

)
const AllergyDrop = () => (
  <Dropdown id="allergySurvey" placeholder='Allergy' fluid multiple selection options={allergyOptions} />

)
const GoalsDrop = () => (
  <Dropdown id="goalSurvey" placeholder='Goals' fluid multiple selection options={goalOptions} />
)

export default class Survey extends React.Component {
    constructor(props) {
        super(props)
        this.state = {};
    }

    onSubmit() {
        let diets = document.getElementById("dietSurvey").getElementsByTagName("a");
        let allergies = document.getElementById("allergySurvey").getElementsByTagName("a");
        let dVal = [];
        let aVal = [];
        for (let i=0;i<diets.length;i++) {
            dVal[i] = diets[i].text
        }
        for (let i=0;i<allergies.length;i++) {
            aVal[i] = allergies[i].text
        }

        var dietInfo = {
            "diets": dVal
        }
        var allergyInfo = {
            "allergies": aVal
        }
        var dData = JSON.stringify(dietInfo);
        var aData = JSON.stringify(allergyInfo);

        var dReq = {
            method: 'POST',
            headers: myHeader,
            body: dData
        }
        var aReq = {
            method: 'POST',
            headers: myHeader,
            body: aData
        }
        fetch(baseurl+'users/diets', dReq)
        .then(() => {
            fetch(baseurl+'users/allergies', aReq)
            .then(() => {
                this.props.history.push('/surveyend')
            })
            .catch((error) => {
                alert("error adding allergies")
            })
        })
        .catch((error) => {
            alert("error adding diets")
        })
    }

    render() {
        //The survey array-- will have to get the popular diet types from yummly
        return(
            <Layout title="taste profile">
                <h2 className="mdl-layout__header mdl-layout__header-row mdl-layout__header--transparent textSecondary textAccent ">select your diet</h2>
                <div className="mdl-grid dietGrid">
                    <DietDrop/>
                </div>
                <h2 className="mdl-layout__header mdl-layout__header-row mdl-layout__header--transparent textSecondary textAccent ">select your allergies</h2>
                <div className="mdl-grid dietGrid">
                    <AllergyDrop/>
                </div>
                <h2 className="mdl-layout__header mdl-layout__header-row mdl-layout__header--transparent textSecondary textAccent ">select your goals</h2>
                <div className="mdl-grid dietGrid">
                    <GoalsDrop/>
                </div>
                <Button onClick={() => this.onSubmit()}floated='right' id="surveyButton" size="huge">Done</Button>
            </Layout>
        );
    }
}