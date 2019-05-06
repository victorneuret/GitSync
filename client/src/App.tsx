import React from 'react';
import { Grommet, Heading, Box } from 'grommet';
import {
  BrowserRouter as Router,
  Route,
  Switch,
} from 'react-router-dom';

import Theme from './Configs/Theme';
import { AuthProvider } from './Contexts/Auth';
import AuthRequired from './Components/AuthRequired';
import AuthNotRequired from './Components/AuthNotRequired';
import NavBar from './Components/NavBar';
import Login from './Components/Login';
import Links from './Components/Links';
import Register from './Components/Register';
import NotFound from './Components/NotFound';
import Repos from './Components/Repos';
import Users from './Components/Users';

const Hello: React.FC = () => (
  <Heading>Hello</Heading>
);

const Admin: React.FC = () => (
  <Heading>Admin zone :)</Heading>
)

const App: React.FC = () => (
  <AuthProvider>
    <Router>
      <Grommet theme={Theme}>
        <NavBar />
        <Box fill align='center' justify='start' pad='large'>
          <Links />
          <Switch>
            <Route exact path='/' component={Hello} />
            <AuthNotRequired path='/register' component={Register} />
            <AuthRequired path='/users' component={Users} />
            <AuthRequired path='/repos' component={Repos} />
            <AuthNotRequired path='/login' component={Login} />
            <AuthRequired path='/admin' component={Admin} />
            <Route component={NotFound} />
          </Switch>
        </Box>
      </Grommet>
    </Router>
  </AuthProvider>
);

export default App;
