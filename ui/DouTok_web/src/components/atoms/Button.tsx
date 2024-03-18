type ButtonProps = {
  buttonText: string;
} & Partial<HTMLButtonElement>;
export const Button = ({ buttonText }: ButtonProps): JSX.Element => {
  return <button>{buttonText}</button>;
};
