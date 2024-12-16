# inboxed

## Descricpión

Inboxed es un sistema de búsqueda y visualización de datos que utiliza ZincSearch como motor de búsqueda, un backend desarrollado en Go, y un frontend implementado en Vue.js. Proporciona funcionalidades de filtrado y ordenamiento para explorar grandes conjuntos de datos.

## Despliegue

Disponible en http://ec2-18-231-252-104.sa-east-1.compute.amazonaws.com/

## Tecnologías utilizadas

- **Backend**: Go
    - Frameworks y librerías destacadas: Chi
- **Frontend**: Vue.js - TypeScript
    - Estilos: TailwindCSS
- **Base de datos**: ZincSearch
- **Infraestructura**: Docker y AWS EC2

## Instalación

1. Clona el repositorio.
2. Crea las variables de entorno. Ej: API_PORT=":8080"
3. Descarga los registros de [Enron](http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz)
4. Crea la carpeta 'data' y extrae los ficheros ahí.
5. Construye y ejecuta el proyecto: 'docker compose up --build'
6. Prueba la aplicación en tu navegador con 'localhost'

## Disclaimer:

- Para el favicon utilicé una imagen de [kropekk_pl](https://pixabay.com/es/users/kropekk_pl-114936/)
