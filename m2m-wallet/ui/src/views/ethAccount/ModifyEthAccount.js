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

const coinType = "Ether";

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

function modifyAccount (resp , organizationID, history) {
  resp.moneyAbbr = coinType;
  return new Promise((resolve, reject) => {
    MoneyStore.modifyMoneyAccount(resp, resp => {
      history.push(`/modify-account/${organizationID}`);
      resolve(resp);
    })
  });
}

function createAccount (req, organizationID, history) {
  //console.log('req.organizationID', req);
  req.moneyAbbr = coinType;
  return new Promise((resolve, reject) => {
    SupernodeStore.addSuperNodeMoneyAccount(req, resp => {
      history.push(`/modify-account/${organizationID}`);
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

      if (org_id == 0) {
        SupernodeStore.getSuperNodeActiveMoneyAccount(coinType, resp => {
          this.setState({
            activeAccount: resp.supernodeActiveAccount,
          });
        });
      }else{
        MoneyStore.getActiveMoneyAccount(coinType, org_id, resp => {
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
      try {
        const isOK = await verifyUser(resp);
        
        if(resp.action === 'modifyAccount' && isOK) {
          const result = await modifyAccount(resp, this.props.match.params.organizationID, this.props.history);
        } 
        if(resp.action === 'createAccount' && isOK) {
          const result = await createAccount(resp, this.props.match.params.organizationID, this.props.history);
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