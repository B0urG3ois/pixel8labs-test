import { useContext, useEffect } from 'react';
import { AppDataContext, AppData } from '../context/AppDataContext';
import { getConfig } from '../config/getConfig';
import { Repository } from '../types/repository';

export const useGetUserRepositories = () => {
  const ctx = useContext(AppDataContext) as AppData;

  useEffect(() => {
    const config = getConfig();
    const url = process.env.NEXT_PUBLIC_BASE_API + '/' + process.env.NEXT_PUBLIC_API_VERSION + '/repositories';

    fetch(url, config)
      .then(async (response) => {
        return await response.json();
      })
      .then((data) => {
        const userRepositories = data.data as Repository[];
        ctx.setUserRepositories(userRepositories);
      })
      .catch((err) => {
        throw err;
      })
  }, []);
}