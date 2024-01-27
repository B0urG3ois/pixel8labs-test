import { type NextPage } from 'next';
import { useGetAccessToken } from '../hooks/getAccessToken';

const Callback: NextPage = () => {
  useGetAccessToken();

  return (
    <div>
      <p>Logging in...</p>
    </div>
  );
};

export default Callback;
