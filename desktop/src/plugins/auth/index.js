import React from 'react'
import { Route } from 'react-router'

import Home from './Home'
import NoMatch from './NoMatch'
import Install from './Install'

export default [
  <Route key="site.home" exact path="/" component={Home}/>,
  <Route key="site.install" path="/install" component={Install}/>,

  <Route key="site.no-match" component={NoMatch}/>
]