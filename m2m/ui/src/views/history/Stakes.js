import React, { Component } from "react";

import Grid from "@material-ui/core/Grid";
import TableCell from "@material-ui/core/TableCell";
import TableRow from "@material-ui/core/TableRow";
import moment from 'moment';
import mtz from 'moment-timezone';

import StakeStore from "../../stores/StakeStore";
import TitleBar from "../../components/TitleBar";

import TableCellExtLink from '../../components/TableCellExtLink';
import TitleBarButton from "../../components/TitleBarButton";
import DataTable from "../../components/DataTable";
import LinkVariant from "mdi-material-ui/LinkVariant";
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
    console.log(123);
    StakeStore.getStakingHistory(this.props.organizationID, offset, limit, data => {
        console.log(data);
        callbackFunc({
            totalCount: parseInt(data.count),
            result: data.stakingHist
          });
      }); 
  }
  
  getRow(obj, index) {
    const url = process.env.REACT_APP_ETHERSCAN_ROPSTEN_HOST + `/tx/${obj.txHash}`;
    
    return(
      <TableRow key={index}>
        <TableCell align={'left'} className={this.props.classes.maxW140} >{obj.stakeAmount}</TableCell>
        <TableCell align={'left'} className={this.props.classes.maxW140}>{obj.start.substring(0,16)}</TableCell>
        <TableCell align={'left'} className={this.props.classes.maxW140}>{obj.end.substring(0,16)}</TableCell>
        <TableCell align={'left'}>{obj.revMonth}</TableCell>
        <TableCell align={'left'}>{obj.networkIncome}</TableCell>
        <TableCell align={'left'}>{obj.monthlyRate}</TableCell>
        <TableCell align={'left'}>{obj.revenue}</TableCell>
        <TableCell align={'left'}>{obj.balance}</TableCell>
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
                <TableCell align={'left'}>Stake Amount</TableCell>
                <TableCell align={'left'}>Start</TableCell>
                <TableCell align={'left'}>End</TableCell>
                <TableCell align={'left'}>Revenue Month</TableCell>
                <TableCell align={'left'}>Network Income</TableCell>
                <TableCell align={'left'}>Monthly Rate</TableCell>
                <TableCell align={'left'}>Revenue</TableCell>
                <TableCell align={'left'}>Balance</TableCell>
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