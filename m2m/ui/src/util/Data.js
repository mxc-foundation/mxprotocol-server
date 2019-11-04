export const INACTIVE = 'Inactive';
export const PRIVATE = 'Private';
export const NETWORK = 'Network';
export const DV_INACTIVE = 'DV_INACTIVE';
export const DV_FREE_GATEWAYS_LIMITED = 'DV_FREE_GATEWAYS_LIMITED';
export const DV_WHOLE_NETWORK = 'DV_WHOLE_NETWORK';
export const GW_INACTIVE = 'GW_INACTIVE';
export const GW_FREE_GATEWAYS_LIMITED = 'GW_FREE_GATEWAYS_LIMITED';
export const GW_WHOLE_NETWORK = 'GW_WHOLE_NETWORK';

export const DV_MODE_OPTION = [
    //{ value: DV_INACTIVE, label: INACTIVE },
    { value: DV_FREE_GATEWAYS_LIMITED, label: PRIVATE },
    { value: DV_WHOLE_NETWORK, label: NETWORK }];

export const GW_MODE_OPTION = [
    { value: GW_INACTIVE, label: INACTIVE },
    { value: GW_FREE_GATEWAYS_LIMITED, label: PRIVATE },
    { value: GW_WHOLE_NETWORK, label: NETWORK }];   
