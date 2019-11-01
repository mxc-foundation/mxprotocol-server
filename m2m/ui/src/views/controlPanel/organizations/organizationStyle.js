import theme from '../../../theme';

const organizationStyle = {
	root: {
		color: '#ffffff'
	},
	TextField: {
		'& input': {
			color: '#FFFFFF'
		}
	},

	card: {
		width: '100%',
		backgroundColor: '#0C027060',
		color: '#ffffff'
	},
	cardTable: {
		'& td': {
			borderBottom: 'none',
			'& span': {
				color: '#00FFD9',
				fontSize: '18px',
				fontWeight: 'bold'
			}
		}
	},

	Button:{
		width:'100%'
	},

	userListSettings:{
		backgroundColor: '#0C027060',
		color: '#ffffff',
		minHeight:'800px',
		height:'100%'
	},
	link: {
        textDecoration: "none",
        fontWeight: "bold",
        fontSize: 12,
        color: theme.palette.textSecondary.main,
        opacity: 0.7,
          "&:hover": {
            opacity: 1,
          }
      },
};

export default organizationStyle;
