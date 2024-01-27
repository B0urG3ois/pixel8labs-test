import { User } from '../../../types/user';
import { Envelope } from '../../atoms/icon/Envelope';
import { UserIcon } from '../../atoms/icon/UserIcon';

interface Props {
  user?: User;
}

export const AboutSection = ({ user }: Props) => {
  return (
    <div className='mt-6'>
      <div>
        <h3 className='text-gray-800 font-bold text-sm'>About</h3>
        <p className='text-gray-800 mt-2'>{user?.bio}</p>
      </div>

      <div className='mt-4'>
        <div className='flex items-center gap-2'>
          <span>
            <Envelope />
          </span>
          <span className='text-gray-600'>{user?.email}</span>
        </div>
        <div className='flex items-center gap-2 mt-2'>
          <span>
            <UserIcon />
          </span>
          <div>
            <span className='text-gray-800 font-bold '>{user?.visitors}</span>
            <span className='text-gray-600 ml-2'>profile visitor</span>
          </div>
        </div>
      </div>
    </div>
  );
};