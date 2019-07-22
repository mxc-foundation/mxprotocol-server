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

class EthAccount extends Component {
  constructor() {
    super();
    this.getPage = this.getPage.bind(this);
    this.getRow = this.getRow.bind(this);
  }

  getPage(limit, offset, callbackFunc) {
    HistoryStore.getChangeMoneyAccountHistory("Ether",this.props.match.params.organizationID, limit, offset, (data) => {
      callbackFunc({
        totalCount: offset + 2 * limit,
        result: data.changeHistory
      });
    });
  }

  getRow(obj, index) {
    return(
      <TableRow key={index}>
        <TableCell>{obj.from}</TableCell>
        <TableCell>{obj.to}</TableCell>
        <TableCell>{obj.createdAt}</TableCell>
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
          <TitleBarTitle title="EthAccount" />
        </TitleBar>
        <Grid item xs={12}>
          <DataTable
            header={
              <TableRow>
                <TableCell>Old Account</TableCell>
                <TableCell>New Account</TableCell>
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

export default EthAccount;
