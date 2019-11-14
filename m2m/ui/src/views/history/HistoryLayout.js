import React, { Component } from "react";
import { Route, Switch, Link, withRouter } from "react-router-dom";

import { withStyles } from "@material-ui/core/styles";
import Grid from '@material-ui/core/Grid';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import Button from "@material-ui/core/Button";

import TitleBar from "../../components/TitleBar";
import TitleBarTitle from "../../components/TitleBarTitle";
import Spinner from "../../components/ScaleLoader"
//import SessionStore from "../../stores/SessionStore";

//import Transactions from "./Transactions";
import StakeStore from "../../stores/StakeStore";
import EthAccount from "./EthAccount";
import Transactions from "./Transactions";
import NetworkActivityHistory from "./NetworkActivityHistory";
import Stakes from "./Stakes";

import styles from "./HistoryStyle";


class HistoryLayout extends Component {
  constructor(props) {
    super(props);
    this.state = {
      tab: 0,
      loading: false,
      admin: false,
    };

    this.onChangeTab = this.onChangeTab.bind(this);
    this.locationToTab = this.locationToTab.bind(this);
  }

  componentDidMount() {
    const prevLoc = this.props.location.search.split('=')[1];
    this.setState({loading:true});
    this.locationToTab(prevLoc);
    this.setState({loading:false});
  }

  componentDidUpdate(oldProps) {
    if (this.props == oldProps) {
      return;
    }

    this.locationToTab();
  }

  onChangeTab(e, v) {
    this.setState({
      tab: v,
    });
  }

  locationToTab(prevLoc) {
    let tab = 0;
    if (window.location.href.endsWith("/eth-account")) {
      tab = 1;
    } else if (window.location.href.endsWith("/network-activity")) {
      tab = 2;
    } else if (window.location.href.endsWith("/stake")) {
      tab = 3;
    }  
    
    this.setState({
      tab,
    });
  }

  unstake = (e) => {
    e.preventDefault();
    const resp = StakeStore.unstake(this.props.match.params.organizationID);
    resp.then((res) => {
    })
  }

  render() {
    const organizationID = this.props.match.params.organizationID;
    
    return(
      <Grid container spacing={24}>
        <Spinner on={this.state.loading}/>
        <Grid item xs={12} className={this.props.classes.divider}>
          <div className={this.props.classes.TitleBar}>
              {/* <TitleBar className={this.props.classes.padding}>
                <TitleBarTitle title="Stake" />
              </TitleBar> */}    
              {/* <Divider light={true}/> */}
              <div className={this.props.classes.between}>
              <TitleBar>
                <TitleBarTitle title="History" />
                {/* <TitleBarTitle component={Link} to="#" title="M2M Wallet" className={this.props.classes.link}/> 
                <TitleBarTitle component={Link} to="#" title="/" className={this.props.classes.link}/>
                <TitleBarTitle component={Link} to="#" title="Devices" className={this.props.classes.link}/> */}
              </TitleBar>
              {this.state.tab === 3 && <div className={this.props.classes.alignCol}>
                <Button color="primary.main" component={Link} to={`/stake/${this.props.match.params.organizationID}/set-stake`} /* onClick={this.handleOpenAXS} */ type="button" disabled={false}>CHECK STAKE</Button>
                <Button variant="outlined" color="inherit" component={Link} to={`/stake/${this.props.match.params.organizationID}/set-stake`} onClick={this.unstake} type="button" disabled={false}>UNSTAKE</Button>
              </div>}
              {/* <TitleBarButton
                label="SET STAKE"
                color="primary"
                to={`/stake/${this.props.match.params.organizationID}/set-stake`}
                classes={this.props.classes}
              /> */}
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
            <Tab label="Transactions" component={Link} to={`/history/${organizationID}/`} />
            <Tab label="ETH Account" component={Link} to={`/history/${organizationID}/eth_account`} />
            <Tab label="Network Activity" component={Link} to={`/history/${organizationID}/network-activity`} />
            <Tab label="Staking" component={Link} to={`/history/${organizationID}/stake`} />
          </Tabs>
        </Grid>

        <Grid item xs={12}>
          <Switch>
            <Route exact path={`${this.props.match.path}/`} render={props => <Transactions organizationID={organizationID} {...props} />} />
            <Route exact path={`${this.props.match.path}/eth_account`} render={props => <EthAccount {...props} />} />
            <Route exact path={`${this.props.match.path}/network-activity`} render={props => <NetworkActivityHistory organizationID={organizationID} {...props} />} />
            <Route exact path={`${this.props.match.path}/stake`} render={props => <Stakes organizationID={organizationID} {...props} />} />
            {/* <Redirect to={`/history/${organizationID}/transactions`} /> */}
          </Switch>
        </Grid>
      </Grid>
    );
  }
}

export default withStyles(styles)(withRouter(HistoryLayout));
