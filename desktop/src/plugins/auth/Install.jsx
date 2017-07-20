import React, { Component } from 'react'
import { connect } from 'react-redux'
import PropTypes from 'prop-types'
import { push } from 'react-router-redux'
import Divider from 'material-ui/Divider';

import Layout from '../../layouts/Application'

class Widget extends Component {
  render() {
    return (
      <Layout breads={[]}>
        <div>install</div>
         <Divider/>
      </Layout>
    );
  }
}


Widget.propTypes = {
  cards: PropTypes.array.isRequired,
  push: PropTypes.func.isRequired,
}


export default connect(
  state => ({
    cards: state.siteInfo.cards ? state.siteInfo.cards : [],
  }),
  {push},
)(Widget)
