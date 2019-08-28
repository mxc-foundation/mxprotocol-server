import React, { Component } from "react";
import { Link } from 'react-router-dom';

import TableCell from '@material-ui/core/TableCell';
import { withStyles } from '@material-ui/core/styles';

import theme from "../theme";


const styles = {
  link: {
    textDecoration: "none",
    color: theme.palette.primary.main,
    cursor: "pointer",
    textDecoration: "none",
    padding: 0,
    fontWeight: "bold",
    fontSize: 14,
    opacity: 0.7,
      "&:hover": {
        opacity: 1,
      }
  },
};


class TableCellLink extends Component {
  render() {
    return(
      <TableCell>
        {this.props.to && <Link className={this.props.classes.link} to={this.props.to}>{this.props.children}</Link>}
        {this.props.onClick && <span className={this.props.classes.link} onClick={this.props.onClick}>{this.props.children}</span>}
      </TableCell>
    );
  }
}

export default withStyles(styles)(TableCellLink);
