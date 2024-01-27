import { useContext } from 'react';
import { AppContext, AppContextType } from '../../../context/AppContext';
import Link from 'next/link';
import Image from 'next/image';
import { ButtonLogin } from '../../atoms/button/ButtonLogin';
import { Menu } from '../menu/Menu';

export const Header = () => {
  const ctx = useContext(AppContext) as AppContextType;

  return (
    <nav className='header flex justify-center shadow-sm bg-white'>
      <div className='flex items-center justify-between py-2 px-4 container'>

        <Link href={'/'}>
          <Image
            src={'/assets/images/logo-full.svg'}
            width={173}
            height={52}
            priority={false}
            alt='Logo'
          />
        </Link>

        <div>{ctx.user?.login === 'octocat' ? <ButtonLogin /> : <Menu />}</div>
      </div>
    </nav>
  );
};