import { useContext } from 'react';
import { AppContext, AppContextType } from '../../../context/AppContext';
import { VisitorSection } from '../../molecules/profile/VisitorSection';

export const ListOfVisitors = () => {
  const ctx = useContext(AppContext) as AppContextType;

  return (
    <div className='mt-10'>
      <h3 className='text-gray-800 font-bold text-sm'>Latest Visitor</h3>
      <div className='flex items-center gap-4 mt-2'>
        <VisitorSection item={ctx.followers?.[0]} />
        <VisitorSection item={ctx.followers?.[1]} />
        <VisitorSection item={ctx.followers?.[2]} />
      </div>
    </div>
  );
};