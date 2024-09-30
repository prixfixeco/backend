const path = require('path');
const withTM = require('next-transpile-modules')([
  '@dinnerdonebetter/models',
  '@dinnerdonebetter/utils',
  '@dinnerdonebetter/api-client',
  '@dinnerdonebetter/logger',
  '@dinnerdonebetter/server-timing',
  '@dinnerdonebetter/tracing',
  '@dinnerdonebetter/next-routes',
  '@dinnerdonebetter/encryption',
]);

module.exports = withTM({
  reactStrictMode: true,
  output: 'standalone',
  env: {
    NEXT_PUBLIC_API_ENDPOINT: 'https://api.dinnerdonebetter.dev', // TODO: make this actually variable
    NEXT_COOKIE_ENCRYPTION_KEY: 'ZOTGz4KEhZFSM6udeESOX5JVqhtEdHdS', // TODO: make this actually variable
    NEXT_BASE64_COOKIE_ENCRYPT_IV: 'S2IwVXVvMW9hSEl4WjQ0ak1NYW50QndMTzJBWDJFV2o=', // TODO: make this actually variable
  },
  experimental: {
    outputFileTracingRoot: path.join(__dirname, '../../'),
  },
});
