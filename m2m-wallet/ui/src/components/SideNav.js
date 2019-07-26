import React, { Component } from "react";
import { Link, withRouter } from "react-router-dom";

import { withStyles } from "@material-ui/core/styles";
import Drawer from '@material-ui/core/Drawer';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';

import Divider from '@material-ui/core/Divider';
import AutocompleteSelect from "./AutocompleteSelect";

import CalendarCheckOutline from "mdi-material-ui/CalendarCheckOutline";
import CreditCard from "mdi-material-ui/CreditCard";
import AccessPoint from "mdi-material-ui/AccessPoint";

import ProfileStore from "../stores/ProfileStore"
import SessionStore from "../stores/SessionStore"
import PageNextOutline from "mdi-material-ui/PageNextOutline";
import PagePreviousOutline from "mdi-material-ui/PagePreviousOutline";
import styles from "./SideNavStyle";



const LinkToLora = ({children, ...otherProps}) => 
<a href={SessionStore.getLoraHostUrl()} {...otherProps}>{children}</a>;

const coinType = 'Ether';

function updateOrganizationList(org_id) {
  /* return new Promise((resolve, reject) => {
    resolve(ProfileStore.getUserOrganizationList(org_id));
  }); */

  return new Promise((resolve, reject) => {
    ProfileStore.getUserOrganizationList(org_id,
      resp => {
        resolve(resp);
      })
  });
}

class SideNav extends Component {
  constructor() {
    super();

    this.state = {
      open: true,
      organization: {},
      options: SessionStore.getOrganizationList(),
      organizationID: '',
      cacheCounter: 0,
    };

    this.onChange = this.onChange.bind(this);
    this.getOrganizationOptions = this.getOrganizationOptions.bind(this);
  }

  handleMXC = () => {
    window.location.replace(`http://mxc.org/`);
  } 

  componentDidMount() {
    const organizationID = SessionStore.getOrganizationID();
   
    this.setState({
      organizationID,
    })
    
    /* SessionStore.on("organizationList.change", () => {
      const organizationID = SessionStore.getOrganizationID();
      const options = SessionStore.getOrganizationList();
      
      this.setState({
        organizationID,
        options
      })
    });   */
  }

  onChange(e) {
    SessionStore.setOrganizationID(e.target.value);
    
    this.setState({
      organizationID: e.target.value
    })
    this.props.history.push(`/withdraw/${e.target.value}`);
  }

  getOrganizationFromLocation() {
    /* const organizationRe = /\/organizations\/(\d+)/g;
    const match = organizationRe.exec(this.props.history.location.pathname);

    if (match !== null && (this.state.organization === null || this.state.organization.id !== match[1])) {
      SessionStore.setOrganizationID(match[1]);
    } */
  }

  getOrganizationOptions(search, callbackFunc) {
    let options = this.state.options;
    return callbackFunc(options);
  }

  selectClicked = async () => {
    const res = await updateOrganizationList(this.state.organizationID);
  }

  render() {
    const { organizationID } = this.state;
    const { pathname } = this.props.location;

    const active = (path) => Boolean(pathname.match(path));
    const selected = (path) => {
      if(Boolean(pathname.match(path))){
        return { primary: this.props.classes.selected };
      }else{
        return {};
      }
    }

    /*setTimeout(function() {
      const select = document.querySelector('.Select .Select-value');
      
      select && select.addEventListener('click', function(event) {
        console.log(event.target, event.currentTarget)
      });
      window.addEventListener('click', function(event) {
        console.log(event.target, event.currentTarget)
        // .Select
      })
    })*/

    return(
      <Drawer
        variant="persistent"
        anchor="left"
        open={this.props.open}
        classes={{paper: this.props.classes.drawerPaper}}
      >
        {this.state.organization && <List className={this.props.classes.static}>
          {/* <ListItem button component={Link} to={`/withdraw/${this.state.organization.id}`}> */}
          <Divider />
          <div>
          <AutocompleteSelect
            id="organizationID"
            margin="none"
            value={organizationID}
            updateOptions={this.selectClicked}
            onChange={this.onChange}//{this.state.organization && 
            getOptions={this.getOrganizationOptions}
            className={this.props.classes.select}
            triggerReload={this.state.cacheCounter}
          />
        </div>
          <ListItem selected={active('/withdraw')} button component={Link} to={`/withdraw/${this.state.organizationID}`}>
            <ListItemIcon className={this.props.classes.iconStyle}>
              <PagePreviousOutline />
            </ListItemIcon>
            <ListItemText classes={selected('/withdraw')} primary="Withdraw" />
          </ListItem>
          <ListItem selected={active('/topup')} button component={Link} to={`/topup/${this.state.organizationID}`}>
            <ListItemIcon>
              <PageNextOutline />
            </ListItemIcon>
            <ListItemText classes={selected('/topup')} primary="Top up" />
          </ListItem>
          <ListItem selected={active('/history')} button component={Link} to={`/history/${this.state.organizationID}`}>
            <ListItemIcon>
              <CalendarCheckOutline />
            </ListItemIcon>
            <ListItemText classes={selected('/history')} primary="History" />
          </ListItem>
          <ListItem selected={active('/modify-account')} button component={Link} to={`/modify-account/${this.state.organizationID}`}>
            <ListItemIcon>
              <CreditCard />
            </ListItemIcon>
            <ListItemText classes={selected('/modify-account')} primary="ETH Account" />
          </ListItem>
              <List className={this.props.classes.card}>
                <ListItem button component={LinkToLora} className={this.props.classes.static}>  
                  <ListItemIcon>
                    <AccessPoint />
                  </ListItemIcon>
                  <ListItemText primary="LoRa Server" />
                </ListItem>
                <Divider />
                <Divider />
                <ListItem>
                  <ListItemText primary="Powered by" />
                  <ListItemIcon>
                    <img src="/logo/mxc_logo.png" className="iconStyle" alt="LoRa Server" onClick={this.handleMXC} />
                  </ListItemIcon>
                </ListItem>
              </List>
        </List>}
      </Drawer>
    );
  }
}

export default withRouter(withStyles(styles)(SideNav));
