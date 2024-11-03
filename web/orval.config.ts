/* eslint-disable @typescript-eslint/no-var-requires */
const path = require('node:path')
const { defineConfig } = require('orval')

// eslint-disable-next-line import/no-default-export
export default defineConfig({
  apiClient: {
    input: './../server/docs/openapi.yaml',
    output: {
      client: 'react-query',
      mode: 'split',
      target: './src/api/generated',
      override: {
        mutator: {
          path: './src/libs/axios/index.ts',
          name: 'useCustomInstance',
          alias: {
            '@': path.resolve(__dirname, './src')
          }
        }
      }
    }
  }
})
