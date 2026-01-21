# Setup Guide

## Prerequisites

- Node.js (v18 or higher)
- npm or yarn
- Git

## Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd admin-frontend
   ```

2. Install dependencies:
   ```bash
   npm install
   # or
   yarn install
   ```

3. Create environment files:
   ```bash
   cp .env.example .env.development
   cp .env.example .env.production
   ```

4. Update environment variables as needed

## Development

1. Start the development server:
   ```bash
   npm run dev
   # or
   yarn dev
   ```

2. Open your browser to `http://localhost:4173`

## Production

1. Build the application:
   ```bash
   npm run build
   # or
   yarn build
   ```

2. Serve the built files:
   ```bash
   npm run preview
   # or
   yarn preview
   ```

## Docker

1. Build the Docker image:
   ```bash
   docker build -t admin-frontend .
   ```

2. Run the container:
   ```bash
   docker run -p 4173:80 admin-frontend
   ```

## Environment Variables

- `REACT_APP_API_BASE_URL`: API base URL
- `REACT_APP_API_VERSION`: API version (default: v1)
- `REACT_APP_API_TIMEOUT`: Request timeout in ms
- `REACT_APP_JWT_COOKIE_NAME`: JWT cookie name

## API Gateway Configuration

The frontend is designed to work with an API gateway. Update your gateway configuration to forward requests:

```
/api/v1/auth/* → auth-service:8001
/api/v1/product/* → product-service:8002
/api/v1/user/* → user-service:8003
# ... etc
```