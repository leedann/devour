import React from "react";
import Header from "./Header.jsx";
import Footer from "./Footer.jsx";
const homeIcons=["shopping_cart", "settings", "info"];
const eventFooter=["Upcoming", "Past"];
const recipeIcons=["search", "filter"];
const planFooter=["Day", "Week"];
const socialIcons=["library_add", "info"];
const budgetFooter=["Dashboard", "Spending"];
const base=["", "info"];
const blank=["",""]

//The layout of the page -- wraps all of the other pages in a header and footer
export default class Layout extends React.Component {
    constructor(props) {
        super(props);
        this.state={
            title: this.props.title,
            page: "",
            footNames: ""
        }
    }
    componentDidMount() {
    }

    //helps the header get its icons
    iconHelper() {
        switch (this.state.title) {
            case "devour":
                return homeIcons;
            case "taste profile":
                return blank;
            case "recipes":
                return recipeIcons; 
            case "social":
                return socialIcons;
            case "recipe book":
                return base;
            case "budget":
                return base;
            case "plan":
                return base;
            default:
                break;
        }
    }

    //helps the footer get its words
    footerHelper() {
        switch (this.state.title) {
            case "devour":
                return blank;
            case "taste profile":
                return blank
            case "recipes":
                return blank
            case "social":
                return eventFooter;
            case "recipe book":
                return blank;
            case "budget":
                return budgetFooter;
            case "plan":
                return planFooter;
            default:
                break;
        }
    }


    render() {
        var headerIcons = this.iconHelper();
        var footieWords = this.footerHelper();
        return (
            <div className="pageWrap mdl-layout mdl-js-layout">
                <Header title={this.state.title} page={headerIcons} />
                {/*the spans here allow the page to fit nicely rather than overlap on the footy*/}
                <span className="pageContent">{ this.props.children }</span>
                <div>
                    <div id="snackBar" className="mdl-js-snackbar mdl-snackbar">
                    <div className="mdl-snackbar__text"></div>
                        <button className="mdl-snackbar__action" type="button"></button>
                    </div>
                </div>
                <Footer title={this.state.title} page={footieWords}/>
            </div>
        );
    }
}


