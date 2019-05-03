import React from 'react';
import { Route, Redirect } from 'react-router-dom';

import { AuthConsumer } from '../Contexts/Auth';

const AuthRequired: React.FC<any> = ({ component: Component, ...rest }) => (
  <AuthConsumer>
    {({ isAuth }) => (
      <Route {...rest} render={(props) => {
        return isAuth ? (
          <Component {...props} />
        ) : (
          <Redirect to='/login' />
        )}
      } />
    )}
  </AuthConsumer>
);

export default AuthRequired;
