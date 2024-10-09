// GENERATED CODE, DO NOT EDIT MANUALLY

import { Axios } from 'axios';

import {
  ValidMeasurementUnitConversion, 
  APIResponse, 
  ValidMeasurementUnitConversionCreationRequestInput, 
} from '@dinnerdonebetter/models';

export async function createValidMeasurementUnitConversion(
  client: Axios,
  
  input: ValidMeasurementUnitConversionCreationRequestInput,
): Promise<  APIResponse <  ValidMeasurementUnitConversion >  >  {
  return new Promise(async function (resolve, reject) {
    const response = await client.post<APIResponse < ValidMeasurementUnitConversion  >  >(`/api/v1/valid_measurement_conversions`, input);
	    if (response.data.error) {
	      reject(new Error(response.data.error.message));
	    }

	    resolve(response.data);
	  });
}