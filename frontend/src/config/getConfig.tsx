export const getConfig = () => {
  let cfg: any = {};
  var localStorage = require('localStorage')
  const token = localStorage.getItem('token')

  if (token !== null) {
    cfg = {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    };
  }

  return cfg;
}