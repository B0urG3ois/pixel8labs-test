export const useRevokeAuth = () => {
  const deauthenticate = async () => {
    var localStorage = require('localStorage')
    const token = localStorage.getItem('token');
    const url = process.env.NEXT_PUBLIC_BASE_API + '/logout';

    try {
      await fetch(url, {
        method: 'POST',
        headers: {
          Authorization: `Bearer ${token}`,
          'content-type': 'application/json',
        },
      });

      localStorage.removeItem('token');
      window.location.href = '/';
    } catch (err: any) {
      throw new Error('Failed to revoke token' + err.message);
    }
  };

  return deauthenticate;
}