import React from "react";

import TextField from '@material-ui/core/TextField';
import i18n, { packageNS } from '../../../i18n';
import Button from "@material-ui/core/Button";
import FormComponent from "../../../classes/FormComponent";
import Form from "../../../components/Form";
import NumberFormat from "react-number-format";

const NumberFormatMXC=(props)=> {
    const { inputRef, onChange, ...other } = props;

    return (
        <NumberFormat
            {...other}
            getInputRef={inputRef}
            onValueChange={(values) => {
                onChange({
                    target: {
                        value: values.value
                    }
                });
            }}
            suffix=" MXC"
        />
    );
}

class SettingsForm extends FormComponent {

    state = {
        downlinkPrice: '',
        percentageShare: '',
        lbWarning: '',
        withdrawFee: ''
    }

    onChange = (event) => {
        const { id, value } = event.target;

        this.setState({
            [id]: value
        });
    }

    clear = () => {
        this.setState({
            downlinkPrice: '',
            percentageShare: '',
            lbWarning: '',
            withdrawFee: ''
        })
    }

    submit = () => {
        this.props.onSubmit({
            action: 'modifySystemSettings',
            downlinkPrice: this.props.downlinkPrice,
            percentageShare: this.props.percentageShare,
            lbWarning: this.props.lbWarning,
            withdrawFee: this.props.withdrawFee
        })
    }

    render() {
        const extraButtons = <>
            <Button  variant="outlined" color="inherit" onClick={this.clear} type="button" disabled={false}>{i18n.t(`${packageNS}:menu.staking.reset`)}</Button>
        </>;

        return(
            <Form
                submitLabel={this.props.submitLabel}
                extraButtons={extraButtons}
                onSubmit={this.onSubmit}
            >
                <TextField
                    id="withdrawFee"
                    label={i18n.t(`${packageNS}:menu.settings.withdraw_fee`)}
                    variant="filled"
                    InputLabelProps={{
                        shrink: true
                    }}
                    InputProps={{
                        inputComponent: NumberFormatMXC
                    }}
                    value={this.state.withdrawFee}
                    /*placeholder={i18n.t(`${packageNS}:menu.settings.type_here`)}*/
                    fullWidth
                    required
                    margin="normal"
                    onChange={(e) => this.handleChange('withdrawFee', e)}
                />

                <TextField
                    id="downlinkPrice"
                    label={i18n.t(`${packageNS}:menu.settings.downlink_price`)}
                    variant="filled"
                    InputLabelProps={{
                        shrink: true
                    }}
                    InputProps={{
                        inputComponent: NumberFormatMXC
                    }}
                    placeholder={i18n.t(`${packageNS}:menu.settings.type_here`)}
                    fullWidth
                    required
                    margin="normal"
                    value={this.state.downlinkPrice}
                    onChange={(e) => this.handleChange('downlinkPrice', e)}
                />

                <TextField
                    id="percentageShare"
                    label={i18n.t(`${packageNS}:menu.settings.percentage_share`)}
                    variant="filled"
                    InputLabelProps={{
                        shrink: true
                    }}
                    InputProps={{
                        inputComponent: NumberFormatMXC
                    }}
                    placeholder={i18n.t(`${packageNS}:menu.settings.type_here`)}
                    fullWidth
                    required
                    margin="normal"
                    value={this.state.percentageShare}
                    onChange={(e) => this.handleChange('percentageShare', e)}
                />

                <TextField
                    id="lbWarning"
                    label={i18n.t(`${packageNS}:menu.settings.low_balance`)}
                    variant="filled"
                    InputLabelProps={{
                        shrink: true
                    }}
                    InputProps={{
                        inputComponent: NumberFormatMXC
                    }}
                    placeholder={i18n.t(`${packageNS}:menu.settings.type_here`)}
                    fullWidth
                    required
                    margin="normal"
                    value={this.state.lbWarning}
                    onChange={(e) => this.handleChange('lbWarning', e)}
                />

                {/* <TitleBarButton
            key={1}
            label="Go to Etherscan.io"
            icon={<LinkI />}
            color="secondary"
            onClick={this.deleteOrganization}
        /> */}

            </Form>
        );
    }
}

export default SettingsForm;