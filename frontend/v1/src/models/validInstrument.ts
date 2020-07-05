import * as Factory from "factory.ts";
import faker from "faker";
import {defaultFactories} from "@/models/fakes";

export class ValidInstrument {
  id: number;
  name: string;
  variant: string;
  description: string;
  icon: string;
  createdOn: number;
  updatedOn?: number;
  archivedOn?: number;

  constructor() {
    this.id = 0;
    this.name = '';
    this.variant = '';
    this.description = '';
    this.createdOn = 0;
    this.icon = '';
  }
}

export function validInstrumentsAreEqual(
  i1: ValidInstrument,
  i2: ValidInstrument,
): boolean {
  return (
    i1.id === i2.id &&
    i1.name === i2.name &&
    i1.variant === i2.variant &&
    i1.description === i2.description &&
    i1.icon === i2.icon &&
    i1.createdOn === i2.createdOn &&
    i1.updatedOn === i2.updatedOn &&
    i1.archivedOn === i2.archivedOn
  );
}

export const fakeValidInstrumentFactory = Factory.Sync.makeFactory<ValidInstrument> ({
  name: Factory.Sync.each(() => faker.lorem.word()),
  variant: Factory.Sync.each(() => faker.lorem.words(3)),
  description: Factory.Sync.each(() => faker.lorem.words(6)),
  icon: Factory.Sync.each(() => ''),
  ...defaultFactories,
});
