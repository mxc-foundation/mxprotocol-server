import React, { Component } from "react";

import Grid from "@material-ui/core/Grid";
import TableCell from "@material-ui/core/TableCell";
import TableRow from "@material-ui/core/TableRow";

import TopupStore from "../../stores/TopupStore";
import TitleBar from "../../components/TitleBar";
import TitleBarTitle from "../../components/TitleBarTitle";
import TitleBarButton from "../../components/TitleBarButton";
import DataTable from "../../components/DataTable";
import Admin from "../../components/Admin";
import { withRouter } from "react-router-dom";
import { withStyles } from "@material-ui/core/styles";

const styles = {
  maxW140: {
    maxWidth: 140,
    //backgroundColor: "#0C0270",
    whiteSpace: 'nowrap', 
    overflow: 'hidden',
    textOverflow: 'ellipsis'
  },
};

class TopupHistory extends Component {
  constructor(props) {
    super(props);
    this.getPage = this.getPage.bind(this);
    this.getRow = this.getRow.bind(this);
  }

  getPage(limit, offset, callbackFunc) {
    TopupStore.getTopUpHistory(this.props.organizationID, offset, limit, data => {
        callbackFunc({
            totalCount: parseInt(data.count),
            result: data.topupHistory
          });
      }); 
  }
  
  getRow(obj, index) {
    return(
      <TableRow key={index}>
        <TableCell className={this.props.classes.maxW140} >{obj.from}</TableCell>
        <TableCell className={this.props.classes.maxW140}>{obj.to}</TableCell>
        <TableCell>{obj.moneyType}</TableCell>
        <TableCell>{obj.amount}</TableCell>
        <TableCell className={this.props.classes.maxW140}>{obj.txHash}</TableCell>
        <TableCell>{obj.createdAt.substring(0, 19)}</TableCell>
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
        <TitleBarTitle title="Topup History" />
        </TitleBar>
        <Grid item xs={12}>
          <DataTable
            header={
              <TableRow>
                <TableCell>From</TableCell>
                <TableCell>To</TableCell>
                <TableCell>Type</TableCell>
                <TableCell>VMXC Amount</TableCell>
                <TableCell>TxHash</TableCell>
                <TableCell>Date</TableCell>
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

export default withStyles(styles)(withRouter(TopupHistory));