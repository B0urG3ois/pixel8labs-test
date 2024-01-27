import { AppContext, AppContextType } from '../context/AppContext';
import { useContext, useEffect } from 'react';
import { getConfig } from '../config/getConfig';
import { User } from '../types/user';

export const useGetUserInfo = () => {
  const ctx = useContext(AppContext) as AppContextType;

  useEffect(() => {
    const config = getConfig();
    const url = process.env.NEXT_PUBLIC_BASE_API + '/user';

    fetch(url, config)
      .then(async (response) => {
        return await response.json();
      })
      .then((data) => {
        const userData = data.data as User;
        ctx.setUser(userData);
      })
      .catch((err) => {
        throw err;
      });
  }, []);
}