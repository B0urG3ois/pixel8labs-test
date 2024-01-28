import { AppDataContext, AppData } from '../context/AppDataContext';
import { useContext, useEffect } from 'react';
import { getConfig } from '../config/getConfig';
import { User } from '../types/user';

export const useGetUserInfo = () => {
  const ctx = useContext(AppDataContext) as AppData;

  useEffect(() => {
    const config = getConfig();
    const url = process.env.NEXT_PUBLIC_BASE_API + '/' + process.env.NEXT_PUBLIC_API_VERSION + '/user';

    fetch(url, config)
      .then(async (response) => {
        return await response.json();
      })
      .then((data) => {
        const userData = data.data as User;
        ctx.setCurrentUser(userData);
      })
      .catch((err) => {
        throw err;
      });
  }, []);
}