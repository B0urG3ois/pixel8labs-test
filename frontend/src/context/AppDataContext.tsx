import React, { ReactNode, useState } from 'react';
import { User } from '../types/user';
import { Followers } from '../types/followers';
import { Repository } from '../types/repository';

export const AppDataContext = React.createContext({});

interface Props {
  children?: ReactNode
}

export interface AppData {
  currentUser?: User;
  setCurrentUser: React.Dispatch<React.SetStateAction<User | undefined>>;
  userFollowers?: Followers[];
  setUserFollowers: React.Dispatch<React.SetStateAction<Followers[] | undefined>>;
  userRepositories?: Repository[];
  setUserRepositories: React.Dispatch<React.SetStateAction<Repository[] | undefined>>;
}
export const AppDataWrapper = ({ children }: Props) => {
  const [currentUser, setCurrentUser] = useState<User>();
  const [userFollowers, setUserFollowers] = useState<Followers[]>();
  const [userRepositories, setUserRepositories] = useState<Repository[]>();

  const data: AppData = {
    currentUser,
    setCurrentUser,
    userFollowers,
    setUserFollowers,
    userRepositories,
    setUserRepositories,
  };

  return <AppDataContext.Provider value={data}>{children}</AppDataContext.Provider>
}