import React, { ReactNode, useState } from 'react';
import { User } from '../types/user';
import { Followers } from '../types/followers';
import { Repository } from '../types/repository';

export const AppContext = React.createContext({});

interface Props {
  children?: ReactNode
}

export interface AppContextType {
  user?: User;
  setUser: React.Dispatch<React.SetStateAction<User | undefined>>;
  followers?: Followers[];
  setFollowers: React.Dispatch<React.SetStateAction<Followers[] | undefined>>;
  repositories?: Repository[];
  setRepositories: React.Dispatch<React.SetStateAction<Repository[] | undefined>>;
}
export const AppProvider = ({ children }: Props) => {
  const [user, setUser] = useState<User>();
  const [followers, setFollowers] = useState<Followers[]>();
  const [repositories, setRepositories] = useState<Repository[]>();

  const data: AppContextType = {
    user,
    setUser,
    followers,
    setFollowers,
    repositories,
    setRepositories,
  };

  return <AppContext.Provider value={data}>{children}</AppContext.Provider>
}