import React, { Component } from "react";

import Grid from "@material-ui/core/Grid";
import TableCell from "@material-ui/core/TableCell";
import TableRow from "@material-ui/core/TableRow";

import HistoryStore from "../../stores/HistoryStore";
import TitleBar from "../../components/TitleBar";
import TitleBarTitle from "../../components/TitleBarTitle";
import TitleBarButton from "../../components/TitleBarButton";
import DataTable from "../../components/DataTable";
import Admin from "../../components/Admin";
import { ETHER } from "../../util/Coin-type";

class WithdrawHistory extends Component {
  constructor() {
    super();
    this.getPage = this.getPage.bind(this);
    this.getRow = this.getRow.bind(this);
  }

  getPage(limit, offset, callbackFunc) {
    HistoryStore.getWithdrawHistory(ETHER, this.props.match.params.organizationID, limit, offset, (data) => {
      callbackFunc({
        totalCount: parseInt(data.count),
        result: data.withdrawHistory
      });
    });
  }

  getRow(obj, index) {
    return(
      <TableRow key={index}>
        {/* <TableCell>{obj.from}</TableCell> */}
        <TableCell>{obj.to}</TableCell>
        <TableCell>{obj.moneyType}</TableCell>
        <TableCell>{obj.amount}</TableCell>
        <TableCell>{obj.withdrawFee}</TableCell>
        <TableCell>{obj.txSentTime}</TableCell>
        <TableCell>{obj.txApprovedTime}</TableCell>
        <TableCell>{obj.txStatus}</TableCell>
        <TableCell>{obj.txHash}</TableCell> 
        {/* <TableCell>{obj.createdAt}</TableCell> */}
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
        <TitleBarTitle title="WithDraw" />
        </TitleBar>
        <Grid item xs={12}>
          <DataTable
            header={
              <TableRow>
                {/* <TableCell>From</TableCell> */}
                <TableCell>To</TableCell>
                <TableCell>Type</TableCell>
                <TableCell>VMXC Amount</TableCell>
                <TableCell>Withdraw Fee</TableCell>
                <TableCell>TxSentTime</TableCell>
                <TableCell>TxApprovedTime</TableCell>
                <TableCell>TxStatus</TableCell>
                <TableCell>TxHash</TableCell>
                {/* <TableCell>Date</TableCell> */}
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

export default WithdrawHistory;
