# GoAuth Project

## Environment Configuration

This project uses different environment variable files to distinguish between development and production configurations.

### Development Environment

In the development environment, the project uses the `.env` file by default. Create a `.env` file in the project root directory with the following content:

```
PORT=3000
DB_URL="host=localhost user=postgres password=your_password dbname=goauth_dev port=5432 sslmode=prefer"
SECRET=dev-secret-key
```

### Production Environment

In the production environment, the project uses the `.env.production` file. Create a `.env.production` file in the project root directory with the following content:

```
PORT=8080
DB_URL="host=production_host user=prod_user password=prod_password dbname=goauth_prod port=5432 sslmode=require"
SECRET=production-secret-key
```

Note: Ensure that you set a strong secret key for the production environment and keep it secure.

### Switching Environments

The project determines the current running environment through the `GO_ENV` environment variable:

- When `GO_ENV` is set to "production", it loads the `.env.production` file.
- When `GO_ENV` is not set or set to any other value, it loads the default `.env` file.

#### Running in Development Environment

```bash
go run main.go
```

#### Running in Production Environment

```bash
GO_ENV=production go run main.go
```
#### Check health of the service by:
<ip>:<port> - if it shows healthy service, then the main go service is running successfully. 



### Important Notes

1. Do not commit `.env` and `.env.production` files containing sensitive information to version control systems.
2. In production, it's recommended to use environment variables or secure configuration management systems to set these values rather than relying on the `.env.production` file.
3. Regularly update the SECRET key to enhance security.
4. Ensure HTTPS is used in the production environment to protect data transmission.
5. Adjust database connection settings as needed, ensuring appropriate security settings are used in the production environment.

## Configuration Item Descriptions

- `PORT`: The port number on which the application listens
- `DB_URL`: Database connection string
- `SECRET`: Key used for JWT signing

Please adjust these configurations according to your specific needs and security requirements.
