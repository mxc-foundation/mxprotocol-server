import React, { Component } from "react";
import { withRouter } from 'react-router-dom';

import Grid from '@material-ui/core/Grid';
import Card from '@material-ui/core/Card';

import { CardContent } from "@material-ui/core";

import TitleBar from "../../components/TitleBar";
import TitleBarTitle from "../../components/TitleBarTitle";
import UserForm from "./UserForm";
import UserStore from "../../stores/UserStore";
import i18n, { packageNS } from '../../i18n';


class CreateUser extends Component {
  constructor() {
    super();
    this.onSubmit = this.onSubmit.bind(this);
  }

  onSubmit(user) {
    UserStore.create(user, user.password, [], resp => {
      this.props.history.push("/users");
    });
  }

  render() {
    return(
      <Grid container spacing={24}>
        <TitleBar>
          <TitleBarTitle title={i18n.t(`${packageNS}:menu.login.users`)} to="/Users" />
          <TitleBarTitle title="/" />
          <TitleBarTitle title={i18n.t(`${packageNS}:menu.login.create`)} />
        </TitleBar>
        <Grid item xs={12}>
          <Card>
            <CardContent>
              <UserForm
                submitLabel={i18n.t(`${packageNS}:menu.login.create_user`)}
                onSubmit={this.onSubmit}
              />
            </CardContent>
          </Card>
        </Grid>
      </Grid>
    );
  }
}

export default withRouter(CreateUser);
