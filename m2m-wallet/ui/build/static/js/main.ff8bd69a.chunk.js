(window.webpackJsonp=window.webpackJsonp||[]).push([[0],{361:function(e,t,n){e.exports=n(739)},738:function(e,t,n){},739:function(e,t,n){"use strict";n.r(t);var a=n(0),i=n.n(a),o=n(27),r=n.n(o),s=(n(366),n(285)),c=n.n(s),l=n(17),u=n(18),h=n(20),p=n(19),g=n(21),m=n(10),b=n(743),d=n(742),f=n(741),O=n(6),j=n.n(O),v=n(129),y=n.n(v),E=n(16),w=n(41),k=n.n(w),S=n(133),C=Object(S.a)(),I=n(44),T=n.n(I),D=Object(E.createMuiTheme)({palette:{primary:T.a}}),N=n(740),A=n(744),x=n(118),z=n.n(x),L=n(119),M=n.n(L),B=n(202),R=n(87),_=n.n(R),F=n(124),U=n.n(F),P=n(51),H=n.n(P),q=n(121),G=n.n(q),J=n(123),W=n.n(J),K=n(122),Q=n.n(K),V=n(289),X=n.n(V),Y=n(290),Z=n.n(Y),$=n(293),ee=n.n($),te=n(291),ne=n.n(te),ae=n(292),ie=n.n(ae),oe=n(130),re=n(287),se=n.n(re),ce=new(n(288).Dispatcher);function le(e){if(e.status>=200&&e.status<300)return e;throw e.json()}function ue(e){void 0===e.response?ce.dispatch({type:"CREATE_NOTIFICATION",notification:{type:"error",message:e.message}}):16===e.response.obj.code?C.push("/login"):ce.dispatch({type:"CREATE_NOTIFICATION",notification:{type:"error",message:e.response.obj.error+" (code: "+e.response.obj.code+")"}})}function he(e){void 0===e.response?ce.dispatch({type:"CREATE_NOTIFICATION",notification:{type:"error",message:e.message}}):ce.dispatch({type:"CREATE_NOTIFICATION",notification:{type:"error",message:e.response.obj.error+" (code: "+e.response.obj.code+")"}})}var pe=new(function(e){function t(){var e;return Object(l.a)(this,t),(e=Object(h.a)(this,Object(p.a)(t).call(this))).client=null,e.user=null,e.organizations=[],e.settings={},e.branding={},e.swagger=se()("/swagger/internal.swagger.json",e.getClientOpts()),e.swagger.then(function(t){e.client=t,null!==e.getToken()&&e.fetchProfile(function(){})}),e}return Object(g.a)(t,e),Object(u.a)(t,[{key:"getClientOpts",value:function(){var e=this;return{requestInterceptor:function(t){return null!==e.getToken()&&(t.headers["Grpc-Metadata-Authorization"]="Bearer "+e.getToken()),t}}}},{key:"setToken",value:function(e){localStorage.setItem("jwt",e)}},{key:"getToken",value:function(){return localStorage.getItem("jwt")}},{key:"getOrganizationID",value:function(){var e=localStorage.getItem("organizationID");return""===e?null:e}},{key:"setOrganizationID",value:function(e){localStorage.setItem("organizationID",e),this.emit("organization.change")}},{key:"getUser",value:function(){return this.user}},{key:"getSettings",value:function(){return this.settings}},{key:"isAdmin",value:function(){return void 0!==this.user&&null!==this.user&&this.user.isAdmin}},{key:"isOrganizationAdmin",value:function(e){for(var t=0;t<this.organizations.length;t++)if(this.organizations[t].organizationID===e)return this.organizations[t].isAdmin}},{key:"login",value:function(e,t){var n=this;this.swagger.then(function(a){a.apis.InternalService.Login({body:e}).then(le).then(function(e){n.setToken(e.obj.jwt),n.fetchProfile(t)}).catch(he)})}},{key:"logout",value:function(e){localStorage.clear(),this.user=null,this.organizations=[],this.settings={},this.emit("change"),e()}},{key:"fetchProfile",value:function(e){var t=this;this.swagger.then(function(n){n.apis.InternalService.Profile({}).then(le).then(function(n){t.user=n.obj.user,void 0!==n.obj.organizations&&(t.organizations=n.obj.organizations),void 0!==n.obj.settings&&(t.settings=n.obj.settings),t.emit("change"),e()}).catch(ue)})}},{key:"globalSearch",value:function(e,t,n,a){this.swagger.then(function(i){i.apis.InternalService.GlobalSearch({search:e,limit:t,offset:n}).then(le).then(function(e){a(e.obj)}).catch(ue)})}},{key:"getBranding",value:function(e){this.swagger.then(function(t){t.apis.InternalService.Branding({}).then(le).then(function(t){e(t.obj)}).catch(ue)})}}]),t}(oe.EventEmitter)),ge={appBar:{zIndex:D.zIndex.drawer+1},menuButton:{marginLeft:-12,marginRight:10},hidden:{display:"none"},flex:{flex:1},logo:{height:32},search:{marginRight:3*D.spacing.unit,color:D.palette.common.white,background:T.a[400],width:450,padding:5,borderRadius:3},avatar:{background:T.a[600],color:D.palette.common.white},chip:{background:T.a[600],color:D.palette.common.white,marginRight:D.spacing.unit,"&:hover":{background:T.a[400]},"&:active":{background:T.a[400]}},iconButton:{color:D.palette.common.white,marginRight:D.spacing.unit}},me=function(e){function t(){var e;return Object(l.a)(this,t),(e=Object(h.a)(this,Object(p.a)(t).call(this))).state={menuAnchor:null,search:""},e.handleDrawerToggle=e.handleDrawerToggle.bind(Object(m.a)(Object(m.a)(e))),e.onMenuOpen=e.onMenuOpen.bind(Object(m.a)(Object(m.a)(e))),e.onMenuClose=e.onMenuClose.bind(Object(m.a)(Object(m.a)(e))),e.onLogout=e.onLogout.bind(Object(m.a)(Object(m.a)(e))),e.onSearchChange=e.onSearchChange.bind(Object(m.a)(Object(m.a)(e))),e.onSearchSubmit=e.onSearchSubmit.bind(Object(m.a)(Object(m.a)(e))),e}return Object(g.a)(t,e),Object(u.a)(t,[{key:"onMenuOpen",value:function(e){this.setState({menuAnchor:e.currentTarget})}},{key:"onMenuClose",value:function(){this.setState({menuAnchor:null})}},{key:"onLogout",value:function(){var e=this;pe.logout(function(){e.props.history.push("/login")})}},{key:"handleDrawerToggle",value:function(){this.props.setDrawerOpen(!this.props.drawerOpen)}},{key:"onSearchChange",value:function(e){this.setState({search:e.target.value})}},{key:"onSearchSubmit",value:function(e){e.preventDefault(),this.props.history.push("/search?search=".concat(encodeURIComponent(this.state.search)))}},{key:"render",value:function(){var e;e=this.props.drawerOpen?i.a.createElement(Z.a,null):i.a.createElement(X.a,null);var t=Boolean(this.state.menuAnchor);return i.a.createElement(z.a,{className:this.props.classes.appBar},i.a.createElement(M.a,null,i.a.createElement(B.a,{color:"inherit","aria-label":"toggle drawer",onClick:this.handleDrawerToggle,className:this.props.classes.menuButton},e),i.a.createElement("div",{className:this.props.classes.flex},i.a.createElement("img",{src:"/logo/logo.png",className:this.props.classes.logo,alt:"LoRa Server"})),i.a.createElement("form",{onSubmit:this.onSearchSubmit},i.a.createElement(H.a,{placeholder:"Search organization, application, gateway or device",className:this.props.classes.search,disableUnderline:!0,value:this.state.search||"",onChange:this.onSearchChange,startAdornment:i.a.createElement(G.a,{position:"start"},i.a.createElement(ne.a,null))})),i.a.createElement("a",{href:"https://www.loraserver.io/lora-app-server/",target:"loraserver-doc"},i.a.createElement(B.a,{className:this.props.classes.iconButton},i.a.createElement(ie.a,null))),i.a.createElement(Q.a,{avatar:i.a.createElement(W.a,null,i.a.createElement(ee.a,null)),label:this.props.user.username,onClick:this.onMenuOpen,classes:{avatar:this.props.classes.avatar,root:this.props.classes.chip}}),i.a.createElement(U.a,{id:"menu-appbar",anchorEl:this.state.menuAnchor,anchorOrigin:{vertical:"top",horizontal:"right"},transformOrigin:{vertical:"top",horizontal:"right"},open:t,onClose:this.onMenuClose},i.a.createElement(_.a,{component:N.a,to:"/users/".concat(this.props.user.id,"/password")},"Change password"),i.a.createElement(_.a,{onClick:this.onLogout},"Logout"))))}}]),t}(a.Component),be=Object(E.withStyles)(ge)(Object(A.a)(me)),de={drawerPaper:{position:"fixed",width:270,paddingTop:9*D.spacing.unit},select:{paddingTop:D.spacing.unit,paddingLeft:3*D.spacing.unit,paddingRight:3*D.spacing.unit,paddingBottom:1*D.spacing.unit}},fe=function(e){function t(){var e;return Object(l.a)(this,t),(e=Object(h.a)(this,Object(p.a)(t).call(this))).state={open:!0,organization:null,cacheCounter:0},e}return Object(g.a)(t,e),Object(u.a)(t,[{key:"componentDidMount",value:function(){}},{key:"render",value:function(){return i.a.createElement(i.a.Fragment,null)}}]),t}(a.Component),Oe=Object(A.a)(Object(E.withStyles)(de)(fe)),je=n(40),ve=n.n(je),ye={footer:{paddingBottom:D.spacing.unit,"& a":{color:D.palette.primary.main,textDecoration:"none"}}},Ee=function(e){function t(){var e;return Object(l.a)(this,t),(e=Object(h.a)(this,Object(p.a)(t).call(this))).state={footer:null},e}return Object(g.a)(t,e),Object(u.a)(t,[{key:"componentDidMount",value:function(){var e=this;pe.getBranding(function(t){""!==t.footer&&e.setState({footer:t.footer})})}},{key:"render",value:function(){return null===this.state.footer?null:i.a.createElement("footer",{className:this.props.classes.footer},i.a.createElement(ve.a,{align:"center",dangerouslySetInnerHTML:{__html:this.state.footer}}))}}]),t}(a.Component),we=Object(E.withStyles)(ye)(Ee),ke=n(125),Se=n.n(ke),Ce=n(120),Ie=n.n(Ce),Te=n(296),De=n.n(Te),Ne=new(function(e){function t(){var e;return Object(l.a)(this,t),(e=Object(h.a)(this,Object(p.a)(t).call(this))).notifications=[],e}return Object(g.a)(t,e),Object(u.a)(t,[{key:"getAll",value:function(){return this.notifications}},{key:"createNotification",value:function(e,t){var n=Date.now();this.notifications.push({id:n,type:e,message:t}),this.emit("change")}},{key:"deleteNotification",value:function(e){var t=null,n=!0,a=!1,i=void 0;try{for(var o,r=this.notifications[Symbol.iterator]();!(n=(o=r.next()).done);n=!0){var s=o.value;s.id===e&&(t=s)}}catch(c){a=!0,i=c}finally{try{n||null==r.return||r.return()}finally{if(a)throw i}}this.notifications.splice(this.notifications.indexOf(t),1),this.emit("change")}},{key:"handleActions",value:function(e){switch(e.type){case"CREATE_NOTIFICATION":this.createNotification(e.notification.type,e.notification.message);break;case"DELETE_NOTIFICATION":this.deleteNotification(e.id)}}}]),t}(oe.EventEmitter));ce.register(Ne.handleActions.bind(Ne));var Ae=Ne,xe=function(e){function t(){var e;return Object(l.a)(this,t),(e=Object(h.a)(this,Object(p.a)(t).call(this))).onClose=e.onClose.bind(Object(m.a)(Object(m.a)(e))),e}return Object(g.a)(t,e),Object(u.a)(t,[{key:"onClose",value:function(e,t){ce.dispatch({type:"DELETE_NOTIFICATION",id:this.props.id})}},{key:"render",value:function(){return i.a.createElement(Se.a,{anchorOrigin:{vertical:"bottom",horizontal:"left"},open:!0,message:i.a.createElement("span",null,this.props.notification.message),autoHideDuration:3e3,onClose:this.onClose,action:[i.a.createElement(Ie.a,{key:"close","aria-label":"Close",color:"inherit",onClick:this.onClose},i.a.createElement(De.a,null))]})}}]),t}(a.Component),ze=function(e){function t(){var e;return Object(l.a)(this,t),(e=Object(h.a)(this,Object(p.a)(t).call(this))).state={notifications:Ae.getAll()},e}return Object(g.a)(t,e),Object(u.a)(t,[{key:"componentDidMount",value:function(){var e=this;Ae.on("change",function(){e.setState({notifications:Ae.getAll()})})}},{key:"render",value:function(){return this.state.notifications.map(function(e,t){return i.a.createElement(xe,{key:e.id,id:e.id,notification:e})})}}]),t}(a.Component),Le=Object(A.a)(ze),Me=n(88),Be=n.n(Me),Re=n(127),_e=n.n(Re),Fe=n(128),Ue=n.n(Fe),Pe=n(89),He=n.n(Pe),qe=n(126),Ge=n.n(qe),Je=function(e){function t(){return Object(l.a)(this,t),Object(h.a)(this,Object(p.a)(t).apply(this,arguments))}return Object(g.a)(t,e),Object(u.a)(t,[{key:"render",value:function(){return i.a.createElement("form",{onSubmit:this.props.onSubmit},this.props.children,i.a.createElement(k.a,{container:!0,justify:"flex-end",className:this.props.classes.formControl},this.props.extraButtons,this.props.submitLabel&&i.a.createElement(Ge.a,{color:"primary",type:"submit",disabled:this.props.disabled},this.props.submitLabel)))}}]),t}(a.Component),We=Object(E.withStyles)({formControl:{paddingTop:24}})(Je),Ke=function(e){function t(){var e;return Object(l.a)(this,t),(e=Object(h.a)(this,Object(p.a)(t).call(this))).state={},e.onChange=e.onChange.bind(Object(m.a)(Object(m.a)(e))),e.onSubmit=e.onSubmit.bind(Object(m.a)(Object(m.a)(e))),e}return Object(g.a)(t,e),Object(u.a)(t,[{key:"onChange",value:function(e){var t=e.target.id.split("."),n=t[t.length-1];t.pop();var a=this.state.object,i=a,o=!0,r=!1,s=void 0;try{for(var c,l=t[Symbol.iterator]();!(o=(c=l.next()).done);o=!0){i=i[c.value]}}catch(u){r=!0,s=u}finally{try{o||null==l.return||l.return()}finally{if(r)throw s}}"checkbox"===e.target.type?i[n]=e.target.checked:"number"===e.target.type?i[n]=parseInt(e.target.value,10):i[n]=e.target.value,this.setState({object:a})}},{key:"onSubmit",value:function(e){e.preventDefault(),this.props.onSubmit(this.state.object)}},{key:"componentDidMount",value:function(){this.setState({object:this.props.object||{}})}},{key:"componentDidUpdate",value:function(e){e.object!==this.props.object&&this.setState({object:this.props.object||{}})}}]),t}(a.Component),Qe={textField:{width:"100%"},link:{"& a":{color:D.palette.primary.main,textDecoration:"none"}}},Ve=function(e){function t(){return Object(l.a)(this,t),Object(h.a)(this,Object(p.a)(t).apply(this,arguments))}return Object(g.a)(t,e),Object(u.a)(t,[{key:"render",value:function(){return void 0===this.state.object?null:i.a.createElement(We,{submitLabel:this.props.submitLabel,onSubmit:this.onSubmit},i.a.createElement(Be.a,{id:"username",label:"Username",margin:"normal",value:this.state.object.username||"",onChange:this.onChange,fullWidth:!0,required:!0}),i.a.createElement(Be.a,{id:"password",label:"Password",type:"password",margin:"normal",value:this.state.object.password||"",onChange:this.onChange,fullWidth:!0,required:!0}))}}]),t}(Ke),Xe=function(e){function t(){var e;return Object(l.a)(this,t),(e=Object(h.a)(this,Object(p.a)(t).call(this))).state={registration:null},e.onSubmit=e.onSubmit.bind(Object(m.a)(Object(m.a)(e))),e}return Object(g.a)(t,e),Object(u.a)(t,[{key:"componentDidMount",value:function(){var e=this;pe.logout(function(){}),pe.getBranding(function(t){""!==t.registration&&e.setState({registration:t.registration})})}},{key:"onSubmit",value:function(e){var t=this;pe.login(e,function(){t.props.history.push("/")})}},{key:"render",value:function(){return i.a.createElement(k.a,{container:!0,justify:"center"},i.a.createElement(k.a,{item:!0,xs:6,lg:4},i.a.createElement(_e.a,null,i.a.createElement(Ue.a,{title:"Login"}),i.a.createElement(He.a,null,i.a.createElement(Ve,{submitLabel:"Login",onSubmit:this.onSubmit})),this.state.registration&&i.a.createElement(He.a,null,i.a.createElement(ve.a,{className:this.props.classes.link,dangerouslySetInnerHTML:{__html:this.state.registration}})))))}}]),t}(a.Component),Ye=Object(E.withStyles)(Qe)(Object(A.a)(Xe)),Ze={root:{flexGrow:1,display:"flex",minHeight:"100vh",flexDirection:"column"},paper:{padding:2*D.spacing.unit,textAlign:"center",color:D.palette.text.secondary},main:{width:"100%",padding:48,paddingTop:115,flex:1},mainDrawerOpen:{paddingLeft:318},footerDrawerOpen:{paddingLeft:270}},$e=function(e){function t(){var e;return Object(l.a)(this,t),(e=Object(h.a)(this,Object(p.a)(t).call(this))).state={user:null,drawerOpen:!1},e.setDrawerOpen=e.setDrawerOpen.bind(Object(m.a)(Object(m.a)(e))),e}return Object(g.a)(t,e),Object(u.a)(t,[{key:"componentDidMount",value:function(){var e=this;pe.on("change",function(){e.setState({user:pe.getUser(),drawerOpen:null!=pe.getUser()})}),this.setState({user:pe.getUser(),drawerOpen:null!=pe.getUser()})}},{key:"setDrawerOpen",value:function(e){this.setState({drawerOpen:e})}},{key:"render",value:function(){var e=null,t=null;return null!==this.state.user&&(e=i.a.createElement(be,{setDrawerOpen:this.setDrawerOpen,drawerOpen:this.state.drawerOpen,user:this.state.user}),t=i.a.createElement(Oe,{open:this.state.drawerOpen,user:this.state.user})),i.a.createElement(b.a,{history:C},i.a.createElement(i.a.Fragment,null,i.a.createElement(y.a,null),i.a.createElement(E.MuiThemeProvider,{theme:D},i.a.createElement("div",{className:this.props.classes.root},e,t,i.a.createElement("div",{className:j()(this.props.classes.main,this.state.drawerOpen&&this.props.classes.mainDrawerOpen)},i.a.createElement(k.a,{container:!0,spacing:24},i.a.createElement(d.a,null,i.a.createElement(f.a,{exact:!0,path:"/",component:Ye}),i.a.createElement(f.a,{exact:!0,path:"/login",component:Ye})))),i.a.createElement("div",{className:this.state.drawerOpen?this.props.classes.footerDrawerOpen:""},i.a.createElement(we,null))),i.a.createElement(Le,null))))}}]),t}(a.Component),et=Object(E.withStyles)(Ze)($e);n(735),n(736),n(737),n(738);c.a.Icon.Default.imagePath="//cdnjs.cloudflare.com/ajax/libs/leaflet/1.0.0/images/",r.a.render(i.a.createElement(et,null),document.getElementById("root"))}},[[361,1,2]]]);
//# sourceMappingURL=main.ff8bd69a.chunk.js.map