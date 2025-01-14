import { GetServerSideProps, GetServerSidePropsContext, GetServerSidePropsResult } from 'next';
import Link from 'next/link';
import { Container, Divider, Grid, List, NumberInput, Space, Text, Title } from '@mantine/core';
import { ReactNode, useState } from 'react';

import {
  ALL_MEAL_COMPONENT_TYPE,
  APIResponse,
  EitherErrorOr,
  IAPIError,
  Meal,
  MealComponent,
} from '@dinnerdonebetter/models';
import { determineAllIngredientsForRecipes, determineAllInstrumentsForRecipes } from '@dinnerdonebetter/utils';
import { ServerTimingHeaderName, ServerTiming } from '@dinnerdonebetter/server-timing';

import { buildServerSideClientOrRedirect } from '../../../src/client';
import { AppLayout } from '../../../src/layouts';
import { serverSideTracer } from '../../../src/tracer';
import { serverSideAnalytics } from '../../../src/analytics';
import { userSessionDetailsOrRedirect } from '../../../src/auth';
import { RecipeInstrumentListComponent } from '../../../src/components';
import { RecipeIngredientListComponent } from '../../../src/components/IngredientList';
import { valueOrDefault } from '../../../src/utils';

declare interface MealPageProps {
  meal: EitherErrorOr<Meal>;
}

export const getServerSideProps: GetServerSideProps = async (
  context: GetServerSidePropsContext,
): Promise<GetServerSidePropsResult<MealPageProps>> => {
  const timing = new ServerTiming();
  const span = serverSideTracer.startSpan('MealPage.getServerSideProps');

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

  const { mealID } = context.query;
  if (!mealID) {
    throw new Error('meal ID is somehow missing!');
  }

  const extractCookieTimer = timing.addEvent('extract cookie');
  const sessionDetails = userSessionDetailsOrRedirect(context.req.cookies);
  if (sessionDetails.redirect) {
    span.end();
    return { redirect: sessionDetails.redirect };
  }
  const userSessionData = sessionDetails.details;
  extractCookieTimer.end();

  if (userSessionData?.userID) {
    const analyticsTimer = timing.addEvent('analytics');
    serverSideAnalytics.page(userSessionData.userID, 'MEAL_PAGE', context, {
      mealID,
      householdID: userSessionData.householdID,
    });
    analyticsTimer.end();
  }

  const fetchRecipeTimer = timing.addEvent('fetch meal');
  const mealPromise = apiClient
    .getMeal(mealID.toString())
    .then((result: APIResponse<Meal>) => {
      span.addEvent(`recipe retrieved`);
      return { data: result.data };
    })
    .catch((error: IAPIError) => {
      span.addEvent('error occurred', { error: error.message });
      return { error };
    })
    .finally(() => {
      fetchRecipeTimer.end();
    });

  const retrievedData = await Promise.all([mealPromise]);
  const [meal] = retrievedData;

  context.res.setHeader(ServerTimingHeaderName, timing.headerValue());

  span.end();
  return {
    props: {
      meal,
    },
  };
};

// https://stackoverflow.com/a/14872766
const ordering: Record<string, number> = {};
for (let i = 0; i < ALL_MEAL_COMPONENT_TYPE.length; i++) {
  ordering[ALL_MEAL_COMPONENT_TYPE[i]] = i;
}

const formatRecipeList = (meal: Meal): ReactNode => {
  const sorted = (meal.components || []).sort(function (a: MealComponent, b: MealComponent) {
    return ordering[a.componentType] - ordering[b.componentType] || a.componentType.localeCompare(b.componentType);
  });

  return sorted.map((c: MealComponent, index: number) => {
    return (
      <List.Item key={index} mb="md">
        <Link href={`/recipes/${c.recipe.id}`}>{c.recipe.name}</Link>
        <em>{c.recipe.description}</em>
      </List.Item>
    );
  });
};

function MealPage(props: MealPageProps) {
  const ogMeal = valueOrDefault(props.meal, new Meal());
  const [mealError] = useState<IAPIError | undefined>(props.meal.error);
  const [meal] = useState(ogMeal);

  const [mealScale, setMealScale] = useState(1.0);

  return (
    <AppLayout title={meal.name} userLoggedIn>
      <Container size="xs">
        {mealError && <Text color="tomato">{mealError.message}</Text>}

        {!mealError && meal.id !== '' && (
          <>
            <Title order={3}>{meal.name}</Title>

            <NumberInput
              mt="sm"
              mb="lg"
              value={mealScale}
              precision={2}
              step={0.25}
              removeTrailingZeros={true}
              description={`this meal normally yields about ${meal.estimatedPortions.min} ${
                meal.estimatedPortions.min === 1 ? 'portion' : 'portions'
              }${
                mealScale === 1.0
                  ? ''
                  : `, but is now set up to yield ${meal.estimatedPortions.min * mealScale}  ${
                      meal.estimatedPortions.min === 1 ? 'portion' : 'portions'
                    }`
              }`}
              onChange={(value: number | undefined) => {
                if (!value) return;

                setMealScale(value);
              }}
            />

            <Divider label="recipes" labelPosition="center" mb="xs" />

            <Grid grow gutter="md">
              <List mt="sm">{formatRecipeList(meal)}</List>
            </Grid>

            <Divider label="resources" labelPosition="center" mb="md" />

            <Grid>
              <Grid.Col span={6}>
                <Title order={6}>Tools:</Title>
                {determineAllInstrumentsForRecipes((meal.components || []).map((x) => x.recipe)).length > 0 && (
                  <RecipeInstrumentListComponent recipes={(meal.components || []).map((x) => x.recipe)} />
                )}
              </Grid.Col>

              <Grid.Col span={6}>
                <Title order={6}>Ingredients:</Title>
                {determineAllIngredientsForRecipes(
                  (meal.components || []).map((x) => {
                    return { scale: mealScale, recipe: x.recipe };
                  }),
                ).length > 0 && (
                  <RecipeIngredientListComponent
                    recipes={(meal.components || []).map((x) => x.recipe)}
                    scale={mealScale}
                  />
                )}
              </Grid.Col>
            </Grid>
          </>
        )}
      </Container>
      <Space h="xl" my="xl" />
    </AppLayout>
  );
}

export default MealPage;
