// GENERATED CODE, DO NOT EDIT MANUALLY

import type { Page, Route } from '@playwright/test';

import { UserNotification } from '@dinnerdonebetter/models';

import { assertClient, assertMethod, ResponseConfig } from './helpers';

export class MockCreateUserNotificationResponseConfig extends ResponseConfig<UserNotification> {
  constructor(status: number = 201, body?: UserNotification) {
    super();

    this.status = status;
    if (this.body) {
      this.body = body;
    }
  }
}

export const mockCreateUserNotification = (resCfg: MockCreateUserNotificationResponseConfig) => {
  return (page: Page) =>
    page.route(
      `**/api/v1/user_notifications`,
      (route: Route) => {
        const req = route.request();

        assertMethod('POST', route);
        assertClient(route);

        route.fulfill(resCfg.fulfill());
      },
      { times: resCfg.times },
    );
};
