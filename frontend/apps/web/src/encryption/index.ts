import { EncryptorDecryptor } from '@dinnerdonebetter/encryption';

export default new EncryptorDecryptor(
  process.env.NEXT_COOKIE_ENCRYPTION_KEY || '',
  process.env.NEXT_BASE64_COOKIE_ENCRYPT_IV || '',
);
