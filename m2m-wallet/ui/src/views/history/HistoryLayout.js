import React, { Component } from "react";
import { Route, Switch, Link, withRouter, Redirect } from "react-router-dom";

import { withStyles } from "@material-ui/core/styles";
import Grid from '@material-ui/core/Grid';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';

import TitleBar from "../../components/TitleBar";
import TitleBarTitle from "../../components/TitleBarTitle";
import Divider from '@material-ui/core/Divider';

//import SessionStore from "../../stores/SessionStore";

import Transactions from "./Transactions";
import EthAccount from "./EthAccount";
import SubScriptions from "./SubScriptions";

import styles from "./HistoryStyle";


class HistoryLayout extends Component {
  constructor() {
    super();
    this.state = {
      tab: 0,
      admin: false,
    };

    this.onChangeTab = this.onChangeTab.bind(this);
    this.locationToTab = this.locationToTab.bind(this);
  }

  componentDidMount() {
    this.locationToTab();
  }

  componentDidUpdate(oldProps) {
    if (this.props.match.url === oldProps.match.url) {
      return;
    }

    this.locationToTab();
  }

  onChangeTab(e, v) {
    this.setState({
      tab: v,
    });
  }

  locationToTab() {
    let tab = 0;

    if (window.location.href.endsWith("/eth_account")) {
      tab = 1;
    } else if (window.location.href.endsWith("/Withdraw")) {
      tab = 2;
    } 

    this.setState({
      tab,
    });
  }

  render() {
    const organizationID = this.props.match.params.organizationID;
    
    return(
      <Grid container spacing={24}>
        <Grid item xs={12} className={this.props.classes.divider}>
          <div className={this.props.classes.TitleBar}>
                <TitleBar className={this.props.classes.padding}>
                  <TitleBarTitle title="History" />
                </TitleBar>
                <Divider light={true}/>
                <div className={this.props.classes.breadcrumb}>
                <TitleBar>
                  <TitleBarTitle component={Link} to="#" title="M2M Wallet" className={this.props.classes.link}/> 
                  <TitleBarTitle title="/" className={this.props.classes.navText}/>
                  <TitleBarTitle component={Link} to="#" title="History" className={this.props.classes.link}/>
                </TitleBar>
                </div>
            </div>
        </Grid>

        <Grid item xs={12}>
          <Tabs
            value={this.state.tab}
            onChange={this.onChangeTab}
            indicatorColor="primary"
            className={this.props.classes.tabs}
            variant="scrollable"
            scrollButtons="auto"
            textColor="primary"
          >
            <Tab label="Transactions" component={Link} to={`/history/${organizationID}/transactions`} />
            <Tab label="ETH Account" component={Link} to={`/history/${organizationID}/eth_account`} />
            <Tab label="Withdraw" component={Link} to={`/history/${organizationID}/Withdraw`} />
            
          </Tabs>
        </Grid>

        <Grid item xs={12}>
          <Switch>
            <Route exact path={`${this.props.match.path}/${organizationID}/transactions`} render={props => <Transactions {...props} />} />
            <Route exact path={`${this.props.match.path}/${organizationID}/eth_account`} render={props => <EthAccount {...props} />} />
            <Route exact path={`${this.props.match.path}/${organizationID}/Withdraw`} render={props => <SubScriptions {...props} />} />

            <Redirect to={`/history/${organizationID}/transactions`} />
          </Switch>
        </Grid>
      </Grid>
    );
  }
}

export default withStyles(styles)(withRouter(HistoryLayout));
