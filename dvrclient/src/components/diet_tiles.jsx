import React from "react";

export default class DietTile extends React.Component {
    constructor(props) {
        super(props)
        this.state={
            selected: []
        }
    }

    
    clickAction(e) {
        e.preventDefault()
        var target = e.currentTarget;
        var selectedHolder = this.state.selected.slice();
        // diet was found, so remove it
        var index = selectedHolder.indexOf(target.textContent);
        if (index > -1) {
            selectedHolder.splice(index, 1);
            target.style.backgroundColor = '#FFF';
        //it was not found so unselect it
        }else {
            target.style.backgroundColor = '#E0E0E0';
            selectedHolder.push(target.textContent)
        }
        this.setState({
            selected: selectedHolder
        })
        // TODO: add redux to save all selections for diets
    }

    render() {
        if (this.props.title === "Everything") {
            return (
                <div className="dietCard mdl-card mdl-shadow--2dp mdl-cell mdl-cell--2-col mdl-cell--2-col-phone mdl-cell--8-col-tablet" onClick={(e) => this.clickAction(e)} >
                    <div className="mdl-card__title dietName">{this.props.title}</div>
                </div>
            );
        }
        return (
            <div className="dietCard mdl-card mdl-shadow--2dp mdl-cell mdl-cell--2-col mdl-cell--1-col-phone mdl-cell--4-col-tablet" onClick={(e) => this.clickAction(e)} >
                <div className="mdl-card__title dietName">{this.props.title}</div>
            </div>
        );
    }
}