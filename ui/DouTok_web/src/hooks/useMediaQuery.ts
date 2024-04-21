import { useState } from 'react';

export const breakpoints = {
  xs: 375,
  sm: 640,
  md: 768,
  legacymobile: 960,
  lg: 1024,
  xl: 1280,
  '2xl': 1536,
};

export const useMediaQuery = () => {
  const width = globalThis.innerWidth;
  const [currentWidth, setCurrentWidth] = useState(width);

  window.onresize = () => {
    setCurrentWidth(window.innerWidth);
  };

  const lg = currentWidth >= breakpoints.lg && currentWidth < breakpoints.xl;
  const sm = currentWidth >= breakpoints.sm && currentWidth < breakpoints.md;
  const xs = currentWidth >= breakpoints.xs && currentWidth < breakpoints.sm;

  return { currentWidth, lg, sm, xs };
};
