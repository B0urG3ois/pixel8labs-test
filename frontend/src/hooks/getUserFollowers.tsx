import { useContext, useEffect } from 'react';
import { AppDataContext, AppData } from '../context/AppDataContext';
import { getConfig } from '../config/getConfig';
import { Followers } from '../types/followers';

export const useGetUserFollowers = () => {
  const ctx = useContext(AppDataContext) as AppData;

  useEffect(() => {
    const config = getConfig();

    if (ctx.currentUser !== undefined) {
      fetch(ctx.currentUser.followers_url, config)
        .then(async (response) => {
          return await response.json();
        })
        .then((followers: Followers[]) => {
          ctx.setUserFollowers(followers);
        })
        .catch((err) => {
          throw err;
        });
    }
  }, [ctx.currentUser]);
}