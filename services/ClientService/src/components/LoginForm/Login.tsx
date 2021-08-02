import React, {
  BaseSyntheticEvent,
  useState
} from 'react';
import { useHistory, } from 'react-router-dom';
import { InputField } from '../../common-components/InputField/InputField';
import {
  initialLoginState,
  isError,
  validateLoginForm
} from './Login.utils';
import styles from './styles.module.scss';
import { apiService } from '../../grpc-web/apiService';
import { token } from '../../utils/token';

export const Login = () => {
  const history = useHistory();
  const [loginForm, setLoginForm] = useState(initialLoginState);
  const [loginErrors, setLoginErrors] = useState(initialLoginState);

  const handleLoginInput = (e: BaseSyntheticEvent) => setLoginForm(prev => ({
    ...prev,
    [e.target.name]: e.target.value
  }));

  const handleLogin = async (e: BaseSyntheticEvent) => {
    e.preventDefault();
    const errors = validateLoginForm(loginErrors, loginForm);
    setLoginErrors(errors);
    if (isError(Object.entries(errors))) return;

    const {email, password} = loginForm;
    try {
      const response = await apiService.Login(email, password);
      token.setAccessToken = response.toObject().accessToken;
      token.setRefreshToken = response.toObject().refreshToken;
      history.push('/');
    } catch (err) {
      console.error(err.code, err.message);
    }
  };

  return (
    <form className={styles.formWrapper} onSubmit={handleLogin}>
      <h2 className={styles.formText}>Login</h2>
      <div className={styles.inputWrapper}>
        <InputField
          label="Email"
          name="email"
          value={loginForm.email}
          handler={handleLoginInput}
        />
        {loginErrors.email && <p className={styles.validationError}>{loginErrors.email}</p>}
        <InputField
          label="Password"
          type="password"
          name="password"
          value={loginForm.password}
          handler={handleLoginInput}
          autoComplete="on"
        />
        {loginErrors.password && <p className={styles.validationError}>{loginErrors.password}</p>}
        <div className={styles.buttonsWrapper}>
          <button type="submit" className={styles.button}>Login</button>
        </div>
      </div>
    </form>
  );
};
