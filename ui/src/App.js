import React, {Fragment} from 'react'
import { Route } from 'react-router-dom'
import './App.css'

import Home from './Views/Home'
import Login from './Views/Login'

const App = () =>
  <Fragment>
    <Route exact path="/" component={Home}/>
    <Route exact path="/login" component={Login}/>
  </Fragment>

export default App
