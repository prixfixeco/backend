// GENERATED CODE, DO NOT EDIT MANUALLY

import { Axios } from 'axios';

import {
  ValidIngredientState, 
  QueryFilter,
  QueryFilteredResult,
  APIResponse, 
} from '@dinnerdonebetter/models'; 

export async function searchForValidIngredientStates(
  client: Axios,
  filter: QueryFilter = QueryFilter.Default(),
  q: string,
	): Promise< QueryFilteredResult< ValidIngredientState >> {
  return new Promise(async function (resolve, reject) {
    const response = await client.get< APIResponse < Array<ValidIngredientState>  >  >(`/api/v1/valid_ingredient_states/search`, {
      params: filter.asRecord(),
    });

    if (response.data.error) {
      reject(new Error(response.data.error.message));
    }

    const result = new QueryFilteredResult<ValidIngredientState>({
      data: response.data.data,
      totalCount: response.data.pagination?.totalCount,
      page: response.data.pagination?.page,
      limit: response.data.pagination?.limit,
    });

    resolve(result);
  });
}