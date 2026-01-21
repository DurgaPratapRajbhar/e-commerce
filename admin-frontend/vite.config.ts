import { defineConfig, loadEnv } from 'vite';
import react from '@vitejs/plugin-react-swc';

export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, process.cwd(), '');
  const adminPort = parseInt(env.ADMIN_FRONTEND_PORT || '5173');

  
  return {
    plugins: [react()],
    server: {
      port: adminPort,
      host: true
    },
    build: {
      outDir: 'dist',
      sourcemap: mode === 'development',
      minify: mode === 'production'
    },
    optimizeDeps: {
      include: ['react', 'react-dom'],
    },
  };
});
