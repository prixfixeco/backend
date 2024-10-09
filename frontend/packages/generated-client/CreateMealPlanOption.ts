// GENERATED CODE, DO NOT EDIT MANUALLY

import { Axios } from 'axios';

import {
  MealPlanOption, 
  APIResponse, 
  MealPlanOptionCreationRequestInput, 
} from '@dinnerdonebetter/models';

export async function createMealPlanOption(
  client: Axios,
  mealPlanID: string,mealPlanEventID: string,
  input: MealPlanOptionCreationRequestInput,
): Promise<  APIResponse <  MealPlanOption >  >  {
  return new Promise(async function (resolve, reject) {
    const response = await client.post<APIResponse < MealPlanOption  >  >(`/api/v1/meal_plans/${mealPlanID}/events/${mealPlanEventID}/options`, input);
	    if (response.data.error) {
	      reject(new Error(response.data.error.message));
	    }

	    resolve(response.data);
	  });
}