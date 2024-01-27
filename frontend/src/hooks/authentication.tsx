export const useAuthentication = () => {
  const url = process.env.NEXT_PUBLIC_BASE_API + '/login';

  const authenticate = async () => {
    try {
      const response = await fetch(url, {
        method: 'POST',
        headers: {
          'content-type': 'application/json',
        },
      });

      const jsonData = await response.json();
      window.location.href = jsonData.data;
    } catch (err: any) {
      throw new Error('Failed to authenticate' + err.message)
    }
  };

  return authenticate;
}