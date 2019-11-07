import React, { Component } from "react";
import { Route, Switch, Link, withRouter } from "react-router-dom";
import {Card,Table,TableBody,TableRow,TableCell,Grid,Tabs,Tab,Divider} from '@material-ui/core';
import { withStyles } from "@material-ui/core/styles";
import TitleBar from "../../../components/TitleBar";
import TitleBarTitle from "../../../components/TitleBarTitle";
import Spinner from "../../../components/ScaleLoader"
//import SessionStore from "../../stores/SessionStore";

//import Transactions from "./Transactions";
import EthAccount from "../../history/EthAccount";
import Transactions from "../../history/Transactions";
import NetworkActivityHistory from "../../history/NetworkActivityHistory";

import topupStore from "../../../stores/TopupStore";
import styles from "./HistoryStyle";


class HistoryLayout extends Component {
  constructor(props) {
    super(props);
    this.state = {
      tab: 0,
      loading: false,
      admin: false,
      income:0
    };

    this.onChangeTab = this.onChangeTab.bind(this);
    this.locationToTab = this.locationToTab.bind(this);
  }

  componentDidMount() {
    this.setState({loading:true});
    this.locationToTab();
    this.setState({loading:false});
    this.getIncome();
  }

  componentDidUpdate(oldProps) {
    if (this.props == oldProps) {
      return;
    }

    this.locationToTab();
  }

  getIncome(){
    topupStore.getIncome(0, resp => {
      this.setState({income:resp.amount});
    });
  }

  onChangeTab(e, v) {
    this.setState({
      tab: v,
    });
  }

  locationToTab() {
    let tab = 0;
    if (window.location.href.endsWith("/eth-account")) {
      tab = 1;
    } else if (window.location.href.endsWith("/network-activity")) {
      tab = 2;
    }  
    
    this.setState({
      tab,
    });
  }

  render() {
    const organizationID = 0;
    
    return(
      <Grid container spacing={24}>
        <Spinner on={this.state.loading}/>
        <Grid item container xs={6} direction="column" className={this.props.classes.divider} padding={12}>
       
                <TitleBar >
                  <TitleBarTitle title="History" />
                </TitleBar>
                 <Divider light={true}/>
                <div className={this.props.classes.breadcrumb}>
                <TitleBar>
                  <TitleBarTitle component={Link} to="#" title="M2M Wallet" className={this.props.classes.link}/> 
                  <TitleBarTitle component={Link} to="#" title="/" className={this.props.classes.link}/>
                  <TitleBarTitle component={Link} to="#" title="Control Panel" className={this.props.classes.link}/>
                  <TitleBarTitle component={Link} to="#" title="/" className={this.props.classes.link}/>
                  <TitleBarTitle component={Link} to="#" title="History" className={this.props.classes.link}/>
                </TitleBar>
                </div> 
          
        </Grid>


        <Grid item container xs={12} justify="space-between" className={this.props.classes.tabsBlock}>
          <Tabs
            value={this.state.tab}
            onChange={this.onChangeTab}
            indicatorColor="primary"
            className={this.props.classes.tabs}
            variant="scrollable"
            scrollButtons="auto"
            textColor="primary"
          >
            <Tab label="Transactions" component={Link} to={`/control-panel/history/`} />
            <Tab label="ETH Account" component={Link} to={`/control-panel/history/eth-account`} />
            <Tab label="Network Activity" component={Link} to={`/control-panel/history/network-activity`} />
          </Tabs>

            <Grid container justify="space-between" alignItems="center" className={this.props.classes.card}>
               <Grid item>Last 24h income</Grid>
              <Grid item align="right"><b>{this.state.income}MXC</b></Grid>
            </Grid>
        
        </Grid>

        <Grid item xs={12}>
          <Switch>
            <Route exact path={`/control-panel/history/`} render={props => <Transactions organizationID={organizationID} {...props} />} />
            <Route exact path={`/control-panel/history/eth-account`} render={props => <EthAccount organizationID={organizationID} {...props} />} />
            <Route exact path={`/control-panel/history/network-activity`} render={props => <NetworkActivityHistory organizationID={organizationID} {...props} />} />
            {/* <Redirect to={`/history/${organizationID}/transactions`} /> */}
          </Switch>
        </Grid>
      </Grid>
    );
  }
}

export default withStyles(styles)(withRouter(HistoryLayout));
