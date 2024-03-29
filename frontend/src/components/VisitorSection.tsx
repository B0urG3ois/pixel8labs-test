import { Followers } from '../types/followers';
import Image from 'next/image';
import Link from 'next/link';

interface Props {
  item?: Followers;
}

export const VisitorSection = ({ item }: Props) => {
  return (
    <Link href={`${item?.html_url}` ?? 'https://github.com'}>
      <Image className='rounded-full' src={item?.avatar_url ?? '/assets/images/avatar.png'} width={56} height={56} alt='Avatar' />
    </Link>
  );
};