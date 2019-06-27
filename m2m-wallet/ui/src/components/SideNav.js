import React, { Component } from "react";
import { withRouter } from "react-router-dom";
import { withStyles } from "@material-ui/core/styles";

import theme from "../theme";


const styles = {
  drawerPaper: {
    position: "fixed",
    width: 270,
    paddingTop: theme.spacing.unit * 9,
  },
  select: {
    paddingTop: theme.spacing.unit,
    paddingLeft: theme.spacing.unit * 3,
    paddingRight: theme.spacing.unit * 3,
    paddingBottom: theme.spacing.unit * 1,
  },
};

class SideNav extends Component {
  constructor() {
    super();

    this.state = {
      open: true,
      organization: null,
      cacheCounter: 0,
    };
  }

  componentDidMount() {
    
  }

  render() {

    return(
      <></>
    );
  }
}

export default withRouter(withStyles(styles)(SideNav));
