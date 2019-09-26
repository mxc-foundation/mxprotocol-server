import React, { Component } from "react";
import { Link, withRouter } from "react-router-dom";

import { withStyles } from "@material-ui/core/styles";
import Drawer from '@material-ui/core/Drawer';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';

import Divider from '@material-ui/core/Divider';
import DropdownMenu from "./DropdownMenu";

import CalendarCheckOutline from "mdi-material-ui/CalendarCheckOutline";
import CreditCard from "mdi-material-ui/CreditCard";
import AccessPoint from "mdi-material-ui/AccessPoint";

import ProfileStore from "../stores/ProfileStore"
import SessionStore from "../stores/SessionStore"
import PageNextOutline from "mdi-material-ui/PageNextOutline";
import PagePreviousOutline from "mdi-material-ui/PagePreviousOutline";
import { getLoraHost } from "../util/M2mUtil";
import styles from "./SideNavStyle";



const LinkToLora = ({children, ...otherProps}) => 
<a href={getLoraHost()} {...otherProps}>{children}</a>;

function updateOrganizationList(orgId) {
  return new Promise((resolve, reject) => {
    ProfileStore.getUserOrganizationList(orgId,
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
      //organization: {},
      organizationID: '',
      cacheCounter: 0,
    };
    this.onChange = this.onChange.bind(this);
  }

  handleMXC = async () => {
    window.location.replace(`http://mxc.org/`);
  } 

  loadData = async () => {
    try {
      const organizationID = SessionStore.getOrganizationID();
      this.setState({
        organizationID,
      })
      this.setState({loading: true})
      
    } catch (error) {
      this.setState({loading: false})
      console.error(error);
      this.setState({ error });
    }
  }
  componentDidMount() {
    this.loadData();
  }

  onChange(e) {
    SessionStore.setOrganizationID(e.target.value);
    
    this.setState({
      organizationID: e.target.value
    })
    
    const currentLocation = this.props.history.location.pathname.split('/')[1];
    this.props.history.push(`/${currentLocation}/${e.target.value}`);
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

    return(
      <Drawer
        variant="persistent"
        anchor="left"
        open={this.props.open}
        classes={{paper: this.props.classes.drawerPaper}}
      >
        {organizationID && <List className={this.props.classes.static}>
          {/* <ListItem button component={Link} to={`/withdraw/${this.state.organization.id}`}> */}
          <div>
            <DropdownMenu default={ this.state.default } onChange={this.onChange} />
          </div>
          {/* <Divider /> */}
          <ListItem selected={active('/withdraw')} button component={Link} to={`/withdraw/${organizationID}`}>
            <ListItemIcon className={this.props.classes.iconStyle}>
              <PagePreviousOutline />
            </ListItemIcon>
            <ListItemText classes={selected('/withdraw')} primary="Withdraw" />
          </ListItem>
          <ListItem selected={active('/topup')} button component={Link} to={`/topup/${organizationID}`}>
            <ListItemIcon>
              <PageNextOutline />
            </ListItemIcon>
            <ListItemText classes={selected('/topup')} primary="Top up" />
          </ListItem>
          <ListItem selected={active('/history')} button component={Link} to={`/history/${organizationID}`}>
            <ListItemIcon>
              <CalendarCheckOutline />
            </ListItemIcon>
            <ListItemText classes={selected('/history')} primary="History" />
          </ListItem>
          <ListItem selected={active('/modify-account')} button component={Link} to={`/modify-account/${organizationID}`}>
            <ListItemIcon>
              <CreditCard />
            </ListItemIcon>
            <ListItemText classes={selected('/modify-account')} primary="ETH Account" />
          </ListItem>
              <List className={this.props.classes.card}>
              <Divider />
                <ListItem button component={LinkToLora} className={this.props.classes.static}>  
                  <ListItemIcon>
                    <AccessPoint />
                  </ListItemIcon>
                  <ListItemText primary="LoRa Server" />
                </ListItem>
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
