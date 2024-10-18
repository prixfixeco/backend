// GENERATED CODE, DO NOT EDIT MANUALLY

import type { Page, Route } from '@playwright/test';

import { ValidMeasurementUnitConversion } from '@dinnerdonebetter/models';

import { assertClient, assertMethod, ResponseConfig } from './helpers';

export class MockUpdateValidMeasurementUnitConversionResponseConfig extends ResponseConfig<ValidMeasurementUnitConversion> {
  validMeasurementUnitConversionID: string;

  constructor(validMeasurementUnitConversionID: string, status: number = 200, body?: ValidMeasurementUnitConversion) {
    super();

    this.validMeasurementUnitConversionID = validMeasurementUnitConversionID;

    this.status = status;
    if (this.body) {
      this.body = body;
    }
  }
}

export const mockUpdateValidMeasurementUnitConversion = (
  resCfg: MockUpdateValidMeasurementUnitConversionResponseConfig,
) => {
  return (page: Page) =>
    page.route(
      `**/api/v1/valid_measurement_conversions/${resCfg.validMeasurementUnitConversionID}`,
      (route: Route) => {
        const req = route.request();

        assertMethod('PUT', route);
        assertClient(route);

        route.fulfill(resCfg.fulfill());
      },
      { times: resCfg.times },
    );
};