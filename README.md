# API de Productos fravega-tech

## Descripción
Esta API permite gestionar productos en una tienda en línea. Ofrece funcionalidades como:
- Crear productos
- Consultar productos con filtros
- Obtener productos por su ID
- Actualizar productos
- Eliminar productos

## Tecnologías
- Lenguaje: Go (Golang)
- Framework: Gin
- Base de Datos: MongoDB
- Documentación: Swagger
- Docker: Para configuración y despliegue de servicios

## Endpoints Principales

### Productos
- **GET /products**: Obtiene todos los productos activos o filtra según los parámetros de consulta.
    Ejemplo localhost:8080/v1/products?min_price=10&max_price=15&categories=posters&categories=deadpool
    categories es un array donde las condiciones se intersectan (&)
- **GET /products/{id}**: Obtiene un producto por su identificador.
- **POST /products**: Crea un nuevo producto.
- **PATCH /products/{id}**: Actualiza campos específicos de un producto.
- **DELETE /products/{id}**: Marca un producto como eliminado.

## Requisitos Previos

- Tener instalado Docker y Docker Compose.
- Puerto `8080` disponible en el host.

## Instalación y Ejecución

1. Clona el repositorio:
   ```bash
   git clone https://github.com/mp02/fravega-tech
   cd tu-repositorio
   ```

2. Levanta los servicios usando Docker Compose:
   ```bash
   docker-compose up --build
   ```

   Esto inicializará los siguientes servicios:
   - MongoDB
   - Contenedor para cargar datos iniciales (mongo-seed)
   - La aplicación en Go

3. Accede a la API:
   - Swagger: [http://localhost:8080/api/doc/index.html](http://localhost:8080/api/doc/index.html)
   - Endpoints directamente en [http://localhost:8080](http://localhost:8080).

## Documentación

La documentación de la API se encuentra disponible en Swagger:
- URL: [http://localhost:8080/api/doc/index.html](http://localhost:8080/api/doc/index.html)

Ejemplo de uso con Swagger:
1. Ve a la URL de Swagger.
2. Explora los endpoints disponibles.
3. Prueba las operaciones directamente desde la interfaz de Swagger.

## Variables de Entorno

El archivo `docker-compose.yml` configura las variables necesarias. Por defecto:

- **MongoDB**:
  - Usuario: `root`
  - Contraseña: `example`
  - Base de datos: `productsdb`

- **Aplicación**:
  - `MONGO_URI`: `mongodb://root:example@mongodb:27017/productsdb?authSource=admin`

## Ejemplo de Request y Response

### Crear Producto (POST /products)

#### Request Body:
```json
{
  "name": "Smartphone",
  "description": "A powerful smartphone with the latest features.",
  "price": 799.99,
  "categories": ["electronics", "phone"],
  "images_url": ["https://example.com/images/smartphone.jpg"]
}
```

#### Response:
```json
{
  "id": "64e7d9f4e8b1b6d9e0c77abc",
  "name": "Smartphone",
  "description": "A powerful smartphone with the latest features.",
  "price": 799.99,
  "categories": ["electronics", "phone"],
  "created_at": "2025-01-15T00:00:00Z",
  "updated_at": "2025-01-15T00:00:00Z",
  "is_deleted": false,
  "images_url": ["https://example.com/images/smartphone.jpg"]
}
```
