import React, { Component } from "react";
import { withRouter, Link } from 'react-router-dom';
import { withStyles } from "@material-ui/core/styles";

import Grid from '@material-ui/core/Grid';
import TitleBar from "../../components/TitleBar";
import TitleBarTitle from "../../components/TitleBarTitle";
import Divider from '@material-ui/core/Divider';
import MoneyStore from "../../stores/MoneyStore";
import SessionStore from "../../stores/SessionStore";
import SupernodeStore from "../../stores/SupernodeStore";
import ModifyEthAccountForm from "./ModifyEthAccountForm";
import NewEthAccountForm from "./NewEthAccountForm";
import styles from "./EthAccountStyle";
import { ETHER } from "../../util/Coin-type";
import { SUPER_ADMIN } from "../../util/M2mUtil";

function verifyUser (resp) {
  const login = {};
  login.username = resp.username;
  login.password = resp.password;

  return new Promise((resolve, reject) => {
    SessionStore.login(login, (resp) => {
      if(resp){
        resolve(resp);
      } else {
        alert("inccorect username or password.");
        return false;
      }
    })
  });
}

function modifyAccount (req, orgId) {
  req.moneyAbbr = ETHER;
  req.orgId = orgId;
  return new Promise((resolve, reject) => {
    MoneyStore.modifyMoneyAccount(req, resp => {
      resolve(resp);
    })
  });
}

function createAccount (req) {
  req.moneyAbbr = ETHER;
  return new Promise((resolve, reject) => {
    SupernodeStore.addSuperNodeMoneyAccount(req, resp => {
      resolve(resp);
    })
  });
}

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
      const org_id = this.props.match.params.organizationID;

      if (org_id === SUPER_ADMIN) {
        SupernodeStore.getSuperNodeActiveMoneyAccount(ETHER, resp => {
          this.setState({
            activeAccount: resp.supernodeActiveAccount,
          });
        });
      }else{
        MoneyStore.getActiveMoneyAccount(ETHER, org_id, resp => {
          this.setState({
            activeAccount: resp.activeAccount,
          });
        });
      }
    }

    componentDidUpdate(oldProps) {
      if (this.props.match.url === oldProps.match.url) {
        return;
      }

      this.loadData();
    }

    onSubmit = async (resp) => {
      const org_id = this.props.match.params.organizationID;
      
      try {
        if(resp.username !== SessionStore.getUsername() ){
          alert('inccorect username or password.');
          return false;
        }
        const isOK = await verifyUser(resp);
        
        if(org_id == 0 && isOK) {
          const res = await createAccount(resp, org_id);
          if(res.status){
            window.location.reload();
          }
        }else{
          const res = await modifyAccount(resp, org_id);
          if(res.status){
            window.location.reload();
          }
        } 
      } catch (error) {
        console.error(error);
        this.setState({ error });
      }
    } 

  render() {
    return(
      <Grid container spacing={24}>
        <Grid item xs={12} className={this.props.classes.divider}>
          <div className={this.props.classes.TitleBar}>
                <TitleBar className={this.props.classes.padding}>
                  <TitleBarTitle title="ETH Account" />
                </TitleBar>
                <Divider light={true}/>
                <div className={this.props.classes.breadcrumb}>
                <TitleBar>
                  <TitleBarTitle component={Link} to="#" title="M2M Wallet" className={this.props.classes.link}/> 
                  <TitleBarTitle title="/" className={this.props.classes.navText}/>
                  <TitleBarTitle component={Link} to="#" title="ETH Account" className={this.props.classes.link}/>
                </TitleBar>
                </div>
            </div>
        </Grid>
        <Grid item xs={6} className={this.props.classes.column}>
          {this.state.activeAccount &&  
            <ModifyEthAccountForm
              submitLabel="Confirm"
              onSubmit={this.onSubmit}
              activeAccount={this.state.activeAccount}
            />
          }
          {!this.state.activeAccount &&  
          <NewEthAccountForm
            submitLabel="Confirm"
            onSubmit={this.onSubmit}
          />
          }
        </Grid>
        <Grid item xs={6}>
          
        </Grid>
      </Grid>
    );
  }
}

export default withStyles(styles)(withRouter(ModifyEthAccount));