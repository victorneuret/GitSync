import React from 'react';

import PersistentState from './PersistentState';

const _empty = (): void => {};

const AuthContext = React.createContext({
  isAuth: false,
  login: _empty,
  logout: _empty
});

const AuthProvider: React.FC = (props) => {
  const [isAuth, setAuth] = PersistentState('isAuth', false);

  return (
    <AuthContext.Provider
      value={{
        isAuth: isAuth,
        login: () => setAuth(true),
        logout: () => setAuth(false)
      }}
    >
      {props.children}
    </AuthContext.Provider>
  );
};

const AuthConsumer = AuthContext.Consumer;

export { AuthConsumer, AuthProvider };
