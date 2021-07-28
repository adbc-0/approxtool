import React, {
  BaseSyntheticEvent,
  useState
} from 'react';
import {
  initialRegisterState,
  isError,
  validateRegisterForm
} from '../LoginForm/Login.utils';
import { InputField } from '../../common-components/InputField/InputField';
import styles from './styles.module.scss';
import { Button } from '../../common-components/Button/Button';
import { mutateUser } from '../UserManager/UserManager.utils';
import { User } from '../../types';
import {
  NewUserRequest,
  Role
} from '../../protos/userservice_pb';
import {
  apiSrv,
  fetchWithAuthRetry
} from '../../grpc-web';
import { token } from '../../utils/token';

interface Props {
  handleClose: () => void;
  setUsers?: React.Dispatch<React.SetStateAction<User[]>>;
}

export const RegisterForm: React.FC<Props> = ({handleClose, setUsers}): JSX.Element => {
  const [registerForm, setRegisterForm] = useState(initialRegisterState);
  const [registerErrors, setRegisterErrors] = useState(initialRegisterState);

  const handleRegisterInput = (e: BaseSyntheticEvent) => setRegisterForm(prev => ({
    ...prev,
    [e.target.name]: e.target.value
  }));

  const handleRegister = async (e: BaseSyntheticEvent) => {
    e.preventDefault();
    const errors = validateRegisterForm(registerErrors, registerForm);
    setRegisterErrors(errors);
    if (isError(Object.entries(errors))) return;

    const request = new NewUserRequest();
    request.setUsername(registerForm.name);
    request.setEmail(registerForm.email);
    request.setPassword(registerForm.password);
    request.setStatus(Role.BASIC_USER);

    try {
      const result = await fetchWithAuthRetry(() => {
        request.setAccesstoken(token.accessToken);
        return apiSrv.createUser(request, null);
      });

      if (setUsers) setUsers(prev => mutateUser.addUser(prev, result.toObject()));
      handleClose();
    } catch (e) {
      console.error(e);
    }
  };

  return (
    <form className={styles.formWrapper} onSubmit={handleRegister}>
      <h2 className={styles.formText}>Register</h2>
      <div className={styles.inputWrapper}>
        <InputField
          label="Email"
          name="email"
          value={registerForm.email}
          handler={handleRegisterInput}
        />
        {registerErrors.email && <p className={styles.validationError}>{registerErrors.email}</p>}
        <InputField
          label="Name"
          name="name"
          value={registerForm.name}
          handler={handleRegisterInput}
        />
        {registerErrors.name && <p className={styles.validationError}>{registerErrors.name}</p>}
        <InputField
          label="Password"
          type="password"
          name="password"
          value={registerForm.password}
          handler={handleRegisterInput}
          autoComplete="on"
        />
        {registerErrors.password && <p className={styles.validationError}>{registerErrors.password}</p>}
        <InputField
          label="Confirm password"
          type="password"
          name="confirmPassword"
          value={registerForm.confirmPassword}
          handler={handleRegisterInput}
          autoComplete="on"
        />
        {registerErrors.confirmPassword && <p className={styles.validationError}>{registerErrors.confirmPassword}</p>}
        <Button text="Create account" type="submit"/>
        <Button text="Close" onClick={handleClose}/>
      </div>
    </form>
  );
};