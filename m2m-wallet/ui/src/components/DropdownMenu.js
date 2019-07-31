import React, { Component } from 'react';
import AsyncSelect from 'react-select/async';
import ProfileStore from '../stores/ProfileStore'
import SessionStore from "../stores/SessionStore";

const getOrgList = (organizations) => {
    let organizationList = null;
    if(organizations){
        organizationList = organizations.map((o, i) => { 
        return {label: o.organizationName, value: o.organizationID}});
    }
    
    return organizationList;
};

const promiseOptions = () =>
  new Promise((resolve, reject) => {
    ProfileStore.getUserOrganizationList(SessionStore.getOrganizationID(),
      resp => {
        resolve(getOrgList(resp.organizations));
      })
  });

export default class WithPromises extends Component {
    constructor() {
        super();
        this.state = {
            selectedValue: null,
            options:[]
        };
    } 

    componentDidMount() {
        promiseOptions().then(options => {
            this.setState({
                options
            })
        })
    }

    onChange = (v) => {
        let value = null;
        if (v !== null) {
            value = v.value;
        }

        this.props.onChange({
            target: {
                id: this.props.id,
                value: value,
            },
        });
    }
    onClick = (v) => {
        console.log('onClick',v);
    }
    render() {
        //console.log("SessionStore: ", SessionStore.getUserOrganizationList());
        const dValue = {label:SessionStore.getOrganizationName(), value: SessionStore.getOrganizationID()}; 
        return (
            <AsyncSelect 
                cacheOptions 
                defaultOptions
                defaultValue={dValue}
                onClick={this.onClick}
                //defaultValue={this.state.value}
                //inputValue={this.state.value}
                onChange={this.onChange}
                loadOptions={promiseOptions}
                //options={this.state.options}
            />
        );
    }
}