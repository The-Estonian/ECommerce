export const Login = async (formData) => {
  try {
    const response = await fetch(`localhost:8080/login`, {
      method: 'POST',
      mode: 'cors',
      cache: 'no-cache',
      referrerPolicy: 'no-referrer',
      redirect: 'follow',
      body: formData,
      credentials: 'include',
    });
    const resp = await response.json();
    console.log(resp);
    return resp;
  } catch (error) {
    console.log('Login error');
    console.log(error);
    return { server: 'fail', error: 'Error 500, Internal server error!' };
  }
};
