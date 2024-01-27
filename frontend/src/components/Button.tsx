import React from 'react';

interface Props {
  className?: string;
  children?: React.ReactNode;
  onClick?: React.MouseEventHandler<HTMLButtonElement> | undefined;
}

export const Button = ({ children, className, onClick }: Props) => {
  return (
    <button onClick={onClick} className={className}>
      {children}
    </button>
  )
}