import { GetServerSideProps, GetServerSidePropsContext, GetServerSidePropsResult } from 'next';
import { Grid, Pagination, Stack, Table, TextInput } from '@mantine/core';
import { AxiosError } from 'axios';
import { formatRelative } from 'date-fns';
import { IconSearch } from '@tabler/icons';
import { useEffect, useState } from 'react';

import { EitherErrorOr, IAPIError, QueryFilter, QueryFilteredResult, User } from '@dinnerdonebetter/models';
import { ServerTiming, ServerTimingHeaderName } from '@dinnerdonebetter/server-timing';
import { buildLocalClient } from '@dinnerdonebetter/api-client';

import { buildServerSideClientOrRedirect } from '../../src/client';
import { AppLayout } from '../../src/layouts';
import { serverSideTracer } from '../../src/tracer';
import { valueOrDefault } from '../../src/utils';

declare interface UsersPageProps {
  pageLoadUsers: EitherErrorOr<QueryFilteredResult<User>>;
}

export const getServerSideProps: GetServerSideProps = async (
  context: GetServerSidePropsContext,
): Promise<GetServerSidePropsResult<UsersPageProps>> => {
  const timing = new ServerTiming();
  const span = serverSideTracer.startSpan('UsersPage.getServerSideProps');

  const clientOrRedirect = buildServerSideClientOrRedirect(context);
  if (clientOrRedirect.redirect) {
    span.end();
    return { redirect: clientOrRedirect.redirect };
  }

  if (!clientOrRedirect.client) {
    // this should never occur if the above state is false
    throw new Error('no client returned');
  }
  const apiClient = clientOrRedirect.client.withSpan(span);

  // TODO: parse context.query as QueryFilter.
  let props!: GetServerSidePropsResult<UsersPageProps>;

  const qf = QueryFilter.deriveFromGetServerSidePropsContext(context.query);
  qf.attachToSpan(span);

  const fetchUsersTimer = timing.addEvent('fetch users');
  await apiClient
    .getUsers(qf)
    .then((res: QueryFilteredResult<User>) => {
      span.addEvent('users retrieved');
      props = {
        props: {
          pageLoadUsers: JSON.parse(JSON.stringify(res)),
        },
      };
    })
    .catch((error: IAPIError) => {
      span.addEvent('error occurred', { error: error.message });
      return { error };
    })
    .finally(() => {
      fetchUsersTimer.end();
    });

  context.res.setHeader(ServerTimingHeaderName, timing.headerValue());

  span.end();
  return props;
};

function UsersPage(props: UsersPageProps) {
  let { pageLoadUsers } = props;

  const ogUsers = valueOrDefault(pageLoadUsers, new QueryFilteredResult<User>());
  const [usersError] = useState<IAPIError | undefined>(pageLoadUsers.error);
  const [users, setUsers] = useState<QueryFilteredResult<User>>(ogUsers);
  const [search, setSearch] = useState('');

  useEffect(() => {
    const apiClient = buildLocalClient();

    if (search.trim().length < 1) {
      const qf = QueryFilter.deriveFromGetServerSidePropsContext({ search });
      apiClient
        .getUsers(qf)
        .then((res: QueryFilteredResult<User>) => {
          setUsers(res);
        })
        .catch((err: AxiosError) => {
          console.error(err);
        });
    } else {
      apiClient
        .searchForUsers(search)
        .then((res: QueryFilteredResult<User>) => {
          setUsers(res);
        })
        .catch((err: AxiosError) => {
          console.error(err);
        });
    }
  }, [search]);

  useEffect(() => {
    const apiClient = buildLocalClient();

    const qf = QueryFilter.deriveFromPage();
    qf.page = users.page;

    apiClient
      .getUsers(qf)
      .then((res: QueryFilteredResult<User>) => {
        setUsers(res);
      })
      .catch((err: AxiosError) => {
        console.error(err);
      });
  }, [users.page]);

  const formatDate = (x: string | undefined): string => {
    return x ? formatRelative(new Date(x), new Date()) : 'never';
  };

  const rows = (users.data || []).map((preparation) => (
    <tr key={preparation.id} style={{ cursor: 'pointer' }}>
      <td>{preparation.id}</td>
      <td>{preparation.username}</td>
      <td>{formatDate(preparation.createdAt)}</td>
      <td>{formatDate(preparation.lastUpdatedAt)}</td>
    </tr>
  ));

  return (
    <AppLayout title="Valid Preparations">
      <Stack>
        <Grid justify="space-between">
          <Grid.Col md="auto" sm={12}>
            <TextInput
              placeholder="Search..."
              icon={<IconSearch size={14} />}
              onChange={(event) => setSearch(event.target.value || '')}
            />
          </Grid.Col>
        </Grid>

        {usersError && <div>{usersError.message}</div>}
        {!usersError && users.data.length > 0 && (
          <>
            <Table mt="xl" striped highlightOnHover withBorder withColumnBorders>
              <thead>
                <tr>
                  <th>ID</th>
                  <th>Username</th>
                  <th>Created At</th>
                  <th>Last Updated At</th>
                </tr>
              </thead>
              <tbody>{rows}</tbody>
            </Table>

            <Pagination
              disabled={search.trim().length > 0}
              position="center"
              page={users.page}
              total={Math.ceil(users.totalCount / users.limit)}
              onChange={(value: number) => {
                setUsers({ ...users, page: value });
              }}
            />
          </>
        )}
      </Stack>
    </AppLayout>
  );
}

export default UsersPage;
