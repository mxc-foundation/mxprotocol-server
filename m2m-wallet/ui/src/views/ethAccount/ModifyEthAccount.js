import React, { Component } from "react";
import { withRouter, Link } from 'react-router-dom';
import { withStyles } from "@material-ui/core/styles";

import Grid from '@material-ui/core/Grid';
import TitleBar from "../../components/TitleBar";
import TitleBarTitle from "../../components/TitleBarTitle";
import Divider from '@material-ui/core/Divider';
import MoneyStore from "../../stores/MoneyStore";
import SessionStore from "../../stores/SessionStore";
import ModifyEthAccountForm from "./ModifyEthAccountForm";
import styles from "./EthAccountStyle";

const coinType = "Ether";

class ModifyEthAccount extends Component {
  constructor() {
      super();
      this.state = {};
      this.loadData = this.loadData.bind(this);
    }
    
    componentDidMount() {
      this.loadData();
    }
    
    loadData() {
      MoneyStore.getActiveMoneyAccount(coinType, this.props.match.params.organizationID, resp => {
        this.setState({
          activeAccount: resp.activeAccount,
        });
      }); 
    }

    onSubmit = (resp) => {
      resp.orgId = this.props.match.params.organizationID;
      resp.money_abbr = coinType;
      
      const login = {};
      login.username = resp.username;
      login.password = resp.password;

      SessionStore.login(login, (response) => {
        if(response === "ok"){
          MoneyStore.modifyMoneyAccount(resp, resp => {
            
          })
        }else{
          alert("inccorect username or password.");
        }
      })
    } 

  render() {
    return(
      <Grid container spacing={24}>
        <Grid item xs={12} className={this.props.classes.divider}>
          <div className={this.props.classes.TitleBar}>
                <TitleBar className={this.props.classes.padding}>
                  <TitleBarTitle title="Eth Account" />
                </TitleBar>
                <Divider light={true}/>
                <div className={this.props.classes.breadcrumb}>
                <TitleBar>
                  <TitleBarTitle component={Link} to="#" title="M2M Wallet" className={this.props.classes.link}/> 
                  <TitleBarTitle title="/" className={this.props.classes.navText}/>
                  <TitleBarTitle component={Link} to="#" title="Eth Account" className={this.props.classes.link}/>
                </TitleBar>
                </div>
            </div>
        </Grid>
        <Grid item xs={6} className={this.props.classes.column}>
          <ModifyEthAccountForm
            submitLabel="Confirm"
            onSubmit={this.onSubmit}
            activeAccount={this.state.activeAccount}
          />
        </Grid>
        <Grid item xs={6}>
        </Grid>
      </Grid>
    );
  }
}

export default withStyles(styles)(withRouter(ModifyEthAccount));