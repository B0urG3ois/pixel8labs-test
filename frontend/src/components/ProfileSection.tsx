import { User } from '../types/user';
import Image from 'next/image';

interface Props {
  user?: User;
}

export const ProfileSection = ({ user }: Props) => {
  const avatarUrl = user?.avatar_url;

  return (
    <div className='flex lg:flex-col items-center'>
      <Image className='hidden rounded-full lg:block' src={`${avatarUrl ?? '/assets/images/profile.png'}`} width={160} height={160} priority={true} alt='Profile photo' />
      <Image className='lg:hidden rounded-full' src={`${avatarUrl ?? '/assets/images/profile.png'}`} width={80} priority={true} height={80} alt='Profile photo' />

      <div className='ml-4 lg:ml-0 lg:text-center'>
        <h2 className='mt-2 text-2xl text-gray-800 font-bold'>
          {user?.name}
        </h2>
        <div className='text-gray-500 -mt-0.5'>@{user?.login}</div>
      </div>
    </div>
  );
};