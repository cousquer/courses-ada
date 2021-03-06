import React, { Component } from "react"
import PropTypes from "prop-types"
import { withStyles, createStyleSheet } from "material-ui/styles"
import Typography from "material-ui/Typography"
import ExpandableInstructors from "./ExpandableInstructors"

const styleSheet = createStyleSheet("Instructors", theme => ({
  button: {
    paddingLeft: 0
  },
  links: {
    color: "#3344dd",
    display: "flex",
    flexDirection: "column",
    borderLeftStyle: "solid",
    borderLeftColor: theme.palette.accent[400],
    borderLeftWidth: "0.3em",
    paddingLeft: "1em"
  },
  link: {
    marginBottom: 10
  },
  teacher: {
    fontSize: 16,
    color: theme.palette.text.primary,
    marginBottom: "0.5em"
  }
}))

class Instructors extends Component {
  state = {
    anchorEl: undefined,
    open: false
  }

  handleClick = event => {
    this.setState({ open: true, anchorEl: event.currentTarget })
  }

  handleRequestClose = () => {
    this.setState({ open: false })
  }

  getInstructor() {
    const classes = this.props.classes
    if (this.props.teachers.length >= 2) {
      return <ExpandableInstructors teachers={this.props.teachers} />
    } else {
      return (
        <div key={this.props.teachers[0].lastName + Math.random()}>
          <Typography
            type="headline"
            component="h3"
            className={classes.teacher}
            tabIndex="0"
            key={this.props.teachers[0].lastName + Math.random()}
          >
            {this.props.teachers[0].firstName +
              " " +
              this.props.teachers[0].lastName}
          </Typography>
          <div
            className={classes.links}
            key={this.props.teachers[0].lastName + Math.random()}
          >
            <a
              className={classes.link}
              target="_blank"
              href="https://oakland.edu"
              tabIndex="0"
              rel="noopener noreferrer"
              key={this.props.teachers[0].lastName + Math.random()}
            >
              {this.props.teachers[0].office}
            </a>
            <a
              target="_blank"
              href="mailto:https://oakland.edu"
              tabIndex="0"
              rel="noopener noreferrer"
              key={this.props.teachers[0].lastName + Math.random()}
            >
              {this.props.teachers[0].email}
            </a>
          </div>
        </div>
      )
    }
  }

  render() {
    return (
      <div>
        {this.getInstructor()}
      </div>
    )
  }
}

Instructors.propTypes = {
  classes: PropTypes.object.isRequired
}

export default withStyles(styleSheet)(Instructors)
