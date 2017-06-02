import React from "react"
import Layout from "./Layout.jsx";
import MoodDown from './helpers/moodDrop.jsx';
import TypeDown from './helpers/typeDrop.jsx';
import { Form, Input, TextArea} from 'semantic-ui-react'
import Datetime from 'react-datetime';
import moment from 'moment';
const myHeader = new Headers();
myHeader.append('Authorization', localStorage.getItem("Authorization"));
const baseurl = 'https://dvrapi.leedann.me/v1/'

//The event creation page
export default class CreateEvent extends React.Component {

    handleSubmit(e) {
        e.preventDefault();
        let name = document.getElementById("form-input-control-event-name");
        let desc = document.getElementById("form-textarea-control-description");
        let times = document.getElementsByClassName("rdt");
        let startTime = times[0].getElementsByTagName("input")[0].value
        let endTime = times[1].getElementsByTagName("input")[0].value
        let mood = document.getElementById("moodDown").getElementsByClassName('text')[0].textContent;
        let type = document.getElementById("typeDown").getElementsByClassName('text')[0].textContent;
        if (startTime || endTime || name || mood || type) {
            startTime = startTime.split('.');
            endTime = endTime.split('.');
            let fmtStart = `${startTime[0]} at${startTime[1].toLowerCase()} (PST)`;
            let fmtEnd = `${endTime[0]} at${endTime[1].toLowerCase()} (PST)`;
            var data = {
                "name": name.value,
                "description": desc.value,
                "startTime": fmtStart,
                "endTime": fmtEnd,
                "type": type,
                "mood": mood
            }

            data = JSON.stringify(data);
            var req = {
                method: 'POST',
                headers: myHeader,
                body: data
            }
            fetch(baseurl+'events', req)
            .then((resp) => {
                
            })
            .catch((error) => {
                alert("There was something wrong with your request");
            })

        }


    }


    render() {
        var yesterday = Datetime.moment().subtract( 1, 'day' );
        var valid = function( current ){
            return current.isAfter( yesterday );
        };
        return(
            <Layout title="social">
                    <span className="titleName mdl-cell--12-col mdl-cell--8-col-phone">create event</span>
                    <div className="eventCard mdl-shadow--2dp mdl-grid">
                    <Form className=" mdl-grid mdl-cell mdl-cell--12-col mdl-cell--8-col-phone mdl-cell--8-col-tablet">
                        <div className="eventName mdl-grid mdl-cell mdl-cell--6-col mdl-cell--8-col-phone mdl-cell--8-col-tablet">
                            <span className="eventHeader mdl-cell--12-col mdl-cell--8-col-phone">Event Name</span>
                            <Form.Field required maxLength='100' width={12} id='form-input-control-event-name' control={Input} placeholder='Name' />
                            <span className="eventHeader mdl-cell--12-col mdl-cell--8-col-phone">Event Description</span>
                            <Form.Field maxLength='500' width={12} id='form-textarea-control-description' control={TextArea} placeholder='Description here' />
                        </div>

                        <div className="eventName mdl-cell mdl-cell--6-col mdl-cell--8-col-phone mdl-cell--8-col-tablet">
                            <span className="eventHeader mdl-cell--12-col mdl-cell--8-col-phone">Event Start:</span>
                            <Datetime isValidDate={ valid } dateFormat='LL.' timeFormat='h:mma' input/>
                            <br/>
                            <span className="eventHeader mdl-cell--12-col mdl-cell--8-col-phone">Event End:</span>
                            <Datetime isValidDate={ valid } dateFormat='LL.' timeFormat='h:mma' input/>
                            <br/>
                            <span className="eventHeader mdl-cell--12-col mdl-cell--8-col-phone">Event Mood:</span><br/>
                            <MoodDown/>
                            <br/>
                            <span className="eventHeader mdl-cell--12-col mdl-cell--8-col-phone">Event Type:</span><br/>
                            <TypeDown/>
                        </div>
                        <div className="eventName mdl-cell mdl-cell--12-col mdl-cell--8-col-phone mdl-cell--8-col-tablet">
                            <Form.Button floated='right' size='huge' onClick={(e) => this.handleSubmit(e)}>Create</Form.Button>
                        </div>
                    </Form>
                        {/*<div className="eventName mdl-grid mdl-cell mdl-cell--6-col mdl-cell--8-col-phone mdl-cell--8-col-tablet">
                            <span className="eventHeader mdl-cell--12-col mdl-cell--8-col-phone">Event Name</span>
                        </div>
                        <div className="eventDesc mdl-grid mdl-cell mdl-cell--6-col mdl-cell--8-col-phone mdl-cell--8-col-tablet">
                            <span className="eventHeader mdl-cell--12-col mdl-cell--8-col-phone">Description</span>
                        </div>
                        <div className="eventDate mdl-grid mdl-cell mdl-cell--12-col mdl-cell--8-col-phone mdl-cell--8-col-tablet">
                            <span className="eventHeader mdl-cell--12-col mdl-cell--8-col-phone">Event Times</span>
                        </div>
                        <div className="eventTypeMood mdl-grid mdl-cell mdl-cell--12-col mdl-cell--8-col-phone mdl-cell--8-col-tablet">
                            
                        </div>*/}
                    </div>
            </Layout>
        );
    }
}