import React from 'react';
import { Grommet, Heading } from 'grommet';
import {
  BrowserRouter as Router,
  Route,
  Link,
  Switch,
} from 'react-router-dom';

import Theme from './Configs/Theme';
import { AuthProvider } from './Contexts/Auth';
import AuthRequired from './Components/AuthRequired';
import AuthNotRequired from './Components/AuthNotRequired';
import NavBar from './Components/NavBar';
import Login from './Components/Login';
import Register from './Components/Register';
import NotFound from './Components/NotFound';

const Hello: React.FC = () => (
  <Heading>Hello world!</Heading>
);

const Admin: React.FC = () => (
  <Heading>Admin zone :)</Heading>
)

const App: React.FC = () => (
  <AuthProvider>
    <Router>
      <Grommet theme={Theme}>
        <NavBar />
        <ul>
          <li>
            <Link to="/">Home</Link>
          </li>
          <li>
            <Link to="/login">Login</Link>
          </li>
          <li>
            <Link to="/register">Register</Link>
          </li>
        </ul>
          <Switch>
            <Route exact path='/' component={Hello} />
            <AuthNotRequired path='/register' component={Register} />
            <AuthNotRequired path='/login' component={Login} />
            <AuthRequired path='/admin' component={Admin} />
            <Route component={NotFound} />
          </Switch>
      </Grommet>
    </Router>
  </AuthProvider>
);

export default App;
