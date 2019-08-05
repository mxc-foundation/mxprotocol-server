import React, { Component } from "react";
import { withRouter, Link } from 'react-router-dom';
import { withStyles } from "@material-ui/core/styles";

import Grid from '@material-ui/core/Grid';
import TitleBar from "../../components/TitleBar";
import TitleBarTitle from "../../components/TitleBarTitle";
import Divider from '@material-ui/core/Divider';
import Spinner from "../../components/ScaleLoader";
import TopupStore from "../../stores/TopupStore";
import TopupForm from "./TopupForm";
import styles from "./TopupStyle"

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

  loadData() {
    this.setState({loading:true});
    TopupStore.getTopUpHistory(this.props.match.params.organizationID, 0, 1, resp => {
      if(resp.status){
        this.setState({
          topupHistory: resp.body.topupHistory[0],
        });
        this.setState({loading:false});
      }else{
        this.setState({loading:false});
      }
      
    }); 
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
            reps={this.state.topupHistory} {...this.props}
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