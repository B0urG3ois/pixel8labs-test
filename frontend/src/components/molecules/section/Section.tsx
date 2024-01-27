import { useContext } from 'react';
import { AppContext, AppContextType } from '../../../context/AppContext';
import { Heading } from '../../atoms/titles/Heading';
import { BadgeMedium } from '../../atoms/badge/Badge';

export const Section = () => {
  const ctx = useContext(AppContext) as AppContextType;

  return (
    <div className='flex items-center gap-3 h-20'>
      <Heading text='Repository' />
      <BadgeMedium text={ctx.repositories?.length.toString() ?? '0'} />
    </div>
  );
};