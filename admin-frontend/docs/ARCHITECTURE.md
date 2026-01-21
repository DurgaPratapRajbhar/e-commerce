# Architecture

## Overview

The admin-frontend follows a modern React architecture with TypeScript, Redux Toolkit for state management, and a component-based structure.

## Tech Stack

- **Framework**: React 19 with Vite
- **Language**: TypeScript
- **State Management**: Redux Toolkit
- **Styling**: Tailwind CSS
- **Routing**: React Router v7
- **HTTP Client**: Axios
- **UI Components**: Custom component library
- **Build Tool**: Vite

## Project Structure

```
src/
├── App.tsx                    # Main application component
├── main.tsx                   # Application entry point
├── assets/                    # Static assets
├── styles/                    # Global styles
├── config/                    # Configuration files
├── types/                     # TypeScript type definitions
├── lib/                       # Third-party library configurations
├── app/                       # Redux store configuration
├── features/                  # Redux slices organized by feature
├── services/                  # API service layer
├── context/                   # React Context providers
├── hooks/                     # Custom React hooks
├── utils/                     # Utility functions
└── components/                # Reusable UI components
```

## Features Organization

Each feature follows this structure:

```
features/{featureName}/
├── {featureName}Slice.ts      # Redux slice for state management
├── {featureName}API.ts        # API calls for the feature
├── {featureName}Thunks.ts     # Async Redux thunks
└── {featureName}Selectors.ts  # State selectors
```

## Component Organization

Components are organized by type:

- `common/`: Reusable UI components (Button, Input, etc.)
- `forms/`: Form components with validation
- `modals/`: Modal dialog components
- `filters/`: Filter and search components
- `editors/`: Rich text editors and complex inputs

## State Management

- **Global State**: Managed with Redux Toolkit in the `features/` directory
- **Local State**: Managed with React hooks within components
- **Context State**: Used for features that don't require persistence across the app

## API Layer

- **Services**: Located in `services/` directory, abstract API calls
- **API Configuration**: Centralized in `config/apiConfig.ts`
- **Error Handling**: Consistent error handling with `apiHandler.ts`

## Authentication Flow

1. User logs in via login form
2. Credentials sent to auth service
3. Server sets HttpOnly cookie with JWT
4. Subsequent requests include cookie automatically
5. Auth context manages user state
6. Protected routes check authentication status

## Environment Configuration

Environment variables are managed in:
- `.env.development`: Development settings
- `.env.production`: Production settings
- `config/env.ts`: Runtime environment configuration

## Internationalization

The application supports i18n through the `lib/i18n.ts` configuration, allowing for multiple language support.

## Testing Strategy

- **Unit Tests**: For utility functions and hooks
- **Component Tests**: For UI components
- **Integration Tests**: For API services and Redux slices
- **E2E Tests**: For critical user flows

## Performance Optimization

- Code splitting with React.lazy and Suspense
- Memoization with React.memo and useMemo
- Proper state normalization in Redux
- Image optimization and lazy loading
- Bundle size optimization with tree shaking

## Security Considerations

- HttpOnly cookies for JWT tokens
- Input validation and sanitization
- CSRF protection where applicable
- Secure API communication (HTTPS in production)
- Proper error message sanitization