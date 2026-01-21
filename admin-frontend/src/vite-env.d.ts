/// <reference types="vite/client" />

/// <reference types="vite/client" />

interface ImportMetaEnv {
    readonly VITE_APP_ADMIN_FRONTED_PORT: string;
    readonly VITE_APP_PRODUCT_API_PORT: string;
  }
  
  interface ImportMeta {
    readonly env: ImportMetaEnv;
  }
  