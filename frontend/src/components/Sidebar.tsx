import { useContext } from 'react';
import { AppDataContext, AppData } from '../context/AppDataContext';
import { ProfileSection } from './ProfileSection';
import { AboutSection } from './AboutSection';
import { ListOfVisitors } from './ListOfVisitors';

export const Sidebar = () => {
  const ctx = useContext(AppDataContext) as AppData;

  return (
    <div className='px-4  flex flex-col justify-center'>
      <ProfileSection user={ctx.currentUser} />
      <AboutSection user={ctx.currentUser} />
      <ListOfVisitors />
    </div>
  );
};