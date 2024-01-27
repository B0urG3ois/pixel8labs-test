import { useContext } from 'react';
import { AppDataContext, AppData } from '../context/AppDataContext';
import { Heading } from './Heading';
import { BadgeMedium } from './Badge';

export const Section = () => {
  const ctx = useContext(AppDataContext) as AppData;

  return (
    <div className='flex items-center gap-3 h-20'>
      <Heading text='Repository' />
      <BadgeMedium text={ctx.userRepositories?.length.toString() ?? '0'} />
    </div>
  );
};