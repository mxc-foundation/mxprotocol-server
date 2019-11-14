import React, { Component } from "react";
import Select from "react-select";
import SessionStore from "../stores/SessionStore";
import { supportedLanguages } from "../i18n";

const customStyles = {
  control: (base, state) => ({
    ...base,
    margin: 20,
    // match with the menu
    borderRadius: state.isFocused ? "3px 3px 0 0" : 3,
    // Overwrittes the different states of border
    borderColor: state.isFocused ? "#00FFD9" : "white",
    // Removes weird border around container
    boxShadow: state.isFocused ? null : null,
    "&:hover": {
      // Overwrittes the different states of border
      borderColor: state.isFocused ? "#00FFD9" : "white"
    }
  }),
  menu: base => ({
    ...base,
    background:"#101c4a",
    // override border radius to match the box
    borderRadius: 0,
    // kill the gap
    marginTop: 0,
    paddingLeft: 20,
    paddingRight: 20,
  }),
  menuList: base => ({
    ...base,
    background: "#1a2d6e",
    // kill the white space on first and last option
    paddingTop: 0,
  }),
  option: base => ({
    ...base,
    // kill the white space on first and last option
    padding: "10px",
    maxWidth: 229,
    whiteSpace: "nowrap", 
    overflow: "hidden",
    textOverflow: "ellipsis"
  }),
};

export default class WithPromises extends Component {
  constructor() {
    super();

    this.state = {
      selectedOption: null,
      options:[]
    };
  } 

  componentDidMount() {
    let selectedOption = null;

    const storedLanguageID = SessionStore.getLanguage() && SessionStore.getLanguage().id;
    const storedLanguageName = SessionStore.getLanguage() && SessionStore.getLanguage().name;

    if (storedLanguageID && storedLanguageName) {
      selectedOption = {
        label: storedLanguageID,
        value: storedLanguageName
      };
    }

    this.setState({
      selectedOption,
      options: supportedLanguages
    })
  }

  onChange = (selectedOption) => {
    if (selectedOption !== null && selectedOption.label !== null && selectedOption.value !== null) {
      this.setState({
        selectedOption
      });
  
      this.props.onChange({
        target: {
          label: selectedOption.label,
          value: selectedOption.value,
        },
      });
    }
  }

  render() {
    const { selectedOption } = this.state;

    return (
      <Select
        styles={customStyles}
        theme={(theme) => ({
          ...theme,
          borderRadius: 4,
          colors: {
            primary25: "#00FFD950",
            primary: "#00FFD950",
          },
        })}
        onChange={this.onChange}
        options={supportedLanguages}
        value={selectedOption}
      />
    );
  }
}
