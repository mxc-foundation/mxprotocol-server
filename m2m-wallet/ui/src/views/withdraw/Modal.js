import React from 'react';
import Button from '@material-ui/core/Button';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';
import WithdrawStore from "../../stores/WithdrawStore";

export default function AlertDialog(props) {
  const { open, onClose } = props

  const agree = () => {
    console.log('props', props);
    const data = props;

    /* WithdrawStore.update(data, resp => {
        props.history.push(`/withdraw/${props.match.params.organizationID}`);
    }); */

    if (onClose) onClose();
  }

  return (
      <Dialog
        open={open}
        onClose={onClose}
        aria-labelledby="alert-dialog-title"
        aria-describedby="alert-dialog-description"
      >
        <DialogTitle id="alert-dialog-title">{"Use Google's location service?"}</DialogTitle>
        <DialogContent>
          <DialogContentText id="alert-dialog-description">
            Let Google help apps determine location. This means sending anonymous location data to
            Google, even when no apps are running.
          </DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button onClick={onClose} color="primary" autoFocus>
            Cancel
          </Button>
          <Button onClick={agree} color="primary" autoFocus>
            Agree
          </Button>
        </DialogActions>
      </Dialog>
  );
}
