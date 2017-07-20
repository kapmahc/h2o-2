import React, { Component } from 'react'
import PropTypes from 'prop-types'
import {injectIntl, intlShape, FormattedMessage} from 'react-intl'
import { connect } from 'react-redux'
import { push } from 'react-router-redux'
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';

import {setLocale} from '../intl'
import {signIn, signOut, refresh} from '../actions'
import {get, _delete} from '../ajax'

class WidgetF extends Component {
  render () {
    <MuiThemeProvider>
      <h1>application</h1>
      <div>{children}</div>
    </MuiThemeProvider>
  }
}

WidgetF.propTypes = {
  children: PropTypes.node.isRequired,
  push: PropTypes.func.isRequired,
  refresh: PropTypes.func.isRequired,
  signIn: PropTypes.func.isRequired,
  signOut: PropTypes.func.isRequired,
  user: PropTypes.object.isRequired,
  info: PropTypes.object.isRequired,
  breads: PropTypes.array.isRequired,
  intl: intlShape.isRequired,
}

export default connect(
  state => ({
    user: state.currentUser,
    info: state.siteInfo,
  }),
  {push, signIn, refresh, signOut},
)(Widget)


const Widget = injectIntl(WidgetF)
