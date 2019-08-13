import React, { Component } from "react";
import { withRouter, Link } from 'react-router-dom';
import { withStyles } from "@material-ui/core/styles";

import Grid from '@material-ui/core/Grid';
import TitleBar from "../../components/TitleBar";
import TitleBarTitle from "../../components/TitleBarTitle";
import Divider from '@material-ui/core/Divider';
import Spinner from "../../components/ScaleLoader";
import TopupStore from "../../stores/TopupStore";
import MoneyStore from "../../stores/MoneyStore";
import TopupForm from "./TopupForm";
import { ETHER } from "../../util/Coin-type"
import styles from "./TopupStyle"

function loadSuperNodeActiveMoneyAccount(organizationID) {
  return new Promise((resolve, reject) => {
      TopupStore.getTopUpDestination(ETHER, organizationID, resp => {
        resolve(resp.activeAccount);
      });
    
  });
}
      
function loadActiveMoneyAccount(organizationID) {
  return new Promise((resolve, reject) => {
    MoneyStore.getActiveMoneyAccount(ETHER, organizationID, resp => {
      resolve(resp.activeAccount);
    });
  });
}

class Topup extends Component {
  constructor(props) {
    super(props);
    this.state = {
      loading: false,
    };
    this.loadData = this.loadData.bind(this);
  }
  
  componentDidMount() {
    this.loadData();
  }
  
  componentDidUpdate(oldProps) {
    if (this.props === oldProps) {
      return;
    }

    this.loadData();
  }

  loadData = async () => {
    try {
      const organizationID = this.props.match.params.organizationID;
      this.setState({loading: true})
      var superNodeAccount = await loadSuperNodeActiveMoneyAccount(organizationID);
      var account = await loadActiveMoneyAccount(organizationID);
      
      const accounts = {};
      accounts.superNodeAccount = superNodeAccount;
      accounts.account = account;

      this.setState({
        accounts
      }); 
      this.setState({loading: false})
    } catch (error) {
      this.setState({loading: false})
      console.error(error);
      this.setState({ error });
    }
  }

  render() {
    return(
      <Grid container spacing={24}>
        <Spinner on={this.state.loading}/>
        <Grid item xs={12} className={this.props.classes.divider}>
          <div className={this.props.classes.TitleBar}>
                <TitleBar className={this.props.classes.padding}>
                  <TitleBarTitle title="Top up" />
                </TitleBar>
                <Divider light={true}/>
                <div className={this.props.classes.breadcrumb}>
                <TitleBar>
                  <TitleBarTitle component={Link} to="#" title="M2M Wallet" className={this.props.classes.link}/> 
                  <TitleBarTitle title="/" className={this.props.classes.navText}/>
                  <TitleBarTitle component={Link} to="#" title="Top up" className={this.props.classes.link}/>
                </TitleBar>
                </div>
            </div>
        </Grid>
        <Grid item xs={6} className={this.props.classes.column}>
          <TitleBarTitle title="Send Tokens" />
          <Divider light={true}/>
          <TopupForm
            reps={this.state.accounts} {...this.props}
            orgId ={this.props.match.params.organizationID} 
          />
            
        </Grid>
        <Grid item xs={6}>
        </Grid>
      </Grid>
    );
  }
}

export default withStyles(styles)(withRouter(Topup));