import React, { Component } from "react";

import Grid from "@material-ui/core/Grid";
import TableCell from "@material-ui/core/TableCell";
import TableRow from "@material-ui/core/TableRow";

import StakeStore from "../../stores/StakeStore";
import TitleBar from "../../components/TitleBar";

import TitleBarButton from "../../components/TitleBarButton";
import DataTable from "../../components/DataTable";
import Admin from "../../components/Admin";
import { withRouter } from "react-router-dom";
import { withStyles } from "@material-ui/core/styles";

const styles = {
  maxW140: {
    maxWidth: 140,
    whiteSpace: 'nowrap', 
    overflow: 'hidden',
    textOverflow: 'ellipsis'
  },
  flex:{
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center'

  }
};

class Stakes extends Component {
  constructor(props) {
    super(props);
    this.getPage = this.getPage.bind(this);
    this.getRow = this.getRow.bind(this);
  }

  getPage(limit, offset, callbackFunc) {
    StakeStore.getStakingHistory(this.props.organizationID, offset, limit, data => {
        callbackFunc({
            totalCount: parseInt(data.count),
            result: data.stakingHist
          });
      }); 
  }
  
  getRow(obj, index) {
    return(
      <TableRow key={index}>
        <TableCell align={'right'} className={this.props.classes.maxW140} >{obj.stakeAmount}</TableCell>
        <TableCell align={'center'} className={this.props.classes.maxW140}>{obj.start.substring(0,16)}</TableCell>
        <TableCell align={'center'} className={this.props.classes.maxW140}>{obj.end.substring(0,16)}</TableCell>
        <TableCell align={'center'}>{obj.revMonth}</TableCell>
        <TableCell align={'right'}>{obj.networkIncome}</TableCell>
        <TableCell align={'right'}>{obj.monthlyRate}</TableCell>
        <TableCell align={'right'}>{obj.revenue}</TableCell>
        <TableCell align={'right'}>{obj.balance}</TableCell>
      </TableRow>
    );
  }

  render() {
    return(
      <Grid container spacing={24}>
        <TitleBar
          buttons={
            <Admin organizationID={this.props.match.params.organizationID}>
              <TitleBarButton
                label="Filter"
                //icon={<Plus />}
              />
            </Admin>
          }
        >
        </TitleBar>
        <Grid item xs={12}>
          <DataTable
            header={
              <TableRow>
                <TableCell align={'right'}>Stake Amount</TableCell>
                <TableCell align={'center'}>Start</TableCell>
                <TableCell align={'center'}>End</TableCell>
                <TableCell align={'center'}>Revenue Month</TableCell>
                <TableCell align={'right'}>Network Income</TableCell>
                <TableCell align={'right'}>Monthly Rate</TableCell>
                <TableCell align={'right'}>Revenue</TableCell>
                <TableCell align={'right'}>Balance</TableCell>
              </TableRow>
            }
            getPage={this.getPage}
            getRow={this.getRow}
          />
        </Grid>
      </Grid>
    );
  }
}

export default withStyles(styles)(withRouter(Stakes));