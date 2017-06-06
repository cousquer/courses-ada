// @flow weak
/* eslint-disable react/no-multi-comp */

import React, { Component } from "react"
import PropTypes from "prop-types"
import { withStyles, createStyleSheet } from "material-ui/styles"
import Paper from "material-ui/Paper"
import Tabs, { Tab } from "material-ui/Tabs"
import Courses from "./Courses"
import Grades from "./Grades"
import { translate, Interpolate } from "react-i18next"
import i18n from "./utils/i18n"

const TabContainer = props => (
  <div style={{ padding: 20 }}>
    {props.children}
  </div>
)

TabContainer.propTypes = {
  children: PropTypes.node.isRequired
}

const styleSheet = createStyleSheet("BasicTabs", theme => ({
  root: {
    flexGrow: 1,
    marginTop: 30
  },
  tabs: {
    height: "0.3em"
  },
  appBar: {
    backgroundColor: theme.palette.primary[500],
    color: theme.palette.getContrastText(theme.palette.primary[500])
  }
}))

class BasicTabs extends Component {
  state = {
    index: 0
  }

  handleChange = (event, index) => {
    this.setState({ index })
  }

  render() {
    const classes = this.props.classes
    const { t } = this.props

    return (
      <Paper className={classes.root}>
        <div className={classes.appBar}>
          <Tabs
            index={this.state.index}
            onChange={this.handleChange}
            indicatorClassName={classes.tabs}
          >
            <Tab label={t("courses", {})} tabIndex="0" />
            <Tab label={t("calendar", {})} tabIndex="0" />
            <Tab label={t("grades", {})} tabIndex="0" />
          </Tabs>
        </div>
        {this.state.index === 0 &&
          <TabContainer>
            <div>
              <Courses
                tabIndex="0"
                currentTermCode={this.props.currentTermCode}
                courses={this.props.courses}
              />
            </div>
          </TabContainer>}
        {this.state.index === 1 &&
          <TabContainer>
            <div tabIndex="0"><p tabIndex="0">Hello Aaron</p></div>
          </TabContainer>}
        {this.state.index === 2 &&
          <TabContainer>
            <Grades tabIndex="0" courses={this.props.courses} />
          </TabContainer>}
      </Paper>
    )
  }
}

BasicTabs.propTypes = {
  classes: PropTypes.object.isRequired
}

export default withStyles(styleSheet) (translate("view", { wait: true })(BasicTabs))