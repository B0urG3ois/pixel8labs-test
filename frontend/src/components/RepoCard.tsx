import { useContext } from 'react';
import { AppDataContext, AppData } from '../context/AppDataContext';
import { ListOfRepos } from './ListOfRepos';

export const RepoCard = () => {
  const ctx = useContext(AppDataContext) as AppData;

  return (
    <div className=''>
      {ctx.userRepositories?.map((item, i) => {
        return <ListOfRepos {...item} key={i} />;
      })}
    </div>
  );
};