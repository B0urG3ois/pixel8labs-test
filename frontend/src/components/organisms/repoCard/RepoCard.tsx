import { useContext } from 'react';
import { AppContext, AppContextType } from '../../../context/AppContext';
import { ListOfRepos } from '../../molecules/listOfRepos/ListOfRepos';

export const RepoCard = () => {
  const ctx = useContext(AppContext) as AppContextType;

  return (
    <div className=''>
      {ctx.repositories?.map((item, i) => {
        return <ListOfRepos {...item} key={i} />;
      })}
    </div>
  );
};