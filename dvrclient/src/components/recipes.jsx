import React from "react";
import Layout from "./Layout.jsx";
import {history} from './app.jsx';
import { Button, Popup } from 'semantic-ui-react'
const baseurl='https://api.yummly.com/v1/api/';
const apiHeader = new Headers();
apiHeader.append("X-Yummly-App-ID", "f3a4ad3c");
apiHeader.append("X-Yummly-App-Key", "8377c0df84a59fd60a40c2cb0ce4a964");
const myHeader = new Headers();
myHeader.append('Authorization', localStorage.getItem("Authorization"));
const base = 'https://dvrapi.leedann.me/v1/'
const allergyVals={
    "Gluten": "&allowedAllergy[]=393^Gluten-Free",
    "Peanut": "&allowedAllergy[]=394^Peanut-Free",
    "Seafood": "&allowedAllergy[]=398^Seafood-Free",
    "Sesame": "&allowedAllergy[]=399^Sesame-Free",
    "Soy": "&allowedAllergy[]=400^Soy-Free",
    "Dairy": "&allowedAllergy[]=396^Dairy-Free",
    "Egg": "&allowedAllergy[]=397^Egg-Free",
    "Sulfite": "&allowedAllergy[]=401^Sulfite-Free",
    "Tree Nut": "&allowedAllergy[]=395^Tree Nut-Free",
    "Wheat": "&allowedAllergy[]=392^Wheat-Free"

}
const dietVals={
    "Lacto vegetarian": "&allowedDiet[]=388^Lacto vegetarian",
    "Ovo vegetarian": "&allowedDiet[]=389^Ovo vegetarian",
    "Pescetarian": "&allowedDiet[]=390^Pescetarian",
    "Vegan": "&allowedDiet[]=386^Vegan",
    "Vegetarian": "&allowedDiet[]=387^Lacto-ovo vegetarian",
}
import LoaderExampleLoader from './loading.jsx';
//the cards for the recipes
class RecipeCard extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            fullDetails: null
        }
    }
    componentWillMount() {
        var data = {
            method: 'GET',
            headers: apiHeader
        }
        fetch(baseurl + `recipe/${this.props.details.id}`, data)
        .then((resp) => resp.json())
        .then(data => {
            this.setState({
                fullDetails: data
            })
        })
    }

    //toggles the info table 
    showInfo() {
        var tab = document.getElementById(this.props.details.id + "info")
        tab.classList.toggle("hidden")
    }

    addRecipe() {
        //means that this is adding to the event rather than book
        if (this.props.event) {
            var data = {
                "recipeName": this.props.details.recipeName
            }
            data = JSON.stringify(data);
            var req = {
                method: 'POST',
                headers: myHeader,
                body: data
            }
            fetch(base+'events/recipes/' + this.props.event, req)
            .catch((error) => {
                alert("There was something wrong with your request");
            })
        }else {
            var req = {
                method: 'POST',
                headers: myHeader,
            }
            fetch(base+'users/recipebook/' + this.props.details.recipeName, req)
            .catch((error) => {
                alert("There was something wrong with your request");
            })
        }
    }

    handleInfo(id, type) {
        var target = document.getElementById(id + type);
        var content = document.getElementById(id + type+"content");
        switch (type) {
            case "nutrition":
                var other = document.getElementById(id + "ingredients");
                var otherContent = document.getElementById(id + "ingredientscontent");
                var last = document.getElementById(id + "detail");
                var lastContent = document.getElementById(id + "detailcontent");
                target.classList.add('active');
                other.classList.remove('active');
                last.classList.remove('active')
                lastContent.classList.add('hidden');
                otherContent.classList.add('hidden');
                content.classList.remove('hidden');
                break;
            case "ingredients":
                other = document.getElementById(id + "nutrition");
                otherContent = document.getElementById(id + "nutritioncontent");
                last = document.getElementById(id + "detail");
                lastContent = document.getElementById(id + "detailcontent");
                target.classList.add('active');
                other.classList.remove('active');
                last.classList.remove('active');
                otherContent.classList.add('hidden');
                lastContent.classList.add('hidden');
                content.classList.remove('hidden');
                break;
            case "detail":
                other = document.getElementById(id + "nutrition");
                otherContent = document.getElementById(id + "nutritioncontent");
                last = document.getElementById(id + "ingredients");
                lastContent = document.getElementById(id + "ingredientscontent");
                target.classList.add('active');
                other.classList.remove('active');
                last.classList.remove('active');
                otherContent.classList.add('hidden');
                lastContent.classList.add('hidden');
                content.classList.remove('hidden');
                break;
            default:
                break;
        }
    }

    //currently splits the ingredients and desc but will have to work with yummly
    render() {
        if (!this.state.fullDetails) {
            return (<span></span>);
        }
        var styles = {
            backgroundImage: `url(${this.state.fullDetails.images[0].hostedLargeUrl})`,
            backgroundSize: 'cover'
        }
        
        return(
            <div className="recipeCard mdl-card mdl-shadow--2dp mdl-cell mdl-cell--12-col mdl-cell--8-col-phone mdl-cell--8-col-tablet">
                <div className="imageContainer mdl-card__title mdl-card--expand" style={styles}>
                    <span className="RText mdl-cell mdl-cell--6-col mdl-cell--2-col-phone mdl-cell--4-col-tablet textAccent">{this.props.details.recipeName}</span>
                        <Popup
                        trigger={<button className="favButton mdl-button mdl-js-button mdl-button--fab mdl-js-ripple-effect mdl-button--accent" onClick={() => this.addRecipe()}>
                        <i className="material-icons">favorite_border</i>
                    </button>}
                        content={`Added ${this.props.details.recipeName}`}
                        on='click'
                        />
                </div>
                <div className="recipeActions mdl-card__actions mdl-card--border mdl-grid" onClick={() => this.showInfo()}>
                    <i className="material-icons">query_builder</i>
                    <span className="recipeName mdl-cell mdl-cell--6-col mdl-cell--2-col-phone mdl-cell--4-col-tablet">
                    {this.state.fullDetails.totalTime}</span>

                    <span className="recipeTime mdl-cell mdl-cell--6-col mdl-cell--2-col-phone mdl-cell--4-col-tablet">{this.props.details.CookTime}</span>
                </div>
                
                <div className="hidden recipeInfo" id={this.props.details.id + "info"}>
                    <div className="ui secondary menu">
                        <a id={this.props.details.id + "nutrition"}className="active item" onClick={() => this.handleInfo(this.props.details.id, "nutrition")}>Nutrition</a>
                        <a id={this.props.details.id + "ingredients"}className="item" onClick={() => this.handleInfo(this.props.details.id, "ingredients")}>Ingredients</a>
                        <a id={this.props.details.id + "detail"}className="item" onClick={() => this.handleInfo(this.props.details.id, "detail")}>Details</a>
                    </div>
                    <div className="ui bottom attached segment">
                        <div id={this.props.details.id + "ingredientscontent"} className="hidden">
                            <ul className="recipeList">
                                {this.state.fullDetails.ingredientLines.map((d, i) =>
                                <li key={this.props.details.id + "ing"+i}>{i+1+`. `+d}</li>
                                
                                )}
                            </ul>
                        </div>
                        <div id={this.props.details.id + "nutritioncontent"} >
                            <ul className="recipeList">
                                {this.state.fullDetails.nutritionEstimates.map((d, i) =>
                                <li key={this.props.details.id + "nut"+i}>{
                                    d.value + d.unit.abbreviation +' '+ (d.description? d.description : 'Fat')
                                    }</li>
                                
                                )}
                            </ul>
                        </div>
                        <div id={this.props.details.id + "detailcontent"} >
                            <ul className="recipeList">
                                <p>For detailed instructions visit: </p>
                                <a href={this.state.fullDetails.source.sourceRecipeUrl} target="_blank">{this.state.fullDetails.source.sourceDisplayName}</a>
                            </ul>
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}

// the page of the recipes
export default class RecipePage extends React.Component {
    constructor(props) {
        super(props)
        this.state={
            event: false,
            diets: [],
            allergies: [],
            restrictions: "",
            start: 0
        }
    }

    componentDidMount() {
        if (this.props.match.params.eventid) {
            var id = this.props.match.params.eventid
            var req2 = {
                method: 'GET',
                headers: myHeader
            }
        
            fetch(base+'events/' + id, req2)
            .then((resp) => resp.json())
            .then(event => {
                var restrict = ""
                for (let i=0;i<event.restrictions.allergies.length;i++) {
                    restrict += allergyVals[event.restrictions.allergies[i]];
                }
                for (let i=0;i<event.restrictions.diets.length;i++) {
                    restrict += dietVals[event.restrictions.diets[i]]
                }
                this.setState({
                    event: this.props.match.params.eventid,
                    restrictions: restrict
                })
            })
            .then(() => {
                this.handleYum()
            })
        }else {
            this.handleYum()
        }
    }

    handleYum() {
        var req = {
            method: 'GET',
            headers: apiHeader
        }
        fetch(`https://api.yummly.com/v1/api/recipes?&maxResult=5&start=${this.state.start}&requirePictures=true`+this.state.restrictions, req)
        .then((resp) => resp.json())
        .then(data => {
            this.setState({
                matches: data.matches
            })
        })
    }

    handleForward(e) {
        e.preventDefault();
        var req = {
            method: 'GET',
            headers: apiHeader
        }
        fetch(`https://api.yummly.com/v1/api/recipes?&maxResult=5&start=${this.state.start +5}&requirePictures=true`+this.state.restrictions, req)
        .then((resp) => resp.json())
        .then(data => {
            this.setState({
                start: this.state.start+5,
                matches: data.matches
            })
        })
    }

    handleBack(e) {
        e.preventDefault();
        var req = {
            method: 'GET',
            headers: apiHeader
        }
        if (this.state.start < 5) {
            fetch(`https://api.yummly.com/v1/api/recipes?&maxResult=5&start=${0}&requirePictures=true`+this.state.restrictions, req)
            .then((resp) => resp.json())
            .then(data => {
                this.setState({
                    start: 0,
                    matches: data.matches
                })
            })
        }else {
            fetch(`https://api.yummly.com/v1/api/recipes?&maxResult=5&start=${this.state.start -5}&requirePictures=true`+this.state.restrictions, req)
            .then((resp) => resp.json())
            .then(data => {
                this.setState({
                    start: this.state.start - 5,
                    matches: data.matches
                })
            })   
        }
    }

    render() {
        if (!this.state.matches) {
            return (<LoaderExampleLoader />);
        }
        //this is the test array-- will change to the prop with api call to recipes
        return(
            <Layout title="recipes">
                <p className="titleName">View Recipes</p>
                <div className="recipeWrapper mdl-grid">
                    <Button.Group className="mdl-cell mdl-cell--12-col mdl-cell--8-col-phone mdl-cell--8-col-tablet">
                        <Button onClick={(e) => this.handleBack(e)} labelPosition='left' icon='left chevron' content='Back' />
                        <Button onClick={(e) => this.handleForward(e)} labelPosition='right' icon='right chevron' content='Forward' />
                    </Button.Group>
                    {this.state.matches? this.state.matches.map((rec) => {
                        return <RecipeCard key={rec.id} event={this.state.event} details={rec}/>
                    }) : ""}
                    <Button.Group widths='5' className="mdl-cell mdl-cell--12-col mdl-cell--8-col-phone mdl-cell--8-col-tablet">
                        <Button onClick={(e) => this.handleBack(e)} labelPosition='left' icon='left chevron' content='Back' />
                        <Button onClick={(e) => this.handleForward(e)} labelPosition='right' icon='right chevron' content='Forward' />
                    </Button.Group>
                </div>
            </Layout>
        );
    }
}