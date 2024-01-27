import { CSSProperties } from 'react';

interface Props {
  className?: string;
  style?: CSSProperties;
}

export const LanguageDot = ({ className, style }: Props) => {
  return (
    <div style={style} className={`${className} w-3 h-3 rounded-full`}></div>
  )
}