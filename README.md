# Problema

Desarrollar un servicio que exponga una API REST para una lista de tareas (TODO) con endpoints para la creación, lectura, actualización y eliminación de tareas.

* El almacenamiento de tareas debe ser en memoria o en alguna base de datos como SQLite.
* La salida de cada endpoint debe ser en formato JSON.
* Armar el Dockerfile
* Opcional: Documentación (OpenAPI), casos de prueba y colección de Postman.

## Restricciones

* No utilizar frameworks ni ORMs.
* No utilizar bases de datos NoSQL.

## Instrucciones

Para correr utilizando docker, ejecutar:

```
docker-compose up
```

Utilizando el flag  `--build` en nuestra primer ejecucion

Para finalizar los contenedores:

```
docker-compose down
```

Adicionalmente, utilizar el flag `-v` para eliminar el volumen asociado.


Localmente (Se utilizo go 1.19):


```
go build -o tasks.api &&
./tasks.api
```

Ejecutar tests unitarios:

```
go test ./...
```

## Documentacion de API y Coleccion de Postman

La API se encuentra documentada en swaggerhub: https://app.swaggerhub.com/apis/PINTOTOMASE/task-api/1.0.0

La coleccion de Postman se puede encontrar en el repositorio para descargarla y poder ejecutar las requests.

## Comentarios acerca de la solucion

- El nombre de la base de datos, el driver, el puerto del server deberian moverse a variables de ambiente.

- Durante la ejecucion se mantiene una unica conexion con la base de datos, tambien es valido conectarnos solo cuando necesitamos realizar alguna query. Esto dependera de las necesidades de nuestra app (por ejemplo, si valoramos performance quizas no queremos tener el overhead de conectarnos muchas veces)


