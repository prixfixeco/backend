import { GetServerSideProps, GetServerSidePropsContext, GetServerSidePropsResult } from 'next';
import { Button, Grid, Pagination, Stack, Table, TextInput } from '@mantine/core';
import { AxiosError } from 'axios';
import { formatRelative } from 'date-fns';
import router from 'next/router';
import { IconSearch } from '@tabler/icons';
import { useEffect, useState } from 'react';

import { EitherErrorOr, IAPIError, QueryFilter, QueryFilteredResult, ServiceSetting } from '@dinnerdonebetter/models';
import { ServerTiming, ServerTimingHeaderName } from '@dinnerdonebetter/server-timing';
import { buildLocalClient } from '@dinnerdonebetter/api-client';

import { buildServerSideClientOrRedirect } from '../../src/client';
import { AppLayout } from '../../src/layouts';
import { serverSideTracer } from '../../src/tracer';
import { valueOrDefault } from '../../src/utils';

declare interface ServiceSettingsPageProps {
  pageLoadServiceSettings: EitherErrorOr<QueryFilteredResult<ServiceSetting>>;
}

export const getServerSideProps: GetServerSideProps = async (
  context: GetServerSidePropsContext,
): Promise<GetServerSidePropsResult<ServiceSettingsPageProps>> => {
  const timing = new ServerTiming();
  const span = serverSideTracer.startSpan('ServiceSettingsPage.getServerSideProps');

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
  let props!: GetServerSidePropsResult<ServiceSettingsPageProps>;

  const qf = QueryFilter.deriveFromGetServerSidePropsContext(context.query);
  qf.attachToSpan(span);

  const fetchServiceSettingsTimer = timing.addEvent('fetch service settings');
  await apiClient
    .getServiceSettings(qf)
    .then((res: QueryFilteredResult<ServiceSetting>) => {
      span.addEvent('service settings retrieved');
      props = {
        props: {
          pageLoadServiceSettings: JSON.parse(JSON.stringify(res)),
        },
      };
    })
    .catch((error: IAPIError) => {
      span.addEvent('error occurred', { error: error.message });
      return { error };
    })
    .finally(() => {
      fetchServiceSettingsTimer.end();
    });

  context.res.setHeader(ServerTimingHeaderName, timing.headerValue());

  span.end();
  return props;
};

function ServiceSettingsPage(props: ServiceSettingsPageProps) {
  let { pageLoadServiceSettings } = props;

  const ogServiceSettings: QueryFilteredResult<ServiceSetting> = valueOrDefault(
    pageLoadServiceSettings,
    new QueryFilteredResult<ServiceSetting>(),
  );
  const [serviceSettingsError] = useState(pageLoadServiceSettings.error);
  const [serviceSettings, setServiceSettings] = useState<QueryFilteredResult<ServiceSetting>>(ogServiceSettings);
  const [search, setSearch] = useState('');

  useEffect(() => {
    const apiClient = buildLocalClient();

    if (search.trim().length < 1) {
      const qf = QueryFilter.deriveFromGetServerSidePropsContext({ search });
      apiClient
        .getServiceSettings(qf)
        .then((res: QueryFilteredResult<ServiceSetting>) => {
          setServiceSettings(res);
        })
        .catch((err: AxiosError) => {
          console.error(err);
        });
    } else {
      apiClient
        .searchForServiceSettings(search)
        .then((res: QueryFilteredResult<ServiceSetting>) => {
          setServiceSettings(res);
        })
        .catch((err: AxiosError) => {
          console.error(err);
        });
    }
  }, [search]);

  useEffect(() => {
    const apiClient = buildLocalClient();

    const qf = QueryFilter.deriveFromPage();
    qf.page = serviceSettings.page;

    apiClient
      .getServiceSettings(qf)
      .then((res: QueryFilteredResult<ServiceSetting>) => {
        setServiceSettings(res);
      })
      .catch((err: AxiosError) => {
        console.error(err);
      });
  }, [serviceSettings.page]);

  const formatDate = (x: string | undefined): string => {
    return x ? formatRelative(new Date(x), new Date()) : 'never';
  };

  const rows = (serviceSettings.data || []).map((setting) => (
    <tr key={setting.id} onClick={() => router.push(`/service_settings/${setting.id}`)} style={{ cursor: 'pointer' }}>
      <td>{setting.name}</td>
      <td>{formatDate(setting.createdAt)}</td>
      <td>{formatDate(setting.lastUpdatedAt)}</td>
    </tr>
  ));

  return (
    <AppLayout title="Service Settings">
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
                router.push('/service_settings/new');
              }}
            >
              Create New
            </Button>
          </Grid.Col>
        </Grid>

        {serviceSettingsError && <div>{serviceSettingsError.message}</div>}
        {!serviceSettingsError && serviceSettings.data.length > 0 && (
          <>
            <Table mt="xl" striped highlightOnHover withBorder withColumnBorders>
              <thead>
                <tr>
                  <th>Name</th>
                  <th>Created At</th>
                  <th>Last Updated At</th>
                </tr>
              </thead>
              <tbody>{rows}</tbody>
            </Table>

            <Pagination
              disabled={search.trim().length > 0}
              position="center"
              page={serviceSettings.page}
              total={Math.ceil(serviceSettings.totalCount / serviceSettings.limit)}
              onChange={(value: number) => {
                setServiceSettings({ ...serviceSettings, page: value });
              }}
            />
          </>
        )}
      </Stack>
    </AppLayout>
  );
}

export default ServiceSettingsPage;
