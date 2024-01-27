import { useContext } from 'react';
import { AppContext, AppContextType } from '../../../context/AppContext';
import { ProfileSection } from '../../molecules/profile/ProfileSection';
import { AboutSection } from '../../molecules/profile/AboutSection';
import { ListOfVisitors } from './ListOfVisitors';

export const Sidebar = () => {
  const ctx = useContext(AppContext) as AppContextType;

  return (
    <div className='px-4  flex flex-col justify-center'>
      <ProfileSection user={ctx.user} />
      <AboutSection user={ctx.user} />
      <ListOfVisitors />
    </div>
  );
};