import React from 'react';
import { gql } from 'apollo-boost';
import { Query, QueryResult } from 'react-apollo';
import { Box, Heading, Image } from 'grommet';

const QUERY = gql`
  query {
    getAllUsers {
      name
      avatarURL
    }
  }
`;

interface IUser {
  id: number;
  name: string;
  login: string;
  email: string;
  avatarURL: string;
  token: string;
  blihUsername: string;
  blihToken: string;
};

interface IResponse {
  getAllUsers: IUser[];
}

const Users: React.FC = () => (
  <Query query={QUERY}>
    {({ loading, error, data } : QueryResult<IResponse>) => {
      if (loading) return <div>Loading...</div>;
      if (error || data === undefined) return <div>Error :(</div>;

      if (data.getAllUsers.length === 0) {
        return (
          <Heading level='2' color='dark-3'>Wow such empty!</Heading>
        )
      }

      return (
        <Box justify='start' fill={true}>
        {data.getAllUsers.map((user, i) => (
          <Box
            key={i}
            width='medium'
            height='small'
            background='brand'
            elevation='medium'
            round='xsmall'
            animation='slideUp'
          >
            <Image fit='cover' src={user.avatarURL} />
            <Heading
              truncate={true}
              level='4'
              alignSelf='center'
              margin='small'
            >
              {user.name}
            </Heading>
          </Box>
        ))}
        </Box>
      );
    }}
  </Query>
);

export default Users;
