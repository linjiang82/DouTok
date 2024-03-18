import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
import { TanStackRouterVite } from '@tanstack/router-vite-plugin';

// https://vitejs.dev/config/
export default defineConfig(({ mode }) => {
  return {
    apiBaseUrl: mode === 'development' ? 'http://localhost:8088' : 'a todo',
    plugins: [react(), TanStackRouterVite()],
  };
});
