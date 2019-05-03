import React from 'react';
import { Box, Heading, Anchor, Button } from 'grommet';

import { AuthConsumer } from '../Contexts/Auth';

const AppBar: React.FC = (props) => (
  <Box
    tag='header'
    direction='row'
    align='center'
    justify='between'
    background='brand'
    pad='small'
    elevation='medium'
    {...props}
  />
);

const NavBar: React.FC = () => (
  <AppBar>
    <Heading level='3' margin='none'>GitSync</Heading>
    <Box
      direction='row'
      justify='between'
      align='center'
      gap='medium'>
      <AuthConsumer>
        {({ isAuth, login, logout }) => isAuth ? (
            <Anchor label='Sign out' onClick={ logout } />
          ) : (
            <>
              <Anchor label='Sign in' onClick={ login } />
              <Button label='Sign up' />
            </>
          )
        }
      </AuthConsumer>
    </Box>
  </AppBar>
);

export default NavBar;
