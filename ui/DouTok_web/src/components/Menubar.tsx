import { Link, Outlet } from '@tanstack/react-router';
import { TanStackRouterDevtools } from '@tanstack/router-devtools';
import { useEffect, useState } from 'react';
import { useMediaQuery } from '../hooks/useMediaQuery';

const routes: Array<[string, string]> = [
  ['Home', '/'],
  ['Log In', '/login'],
  ['ABC', '/login1'],
  ['DEF', '/login2'],
  ['XYZ', '/login3'],
];

export const Menubar = (): JSX.Element => {
  const [menuOpen, setMenuOpen] = useState(false);
  const { xs } = useMediaQuery();

  const toggleMenuStatus = () => {
    if (xs) {
      setMenuOpen((prev) => !prev);
    }
  };

  useEffect(() => {
    return setMenuOpen(false);
  }, [xs]);

  return (
    <>
      <div className='absolute left-0 top-0'>
        <div
          className={`flex flex-col gap-2 transition-all duration-1000 sm:flex-row ${menuOpen ? 'open group max-h-48 overflow-hidden' : 'max-h-8'} mx-4 my-4 `}
          onClick={toggleMenuStatus}
        >
          {routes.map(([title, url]) => (
            <Link
              key={title}
              to={url}
              className='[&.active]:white hidden group-[.open]:block md:block [&.active]:block [&.active]:bg-blue-50'
            >
              {title}
            </Link>
          ))}
        </div>
        <hr />
        <Outlet />
      </div>
      <TanStackRouterDevtools />
    </>
  );
};
