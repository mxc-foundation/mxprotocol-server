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
            const sa = options.find(opt => opt.label.toLowerCase() === 'super_admin')
            
            this.setState({
                options,
                value: SessionStore.getOrganizationID() === "0"
                  ? options[2] : options[0]
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
    
    render() {
        //console.log("SessionStore: ", SessionStore.getUserOrganizationList());
        return (
            <AsyncSelect 
                cacheOptions 
                defaultOptions
                //defaultValue={}
                //inputValue={this.state.value}
                onChange={this.onChange}
                loadOptions={promiseOptions}
                //options={this.state.options}
            />
        );
    }
}