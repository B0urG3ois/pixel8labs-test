import { useContext, useEffect } from 'react';
import { AppContext, AppContextType } from '../context/AppContext';
import { getConfig } from '../config/getConfig';
import { Repository } from '../types/repository';

export const useGetUserRepositories = () => {
  const ctx = useContext(AppContext) as AppContextType;

  useEffect(() => {
    const config = getConfig();
    const url = process.env.NEXT_PUBLIC_BASE_API + '/repositories';

    fetch(url, config)
      .then(async (response) => {
        return await response.json();
      })
      .then((data) => {
        const userRepositories = data.data as Repository[];
        ctx.setRepositories(userRepositories);
      })
      .catch((err) => {
        throw err;
      })
  }, []);
}