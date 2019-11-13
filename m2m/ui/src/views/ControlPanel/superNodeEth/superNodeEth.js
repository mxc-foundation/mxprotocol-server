import React, { Component } from "react";
import { withRouter, Link } from 'react-router-dom';
import { withStyles } from "@material-ui/core/styles";
import Grid from '@material-ui/core/Grid';
import TitleBar from "../../../components/TitleBar";
import TitleBarTitle from "../../../components/TitleBarTitle";
import MoneyStore from "../../../stores/MoneyStore";
import SessionStore from "../../../stores/SessionStore";
import SupernodeStore from "../../../stores/SupernodeStore";
import ModifyEthAccountForm from "../../ethAccount/ModifyEthAccountForm";
import NewEthAccountForm from "../../ethAccount/NewEthAccountForm";
import styles from "./superNodeEthStyle"
import { ETHER } from "../../../util/Coin-type";
import { SUPER_ADMIN } from "../../../util/M2mUtil";


class SuperNodeEth extends Component {
    constructor(props) {
    super(props);
    this.state = {
      activeAccount:'0'
    };
    this.loadData = this.loadData.bind(this);

    }
  
    componentDidMount() {
    this.loadData();
    }


  
    loadData() {
      SupernodeStore.getSuperNodeActiveMoneyAccount(ETHER, SUPER_ADMIN, resp => {
        this.setState({
          activeAccount: resp.supernodeActiveAccount,
        });
      });

    }

    componentDidUpdate(oldProps) {
        if (this.props === oldProps) {
            return;
        }

        this.loadData();
    }


    verifyUser(resp) {
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

    modifyAccount(req, orgId) {
        req.moneyAbbr = ETHER;
        return new Promise((resolve, reject) => {
            SupernodeStore.addSuperNodeMoneyAccount(req, orgId, resp => {
                resolve(resp);
            })
        });
    }

    onSubmit = async (resp) => {
      
      try {
        if(resp.username !== SessionStore.getUsername() ){
          alert('inccorect username or password.');
          return false;
        }
        const isOK = await this.verifyUser(resp);
        
        if(isOK) {
          const res = await this.modifyAccount(resp, SUPER_ADMIN);
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
                      <TitleBarTitle title="Control Panel" />
                    </TitleBar>

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

export default withStyles(styles)(withRouter(SuperNodeEth));