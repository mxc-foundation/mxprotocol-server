import React, { Component } from "react";
import { withRouter } from 'react-router-dom';

import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import { withStyles } from "@material-ui/core/styles";
import Typography from '@material-ui/core/Typography';
import SessionStore from "../stores/SessionStore";
import Avatar from '@material-ui/core/Avatar';
import Chip from '@material-ui/core/Chip';
import AccountCircle from "mdi-material-ui/AccountCircle";
import { IconButton } from "@material-ui/core";
import HelpCircle from "mdi-material-ui/HelpCircle";

import WithdrawStore from "../stores/WithdrawStore";
import WalletStore from "../stores/WalletStore";
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import Wallet from "mdi-material-ui/WalletOutline";
import styles from "./TopNavStyle"
import { SUPER_ADMIN } from "../util/M2mUtil";

import DropdownMenuLanguage from "./DropdownMenuLanguage";

function getWalletBalance() {
  var organizationId = SessionStore.getOrganizationID();
  if (organizationId === undefined) {
    return null;
  }

  if (SessionStore.isAdmin()) {
    organizationId = SUPER_ADMIN
  }

  return new Promise((resolve, reject) => {
    WalletStore.getWalletBalance(organizationId, resp => {
      return resolve(resp);
    });
  });
}

class TopNav extends Component {
  constructor() {
    super();

    this.state = {
      menuAnchor: null,
      balance: null,
      search: ""
    };

    this.handleDrawerToggle = this.handleDrawerToggle.bind(this);
    this.onLanguageChange = this.onLanguageChange.bind(this);
    this.onMenuOpen = this.onMenuOpen.bind(this);
    this.onMenuClose = this.onMenuClose.bind(this);
    this.onLogout = this.onLogout.bind(this);
    this.onSearchChange = this.onSearchChange.bind(this);
    this.onSearchSubmit = this.onSearchSubmit.bind(this);
  }

  componentDidMount() {
    this.loadData();

    SessionStore.on("organization.change", () => {
      this.loadData();
    });
    WithdrawStore.on("withdraw", () => {
      this.loadData();
    });
  }

  loadData = async () => {
    try {
      var result = await getWalletBalance();
      this.setState({ balance: result.balance });

    } catch (error) {
      console.error(error);
      this.setState({ error });
    }
  }

  onLanguageChange(e) {
    const newLanguageID = e.target.label;
    const newLanguageName = e.target.value;
    const newLanguageCode = e.target.code;

    const newLanguageState = {
      id: newLanguageID,
      name: newLanguageName,
      code: newLanguageCode
    }

    this.props.onChangeLanguage(newLanguageState);
  }

  onMenuOpen(e) {
    this.setState({
      menuAnchor: e.currentTarget,
    });
  }

  onMenuClose() {
    this.setState({
      menuAnchor: null,
    });
  }

  onLogout() {
    SessionStore.logout(() => {
      this.props.history.push("/login");
    });
  }

  handleDrawerToggle() {
    this.props.setDrawerOpen(!this.props.drawerOpen);
  }

  onSearchChange(e) {
    this.setState({
      search: e.target.value,
    });
  }

  onSearchSubmit(e) {
    e.preventDefault();
    this.props.history.push(`/search?search=${encodeURIComponent(this.state.search)}`);
  }

  render() {
    /* let drawerIcon;
    if (!this.props.drawerOpen) {
      drawerIcon = <MenuIcon />;
    } else {
      drawerIcon = <Backburger />;
    } */
    const { balance } = this.state;

    //const open = Boolean(this.state.menuAnchor);

    const balanceEl = balance === null ? 
      <span className="color-gray">(no org selected)</span> : 
      balance + " MXC";

    return (
      <AppBar className={this.props.classes.appBar}>
        <Toolbar>
          {/* <IconButton
            color="inherit"
            aria-label="toggle drawerPaper"
            onClick={this.handleDrawerToggle}
            className={this.props.classes.menuButton}
          >
            {drawerIcon}
          </IconButton> */}

          <div className={this.props.classes.flex}>
            <img src="/logo/m2m_logo.png" className={this.props.classes.logo} alt="LoRa Server" />
          </div>

          {/* <div className={this.props.classes.flex}>
            <Typography type="body2" style={{ color: '#FFFFFF', fontFamily: 'Montserrat', fontWeight: 'bold', fontSize: '22px' }} >M2M Wallet</Typography>
          </div> */}

          <List>
            <ListItem>
              <ListItemIcon >
                <Wallet color="primary" className={this.props.classes.iconStyle} />
              </ListItemIcon>
              <ListItemText primary={ balanceEl } classes={{ primary: this.props.classes.noPadding }}/>
            </ListItem>
          </List>

          <Chip
            avatar={
              <Avatar>
                <AccountCircle />
              </Avatar>
            }
            label={this.props.username}
            onClick={this.onMenuOpen}
            classes={{
              avatar: this.props.classes.avatar,
              root: this.props.classes.chip,
            }}
          />

          <div>
            <DropdownMenuLanguage onChange={this.onLanguageChange} />
          </div>

          <a href="https://www.mxc.org/support" target="mxc-support">
            <IconButton className={this.props.classes.iconButton}>
              <HelpCircle />
            </IconButton>
          </a>
        </Toolbar>
      </AppBar>
    );
  }
}

export default withStyles(styles)(withRouter(TopNav));
