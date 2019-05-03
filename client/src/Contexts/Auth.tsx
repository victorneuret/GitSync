import React from 'react';

const _empty = (): void => {};

const AuthContext = React.createContext({
  isAuth: false,
  login: _empty,
  logout: _empty
});

class AuthProvider extends React.Component {
  state = {
    isAuth: false ,
    login: () => this.setState({ isAuth: true }),
    logout: () => this.setState({ isAuth: false })
  };

  render() {
    return (
      <AuthContext.Provider value={this.state}>
        {this.props.children}
      </AuthContext.Provider>
    );
  };
};

const AuthConsumer = AuthContext.Consumer;

export { AuthConsumer, AuthProvider };
