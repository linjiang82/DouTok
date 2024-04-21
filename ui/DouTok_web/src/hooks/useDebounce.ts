let timer: number;

export const useDebounce = <T extends unknown[]>(
  fn: (...args: T) => void,
  delay: number
) => {
  const debouncedFn = (...args: Parameters<typeof fn>) => {
    if (timer) {
      clearTimeout(timer);
    }

    timer = setTimeout(() => {
      fn(...args);
    }, delay);
  };
  return debouncedFn;
};
