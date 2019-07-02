import React, { Component } from "react";

import Grid from "@material-ui/core/Grid";
import TableCell from "@material-ui/core/TableCell";
import TableRow from "@material-ui/core/TableRow";

//import HistoryStore from "../../stores/HistoryStore";
import HistoryStore from "../../stores/HistoryStore";
import TitleBar from "../../components/TitleBar";
import TitleBarTitle from "../../components/TitleBarTitle";
import TableCellLink from "../../components/TableCellLink";
import TitleBarButton from "../../components/TitleBarButton";
import DataTable from "../../components/DataTable";
import Admin from "../../components/Admin";

class Transactions extends Component {
  constructor() {
    super();
    this.getPage = this.getPage.bind(this);
    this.getRow = this.getRow.bind(this);
  }

  getPage(limit, offset, callbackFunc) {
    //HistoryStore.list("", this.props.match.params.organizationID, limit, offset, callbackFunc);
    HistoryStore.GetTopUpHistory(this.props.match.params.organizationID, limit, offset, (data) => {
      callbackFunc({
        totalCount: offset + 2 * limit,
        result: data
      });
    });
  }

  getRow(obj) {
    return(
      <TableRow key={obj.id}>
        <TableCell>{obj.date}</TableCell>
        <TableCell>{obj.from}</TableCell>
        <TableCell>{obj.to}</TableCell>
        <TableCell>{obj.value}</TableCell>
        <TableCell>{obj.balance}</TableCell>
        <TableCell>{obj.txHash}</TableCell>
        <TableCell>{obj.status}</TableCell>
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
          <TitleBarTitle title="Transactions" />
        </TitleBar>
        <Grid item xs={12}>
          <DataTable
            header={
              <TableRow>
                <TableCell>From</TableCell>
                <TableCell>To</TableCell>
                <TableCell>Value</TableCell>
                <TableCell>Date</TableCell>
                <TableCell>Balance</TableCell>
                <TableCell>Tx hash</TableCell>
                <TableCell>Status</TableCell>
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

export default Transactions;
