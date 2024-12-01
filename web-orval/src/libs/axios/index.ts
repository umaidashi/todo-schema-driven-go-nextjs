//NOTE: https://orval.dev/reference/configuration/output#mutator
// use-custom-instance.ts

import { API_BASE_URL } from '@/constants/env'
import Axios, { type AxiosError, type AxiosRequestConfig } from 'axios'

export const AXIOS_INSTANCE = Axios.create({ baseURL: API_BASE_URL, timeout: 30000 })

export const useCustomInstance = <T>(): ((config: AxiosRequestConfig) => Promise<T>) => {
  return async (config: AxiosRequestConfig) => {
    const promise = AXIOS_INSTANCE({
      ...config
    }).then(({ data }) => data)

    return promise
  }
}

export type ErrorType<Error> = AxiosError<Error>
