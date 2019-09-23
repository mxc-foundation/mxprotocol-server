import React, { Component } from "react";
import {Router} from "react-router-dom";
import { Route, Switch, Redirect } from 'react-router-dom';
import classNames from "classnames";
import { BackToLora } from "./util/M2mUtil";

import CssBaseline from "@material-ui/core/CssBaseline";
import { MuiThemeProvider, withStyles } from "@material-ui/core/styles";
import Grid from '@material-ui/core/Grid';

import history from "./history";
import theme from "./theme";

import TopNav from "./components/TopNav";
import TopBanner from "./components/TopBanner";
import SideNav from "./components/SideNav";
import Footer from "./components/Footer";
import Notifications from "./components/Notifications";
import SessionStore from "./stores/SessionStore";
//import ProfileStore from "./stores/ProfileStore";

// search
//import Search from "./views/search/Search";

//M2M Wallet
import Topup from "./views/topup/Topup"
import Withdraw from "./views/withdraw/Withdraw"
import HistoryLayout from "./views/history/HistoryLayout"
import ModifyEthAccount from "./views/ethAccount/ModifyEthAccount"


const drawerWidth = 270;

const styles = {
  outerRoot: {
    width: '100%',
    background: "#311b92",
  },
  root: {
    //width: '1024px',
    flexGrow: 1,
    margin: 'auto',
    display: "flex",
    minHeight: "100vh",
    flexDirection: "column",
    backgroundColor: theme.palette.darkBG.main,
    background: theme.palette.secondary.secondary,
    fontFamily: "Montserrat",
  },
  input: {
    color: theme.palette.textPrimary.main,
  },
  paper: {
    padding: theme.spacing.unit * 2,
    textAlign: 'center',
    color: theme.palette.text.secondary,
  },
  main: {
    width: "100%",
    padding: 2 * 24,
    paddingTop: 115,
    flex: 1,
  },

  mainDrawerOpen: {
    paddingLeft: drawerWidth + (2 * 24),
  },
  footerDrawerOpen: {
    paddingLeft: drawerWidth,
  },
  color: {
    backgroundColor: theme.palette.darkBG.main,
  },
};

class RedirectedFromLora extends Component {
  constructor(...args) {
    super(...args);

    this.state = {}
  }
  
  formatNumber(number) {
    return number.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",");
  }

  render() {
    const { match: { params: { data: dataString } }} = this.props;

    const data = JSON.parse(decodeURIComponent(dataString) || '{}');
    const { path } = data;
    SessionStore.initProfile(data);
    
    return <Redirect to={path} />; 
  }
}

class App extends Component {
  constructor() {
    super();

    this.state = {
      user: true,
      drawerOpen: true,
    };

    this.setDrawerOpen = this.setDrawerOpen.bind(this);
  }

  componentDidMount() {
    SessionStore.on("change", () => {
      this.setState({
        user: SessionStore.getUser(),
        drawerOpen: SessionStore.getUser() != null,
      });
    });

    this.setState({
      drawerOpen: true
    });
  }

  setDrawerOpen(state) {
    this.setState({
      drawerOpen: state,
    });
  }

  rerenderApp = () => {
    this.forceUpdate();
  }

  render() {
    let topNav = null;
    let sideNav = null;
    let topbanner = null;

    if (this.state.user !== null) {
      topNav = <TopNav setDrawerOpen={this.setDrawerOpen} drawerOpen={this.state.drawerOpen} username={SessionStore.getUsername()} />;
      topbanner = <TopBanner setDrawerOpen={this.setDrawerOpen} drawerOpen={this.state.drawerOpen} user={this.state.user} organizationId={this.state.organizationId}/>;
      sideNav = <SideNav initProfile={SessionStore.initProfile} open={this.state.drawerOpen} organizationID={SessionStore.getOrganizationID()}/>
    }
    
    return (
      <Router history={history}>
        <React.Fragment>
          <CssBaseline />
          <MuiThemeProvider theme={theme}>
            <div className={this.props.classes.outerRoot}>
            <div className={this.props.classes.root}>
              {topNav}
              {topbanner}
              {sideNav}
              <div className={classNames(this.props.classes.main, this.state.drawerOpen && this.props.classes.mainDrawerOpen)}>
                <Grid container spacing={24}>
                  <Switch>
                    <Route path="/j/:data" component={RedirectedFromLora} />
                    <Route path="/withdraw/:organizationID" component={Withdraw} />
                    <Route path="/topup/:organizationID" component={Topup} />
                    <Route path="/history/:organizationID" component={HistoryLayout} />
                    <Route path="/modify-account/:organizationID" component={ModifyEthAccount} />

                    <Route render={BackToLora} />
                  </Switch>
                </Grid>
              </div>
              <div className={this.state.drawerOpen ? this.props.classes.footerDrawerOpen : ""}>
                <Footer />
              </div>
            </div>
            <Notifications />
            </div>
          </MuiThemeProvider>
        </React.Fragment>
      </Router>
    );
  }
}

export default withStyles(styles)(App);
