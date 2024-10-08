// Code generated by gen_typescript. DO NOT EDIT.

export interface IUserNotification {
  createdAt: NonNullable<string>;
  lastUpdatedAt?: string;
  id: NonNullable<string>;
  content: NonNullable<string>;
  status: NonNullable<string>;
  belongsToUser: NonNullable<string>;
}

export class UserNotification implements IUserNotification {
  createdAt: NonNullable<string> = '1970-01-01T00:00:00Z';
  lastUpdatedAt?: string;
  id: NonNullable<string> = '';
  content: NonNullable<string> = '';
  status: NonNullable<string> = '';
  belongsToUser: NonNullable<string> = '';

  constructor(input: Partial<UserNotification> = {}) {
    this.createdAt = input.createdAt ?? '1970-01-01T00:00:00Z';
    this.lastUpdatedAt = input.lastUpdatedAt;
    this.id = input.id ?? '';
    this.content = input.content ?? '';
    this.status = input.status ?? '';
    this.belongsToUser = input.belongsToUser ?? '';
  }
}

export interface IUserNotificationCreationRequestInput {
  content: NonNullable<string>;
  status: NonNullable<string>;
  belongsToUser: NonNullable<string>;
}

export class UserNotificationCreationRequestInput implements IUserNotificationCreationRequestInput {
  content: NonNullable<string> = '';
  status: NonNullable<string> = '';
  belongsToUser: NonNullable<string> = '';

  constructor(input: Partial<UserNotificationCreationRequestInput> = {}) {
    this.content = input.content ?? '';
    this.status = input.status ?? '';
    this.belongsToUser = input.belongsToUser ?? '';
  }
}

export interface IUserNotificationUpdateRequestInput {
  status?: string;
}

export class UserNotificationUpdateRequestInput implements IUserNotificationUpdateRequestInput {
  status?: string;

  constructor(input: Partial<UserNotificationUpdateRequestInput> = {}) {
    this.status = input.status;
  }
}
