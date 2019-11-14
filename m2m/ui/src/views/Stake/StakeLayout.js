import React, { Component } from "react";
import { Link } from "react-router-dom";
import Divider from '@material-ui/core/Divider';
import Grid from "@material-ui/core/Grid";
import TitleBar from "../../components/TitleBar";
import TitleBarTitle from "../../components/TitleBarTitle";
import DeviceStore from "../../stores/DeviceStore.js";
import WalletStore from "../../stores/WalletStore.js";
import GatewayStore from "../../stores/GatewayStore.js";
import TitleBarButton from "../../components/TitleBarButton";
import Button from "@material-ui/core/Button";
import StakeStore from "../../stores/StakeStore";
import Typography from '@material-ui/core/Typography';
//import WithdrawBalanceInfo from "./WithdrawBalanceInfo";
import { withRouter } from "react-router-dom";
import { withStyles } from "@material-ui/core/styles";
import styles from "./StakeStyle"
import { CONFIRMATION, CONFIRMATION_TEXT, INVALID_ACCOUNT, INVALID_AMOUNT } from "../../util/Messages"
import { EXT_URL_STAKE } from "../../util/Data"

function doIHaveGateway(orgId) {
  return new Promise((resolve, reject) => {
    GatewayStore.getGatewayList(orgId, 0, 1, data => {
      resolve(parseInt(data.count));
    });
  });
}  

function getDlPrice(orgId) {
  return new Promise((resolve, reject) => {
    WalletStore.getDlPrice(orgId, resp => {
      resolve(resp.downLinkPrice);
    });
  });
}

class StakeLayout extends Component {
  constructor(props) {
    super(props);
    this.state = {
      loading: false,
      isFirst: true
    };
  }

  loadData = async () => {
    const resp = StakeStore.getStakingHistory(this.props.match.params.organizationID, 0, 1);
    resp.then((res) => {
      let amount = 0;
      let isFirst = true;
      if( res.stakingHist.length > 0){
        this.props.history.push(`/stake/${this.props.match.params.organizationID}/set-stake`);
      }
      this.setState({
        amount,
        isFirst
      })
    })
  }

  componentWillMount(){
    this.loadData();
  }

  componentDidMount() {
    //this.loadData();
  }

  componentDidUpdate(oldProps) {
    if (this.props === oldProps) {
      return;
    }
    this.loadData();
  }
  
  onSubmit = (e, apiWithdrawReqRequest) => {
    e.preventDefault();
  }

  render() {
    return (
      <Grid container spacing={24} className={this.props.classes.backgroundColor}>
        <Grid item xs={12} className={this.props.classes.divider}>
          <div className={this.props.classes.TitleBar}>
              {/* <TitleBar className={this.props.classes.padding}>
                <TitleBarTitle title="Stake" />
              </TitleBar> */}    
              {/* <Divider light={true}/> */}
              <div className={this.props.classes.between}>
              <TitleBar>
                <TitleBarTitle title="Stake" />
                {/* <TitleBarTitle component={Link} to="#" title="M2M Wallet" className={this.props.classes.link}/> 
                <TitleBarTitle component={Link} to="#" title="/" className={this.props.classes.link}/>
                <TitleBarTitle component={Link} to="#" title="Devices" className={this.props.classes.link}/> */}
              </TitleBar>
              <Button color="primary.main" component={Link} to={`/stake/${this.props.match.params.organizationID}/set-stake`} /* onClick={this.handleOpenAXS} */ type="button" disabled={false}>SET STAKE</Button>
              {/* <TitleBarButton
                label="SET STAKE"
                color="primary"
                to={`/stake/${this.props.match.params.organizationID}/set-stake`}
                classes={this.props.classes}
              /> */}
              </div>
          </div>
        </Grid>
        <Grid item xs={12} className={this.props.classes.divider}>
          <Grid item xs={6}>
                <div className={this.props.classes.infoBox}>
                  <p>Staking enhances data trade by giving all 
                  holders a fair way to take part in the network.</p>
                  <div className={this.props.classes.between}>
                    <Typography className={this.props.classes.title} gutterBottom>
                      DISMISS
                    </Typography>
                    <TitleBarTitle component={Link} to={EXT_URL_STAKE} title="LEARN MORE" />
                  </div>
                </div>
          </Grid>
        </Grid>
      </Grid>
    );
  }
}

export default withStyles(styles)(withRouter(StakeLayout));