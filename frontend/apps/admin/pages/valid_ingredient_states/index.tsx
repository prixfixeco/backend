import { GetServerSideProps, GetServerSidePropsContext, GetServerSidePropsResult } from 'next';
import { Button, Grid, Pagination, Stack, Table, TextInput } from '@mantine/core';
import { AxiosError } from 'axios';
import { formatRelative } from 'date-fns';
import router from 'next/router';
import { IconSearch } from '@tabler/icons';
import { useEffect, useState } from 'react';

import {
  EitherErrorOr,
  IAPIError,
  QueryFilter,
  QueryFilteredResult,
  ValidIngredientState,
} from '@dinnerdonebetter/models';
import { ServerTiming, ServerTimingHeaderName } from '@dinnerdonebetter/server-timing';
import { buildLocalClient } from '@dinnerdonebetter/api-client';

import { buildServerSideClientOrRedirect } from '../../src/client';
import { AppLayout } from '../../src/layouts';
import { serverSideTracer } from '../../src/tracer';
import { valueOrDefault } from '../../src/utils';

declare interface ValidIngredientStatesPageProps {
  pageLoadValidIngredientStates: EitherErrorOr<QueryFilteredResult<ValidIngredientState>>;
}

export const getServerSideProps: GetServerSideProps = async (
  context: GetServerSidePropsContext,
): Promise<GetServerSidePropsResult<ValidIngredientStatesPageProps>> => {
  const timing = new ServerTiming();
  const span = serverSideTracer.startSpan('ValidIngredientStatesPage.getServerSideProps');

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
  let props!: GetServerSidePropsResult<ValidIngredientStatesPageProps>;

  const qf = QueryFilter.deriveFromGetServerSidePropsContext(context.query);
  qf.attachToSpan(span);

  const fetchValidIngredientStatesTimer = timing.addEvent('fetch valid ingredient states');
  await apiClient
    .getValidIngredientStates(qf)
    .then((res: QueryFilteredResult<ValidIngredientState>) => {
      span.addEvent('valid ingredientStates retrieved');
      props = {
        props: {
          pageLoadValidIngredientStates: JSON.parse(JSON.stringify(res)),
        },
      };
    })
    .catch((error: IAPIError) => {
      span.addEvent('error occurred', { error: error.message });
      return { error };
    })
    .finally(() => {
      fetchValidIngredientStatesTimer.end();
    });

  context.res.setHeader(ServerTimingHeaderName, timing.headerValue());

  span.end();
  return props;
};

function ValidIngredientStatesPage(props: ValidIngredientStatesPageProps) {
  let { pageLoadValidIngredientStates } = props;

  const ogValidINgredientStates = valueOrDefault(
    pageLoadValidIngredientStates,
    new QueryFilteredResult<ValidIngredientState>(),
  );
  const [validIngredientStatesError] = useState<IAPIError | undefined>(pageLoadValidIngredientStates.error);
  const [validIngredientStates, setValidIngredientStates] =
    useState<QueryFilteredResult<ValidIngredientState>>(ogValidINgredientStates);
  const [search, setSearch] = useState('');

  useEffect(() => {
    const apiClient = buildLocalClient();

    if (search.trim().length < 1) {
      const qf = QueryFilter.deriveFromGetServerSidePropsContext({ search });
      apiClient
        .getValidIngredientStates(qf)
        .then((res: QueryFilteredResult<ValidIngredientState>) => {
          setValidIngredientStates(res);
        })
        .catch((err: AxiosError) => {
          console.error(err);
        });
    } else {
      apiClient
        .searchForValidIngredientStates(search)
        .then((res: QueryFilteredResult<ValidIngredientState>) => {
          setValidIngredientStates(res);
        })
        .catch((err: AxiosError) => {
          console.error(err);
        });
    }
  }, [search]);

  useEffect(() => {
    const apiClient = buildLocalClient();

    const qf = QueryFilter.deriveFromPage();
    qf.page = validIngredientStates.page;

    apiClient
      .getValidIngredientStates(qf)
      .then((res: QueryFilteredResult<ValidIngredientState>) => {
        setValidIngredientStates(res);
      })
      .catch((err: AxiosError) => {
        console.error(err);
      });
  }, [validIngredientStates.page]);

  const formatDate = (x: string | undefined): string => {
    return x ? formatRelative(new Date(x), new Date()) : 'never';
  };

  const rows = (validIngredientStates.data || []).map((ingredientState) => (
    <tr
      key={ingredientState.id}
      onClick={() => router.push(`/valid_ingredient_states/${ingredientState.id}`)}
      style={{ cursor: 'pointer' }}
    >
      <td>{ingredientState.id}</td>
      <td>{ingredientState.name}</td>
      <td>{ingredientState.pastTense}</td>
      <td>{ingredientState.slug}</td>
      <td>{formatDate(ingredientState.createdAt)}</td>
      <td>{formatDate(ingredientState.lastUpdatedAt)}</td>
    </tr>
  ));

  return (
    <AppLayout title="Valid IngredientStates">
      <Stack>
        <Grid justify="space-between">
          <Grid.Col md="auto" sm={12}>
            <TextInput
              placeholder="Search..."
              icon={<IconSearch size={14} />}
              onChange={(event) => setSearch(event.target.value || '')}
            />
          </Grid.Col>
          <Grid.Col md="content" sm={12}>
            <Button
              onClick={() => {
                router.push('/valid_ingredient_states/new');
              }}
            >
              Create New
            </Button>
          </Grid.Col>
        </Grid>

        {validIngredientStatesError && <p>{validIngredientStatesError.message}</p>}
        {!validIngredientStatesError && validIngredientStates.data && (
          <>
            <Table mt="xl" striped highlightOnHover withBorder withColumnBorders>
              <thead>
                <tr>
                  <th>ID</th>
                  <th>Name</th>
                  <th>Past Tense</th>
                  <th>Slug</th>
                  <th>Created At</th>
                  <th>Last Updated At</th>
                </tr>
              </thead>
              <tbody>{rows}</tbody>
            </Table>

            <Pagination
              disabled={search.trim().length > 0}
              position="center"
              page={validIngredientStates.page}
              total={Math.ceil(validIngredientStates.totalCount / validIngredientStates.limit)}
              onChange={(value: number) => {
                setValidIngredientStates({ ...validIngredientStates, page: value });
              }}
            />
          </>
        )}
      </Stack>
    </AppLayout>
  );
}

export default ValidIngredientStatesPage;
