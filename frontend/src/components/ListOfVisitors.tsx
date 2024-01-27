import { useContext } from 'react';
import { AppDataContext, AppData } from '../context/AppDataContext';
import { VisitorSection } from './VisitorSection';

export const ListOfVisitors = () => {
  const ctx = useContext(AppDataContext) as AppData;

  return (
    <div className='mt-10'>
      <h3 className='text-gray-800 font-bold text-sm'>Latest Visitor</h3>
      <div className='flex items-center gap-4 mt-2'>
        <VisitorSection item={ctx.userFollowers?.[0]} />
        <VisitorSection item={ctx.userFollowers?.[1]} />
        <VisitorSection item={ctx.userFollowers?.[2]} />
      </div>
    </div>
  );
};