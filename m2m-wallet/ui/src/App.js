import React, { Component } from "react";
import {Router} from "react-router-dom";
import { Route, Switch } from 'react-router-dom';
import classNames from "classnames";

import CssBaseline from "@material-ui/core/CssBaseline";
import { MuiThemeProvider, withStyles } from "@material-ui/core/styles";
import Grid from '@material-ui/core/Grid';

import history from "./history";
import theme from "./theme";

import TopNav from "./components/TopNav";
import SideNav from "./components/SideNav";
import Footer from "./components/Footer";
import Notifications from "./components/Notifications";
import SessionStore from "./stores/SessionStore";

// search
//import Search from "./views/search/Search";

//M2M Wallet
import Topup from "./views/m2m-wallet/Topup"
import Withdraw from "./views/m2m-wallet/Withdraw"
import HistoryLayout from "./views/m2m-wallet/HistoryLayout"
import ModifyEthAccount from "./views/m2m-wallet/ModifyEthAccount"

const drawerWidth = 270;

const styles = {
  root: {
    flexGrow: 1,
    display: "flex",
    minHeight: "100vh",
    flexDirection: "column",
    backgroundColor: "#090046",
    background: "#311b92",
    fontFamily: "Montserrat",
  },
  input: {
    color: '#FFFFFF',
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
    backgroundColor: "#090046",
  },
};


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

    /* this.setState({
      user: SessionStore.getUser(),
      drawerOpen: SessionStore.getUser() != null,
    }); */
    this.setState({
      drawerOpen: true,
    });
  }

  setDrawerOpen(state) {
    this.setState({
      drawerOpen: state,
    });
  }

  render() {
    let topNav = null;
    let sideNav = null;

    if (this.state.user !== null) {
      topNav = <TopNav setDrawerOpen={this.setDrawerOpen} drawerOpen={this.state.drawerOpen} user={this.state.user} />;
      sideNav = <SideNav open={this.state.drawerOpen} user={this.state.user} />
    }

    return (
      <Router history={history}>
        <React.Fragment>
          <CssBaseline />
          <MuiThemeProvider theme={theme}>
            <div className={this.props.classes.root}>
              {topNav}
              {sideNav}
              <div className={classNames(this.props.classes.main, this.state.drawerOpen && this.props.classes.mainDrawerOpen)}>
                <Grid container spacing={24}>
                  <Switch>
                    <Route exact path="/" component={Withdraw} />
                    {/* <Route exact path="/withdraw/:organizationID(\d+)" component={Withdraw} /> */}
                    <Route exact path="/withdraw" component={Withdraw} />
                    <Route exact path="/topup" component={Topup} />
                    <Route path="/history" component={HistoryLayout} />
                    <Route exact path="/modify-account" component={ModifyEthAccount} />
                    
                  </Switch>
                </Grid>
              </div>
              <div className={this.state.drawerOpen ? this.props.classes.footerDrawerOpen : ""}>
                <Footer />
              </div>
            </div>
            <Notifications />
          </MuiThemeProvider>
        </React.Fragment>
      </Router>
    );
  }
}

export default withStyles(styles)(App);
