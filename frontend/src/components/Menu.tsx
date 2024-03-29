import Image from 'next/image';
import Link from 'next/link';
import { Fragment, useContext } from 'react';
import { AppDataContext, AppData } from '../context/AppDataContext';
import { BurgerBars } from './BurgerBars';
import { useRevokeAuth } from '../hooks/revokeAuth';
import { Menu as HeadlessMenu, Transition } from '@headlessui/react';

const MenuBtn = () => {
  const ctx = useContext(AppDataContext) as AppData;
  const avatarUrl = ctx.currentUser?.avatar_url;

  return (
    <div className='flex hover:cursor-pointer hover:bg-gray-50 items-center px-4 py-2 gap-4 border rounded-full'>
      <Image
        className='rounded-full'
        src={`${avatarUrl ?? '/assets/images/avatar.png'}`}
        width={40}
        height={40}
        alt='Avatar'
      />
      <span className='text-gray-800 hover:text-gray-500'>
        <BurgerBars />
      </span>
    </div>
  );
};

export const Menu = () => {
  const ctx = useContext(AppDataContext) as AppData;
  const avatarUrl = ctx.currentUser?.avatar_url;

  const logout = useRevokeAuth();
  const handleLogout = () => {
    logout();
  };

  return (
    <HeadlessMenu as={'div'} className={'relative'}>
      <HeadlessMenu.Button>
        <MenuBtn />
      </HeadlessMenu.Button>
      <Transition
        as={Fragment}
        enter='transition ease-out duration-100'
        enterFrom='transform opacity-0 scale-95'
        enterTo='transform opacity-100 scale-100'
        leave='transition ease-in duration-75'
        leaveFrom='transform opacity-100 scale-100'
        leaveTo='transform opacity-0 scale-95'
      >
        <HeadlessMenu.Items
          className={
            'absolute border mt-1  min-w-[240px] right-0 bg-white shadow-lg rounded-lg pt-2 pb-1 flex flex-col'
          }
        >
          <HeadlessMenu.Item>
            <Link
              href={'/'}
              className='py-[10px] border-b whitespace-nowrap font-medium text-gray-600 text-sm'
            >
              <div className='flex gap-3 items-center px-4'>
                <div className='w-10 h-10'>
                  <Image
                    className='rounded-full'
                    src={`${avatarUrl ?? '/assets/images/avatar.png'}`}
                    width={40}
                    height={40}
                    alt='Avatar'
                  />
                </div>
                <div>
                  <h3 className='text-sm text-gray-600 font-semibold whitespace-nowrap'>
                    {ctx.currentUser?.name}
                  </h3>
                  <p className='whitespace-nowrap text-xs font-normal text-gray-500'>
                    {ctx.currentUser?.email}
                  </p>
                </div>
              </div>
            </Link>
          </HeadlessMenu.Item>
          <HeadlessMenu.Item>
            <Link
              href={`https://github.com/${ctx.currentUser?.login}`}
              className='px-4 py-[10px] font-medium text-gray-600 hover:text-gray-500 text-sm'
            >
              View profile
            </Link>
          </HeadlessMenu.Item>
          <HeadlessMenu.Item>
            <button
              onClick={handleLogout}
              className='text-left px-4 py-[10px] font-medium text-gray-600 hover:text-gray-500 text-sm border-t'
            >
              Log out
            </button>
          </HeadlessMenu.Item>
        </HeadlessMenu.Items>
      </Transition>
    </HeadlessMenu>
  );
};