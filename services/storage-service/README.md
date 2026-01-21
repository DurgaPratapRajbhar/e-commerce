# Storage Service

A microservice for handling file uploads, storage, and serving in the e-commerce platform.

## Features

- File upload with validation
- Image processing (resize and thumbnails)
- File serving
- File deletion
- Category and subcategory organization
- No database required (filesystem-based)

## API Endpoints

### Public Endpoints
- `GET /health` - Health check
- `GET /public/serve/*filepath` - Serve public files
- `GET /public/thumbnail/*filepath` - Serve public thumbnails

### Protected Endpoints (require authentication)
- `POST /upload/:category/:subcategory` - Upload files
- `GET /serve/*filepath` - Serve files
- `GET /thumbnail/*filepath` - Serve thumbnails
- `DELETE /delete/*filepath` - Delete files

## Categories & Subcategories

- **products**
  - images (full, medium, thumbnails)
  - videos
- **categories**
  - banners
- **users**
  - avatars
- **documents**
  - invoices
  - receipts

## Environment Variables

```env
STORAGE_SERVICE_PORT=8009
STORAGE_UPLOAD_PATH=../../static-assets
APP_ENV=development
```

## Running the Service

```bash
cd services/storage-service
go mod tidy
go run cmd/main.go
```

## File Structure

The service stores files in the `static-assets` directory organized by category and subcategory.