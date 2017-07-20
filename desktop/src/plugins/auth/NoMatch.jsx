import React, { Component } from 'react'
import {FormattedMessage} from 'react-intl'

import fail from '../../assets/fail.png'
import Layout from '../../layouts/Application'

class Widget extends Component {
  render() {
    return (
      <Layout breads={[]}>
        <div>
          <img title={<FormattedMessage id="errors.no-match"/>} alt="fail" width="100%" src={fail} />
        </div>
      </Layout>
    );
  }
}

export default Widget;
