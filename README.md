Diego Alexis Laris Valenzuela 200607

Modelo de Usuario:
Se define una estructura llamada User con campos como Name y Email. Esta estructura representa un registro de usuario en la base de datos.

Conexión a la Base de Datos:
La aplicación se conecta a una base de datos MySQL local llamada crud_app utilizando GORM, una biblioteca ORM (Object-Relational Mapping).

Migración de la Base de Datos:
La función db.AutoMigrate(&User{}) se asegura de que la estructura de la base de datos coincida con la estructura del modelo User. Si la tabla no existe, la crea.

Rutas para CRUD:
La aplicación define rutas HTTP para realizar operaciones CRUD:

-GET /users: Obtiene todos los usuarios de la base de datos y los muestra como respuesta JSON.

-GET /users/:id: Obtiene un usuario por su ID y lo muestra como respuesta JSON. Si no se encuentra, devuelve un error.

-POST /users: Crea un nuevo usuario en la base de datos utilizando datos proporcionados en una solicitud JSON.

-PUT /users/:id: Actualiza un usuario existente por su ID con datos proporcionados en una solicitud JSON. Si no se encuentra, devuelve un error.

-DELETE /users/:id: Elimina un usuario por su ID. Si no se encuentra, devuelve un error.

Manejo de Solicitudes y Respuestas:
Para las operaciones POST y PUT, se analiza el cuerpo de la solicitud JSON para obtener los datos del usuario. Luego, GORM se utiliza para interactuar con la base de datos y realizar las operaciones CRUD.

Ejecución del Servidor:
La aplicación se ejecuta en el puerto 8080, lo que permite que responda a las solicitudes HTTP entrantes.

En resumen, este código crea una API web que te permite realizar las operaciones básicas de creación, lectura, actualización y eliminación de registros de usuarios en una base de datos MySQL utilizando el framework Gin y la biblioteca GORM en Go.