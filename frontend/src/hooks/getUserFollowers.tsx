import { useContext, useEffect } from 'react';
import { AppContext, AppContextType } from '../context/AppContext';
import { getConfig } from '../config/getConfig';
import { Followers } from '../types/followers';

export const useGetUserFollowers = () => {
  const ctx = useContext(AppContext) as AppContextType;

  useEffect(() => {
    const config = getConfig();

    if (ctx.user !== undefined) {
      fetch(ctx.user.followers_url, config)
        .then(async (response) => {
          return await response.json();
        })
        .then((followers: Followers[]) => {
          ctx.setFollowers(followers);
        })
        .catch((err) => {
          throw err;
        });
    }
  }, [ctx.user]);
}