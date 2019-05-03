import React from 'react';
import { Route, Redirect } from 'react-router-dom';

import { AuthConsumer } from '../Contexts/Auth';

const AuthNotRequired: React.FC<any> = ({ component: Component, ...rest }) => (
  <AuthConsumer>
    {({ isAuth }) => (
      <Route {...rest} render={(props) => {
        return isAuth ? (
          <Redirect to='/' />
          ) : (
          <Component {...props} />
        )}
      } />
    )}
  </AuthConsumer>
);

export default AuthNotRequired;
