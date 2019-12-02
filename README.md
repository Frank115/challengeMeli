# Challenge Meli
## Descripción
El challenge consistió en realizar una API REST de Items, en donde hubo que conectar la API con nginx, mysql y además agregarle algunos endpoints.

## Problemas que surgieron
* Hacer cambios el archivo docker-compose para que la api y el nginx escuchen en los puertos correctos y estén conectados en la misma red.
* En el nginx hubo que agregar en el location un proxy pass para que redirija el tráfico a la API.
* Cambios en las querys para adaptarlas de sqlite a mysql, y sus drivers en go.

## Cosas a mejorar
* Un tipo Error para poder discernir en el controller que tipo de error se encuentra y mostrar un status más acorde.
* Logear errores.
* El archivo docker-compose podría correr el schema.sql al correr la base, en vez de hacerlo en el código, que como ya estaba comentado, es mala práctica. (Lo intenté pero no pude :S )
