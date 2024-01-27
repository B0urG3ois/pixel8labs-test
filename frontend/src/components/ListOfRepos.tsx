import Link from 'next/link';
import { Repository } from '../types/repository';
import { BadgeSmall } from './Badge';
import { LanguageDot } from './LanguageDot';
import { shortMonthAbbreviations } from '../common/constant/shortMonthAbbreviations';
import { programmingLanguageColor } from '../common/constant/programmingLanguage';

export const ListOfRepos = ({
  name,
  description,
  language,
  updated_at: updatedAt,
  private: isPrivate,
  url: URL,
}: Repository) => {
  const date = new Date(updatedAt);
  const month = shortMonthAbbreviations[date.getMonth()];
  const day = date.getDate();
  const formattedDate = `updated on ${month} ${day}`;

  return (
    <Link href={`${URL}`}>
      <div className='bg-gray-25 border p-6 rounded-lg mb-6'>
        <div className='mb-2 flex items-center gap-3 flex-wrap'>
          <h3 className='font-bold text-gray-800'>{name}</h3>
          {!isPrivate ? (
            <BadgeSmall text='public' className='bg-indigo-50 text-badge-text-public font-medium' />
          ) : (
            <BadgeSmall text='private' className='bg-private text-gray-700 font-medium' />
          )}
        </div>

        <p className='text-gray-800 text-sm'>{description}</p>

        <div className='mt-6 flex items-center gap-6 text-xs'>
          {language !== '' ? (
            <div className='flex items-center gap-1'>
              <LanguageDot style={{ backgroundColor: programmingLanguageColor[language] }} className='mb-[1.2px]' />
              <span className='text-gray-800'>{language}</span>
            </div>
          ) : null}
          <div className='text-gray-500'>{formattedDate}</div>
        </div>
      </div>
    </Link>
  );
};