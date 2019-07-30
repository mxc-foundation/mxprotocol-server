import React, { Component } from 'react';
import AsyncSelect from 'react-select/async';
import ProfileStore from '../stores/ProfileStore'

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
    ProfileStore.getUserOrganizationList("",
      resp => {
        resolve(getOrgList(resp.organizations));
      })
  });

export default class WithPromises extends Component {
    constructor() {
        super();
        this.state = {
            
        };
    } 

    sdfaf = () => {
        if(Object.getOwnPropertyNames(this.props.default).length !== 0){
            this.setState({value: this.props.default});
            return this.props.default;
        }
    }
    componentDidMount() {
        const option = promiseOptions();
        option.then((resp)=>{
            this.setState({value:resp[0]})
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
        return (
            <AsyncSelect 
                cacheOptions 
                defaultOptions 
                defaultValue={this.state.value}
                onChange={this.onChange}
                loadOptions={promiseOptions} 
            />
        );
    }
}