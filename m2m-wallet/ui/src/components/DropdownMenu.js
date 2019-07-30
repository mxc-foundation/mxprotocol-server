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
          default: null,
          modal: null
        };
    } 

    componentDidMount() {
        console.log('componentDidMount', this.props.default)
        this.setState({
            default:{ value: 'ocean', label: 'Ocean', color: '#00B8D9', isFixed: true }
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
                //defaultValue={this.props.default}
                onChange={this.onChange}
                loadOptions={promiseOptions} 
            />
        );
  }
}