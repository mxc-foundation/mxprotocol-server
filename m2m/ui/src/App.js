import React, { Component, Suspense } from "react";
import {Router} from "react-router-dom";
import { Route, Switch, Redirect } from 'react-router-dom';
import classNames from "classnames";
import { BackToLora } from "./util/M2mUtil";

import CssBaseline from "@material-ui/core/CssBaseline";
import { MuiThemeProvider, withStyles } from "@material-ui/core/styles";
import Grid from '@material-ui/core/Grid';

import history from "./history";
import theme from "./theme";
import i18n, { packageNS, DEFAULT_LANGUAGE, SUPPORTED_LANGUAGES } from "./i18n";

import TopNav from "./components/TopNav";
import TopBanner from "./components/TopBanner";
//import SideNav from "./components/SideNav"; [edit]
import Sidebar from "./components/Sidebar";
import Topbar from "./components/Topbar";
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
import SuperNodeEth from "./views/ControlPanel/superNodeEth/superNodeEth"
import DeviceLayout from "./views/device/DeviceLayout";
import GatewayLayout from "./views/gateway/GatewayLayout";
import StakeLayout from "./views/Stake/StakeLayout";
import SetStake from "./views/Stake/SetStake";

import SuperAdminWithdraw from "./views/ControlPanel/withdraw/withdraw"
import SupernodeHistory from "./views/ControlPanel/history/History"
import SystemSettings from "./views/ControlPanel/settings/settings"

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
      language: null
    };

    this.setDrawerOpen = this.setDrawerOpen.bind(this);
  }
k
  componentDidMount() {
    SessionStore.on("change", () => {
      this.setState({
        user: SessionStore.getUser(),
        drawerOpen: SessionStore.getUser() != null,
        language: SessionStore.getLanguage()
      });
    });

    const storedLanguageID = SessionStore.getLanguage() && SessionStore.getLanguage().id;

    if (!storedLanguageID && !i18n.language) {
      i18n.changeLanguage(DEFAULT_LANGUAGE.id, (err, t) => {
        if (err) {
          console.error(`Error setting default language to English: `, err);
        }
      });
    }

    const i18nLanguage = SUPPORTED_LANGUAGES.find(el => el.id === i18n.language);

    // Add the saved i18n language back into Local Storage if it is lost after a page refresh on Login component
    if (!storedLanguageID && i18n.language) {
      SessionStore.setLanguage(i18nLanguage);
    }

    // Language stored in Local Storage persists and takes precedence over i18n language
    if (storedLanguageID && i18n.language !== storedLanguageID) {
      i18n.changeLanguage(storedLanguageID, (err, t) => {
        if (err) {
          console.error(`Error loading language ${storedLanguageID}: `, err);
        }
      });
    }

    this.setState({
      drawerOpen: true,
      language: storedLanguageID ? SessionStore.getLanguage() : i18nLanguage
    });
  }

  onChangeLanguage = (newLanguage) => {
    SessionStore.setLanguage(newLanguage);

    i18n.changeLanguage(newLanguage.id, (err, t) => {
      if (err) {
        console.error(`Error loading language ${newLanguage.id}: `, err);
      }
    });

    this.setState({
      language: newLanguage
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

    /**
     * toggle Menu
     */
    toggleMenu = (e) => {
        e.preventDefault();
        this.setState({ isCondensed: !this.state.isCondensed });
    }

    /**
     * Toggle right side bar
     */
    toggleRightSidebar = () => {
        document.body.classList.toggle("right-bar-enabled");
    }

  render() {
    const { language } = this.state;

    let topNav = null;
    let sideNav = null;
    let topbanner = null;

    if (this.state.user !== null) {
      topNav = (
        <TopNav
          drawerOpen={this.state.drawerOpen}
          language={language}
          onChangeLanguage={this.onChangeLanguage}
          setDrawerOpen={this.setDrawerOpen}
          username={SessionStore.getUsername()}
        />
      );
      //topbanner = <TopBanner setDrawerOpen={this.setDrawerOpen} drawerOpen={this.state.drawerOpen} user={this.state.user} organizationId={this.state.organizationId}/>; [edit]
      //sideNav = <SideNav initProfile={SessionStore.initProfile} open={this.state.drawerOpen} organizationID={SessionStore.getOrganizationID()}/>; [edit]
      topbanner = <Topbar rightSidebarToggle={this.toggleRightSidebar} menuToggle={this.toggleMenu} {...this.props} />;
      sideNav = <Sidebar isCondensed={this.state.isCondensed} {...this.props} />;
    }
    
    return (
      <Router history={history}>
        <React.Fragment>
          <CssBaseline />
          <MuiThemeProvider theme={theme}>
            {/* <div className={this.props.classes.outerRoot}>
            <div className={this.props.classes.root}> [edit]*/}
            <div className="app">
                <div id="wrapper">
              {topNav}
              {topbanner}
              {sideNav}
              {topNav}
              <div className={classNames(this.props.classes.main, this.state.drawerOpen && this.props.classes.mainDrawerOpen)}>
                <Grid container spacing={24}>
                  <Switch>
                    <Route path="/j/:data" component={RedirectedFromLora} />
                    <Route path="/withdraw/:organizationID" component={Withdraw} />
                    <Route path="/topup/:organizationID" component={Topup} />
                    <Route path="/history/:organizationID" component={HistoryLayout} />
                    <Route path="/modify-account/:organizationID" component={ModifyEthAccount} />
                    <Route path="/device/:organizationID" component={DeviceLayout} />
                    <Route path="/gateway/:organizationID" component={GatewayLayout} />
                    <Route exact path="/stake/:organizationID" component={StakeLayout} />
                    <Route exact path="/stake/:organizationID/set-stake" component={SetStake} />
                    <Route path="/control-panel/modify-account" component={SuperNodeEth} />
                    <Route path="/control-panel/withdraw" component={SuperAdminWithdraw} />
                    <Route path="/control-panel/history" component={SupernodeHistory} />
                    <Route path="/control-panel/system-settings" component={SystemSettings} />
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
