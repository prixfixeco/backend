// GENERATED CODE, DO NOT EDIT MANUALLY

import { Axios } from 'axios';

import {
  User, 
  APIResponse, 
} from '@dinnerdonebetter/models'; 

export async function archiveUser(
  client: Axios,
  userID: string,
	): Promise<  APIResponse <  User >    >   {
  return new Promise(async function (resolve, reject) {
    const response = await client.delete< APIResponse < User  >  >(`/api/v1/users/${userID}`, {});

    if (response.data.error) {
      reject(new Error(response.data.error.message));
    }

    resolve(response.data);
  });
}