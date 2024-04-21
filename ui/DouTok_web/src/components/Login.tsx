import * as Form from '@radix-ui/react-form';
import { useQuery } from '@tanstack/react-query';
import React, { useCallback, useEffect, useState } from 'react';
import { useDebounce } from '../hooks/useDebounce';

export const Login = (): JSX.Element => {
  const [submitClicked, setSubmitClicked] = useState(false);
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const onChangeHandler = (
    inputType: string,
    e: React.ChangeEvent<HTMLInputElement>
  ) => {
    switch (inputType) {
      case 'email':
        setEmail(e.target?.value);
        break;

      case 'password':
        setPassword(e.target?.value);
        break;

      default:
        break;
    }
  };
  const debouncedHandler = useDebounce(onChangeHandler, 500);
  const { data } = useQuery({
    queryKey: ['loginData'],
    queryFn: () =>
      new Promise((resolve) => {
        setTimeout(() => resolve({ auth: true }), 1000);
      }),
    enabled: submitClicked,
  });
  useEffect(() => {
    console.log('ljxx', data);
  }, [data]);
  const onSubmit = useCallback((e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setSubmitClicked(true);
  }, []);
  return (
    <Form.Root className='FormRoot' onSubmit={onSubmit}>
      <Form.Field className='FormField' name='email'>
        <div
          style={{
            display: 'flex',
            alignItems: 'baseline',
            justifyContent: 'space-between',
          }}
        >
          <Form.Label className='FormLabel'>Email</Form.Label>
          <Form.Message className='FormMessage' match='valueMissing'>
            Please enter your email
          </Form.Message>
          <Form.Message className='FormMessage' match='typeMismatch'>
            Please provide a valid email
          </Form.Message>
        </div>
        <Form.Control asChild>
          <input
            className='Input'
            type='email'
            required
            onChange={(e) => debouncedHandler('email', e)}
          />
        </Form.Control>
      </Form.Field>
      <Form.Field className='FormField' name='password'>
        <div
          style={{
            display: 'flex',
            alignItems: 'baseline',
            justifyContent: 'space-between',
          }}
        >
          <Form.Label className='FormLabel'>Password</Form.Label>
          <Form.Message className='FormMessage' match='valueMissing'>
            Please enter the password
          </Form.Message>
        </div>
        <Form.Control asChild>
          <input
            className='Input'
            type='password'
            required
            onChange={(e) => debouncedHandler('password', e)}
          ></input>
        </Form.Control>
      </Form.Field>
      <Form.Submit asChild>
        <button className='Button' style={{ marginTop: 10 }}>
          Login
        </button>
      </Form.Submit>
    </Form.Root>
  );
};
