// GENERATED CODE, DO NOT EDIT MANUALLY

import { Axios } from 'axios';

import { Household, APIResponse, HouseholdCreationRequestInput } from '@dinnerdonebetter/models';

export async function createHousehold(
  client: Axios,

  input: HouseholdCreationRequestInput,
): Promise<APIResponse<Household>> {
  return new Promise(async function (resolve, reject) {
    const response = await client.post<APIResponse<Household>>(`/api/v1/households`, input);
    if (response.data.error) {
      reject(new Error(response.data.error.message));
    }

    resolve(response.data);
  });
}
