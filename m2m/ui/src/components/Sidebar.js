import React, { Component } from 'react';
import { Link, withRouter } from 'react-router-dom';
import { UncontrolledDropdown, DropdownMenu, DropdownItem, DropdownToggle } from 'reactstrap';
import { getLoraHost } from "../util/M2mUtil";
import DropdownMenu2 from "./DropdownMenu";
import ProfileDropdown from './ProfileDropdown';
import PerfectScrollbar from 'react-perfect-scrollbar';
import 'react-perfect-scrollbar/dist/css/styles.css';
import MetisMenu from 'metismenujs/dist/metismenujs';
import profilePic from '../assets/images/users/user-1.jpg';
import Divider from '@material-ui/core/Divider';
import SessionStore from '../stores/SessionStore';
import i18n, { packageNS } from '../i18n';

const ProfileMenus = [{
    label: 'My Account',
    icon: 'fe-user',
    redirectTo: "/",
  },
  {
    label: 'Settings',
    icon: 'fe-settings',
    redirectTo: "/"
  },
  {
    label: 'Lock Screen',
    icon: 'fe-lock',
    redirectTo: "/"
  },
  {
    label: 'Logout',
    icon: 'fe-log-out',
    redirectTo: "/logout",
    hasDivider: true
  }]

const LinkToLPWAN = ({children, ...otherProps}) => 
  <a href={getLoraHost()} {...otherProps}>{children}</a>;

const SideNavContent = (props) => {
    
    return <React.Fragment>

        <div id="sidebar-menu">
            <ul className="metismenu" id="side-menu">
                <li className="menu-title">Navigation</li>

                <li>
                    {/* <DropdownMenu default={ this.state.default } onChange={this.onChange} /> [edit] */}
                    <DropdownMenu2 />
                </li> 

                {/* <li>
                    <Link to="/dashboard" className="waves-effect side-nav-link-ref">
                        <i className="mdi mdi-view-dashboard"></i>
                        <span> Dashboard </span>
                    </Link>
                </li> */}
               
                <li>
                    <Link to={`/withdraw/${props.orgId}`} className="waves-effect side-nav-link-ref">
                        {/* <i className="mdi mdi-cloud-print-outline"></i> */}
                        <i className="ti-cloud-down"></i>
                        <span> {i18n.t(`${packageNS}:menu.withdraw.withdraw`)} </span>
                    </Link>
                </li>

                <li>
                    <Link to={`/topup/${props.orgId}`} className="waves-effect side-nav-link-ref">
                        {/* <i className="mdi mdi-inbox-arrow-down"></i>
                        <i className="mdi mdi-basket-fill"></i> */}
                        <i className="ti-cloud-up"></i>
                        <span> {i18n.t(`${packageNS}:menu.topup.topup`)} </span>
                    </Link>
                </li>

                <li>
                    <Link to={`/history/${props.orgId}`} className="waves-effect side-nav-link-ref">
                        <i className="mdi mdi-history"></i>
                        <span> {i18n.t(`${packageNS}:menu.history.history`)} </span>
                    </Link>
                </li>

                <li>
                    <Link to={`/modify-account/${props.orgId}`} className="waves-effect side-nav-link-ref">
                        <i className="mdi mdi-ethereum"></i>
                        <span> {i18n.t(`${packageNS}:menu.eth_account.eth_account`)} </span>
                    </Link>
                </li>

                <li>
                    <Link to={`/device/${props.orgId}`} className="waves-effect side-nav-link-ref">
                        <i className="mdi mdi-network-outline"></i>
                        <span> {i18n.t(`${packageNS}:menu.devices.devices`)} </span>
                    </Link>
                </li>

                <li>
                    <Link to={`/gateway/${props.orgId}`} className="waves-effect side-nav-link-ref">
                        <i className="mdi mdi-remote"></i>
                        <span> {i18n.t(`${packageNS}:menu.gateways.gateways`)} </span>
                    </Link>
                </li>

                <li>
                    <Link to={`/stake/${props.orgId}`} className="waves-effect side-nav-link-ref">
                        <i className="mdi mdi-vote"></i>
                        <span> {i18n.t(`${packageNS}:menu.staking.staking`)} </span>
                    </Link>
                </li>

                <li>
                    <Divider />
                </li>

                <li>
                    <Link to={`/stake/${props.orgId}`} className="waves-effect side-nav-link-ref">
                        <i className="mdi mdi-view-dashboard"></i>
                        <span> {i18n.t(`${packageNS}:menu.nb_iot_server`)} </span>
                    </Link>
                </li>

                <li>
                    <Link to={'/ext'} className="waves-effect side-nav-link-ref">
                        <i className="mdi mdi-view-dashboard"></i>
                        <span> {i18n.t(`${packageNS}:menu.lpwan_server`)} </span>
                    </Link>
                </li>

                <li>
                    <Link to={`/stake/${props.orgId}`} className="waves-effect side-nav-link-ref">
                        <i className="mdi mdi-view-dashboard"></i>
                        <span> {i18n.t(`${packageNS}:menu.powered_by`)} </span>
                    </Link>
                </li>
                
            </ul>
        </div>
        <div className="clearfix"></div>
    </React.Fragment>
}

const UserProfile = () => {
    return <React.Fragment>
        <div className="user-box text-center">
            <img src={profilePic} alt="user-img" title="Nik Patel" className="rounded-circle img-thumbnail avatar-lg" />
            <UncontrolledDropdown>
                <DropdownToggle caret tag="a" className="text-dark dropdown-toggle h5 mt-2 mb-1 d-block">
                    Nik Patel
                </DropdownToggle>
                <DropdownMenu className="user-pro-dropdown">
                    <DropdownItem>
                        <i className="fe-user mr-1"></i>
                        <span>My Account</span>
                    </DropdownItem>
                    <DropdownItem>
                        <i className="fe-settings mr-1"></i>
                        <span>Settings</span>
                    </DropdownItem>
                    <DropdownItem>
                        <i className="fe-lock mr-1"></i>
                        <span>Lock Screen</span>
                    </DropdownItem>
                    <DropdownItem>
                        <i className="fe-log-out mr-1"></i>
                        <span>Logout</span>
                    </DropdownItem>
                </DropdownMenu>
            </UncontrolledDropdown>

            <p className="text-muted">Admin Head</p>
            <ul className="list-inline">
                {/* <li className="list-inline-item">
                    <Link to="/" className="text-muted">
                        <i className="mdi mdi-settings"></i>
                    </Link>
                </li> [edit]*/}

                <li className="list-inline-item">
                    <Link to="/logout" className="text-custom">
                        <i className="mdi mdi-power"></i>
                    </Link>
                </li>
            </ul>
        </div>
    </React.Fragment>
}


class Sidebar extends Component {
    constructor(props) {
        super(props);
        this.handleOtherClick = this.handleOtherClick.bind(this);
        this.initMenu = this.initMenu.bind(this);
    }

    /**
     * Bind event
     */
    componentWillMount = () => {
        document.addEventListener('mousedown', this.handleOtherClick, false);
    }


    /**
     * 
     */
    componentDidMount = () => {
        this.initMenu();
    }

    /**
     * Component did update
     */
    componentDidUpdate = (prevProps) => {
        if (this.props.isCondensed !== prevProps.isCondensed) {
            if (prevProps.isCondensed) {
                document.body.classList.remove("sidebar-enable");
                document.body.classList.remove("enlarged");
            } else {
                document.body.classList.add("sidebar-enable");
                const isSmallScreen = window.innerWidth < 768;
                if (!isSmallScreen) {
                    document.body.classList.add("enlarged");
                }
            }

            this.initMenu();
        }
    }

    /**
     * Bind event
     */
    componentWillUnmount = () => {
        document.removeEventListener('mousedown', this.handleOtherClick, false);
    }

    /**
     * Handle the click anywhere in doc
     */
    handleOtherClick = (e) => {
        if (this.menuNodeRef.contains(e.target))
            return;
        // else hide the menubar
        document.body.classList.remove('sidebar-enable');
    }

    /**
     * Init the menu
     */
    initMenu = () => {
        // render menu
        new MetisMenu("#side-menu");
        var links = document.getElementsByClassName('side-nav-link-ref');
        var matchingMenuItem = null;
        console.log('links: ', links);
        console.log('pathname: ', this.props);
        
        for (var i = 0; i < links.length; i++) {
            if (this.props.location.pathname === links[i].pathname) {
                matchingMenuItem = links[i];
                break;
            }
        }

        if (matchingMenuItem) {
            matchingMenuItem.classList.add('active');
            var parent = matchingMenuItem.parentElement;

            /**
             * TODO: This is hard coded way of expading/activating parent menu dropdown and working till level 3. 
             * We should come up with non hard coded approach
             */
            if (parent) {
                parent.classList.add('active');
                const parent2 = parent.parentElement;
                if (parent2) {
                    parent2.classList.add('in');
                }
                const parent3 = parent2.parentElement;
                if (parent3) {
                    parent3.classList.add('active');
                    var childAnchor = parent3.querySelector('.has-dropdown');
                    if (childAnchor) childAnchor.classList.add('active');
                }

                const parent4 = parent3.parentElement;
                if (parent4)
                    parent4.classList.add('in');
                const parent5 = parent4.parentElement;
                if (parent5)
                    parent5.classList.add('active');
            }
        }
    }

    render() {
        const isCondensed = this.props.isCondensed || false;
        const orgId = SessionStore.getOrganizationID();
        return (
            <React.Fragment>
                <div className='left-side-menu' ref={node => this.menuNodeRef = node}>
                    {!isCondensed && <PerfectScrollbar><UserProfile /><SideNavContent orgId={orgId} /></PerfectScrollbar>}
                    {isCondensed && <UserProfile /> && <SideNavContent />}
                </div>
            </React.Fragment>
        );
    }
}

export default withRouter(Sidebar);
//export default connect()(Sidebar); [edit]