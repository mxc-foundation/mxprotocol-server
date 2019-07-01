import React, { Component } from "react";
import { withRouter } from 'react-router-dom';
import { withStyles } from "@material-ui/core/styles";

import Grid from '@material-ui/core/Grid';
import Card from '@material-ui/core/Card';
import CardContent from "@material-ui/core/CardContent";
import TitleBar from "../../components/TitleBar";
import TitleBarTitle from "../../components/TitleBarTitle";
import Divider from '@material-ui/core/Divider';
//import TopupStore from "../../stores/TopupStore";
import TopupForm from "./TopupForm";
import theme from "../../theme";

const styles = {
  tabs: {
    borderBottom: "1px solid " + theme.palette.divider,
    height: "49px",
  },
  navText: {
    fontSize: 14,
  },
  TitleBar: {
    height: 115,
    width: '50%',
    light: true,
    display: 'flex',
    flexDirection: 'column'
  },
  card: {
    minWidth: 180,
    width: 220,
    backgroundColor: "#0C0270",
  },
  divider: {
    padding: 0,
    color: '#FFFFFF',
    width: '100%',
  },
  padding: {
    padding: 0,
  },
  column: {
    display: 'flex',
    flexDirection: 'column',
  },
};

class Topup extends Component {
    constructor() {
        super();
        this.state = {txinfo:{sendToken:'1232345.2134',from:'mxc',to:'chx'}};
        this.loadData = this.loadData.bind(this);
      }
    
    componentDidMount() {
      this.loadData();
    }
    
    loadData() {
      /* TopupStore.get(this.props.match.params.organizationID, resp => {
        console.log(resp)
        this.setState({
          organization: resp,
        });
      }); */
    }
    /* 
      componentDidUpdate(prevProps) {
        if (prevProps === this.props) {
          return;
        }
    
        this.loadData();
      }
    
    
      deleteOrganization() {
        if (window.confirm("Are you sure you want to delete this organization?")) {
          OrganizationStore.delete(this.props.match.params.organizationID, () => {
            this.props.history.push("/topup");
          });
        }
      }

    onSubmit(organization) {
        OrganzationStore.update(organization, resp => {
        this.props.history.push("/topup");
        });
    } */

  render() {
    console.log(this.state.txinfo);
    return(
      <Grid container spacing={24}>
        <Grid item xs={12} className={this.props.classes.divider}>
          <div className={this.props.classes.TitleBar}>
              <TitleBar className={this.props.classes.padding}>
                <TitleBarTitle title="Topup" />
              </TitleBar>
              <Divider light={true}/>
              <TitleBar>
                <TitleBarTitle title="M2M Wallet" className={this.props.classes.navText}/>
                <TitleBarTitle title="/" className={this.props.classes.navText}/>
                <TitleBarTitle title="Topup" className={this.props.classes.navText}/>
              </TitleBar>
          </div>
        </Grid>
        <Grid item xs={6} className={this.props.classes.column}>
          <TitleBarTitle title="Send Tokens" />
          <Divider light={true}/>
          <TopupForm
            txinfo={this.state.txinfo}
          />
            
        </Grid>
        <Grid item xs={6}>
        </Grid>
      </Grid>
    );
  }
}

export default withStyles(styles)(withRouter(Topup));