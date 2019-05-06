import React from 'react';
import { Link, withRouter, RouteComponentProps, LinkProps } from 'react-router-dom';
import { Box, Anchor } from 'grommet';

interface activeLinkProps extends LinkProps {
  path: string;
};

const defaultLinkStyle: React.CSSProperties = {
  textDecoration: 'none'
};

const ActiveLink: React.FC<activeLinkProps> = ({ to, title, path, ...rest}) => (
  <Link to={to} style={defaultLinkStyle} {...rest}>
    <Anchor as='span' label={title} color={path === to ? 'dark-6' : 'brand'} />
  </Link>
);

const Links: React.FC<RouteComponentProps> = ({ location }) => (
  <Box direction='row' align='center' justify='between' gap='medium'>
    <ActiveLink to='/profile' title='profile' path={location.pathname} />
    <ActiveLink to='/users' title='users' path={location.pathname} />
    <ActiveLink to='/repos' title='repos' path={location.pathname} />
  </Box>
);

export default withRouter(Links);
