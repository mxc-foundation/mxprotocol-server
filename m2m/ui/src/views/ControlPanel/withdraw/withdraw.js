import React, { Component } from "react";
import { Link, withRouter } from "react-router-dom";

import Grid from "@material-ui/core/Grid";
import TitleBarTitle from "../../../components/TitleBarTitle";
import { withStyles } from "@material-ui/core/styles";
import styles from "../../withdraw/WithdrawStyle"
import Button from "@material-ui/core/Button";

class SuperAdminWithdraw extends Component {
    constructor(props) {
        super(props);
        this.state = {};
    }

    loadData = async () => {

    };

    componentDidMount() {
        this.loadData();
    }

    componentDidUpdate(oldProps) {
        if (this.props === oldProps) {
            return;
        }
        this.loadData();
    }

    render() {
        return (
            <Grid container spacing={24} className={this.props.classes.backgroundColor}>
            <Grid item xs={12} className={this.props.classes.divider}>
                <div className={this.props.classes.TitleBar}>
                    <TitleBarTitle title="Withdraw" />
                </div>
            </Grid>

            <Grid item xs={6}>
                <a href={"https://www.google.de"} target="_blank">
                    Coming soon, click here for more information.
                </a>
            </Grid>

            </Grid>
        );
    }
}

export default withStyles(styles)(withRouter(SuperAdminWithdraw));