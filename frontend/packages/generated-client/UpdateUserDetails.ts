// GENERATED CODE, DO NOT EDIT MANUALLY

import { Axios } from 'axios';

import { User, APIResponse, UserDetailsUpdateRequestInput } from '@dinnerdonebetter/models';

export async function updateUserDetails(
  client: Axios,

  input: UserDetailsUpdateRequestInput,
): Promise<APIResponse<User>> {
  return new Promise(async function (resolve, reject) {
    const response = await client.put<APIResponse<User>>(`/api/v1/users/details`, input);
    if (response.data.error) {
      reject(new Error(response.data.error.message));
    }

    resolve(response.data);
  });
}
