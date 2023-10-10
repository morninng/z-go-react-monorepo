import { useCallback } from 'react';


import {AdminApi} from '../api/api'


export const useApiTask = () => {

  const getTasks =
    async (): Promise<any> => {
      try {
        const resp = await new AdminApi().getTasks();
        return resp.data;
      } catch (err) {
        console.error(err);
        return null
      }
    }

  return { getTasks };
};
