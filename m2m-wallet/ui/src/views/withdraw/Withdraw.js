import React, { Component } from "react";
import { Link } from "react-router-dom";

import Grid from "@material-ui/core/Grid";
import TitleBar from "../../components/TitleBar";
import TitleBarTitle from "../../components/TitleBarTitle";
import MoneyStore from "../../stores/MoneyStore";
import WithdrawStore from "../../stores/WithdrawStore";
import WalletStore from "../../stores/WalletStore";
import WithdrawForm from "./WithdrawForm";
import Modal from "./Modal";
import WithdrawBalanceInfo from "./WithdrawBalanceInfo";
import { withRouter } from "react-router-dom";
import { withStyles } from "@material-ui/core/styles";
import Divider from '@material-ui/core/Divider';
import theme from "../../theme";
//import { promises } from "fs";

const styles = {
/*   card: {
    minWidth: 180,
    width: 220,
    backgroundColor: "#0C0270",
  }, */
  flex: {
    display: 'flex',
    flexDirection: 'column'
  },
  title: {
    color: '#FFFFFF',
    fontSize: 14,
    padding: 6,
  },
  balance: {
    fontSize: 24,
    color: '#FFFFFF',
    textAlign: 'center',
  },
  newBalance: {
    fontSize: 24,
    textAlign: 'center',
    color: theme.palette.primary.main,
  },
  navText: {
    fontSize: 14,
  },
  pos: {
    marginBottom: 12,
    color: '#FFFFFF',
    textAlign: 'right',
  },
  TitleBar: {
    height: 115,
    width: '50%',
    light: true,
    display: 'flex',
    flexDirection: 'column'
  },
  divider: {
    padding: 0,
    color: '#FFFFFF',
    width: '100%',
  },
  padding: {
    padding: 0,
  },
  between: {
    display: 'flex',
    justifyContent:'spaceBetween'
  },
  link: {
    textDecoration: "none",
    fontWeight: "bold",
    fontSize: 12,
    color: theme.palette.textSecondary.main,
    opacity: 0.7,
      "&:hover": {
        opacity: 1,
      }
  },
};

const coinType = "Ether";

function formatNumber(number) {
  return number.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",");
}

function loadWithdrawData(coinType) {
  return new Promise((resolve, reject) => {
    WithdrawStore.getWithdrawFee(coinType,
      resp => {
        Object.keys(resp).forEach(attr => {
          const value = resp[attr];

          if (typeof value === 'number') {
            resp[attr] = formatNumber(value);
          }
        });
        resp.moneyAbbr = coinType;
        resolve(resp);
      })
  });
}

function loadCurrentAccount(coinType, organizationID) {
  return new Promise((resolve, reject) => {
    MoneyStore.getActiveMoneyAccount(coinType, organizationID, 
      resp => {
        resolve(resp);
      })
  });
}

function loadWalletBalance(organizationID) {
  return new Promise((resolve, reject) => {
    WalletStore.getWalletBalance(organizationID,
      resp => {
        Object.keys(resp).forEach(attr => {
          const value = resp[attr];
  
          if (typeof value === 'number') {
            resp[attr] = formatNumber(value);
          }
        });
        resolve(resp);
      });
  });
}

class Withdraw extends Component {
  constructor() {
    super();
    this.state = {
      modal: null
    };
  }

  loadData = async () => {
    try {
      var result = await loadWithdrawData(coinType);
      var wallet = await loadWalletBalance(this.props.match.params.organizationID);
      var account = await loadCurrentAccount(coinType, this.props.match.params.organizationID);
      
      const txinfo = {};
      txinfo.withdrawFee = result.withdrawFee;
      txinfo.balance = wallet.balance;
      txinfo.account = account.activeAccount;

      this.setState({
        txinfo
      });
    } catch (error) {
      console.error(error);
      this.setState({ error });
    }
  }

  componentDidMount() {
    this.loadData();
  }

  componentDidUpdate(prevProps) {
    if (prevProps === this.props) {
      return;
    }
  }

  showModal(modal) {
    this.setState({ modal });
  }

  onSubmit = (e, apiWithdrawReqRequest) => {
    e.preventDefault();
    this.showModal(apiWithdrawReqRequest);
  }

  handleCloseModal = () => {
    this.setState({
      modal: null
    })
  }

  onConfirm = (data) => {
    WithdrawStore.WithdrawReq(data, resp => {
      console.log('WithdrawReq',resp)
    });
  }

  render() {
    return (
      <Grid container spacing={24} className={this.props.classes.backgroundColor}>
        {this.state.modal && 
          <Modal title="title" description="description" onClose={this.handleCloseModal} open={!!this.state.modal} data={this.state.modal} onConfirm={this.onConfirm} />}
        <Grid item xs={12} className={this.props.classes.divider}>
          <div className={this.props.classes.TitleBar}>
              <TitleBar className={this.props.classes.padding}>
                <TitleBarTitle title="Withdraw" />
              </TitleBar>
              <Divider light={true}/>
              <div className={this.props.classes.breadcrumb}>
              <TitleBar>
                <TitleBarTitle component={Link} to="#" title="M2M Wallet" className={this.props.classes.link}/> 
                <TitleBarTitle title="/" className={this.props.classes.navText}/>
                <TitleBarTitle component={Link} to="#" title="Withdraw" className={this.props.classes.link}/>
              </TitleBar>
              </div>
          </div>

        </Grid>
        <Grid item xs={6} className={this.props.classes.divider}></Grid>
        <Grid item xs={12} className={this.props.classes.divider}>

        </Grid>
        <Grid item xs={6}>
          <WithdrawForm
            submitLabel="Withdraw"
            txinfo={this.state.txinfo} {...this.props}
            onSubmit={this.onSubmit}
          />
        </Grid>
        <Grid item xs={2}>

        </Grid>
{/*         <Grid item xs={3}>
          <WithdrawBalanceInfo
            txinfo={this.state.txinfo} {...this.props}
          />
          
        </Grid> */}
      </Grid>
    );
  }
}

export default withStyles(styles)(withRouter(Withdraw));