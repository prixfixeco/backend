import { GetServerSideProps, GetServerSidePropsContext, GetServerSidePropsResult } from 'next';
import { useForm, zodResolver } from '@mantine/form';
import {
  ActionIcon,
  Autocomplete,
  AutocompleteItem,
  Button,
  Center,
  Container,
  Divider,
  Grid,
  Group,
  NumberInput,
  Pagination,
  Space,
  Switch,
  Table,
  Text,
  TextInput,
  Title,
} from '@mantine/core';
import { AxiosError } from 'axios';
import { z } from 'zod';
import { useEffect, useState } from 'react';
import Link from 'next/link';
import { useRouter } from 'next/router';
import { IconTrash } from '@tabler/icons';

import {
  APIResponse,
  EitherErrorOr,
  IAPIError,
  QueryFilteredResult,
  ValidIngredient,
  ValidIngredientMeasurementUnit,
  ValidIngredientMeasurementUnitCreationRequestInput,
  ValidIngredientPreparation,
  ValidIngredientPreparationCreationRequestInput,
  ValidIngredientState,
  ValidIngredientStateIngredient,
  ValidIngredientStateIngredientCreationRequestInput,
  ValidIngredientUpdateRequestInput,
  ValidMeasurementUnit,
  ValidPreparation,
} from '@dinnerdonebetter/models';
import { ServerTiming, ServerTimingHeaderName } from '@dinnerdonebetter/server-timing';
import { buildLocalClient } from '@dinnerdonebetter/api-client';

import { AppLayout } from '../../../src/layouts';
import { buildServerSideClientOrRedirect } from '../../../src/client';
import { serverSideTracer } from '../../../src/tracer';
import { inputSlug } from '../../../src/schemas';
import { valueOrDefault } from '../../../src/utils';

declare interface ValidIngredientPageProps {
  pageLoadMeasurementUnits: EitherErrorOr<QueryFilteredResult<ValidIngredientMeasurementUnit>>;
  pageLoadIngredientPreparations: EitherErrorOr<QueryFilteredResult<ValidIngredientPreparation>>;
  pageLoadValidIngredientStates: EitherErrorOr<QueryFilteredResult<ValidIngredientStateIngredient>>;
  pageLoadValidIngredient: EitherErrorOr<ValidIngredient>;
}

export const getServerSideProps: GetServerSideProps = async (
  context: GetServerSidePropsContext,
): Promise<GetServerSidePropsResult<ValidIngredientPageProps>> => {
  const timing = new ServerTiming();
  const span = serverSideTracer.startSpan('ValidIngredientPage.getServerSideProps');

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

  const { validIngredientID } = context.query;
  if (!validIngredientID) {
    throw new Error('valid ingredient ID is somehow missing!');
  }

  const fetchValidIngredientTimer = timing.addEvent('fetch valid ingredient');
  const pageLoadValidIngredientPromise = apiClient
    .getValidIngredient(validIngredientID.toString())
    .then((result: APIResponse<ValidIngredient>) => {
      span.addEvent('valid ingredient retrieved');
      return { data: result.data };
    })
    .catch((error: IAPIError) => {
      span.addEvent('error occurred', { error: error.message });
      return { error };
    })
    .finally(() => {
      fetchValidIngredientTimer.end();
    });

  const fetchMeasurementUnitsTimer = timing.addEvent('fetch valid measurement units fro ingredient');
  const pageLoadMeasurementUnitsPromise = apiClient
    .getValidIngredientMeasurementUnitsByIngredient(validIngredientID.toString())
    .then((res: QueryFilteredResult<ValidIngredientMeasurementUnit>) => {
      span.addEvent('valid ingredient measurement units retrieved');
      return { data: res };
    })
    .catch((error: IAPIError) => {
      span.addEvent('error occurred', { error: error.message });
      return { error };
    })
    .finally(() => {
      fetchMeasurementUnitsTimer.end();
    });

  const fetchIngredientPreparationsTimer = timing.addEvent('fetch valid ingredient preparations for ingredient');
  const pageLoadIngredientPreparationsPromise = apiClient
    .getValidIngredientPreparationsByIngredient(validIngredientID.toString())
    .then((res: QueryFilteredResult<ValidIngredientPreparation>) => {
      span.addEvent('valid ingredient preparations retrieved');
      return { data: res };
    })
    .catch((error: IAPIError) => {
      span.addEvent('error occurred', { error: error.message });
      return { error };
    })
    .finally(() => {
      fetchIngredientPreparationsTimer.end();
    });

  const fetchValidIngredientStatesTimer = timing.addEvent('fetch valid ingredient states for ingredient');
  const pageLoadValidIngredientStatesPromise = apiClient
    .getValidIngredientStateIngredientsByIngredient(validIngredientID.toString())
    .then((res: QueryFilteredResult<ValidIngredientStateIngredient>) => {
      span.addEvent('valid ingredient states retrieved');
      return { data: res };
    })
    .catch((error: IAPIError) => {
      span.addEvent('error occurred', { error: error.message });
      return { error };
    })
    .finally(() => {
      fetchValidIngredientStatesTimer.end();
    });

  const [
    pageLoadValidIngredient,
    pageLoadMeasurementUnits,
    pageLoadIngredientPreparations,
    pageLoadValidIngredientStates,
  ] = await Promise.all([
    pageLoadValidIngredientPromise,
    pageLoadMeasurementUnitsPromise,
    pageLoadIngredientPreparationsPromise,
    pageLoadValidIngredientStatesPromise,
  ]);

  context.res.setHeader(ServerTimingHeaderName, timing.headerValue());

  span.end();
  return {
    props: {
      pageLoadValidIngredient: JSON.parse(JSON.stringify(pageLoadValidIngredient)),
      pageLoadMeasurementUnits: JSON.parse(JSON.stringify(pageLoadMeasurementUnits)),
      pageLoadIngredientPreparations: JSON.parse(JSON.stringify(pageLoadIngredientPreparations)),
      pageLoadValidIngredientStates: JSON.parse(JSON.stringify(pageLoadValidIngredientStates)),
    },
  };
};

const validIngredientUpdateFormSchema = z.object({
  name: z.string().trim().min(1, 'name is required'),
  pluralName: z.string().trim().min(1, 'plural name is required'),
  slug: inputSlug,
});

function ValidIngredientPage(props: ValidIngredientPageProps) {
  const router = useRouter();

  const {
    pageLoadValidIngredient,
    pageLoadMeasurementUnits,
    pageLoadIngredientPreparations,
    pageLoadValidIngredientStates,
  } = props;

  const ogValidIngredient: ValidIngredient = valueOrDefault(pageLoadValidIngredient, new ValidIngredient());

  const apiClient = buildLocalClient();
  const [validIngredient, setValidIngredient] = useState<ValidIngredient>(ogValidIngredient);
  const [validIngredientError] = useState<IAPIError | undefined>(pageLoadValidIngredient.error);
  const [originalValidIngredient, setOriginalValidIngredient] = useState<ValidIngredient>(ogValidIngredient);

  const [newMeasurementUnitForIngredientInput, setNewMeasurementUnitForIngredientInput] =
    useState<ValidIngredientMeasurementUnitCreationRequestInput>(
      new ValidIngredientMeasurementUnitCreationRequestInput({
        validIngredientID: validIngredient.id,
        allowableQuantity: { min: 0.01 },
      }),
    );
  const [measurementUnitQuery, setMeasurementUnitQuery] = useState('');

  const ogValidIngredientMeasurementUnits: QueryFilteredResult<ValidIngredientMeasurementUnit> = valueOrDefault(
    pageLoadMeasurementUnits,
    new QueryFilteredResult<ValidIngredientMeasurementUnit>(),
  );

  const [measurementUnitsForIngredient, setMeasurementUnitsForIngredient] = useState<
    QueryFilteredResult<ValidIngredientMeasurementUnit>
  >(ogValidIngredientMeasurementUnits);
  const [measurementUnitsForIngredientError] = useState<IAPIError | undefined>(pageLoadMeasurementUnits.error);
  const [suggestedMeasurementUnits, setSuggestedMeasurementUnits] = useState<ValidMeasurementUnit[]>([]);

  useEffect(() => {
    if (measurementUnitQuery.length <= 2) {
      setSuggestedMeasurementUnits([]);
      return;
    }

    const apiClient = buildLocalClient();
    apiClient
      .searchForValidMeasurementUnits(measurementUnitQuery)
      .then((res: QueryFilteredResult<ValidMeasurementUnit>) => {
        const newSuggestions = (res.data || []).filter((mu: ValidMeasurementUnit) => {
          return !(measurementUnitsForIngredient.data || []).some((vimu: ValidIngredientMeasurementUnit) => {
            return vimu.measurementUnit.id === mu.id;
          });
        });

        setSuggestedMeasurementUnits(newSuggestions);
      })
      .catch((err: AxiosError) => {
        console.error(err);
      });
  }, [measurementUnitQuery, measurementUnitsForIngredient.data]);

  const [newPreparationForIngredientInput, setNewPreparationForIngredientInput] =
    useState<ValidIngredientPreparationCreationRequestInput>(
      new ValidIngredientPreparationCreationRequestInput({
        validIngredientID: validIngredient.id,
      }),
    );
  const [preparationQuery, setPreparationQuery] = useState('');

  const ogValidIngredientPreparations: QueryFilteredResult<ValidIngredientPreparation> = valueOrDefault(
    pageLoadIngredientPreparations,
    new QueryFilteredResult<ValidIngredientPreparation>(),
  );

  const [preparationsForIngredient, setPreparationsForIngredient] =
    useState<QueryFilteredResult<ValidIngredientPreparation>>(ogValidIngredientPreparations);

  const [preparationsForIngredientError] = useState<IAPIError | undefined>(pageLoadIngredientPreparations.error);
  const [suggestedPreparations, setSuggestedPreparations] = useState<Array<ValidPreparation>>([
    new ValidPreparation({ name: 'blah' }),
  ]);

  useEffect(() => {
    if (preparationQuery.length <= 2) {
      setSuggestedPreparations([]);
      return;
    }

    const apiClient = buildLocalClient();
    apiClient
      .searchForValidPreparations(preparationQuery)
      .then((res: QueryFilteredResult<ValidPreparation>) => {
        const newSuggestions = (res.data || []).filter((mu: ValidPreparation) => {
          return !(preparationsForIngredient.data || []).some((vimu: ValidIngredientPreparation) => {
            return vimu.preparation.id === mu.id;
          });
        });

        setSuggestedPreparations(newSuggestions);
      })
      .catch((err: AxiosError) => {
        console.error(err);
      });
  }, [preparationQuery, preparationsForIngredient.data]);

  const [newIngredientStateForIngredientInput, setNewIngredientStateForIngredientInput] =
    useState<ValidIngredientStateIngredientCreationRequestInput>(
      new ValidIngredientStateIngredientCreationRequestInput({
        validIngredientID: validIngredient.id,
      }),
    );
  const [ingredientStateQuery, setIngredientStateQuery] = useState('');

  const ogValidIngredientStates: QueryFilteredResult<ValidIngredientStateIngredient> = valueOrDefault(
    pageLoadValidIngredientStates,
    new QueryFilteredResult<ValidIngredientStateIngredient>(),
  );

  const [ingredientStatesForIngredientError] = useState<IAPIError | undefined>(pageLoadValidIngredientStates.error);
  const [ingredientStatesForIngredient, setIngredientStatesForIngredient] =
    useState<QueryFilteredResult<ValidIngredientStateIngredient>>(ogValidIngredientStates);
  const [suggestedIngredientStates, setSuggestedIngredientStates] = useState<ValidIngredientState[]>([]);

  useEffect(() => {
    if (ingredientStateQuery.length <= 2) {
      setSuggestedIngredientStates([]);
      return;
    }

    const apiClient = buildLocalClient();
    apiClient
      .searchForValidIngredientStates(ingredientStateQuery)
      .then((res: QueryFilteredResult<ValidIngredientState>) => {
        const newSuggestions = res.data.filter((mu: ValidIngredientState) => {
          return !(ingredientStatesForIngredient.data || []).some((vimu: ValidIngredientStateIngredient) => {
            return vimu.ingredientState.id === mu.id;
          });
        });

        setSuggestedIngredientStates(newSuggestions);
      })
      .catch((err: AxiosError) => {
        console.error(err);
      });
  }, [ingredientStateQuery, ingredientStatesForIngredient.data]);

  const updateForm = useForm({
    initialValues: validIngredient,
    validate: zodResolver(validIngredientUpdateFormSchema),
  });

  const dataHasChanged = (): boolean => {
    return (
      originalValidIngredient.name !== updateForm.values.name ||
      originalValidIngredient.pluralName !== updateForm.values.pluralName ||
      originalValidIngredient.description !== updateForm.values.description ||
      originalValidIngredient.warning !== updateForm.values.warning ||
      originalValidIngredient.iconPath !== updateForm.values.iconPath ||
      originalValidIngredient.containsDairy !== updateForm.values.containsDairy ||
      originalValidIngredient.containsPeanut !== updateForm.values.containsPeanut ||
      originalValidIngredient.containsTreeNut !== updateForm.values.containsTreeNut ||
      originalValidIngredient.containsEgg !== updateForm.values.containsEgg ||
      originalValidIngredient.containsWheat !== updateForm.values.containsWheat ||
      originalValidIngredient.containsShellfish !== updateForm.values.containsShellfish ||
      originalValidIngredient.containsSesame !== updateForm.values.containsSesame ||
      originalValidIngredient.containsFish !== updateForm.values.containsFish ||
      originalValidIngredient.containsGluten !== updateForm.values.containsGluten ||
      originalValidIngredient.animalFlesh !== updateForm.values.animalFlesh ||
      originalValidIngredient.isLiquid !== updateForm.values.isLiquid ||
      originalValidIngredient.containsSoy !== updateForm.values.containsSoy ||
      originalValidIngredient.animalDerived !== updateForm.values.animalDerived ||
      originalValidIngredient.restrictToPreparations !== updateForm.values.restrictToPreparations ||
      originalValidIngredient.containsAlcohol !== updateForm.values.containsAlcohol ||
      originalValidIngredient.storageTemperatureInCelsius.min !== updateForm.values.storageTemperatureInCelsius.min ||
      originalValidIngredient.storageTemperatureInCelsius.max !== updateForm.values.storageTemperatureInCelsius.max ||
      originalValidIngredient.slug !== updateForm.values.slug ||
      originalValidIngredient.shoppingSuggestions !== updateForm.values.shoppingSuggestions
    );
  };

  const submit = async () => {
    const validation = updateForm.validate();
    if (validation.hasErrors) {
      console.error(validation.errors);
      return;
    }

    const submission = new ValidIngredientUpdateRequestInput({
      name: updateForm.values.name,
      pluralName: updateForm.values.pluralName,
      description: updateForm.values.description,
      warning: updateForm.values.warning,
      iconPath: updateForm.values.iconPath,
      containsDairy: updateForm.values.containsDairy,
      containsPeanut: updateForm.values.containsPeanut,
      containsTreeNut: updateForm.values.containsTreeNut,
      containsEgg: updateForm.values.containsEgg,
      containsWheat: updateForm.values.containsWheat,
      containsShellfish: updateForm.values.containsShellfish,
      containsSesame: updateForm.values.containsSesame,
      containsFish: updateForm.values.containsFish,
      containsGluten: updateForm.values.containsGluten,
      animalFlesh: updateForm.values.animalFlesh,
      isLiquid: updateForm.values.isLiquid,
      containsSoy: updateForm.values.containsSoy,
      animalDerived: updateForm.values.animalDerived,
      restrictToPreparations: updateForm.values.restrictToPreparations,
      containsAlcohol: updateForm.values.containsAlcohol,
      storageTemperatureInCelsius: updateForm.values.storageTemperatureInCelsius,
      slug: updateForm.values.slug,
      shoppingSuggestions: updateForm.values.shoppingSuggestions,
    });

    await apiClient
      .updateValidIngredient(validIngredient.id, submission)
      .then((result: APIResponse<ValidIngredient>) => {
        updateForm.setValues(result.data);
        setValidIngredient(result.data);
        setOriginalValidIngredient(result.data);
      })
      .catch((err) => {
        console.error(err);
      });
  };

  return (
    <AppLayout title="Valid Ingredient">
      <Container size="sm">
        {validIngredientError && (
          <Text color="tomato">Error fetching valid ingredient: {validIngredientError.message}</Text>
        )}

        {!validIngredientError && (
          <form onSubmit={updateForm.onSubmit(submit)}>
            <TextInput label="Name" placeholder="thing" {...updateForm.getInputProps('name')} />
            <TextInput label="Slug" placeholder="thing" {...updateForm.getInputProps('slug')} />
            <TextInput label="Plural Name" placeholder="things" {...updateForm.getInputProps('pluralName')} />
            <TextInput
              label="Description"
              placeholder="stuff about things"
              {...updateForm.getInputProps('description')}
            />
            <TextInput label="Warning" placeholder="warning" {...updateForm.getInputProps('warning')} />
            <NumberInput
              label="Min Storage Temp (C°)"
              {...updateForm.getInputProps('storageTemperatureInCelsius.min')}
            />
            <NumberInput
              label="Max Storage Temp (C°)"
              {...updateForm.getInputProps('storageTemperatureInCelsius.max')}
            />

            <Switch
              checked={updateForm.values.containsDairy}
              label="Contains Dairy"
              {...updateForm.getInputProps('containsDairy')}
            />
            <Switch
              checked={updateForm.values.containsPeanut}
              label="Contains Peanut"
              {...updateForm.getInputProps('containsPeanut')}
            />
            <Switch
              checked={updateForm.values.containsTreeNut}
              label="Contains Tree Nut"
              {...updateForm.getInputProps('containsTreeNut')}
            />
            <Switch
              checked={updateForm.values.containsEgg}
              label="Contains Egg"
              {...updateForm.getInputProps('containsEgg')}
            />
            <Switch
              checked={updateForm.values.containsWheat}
              label="Contains Wheat"
              {...updateForm.getInputProps('containsWheat')}
            />
            <Switch
              checked={updateForm.values.containsShellfish}
              label="Contains Shellfish"
              {...updateForm.getInputProps('containsShellfish')}
            />
            <Switch
              checked={updateForm.values.containsSesame}
              label="Contains Sesame"
              {...updateForm.getInputProps('containsSesame')}
            />
            <Switch
              checked={updateForm.values.containsFish}
              label="Contains Fish"
              {...updateForm.getInputProps('containsFish')}
            />
            <Switch
              checked={updateForm.values.containsGluten}
              label="Contains Gluten"
              {...updateForm.getInputProps('containsGluten')}
            />
            <Switch
              checked={updateForm.values.containsSoy}
              label="Contains Soy"
              {...updateForm.getInputProps('containsSoy')}
            />
            <Switch
              checked={updateForm.values.containsAlcohol}
              label="Contains Alcohol"
              {...updateForm.getInputProps('containsAlcohol')}
            />
            <Switch
              checked={updateForm.values.animalFlesh}
              label="Animal Flesh"
              {...updateForm.getInputProps('animalFlesh')}
            />
            <Switch
              checked={updateForm.values.animalDerived}
              label="Animal Derived"
              {...updateForm.getInputProps('animalDerived')}
            />
            <Switch checked={updateForm.values.isLiquid} label="Liquid" {...updateForm.getInputProps('isLiquid')} />
            <Switch
              checked={updateForm.values.restrictToPreparations}
              label="Restrict To Preparations"
              {...updateForm.getInputProps('restrictToPreparations')}
            />

            <Group position="center">
              <Button type="submit" mt="sm" fullWidth disabled={!dataHasChanged()}>
                Submit
              </Button>
              <Button
                type="submit"
                color="red"
                fullWidth
                onClick={() => {
                  if (confirm('Are you sure you want to delete this valid ingredient?')) {
                    apiClient.archiveValidIngredient(validIngredient.id).then(() => {
                      router.push('/valid_ingredients');
                    });
                  }
                }}
              >
                Delete
              </Button>
            </Group>
          </form>
        )}

        {/*

        INGREDIENT MEASUREMENT UNITS

        */}

        <Space h="xl" />
        <Divider />
        <Space h="xl" />

        <form>
          <Center>
            <Title order={4}>Measurement Units</Title>
          </Center>

          {measurementUnitsForIngredientError && (
            <Text color="tomato">Error fetching measurement units: {measurementUnitsForIngredientError.message}</Text>
          )}

          {!measurementUnitsForIngredientError &&
            measurementUnitsForIngredient.data &&
            measurementUnitsForIngredient.data.length !== 0 && (
              <>
                <Table mt="xl" withColumnBorders>
                  <thead>
                    <tr>
                      <th>Min Allowed</th>
                      <th>Max Allowed</th>
                      <th></th>
                    </tr>
                  </thead>
                  <tbody>
                    {measurementUnitsForIngredient.data.map((measurementUnit: ValidIngredientMeasurementUnit) => {
                      return (
                        <tr key={measurementUnit.id}>
                          <td>
                            {measurementUnit.allowableQuantity.min}{' '}
                            <Link href={`/valid_measurement_units/${measurementUnit.id}`}>
                              {measurementUnit.allowableQuantity.min === 1
                                ? measurementUnit.measurementUnit.name
                                : measurementUnit.measurementUnit.pluralName}
                            </Link>
                          </td>
                          <td>
                            {measurementUnit.allowableQuantity.max}{' '}
                            <Link href={`/valid_measurement_units/${measurementUnit.id}`}>
                              {measurementUnit.allowableQuantity.min === 1
                                ? measurementUnit.measurementUnit.name
                                : measurementUnit.measurementUnit.pluralName}
                            </Link>
                          </td>
                          <td>
                            <Center>
                              <ActionIcon
                                variant="outline"
                                aria-label="remove valid ingredient measurement unit"
                                onClick={async () => {
                                  await apiClient
                                    .archiveValidIngredientMeasurementUnit(measurementUnit.id)
                                    .then(() => {
                                      setMeasurementUnitsForIngredient({
                                        ...measurementUnitsForIngredient,
                                        data: measurementUnitsForIngredient.data.filter(
                                          (x: ValidIngredientMeasurementUnit) => x.id !== measurementUnit.id,
                                        ),
                                      });
                                    })
                                    .catch((error) => {
                                      console.error(error);
                                    });
                                }}
                              >
                                <IconTrash size="md" color="tomato" />
                              </ActionIcon>
                            </Center>
                          </td>
                        </tr>
                      );
                    })}
                  </tbody>
                </Table>

                <Space h="xs" />

                <Pagination
                  disabled={
                    Math.ceil(measurementUnitsForIngredient.totalCount / measurementUnitsForIngredient.limit) <=
                    measurementUnitsForIngredient.page
                  }
                  position="center"
                  page={measurementUnitsForIngredient.page}
                  total={Math.ceil(measurementUnitsForIngredient.totalCount / measurementUnitsForIngredient.limit)}
                  onChange={(value: number) => {
                    setMeasurementUnitsForIngredient({ ...measurementUnitsForIngredient, page: value });
                  }}
                />
              </>
            )}

          <Grid>
            <Grid.Col span="auto">
              <Autocomplete
                placeholder="grams"
                label="Measurement Unit"
                value={measurementUnitQuery}
                onChange={setMeasurementUnitQuery}
                onItemSubmit={async (item: AutocompleteItem) => {
                  const selectedMeasurementUnit = suggestedMeasurementUnits.find(
                    (x: ValidMeasurementUnit) => x.pluralName === item.value,
                  );

                  if (!selectedMeasurementUnit) {
                    console.error(`selectedMeasurementUnit not found for item ${item.value}}`);
                    return;
                  }

                  setNewMeasurementUnitForIngredientInput({
                    ...newMeasurementUnitForIngredientInput,
                    validMeasurementUnitID: selectedMeasurementUnit.id,
                  });
                }}
                data={suggestedMeasurementUnits.map((x: ValidMeasurementUnit) => {
                  return { value: x.pluralName, label: x.pluralName };
                })}
              />
            </Grid.Col>
            <Grid.Col span={2}>
              <NumberInput
                value={newMeasurementUnitForIngredientInput.allowableQuantity.min}
                label="Min. Qty"
                precision={2}
                onChange={(value: number) =>
                  setNewMeasurementUnitForIngredientInput({
                    ...newMeasurementUnitForIngredientInput,
                    allowableQuantity: {
                      ...newMeasurementUnitForIngredientInput.allowableQuantity,
                      min: value,
                    },
                  })
                }
              />
            </Grid.Col>
            <Grid.Col span={2}>
              <NumberInput
                value={newMeasurementUnitForIngredientInput.allowableQuantity.max}
                label="Max. Qty"
                precision={2}
                onChange={(value: number) =>
                  setNewMeasurementUnitForIngredientInput({
                    ...newMeasurementUnitForIngredientInput,
                    allowableQuantity: {
                      ...newMeasurementUnitForIngredientInput.allowableQuantity,
                      max: value,
                    },
                  })
                }
              />
            </Grid.Col>
            <Grid.Col span="auto">
              <TextInput
                label="Notes"
                value={newMeasurementUnitForIngredientInput.notes}
                onChange={(event) =>
                  setNewMeasurementUnitForIngredientInput({
                    ...newMeasurementUnitForIngredientInput,
                    notes: event.target.value,
                  })
                }
              />
            </Grid.Col>
            <Grid.Col span={2}>
              <Button
                mt="xl"
                disabled={
                  newMeasurementUnitForIngredientInput.validIngredientID === '' ||
                  newMeasurementUnitForIngredientInput.validMeasurementUnitID === '' ||
                  newMeasurementUnitForIngredientInput.allowableQuantity.min === 0
                }
                onClick={async () => {
                  await apiClient
                    .createValidIngredientMeasurementUnit(newMeasurementUnitForIngredientInput)
                    .then((res: APIResponse<ValidIngredientMeasurementUnit>) => {
                      // the returned value doesn't have enough information to put it in the list, so we have to fetch it
                      apiClient
                        .getValidIngredientMeasurementUnit(res.data.id)
                        .then((res: APIResponse<ValidIngredientMeasurementUnit>) => {
                          setMeasurementUnitsForIngredient({
                            ...measurementUnitsForIngredient,
                            data: [...measurementUnitsForIngredient.data, res.data],
                          });

                          setNewMeasurementUnitForIngredientInput(
                            new ValidIngredientMeasurementUnitCreationRequestInput({
                              validIngredientID: validIngredient.id,
                              validMeasurementUnitID: '',
                              allowableQuantity: {
                                ...newMeasurementUnitForIngredientInput.allowableQuantity,
                                min: 0.01,
                              },
                            }),
                          );

                          setMeasurementUnitQuery('');
                        });
                    })
                    .catch((error) => {
                      console.error(error);
                    });
                }}
              >
                Save
              </Button>
            </Grid.Col>
          </Grid>
        </form>

        {/*

        INGREDIENT PREPARATIONS

        */}

        <Space h="xl" />
        <Divider />
        <Space h="xl" />

        <form>
          <Center>
            <Title order={4}>Preparations</Title>
          </Center>

          {preparationsForIngredientError && (
            <Text color="tomato">Error fetching preparations: {preparationsForIngredientError.message}</Text>
          )}

          {(preparationsForIngredient.data || []).length !== 0 && (
            <>
              <Table mt="xl" withColumnBorders>
                <thead>
                  <tr>
                    <th>Name</th>
                    <th>Notes</th>
                    <th />
                  </tr>
                </thead>
                <tbody>
                  {(preparationsForIngredient.data || []).map(
                    (validIngredientPreparation: ValidIngredientPreparation) => {
                      return (
                        <tr key={validIngredientPreparation.id}>
                          <td>{validIngredientPreparation.preparation.name}</td>
                          <td>{validIngredientPreparation.notes}</td>
                          <td>
                            <Center>
                              <ActionIcon
                                variant="outline"
                                aria-label="remove valid ingredient preparation"
                                onClick={async () => {
                                  await apiClient
                                    .archiveValidIngredientPreparation(validIngredientPreparation.id)
                                    .then(() => {
                                      setPreparationsForIngredient({
                                        ...preparationsForIngredient,
                                        data: (preparationsForIngredient.data || []).filter(
                                          (x: ValidIngredientPreparation) => x.id !== validIngredientPreparation.id,
                                        ),
                                      });
                                    })
                                    .catch((error) => {
                                      console.error(error);
                                    });
                                }}
                              >
                                <IconTrash size="md" color="tomato" />
                              </ActionIcon>
                            </Center>
                          </td>
                        </tr>
                      );
                    },
                  )}
                </tbody>
              </Table>

              <Space h="xs" />

              <Pagination
                disabled={
                  Math.ceil(preparationsForIngredient.totalCount / preparationsForIngredient.limit) <=
                  preparationsForIngredient.page
                }
                position="center"
                page={preparationsForIngredient.page}
                total={Math.ceil(preparationsForIngredient.totalCount / preparationsForIngredient.limit)}
                onChange={(value: number) => {
                  setPreparationsForIngredient({ ...preparationsForIngredient, page: value });
                }}
              />
            </>
          )}

          <Grid>
            <Grid.Col span="auto">
              <Autocomplete
                placeholder="mince"
                label="Preparation"
                value={preparationQuery}
                onChange={setPreparationQuery}
                onItemSubmit={async (item: AutocompleteItem) => {
                  const selectedValidIngredientPreparation = suggestedPreparations.find(
                    (x: ValidPreparation) => x.name === item.value,
                  );

                  if (!selectedValidIngredientPreparation) {
                    console.error(`selectedValidIngredientPreparation not found for item ${item.value}}`);
                    return;
                  }

                  setNewPreparationForIngredientInput({
                    ...newPreparationForIngredientInput,
                    validPreparationID: selectedValidIngredientPreparation.id,
                  });
                }}
                data={suggestedPreparations.map((x: ValidPreparation) => {
                  return { value: x.name, label: x.name };
                })}
              />
            </Grid.Col>
            <Grid.Col span="content">
              <TextInput
                label="Notes"
                value={newPreparationForIngredientInput.notes}
                onChange={(event) =>
                  setNewPreparationForIngredientInput({
                    ...newPreparationForIngredientInput,
                    notes: event.target.value,
                  })
                }
              />
            </Grid.Col>
            <Grid.Col span="auto">
              <Button
                mt="xl"
                disabled={
                  newPreparationForIngredientInput.validIngredientID === '' ||
                  newPreparationForIngredientInput.validPreparationID === ''
                }
                onClick={() => {
                  apiClient
                    .createValidIngredientPreparation(newPreparationForIngredientInput)
                    .then((res: APIResponse<ValidIngredientPreparation>) => {
                      // the returned value doesn't have enough information to put it in the list, so we have to fetch it
                      apiClient
                        .getValidIngredientPreparation(res.data.id)
                        .then((res: APIResponse<ValidIngredientPreparation>) => {
                          setPreparationsForIngredient({
                            ...preparationsForIngredient,
                            data: [...(preparationsForIngredient.data || []), res.data],
                          });

                          setPreparationQuery('');
                          setNewPreparationForIngredientInput(
                            new ValidIngredientPreparationCreationRequestInput({
                              validIngredientID: validIngredient.id,
                              validPreparationID: '',
                              notes: '',
                            }),
                          );
                        });
                    })
                    .catch((error) => {
                      console.error(error);
                    });
                }}
              >
                Save
              </Button>
            </Grid.Col>
          </Grid>
        </form>

        {/*

        INGREDIENT STATES

        */}

        <Space h="xl" />
        <Divider />
        <Space h="xl" />

        <form>
          <Center>
            <Title order={4}>States</Title>
          </Center>

          {ingredientStatesForIngredientError && (
            <Text color="tomato">Error fetching ingredient states: {ingredientStatesForIngredientError.message}</Text>
          )}

          {ingredientStatesForIngredient.data && ingredientStatesForIngredient.data.length !== 0 && (
            <>
              <Table mt="xl" withColumnBorders>
                <thead>
                  <tr>
                    <th>Name</th>
                    <th />
                  </tr>
                </thead>
                <tbody>
                  {(ingredientStatesForIngredient.data || []).map(
                    (validIngredientStateIngredient: ValidIngredientStateIngredient) => {
                      return (
                        <tr key={validIngredientStateIngredient.id}>
                          <td>{validIngredientStateIngredient.ingredientState.name}</td>
                          <td>
                            <Center>
                              <ActionIcon
                                variant="outline"
                                aria-label="remove valid ingredient ingredientState"
                                onClick={async () => {
                                  await apiClient
                                    .archiveValidIngredientStateIngredient(validIngredientStateIngredient.id)
                                    .then(() => {
                                      setIngredientStatesForIngredient({
                                        ...ingredientStatesForIngredient,
                                        data: ingredientStatesForIngredient.data.filter(
                                          (x: ValidIngredientStateIngredient) =>
                                            x.id !== validIngredientStateIngredient.id,
                                        ),
                                      });
                                    })
                                    .catch((error) => {
                                      console.error(error);
                                    });
                                }}
                              >
                                <IconTrash size="md" color="tomato" />
                              </ActionIcon>
                            </Center>
                          </td>
                        </tr>
                      );
                    },
                  )}
                </tbody>
              </Table>

              <Space h="xs" />

              <Pagination
                disabled={
                  Math.ceil(ingredientStatesForIngredient.totalCount / ingredientStatesForIngredient.limit) <=
                  ingredientStatesForIngredient.page
                }
                position="center"
                page={ingredientStatesForIngredient.page}
                total={Math.ceil(ingredientStatesForIngredient.totalCount / ingredientStatesForIngredient.limit)}
                onChange={(value: number) => {
                  setIngredientStatesForIngredient({ ...ingredientStatesForIngredient, page: value });
                }}
              />
            </>
          )}

          <Grid>
            <Grid.Col span="auto">
              <Autocomplete
                placeholder="fragrant"
                label="Ingredient State"
                value={ingredientStateQuery}
                onChange={setIngredientStateQuery}
                onItemSubmit={async (item: AutocompleteItem) => {
                  const selectedValidIngredientStateIngredient = suggestedIngredientStates.find(
                    (x: ValidIngredientState) => x.name === item.value,
                  );

                  if (!selectedValidIngredientStateIngredient) {
                    console.error(`selectedValidIngredientStateIngredient not found for item ${item.value}}`);
                    return;
                  }

                  setNewIngredientStateForIngredientInput({
                    ...newIngredientStateForIngredientInput,
                    validIngredientStateID: selectedValidIngredientStateIngredient.id,
                  });
                }}
                data={suggestedIngredientStates.map((x: ValidIngredientState) => {
                  return { value: x.name, label: x.name };
                })}
              />
            </Grid.Col>
            <Grid.Col span="content">
              <TextInput
                label="Notes"
                value={newIngredientStateForIngredientInput.notes}
                onChange={(event) =>
                  setNewIngredientStateForIngredientInput({
                    ...newIngredientStateForIngredientInput,
                    notes: event.target.value,
                  })
                }
              />
            </Grid.Col>
            <Grid.Col span="auto">
              <Button
                mt="xl"
                disabled={
                  newIngredientStateForIngredientInput.validIngredientID === '' ||
                  newIngredientStateForIngredientInput.validIngredientStateID === ''
                }
                onClick={() => {
                  apiClient
                    .createValidIngredientStateIngredient(newIngredientStateForIngredientInput)
                    .then((res: APIResponse<ValidIngredientStateIngredient>) => {
                      // the returned value doesn't have enough information to put it in the list, so we have to fetch it
                      apiClient
                        .getValidIngredientStateIngredient(res.data.id)
                        .then((res: APIResponse<ValidIngredientStateIngredient>) => {
                          setIngredientStatesForIngredient({
                            ...ingredientStatesForIngredient,
                            data: [...ingredientStatesForIngredient.data, res.data],
                          });

                          setPreparationQuery('');
                          setNewPreparationForIngredientInput(
                            new ValidIngredientPreparationCreationRequestInput({
                              validIngredientID: validIngredient.id,
                              validPreparationID: '',
                              notes: '',
                            }),
                          );
                        });
                    })
                    .catch((error) => {
                      console.error(error);
                    });
                }}
              >
                Save
              </Button>
            </Grid.Col>
          </Grid>
        </form>
      </Container>
    </AppLayout>
  );
}

export default ValidIngredientPage;
