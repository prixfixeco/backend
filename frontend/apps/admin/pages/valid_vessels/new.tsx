import { useRouter } from 'next/router';
import { useEffect, useState } from 'react';
import { useForm, zodResolver } from '@mantine/form';
import {
  Autocomplete,
  AutocompleteItem,
  Button,
  Container,
  Group,
  NumberInput,
  Switch,
  TextInput,
} from '@mantine/core';
import { z } from 'zod';
import { AxiosError } from 'axios';

import {
  APIResponse,
  QueryFilteredResult,
  ValidMeasurementUnit,
  ValidVessel,
  ValidVesselCreationRequestInput,
} from '@dinnerdonebetter/models';
import { buildLocalClient } from '@dinnerdonebetter/api-client';

import { AppLayout } from '../../src/layouts';
import { inputSlug } from '../../src/schemas';

const validVesselCreationFormSchema = z.object({
  name: z.string().trim().min(1, 'name is required'),
  pluralName: z.string().trim().min(1, 'plural name is required'),
  slug: inputSlug,
});

export default function ValidVesselCreator(): JSX.Element {
  const router = useRouter();

  const [measurementUnitQuery, setMeasurementUnitQuery] = useState('');
  const [suggestedMeasurementUnits, setSuggestedMeasurementUnits] = useState([] as ValidMeasurementUnit[]);
  const apiClient = buildLocalClient();

  const creationForm = useForm({
    initialValues: {
      capacityUnitID: '',
      iconPath: '',
      pluralName: '',
      description: '',
      name: '',
      slug: '',
      shape: '',
      widthInMillimeters: 0,
      lengthInMillimeters: 0,
      heightInMillimeters: 0,
      capacity: 0,
      includeInGeneratedInstructions: true,
      displayInSummaryLists: true,
      usableForStorage: true,
    },
    validate: zodResolver(validVesselCreationFormSchema),
  });

  useEffect(() => {
    if (measurementUnitQuery.length <= 2) {
      setSuggestedMeasurementUnits([]);
      return;
    }

    const apiClient = buildLocalClient();
    apiClient
      .searchForValidMeasurementUnits(measurementUnitQuery)
      .then((res: QueryFilteredResult<ValidMeasurementUnit>) => {
        setSuggestedMeasurementUnits(res.data || []);
      })
      .catch((err: AxiosError) => {
        console.error(err);
      });
  }, [measurementUnitQuery]);

  const submit = async () => {
    const validation = creationForm.validate();
    if (validation.hasErrors) {
      console.error(validation.errors);
      return;
    }

    const submission = new ValidVesselCreationRequestInput({
      capacityUnitID: creationForm.values.capacityUnitID,
      iconPath: creationForm.values.iconPath,
      pluralName: creationForm.values.pluralName,
      description: creationForm.values.description,
      name: creationForm.values.name,
      slug: creationForm.values.slug,
      shape: creationForm.values.shape,
      widthInMillimeters: creationForm.values.widthInMillimeters,
      lengthInMillimeters: creationForm.values.lengthInMillimeters,
      heightInMillimeters: creationForm.values.heightInMillimeters,
      capacity: creationForm.values.capacity,
      includeInGeneratedInstructions: creationForm.values.includeInGeneratedInstructions,
      displayInSummaryLists: creationForm.values.displayInSummaryLists,
      usableForStorage: creationForm.values.usableForStorage,
    });

    await apiClient
      .createValidVessel(submission)
      .then((result: APIResponse<ValidVessel>) => {
        if (result) {
          router.push(`/valid_vessels/${result.data.id}`);
        }
      })
      .catch((err) => {
        console.error(err);
      });
  };

  return (
    <AppLayout title="Create New Valid Vessel">
      <Container size="sm">
        <form onSubmit={creationForm.onSubmit(submit)}>
          <TextInput label="Name" placeholder="thing" {...creationForm.getInputProps('name')} />
          <TextInput label="Plural Name" placeholder="things" {...creationForm.getInputProps('pluralName')} />
          <TextInput label="Slug" placeholder="thing" {...creationForm.getInputProps('slug')} />
          <TextInput
            label="Description"
            placeholder="stuff about things"
            {...creationForm.getInputProps('description')}
          />
          <TextInput label="Icon path" placeholder="thing" {...creationForm.getInputProps('iconPath')} />

          <TextInput label="Shape" placeholder="thing" {...creationForm.getInputProps('shape')} />

          <NumberInput label="Capacity" {...creationForm.getInputProps('capacity')} />
          <Autocomplete
            label="Capacity Unit"
            placeholder="grams"
            value={measurementUnitQuery}
            onChange={setMeasurementUnitQuery}
            onItemSubmit={async (item: AutocompleteItem) => {
              const selectedValidMeasurmentUnit = suggestedMeasurementUnits.find(
                (x: ValidMeasurementUnit) => x.name === item.value,
              );

              if (!selectedValidMeasurmentUnit) {
                console.error(`selectedValidMeasurementUnitIngredient not found for item ${item.value}}`);
                return;
              }

              creationForm.setFieldValue('capacityUnitID', selectedValidMeasurmentUnit.id);
              setMeasurementUnitQuery(selectedValidMeasurmentUnit.pluralName);
            }}
            data={suggestedMeasurementUnits.map((x: ValidMeasurementUnit) => {
              return { value: x.name, label: x.pluralName };
            })}
          />

          <NumberInput label="Width (mm)" precision={2} {...creationForm.getInputProps('widthInMillimeters')} />
          <NumberInput label="Length (mm)" precision={2} {...creationForm.getInputProps('lengthInMillimeters')} />
          <NumberInput label="Height (mm)" precision={2} {...creationForm.getInputProps('heightInMillimeters')} />

          <Switch
            checked={creationForm.values.displayInSummaryLists}
            label="Display in summary lists"
            {...creationForm.getInputProps('displayInSummaryLists')}
          />

          <Switch
            checked={creationForm.values.includeInGeneratedInstructions}
            label="Include in generated instructions"
            {...creationForm.getInputProps('includeInGeneratedInstructions')}
          />

          <Switch
            checked={creationForm.values.usableForStorage}
            label="Usable for storage"
            {...creationForm.getInputProps('usableForStorage')}
          />

          <Group position="center">
            <Button type="submit" mt="sm" fullWidth>
              Submit
            </Button>
          </Group>
        </form>
      </Container>
    </AppLayout>
  );
}
