import { useSearchParams } from 'next/navigation';
import { useEffect } from 'react';
import { LoginResponse } from '../types/login';

export const useGetAccessToken = () => {
  const params = useSearchParams()

  useEffect(() => {
    var localStorage = require('localStorage')
    const url = process.env.NEXT_PUBLIC_BASE_API + '/callback?code=' + params.get('code');

    if (params.get('code') != null) {
      fetch(url)
        .then(async (response) => {
          return await response.json();
        })
        .then((data) => {
          const loginData = data.data as LoginResponse;
          localStorage.setItem('token', loginData.access_token);
        })
        .then(() => {
          window.location.href = '/';
        })
        .catch((error) => {
          throw error;
        });
    }
  }, [params]);
}