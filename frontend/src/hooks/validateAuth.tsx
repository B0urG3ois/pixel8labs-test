import { useRouter} from 'next/router';
import { useEffect } from 'react';

export const useValidateAuth = () => {
  const r = useRouter();

  useEffect(() => {
    var localStorage = require('localStorage')
    const token = localStorage.getItem('token')

    if (token === null) {
      r.push('/octocat')
    }
  }, []);
}