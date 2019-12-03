import React, { Component } from "react";
import { withRouter, Link } from 'react-router-dom';
import { withStyles } from "@material-ui/core/styles";

import Grid from '@material-ui/core/Grid';
import i18n, { packageNS } from '../../i18n';
import TitleBar from "../../components/TitleBar";
import TitleBarTitle from "../../components/TitleBarTitle";
import Divider from '@material-ui/core/Divider';
import Spinner from "../../components/ScaleLoader";
import TopupStore from "../../stores/TopupStore";
import MoneyStore from "../../stores/MoneyStore";
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import TopupForm from "./TopupForm";
import { ETHER } from "../../util/Coin-type"
import InfoCard from "./InfoCard"; 
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
        <Grid item xs={12} md={12} lg={12} className={this.props.classes.divider}>
          <div className={this.props.classes.TitleBar}>
              <TitleBar className={this.props.classes.padding}>
                <TitleBarTitle title={i18n.t(`${packageNS}:menu.topup.topup`)} />
              </TitleBar>
          </div>
        </Grid>
        <Grid item xs={12} md={12} lg={6} className={this.props.classes.column}>
          <Card className={this.props.classes.card}>
            <CardContent>
              <TopupForm
                reps={this.state.accounts} {...this.props}
                orgId ={this.props.match.params.organizationID} 
              />
            </CardContent>
          </Card>
        </Grid>
        <Grid item xs={12} md={12} lg={6} className={this.props.classes.column}>
          <InfoCard orgId={this.props.match.params.organizationID} />
        </Grid>
      </Grid>
    );
  }
}

export default withStyles(styles)(withRouter(Topup));