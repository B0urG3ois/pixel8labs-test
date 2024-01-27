import '../common/styles/app.scss';
import { Inter } from 'next/font/google';
import { AppProps } from 'next/app';
import { AppDataWrapper } from '../context/AppDataContext';

const inter = Inter({ subsets: ['latin'] });

export default function App({ Component, pageProps }: AppProps) {
  return (
    <>
      <style jsx global>
        {`html {
          font-family: ${inter.style.fontFamily};
        }`
        }
      </style>
      <AppDataWrapper>
        <Component {...pageProps} />
      </AppDataWrapper>
    </>
  );
}