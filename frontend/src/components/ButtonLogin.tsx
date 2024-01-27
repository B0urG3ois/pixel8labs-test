import { Inter } from 'next/font/google';
import { useAuthentication } from '../hooks/authentication';
import { Button } from './Button';

const inter = Inter({ subsets: ['latin'] });

export const ButtonLogin = () => {
  const handlerAuth = useAuthentication();
  const handleClick = () => {
    handlerAuth().catch((err) => {
      console.log(err);
    });
  };

  return (
    <Button onClick={handleClick} className={`${inter.className} font-semibold text-white bg-custom-pink rounded-lg px-[18px] py-[10px]`}>
      Login with GitHub
    </Button>
  );
};
