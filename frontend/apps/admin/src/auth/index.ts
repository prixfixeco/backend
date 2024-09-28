import { AxiosResponse } from 'axios';
import { serialize } from 'cookie';
import { NextApiRequestCookies } from 'next/dist/server/api-utils';

import { webappCookieName } from '../constants';
import { serverSideTracer } from '../tracer';

export interface sessionAuth {
  userID: string;
  householdID: string;
}

export function processWebappCookieHeader(result: AxiosResponse, userID: string, householdID: string): string[] {
  const span = serverSideTracer.startSpan('processWebappCookieHeader');

  let modifiedAPICookie = result.headers['set-cookie']?.[0] ?? '';
  if (!modifiedAPICookie) {
    throw new Error('missing cookie header');
  }

  if (process.env.REWRITE_COOKIE_HOST_FROM && process.env.REWRITE_COOKIE_HOST_TO) {
    modifiedAPICookie = modifiedAPICookie.replace(
      process.env.REWRITE_COOKIE_HOST_FROM,
      process.env.REWRITE_COOKIE_HOST_TO,
    );
    span.addEvent('cookie host rewritten');
  }

  if (process.env.REWRITE_COOKIE_SECURE === 'true') {
    modifiedAPICookie = modifiedAPICookie.replace('Secure; ', '');
    span.addEvent('secure setting rewritten in cookie');
  }

  const webappCookie = serialize(
    webappCookieName,
    Buffer.from(JSON.stringify({ userID, householdID } as sessionAuth), 'ascii').toString('base64'),
    { path: '/' },
  );

  span.end();
  return [modifiedAPICookie, webappCookie];
}

export const extractUserInfoFromCookie = (cookies: NextApiRequestCookies): sessionAuth => {
  return JSON.parse(Buffer.from(cookies[webappCookieName] || 'e30=', 'base64').toString('ascii')) as sessionAuth;
};
