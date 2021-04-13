<img src="/docs/resources/go-notes-images.png" align="right"/>

# Go-Notes

![GitHub release (latest by date)](https://img.shields.io/github/v/release/Golang-EC/go-notes?style=social)
[![codecov](https://codecov.io/gh/Golang-EC/go-notes/branch/master/graph/badge.svg?token=0K5Q36Y0C4)](https://codecov.io/gh/Golang-EC/go-notes)


> Un proyecto del [curso-tdd](https://jj.github.io/curso-tdd) o curso de desarrollo ágil.


Este proyecto es un MVP que resuelve el siguiente problema:

## [Notas](https://jj.github.io/curso-tdd/problemas/notas.html)

**Problema**: Almacenar las notas en diferentes prácticas de una asignatura, junto con comentarios adicionales y posibles anotaciones que se hayan podido hacer sobre las mismas.

**Solución**: Crear un API REST en el cual el docente pueda publicar las notas de sus alumnos junto a sus comentarios y anotaciones respectivos. Estas notas serán publicadas para que los estudiantes puedan revisarlas en cualquier momento mediante una web estática.

## Herramientas y tecnologías que usaremos

- Lenguaje de programación: [Golang](https://golang.org/)
- Servicio de Log: [log](https://golang.org/pkg/log/)
- DataFiles, archivos JSON para almacenamiento


## Milestones del proyecto

Hemos planificado los siguientes hitos para llevar a cabo un mvp del proyecto 
1. [Definiendo algunas clases y sus atributos](https://github.com/Golang-EC/go-notes/projects/1#column-13357999)
2. [Definiendo ciertas funcionalidades esenciales](https://github.com/Golang-EC/go-notes/projects/1#column-13358053)

## Ejecutando pruebas
Para echar a andar los test ejecutamos el comando mencionado acontinuación, el cual devolverá por terminal el resultado de dichas pruebas con sus aserciones. Esto hará ejecutar el comando (`go test`) para el archivo GoNotes_test.go .
```
make test
```

## Correr test de cobertura | code coverage

Podemos ejecutar los test de cobertura usando la herramienta interna de golang la cual esta configurada para ejecutarse en la siguiente tarea del make task runner.
```shell
make coverage
```
Esto mostrara un archivo de salida _coverage.out_

---

| Equipo                                               |  |
| ---------------------------------------------------- |----------|
| [Kevin Campuzano](https://github.com/Kevincamp)      | [![LinkedIn](https://img.shields.io/badge/-LinkedIn-222222?style=flat-square&logo=linkedin&logoColor=white&link=https://www.linkedin.com/in/kevin-campuzano-castillo-42294966/)](https://www.linkedin.com/in/kevin-campuzano-castillo-42294966/)
| [Richard Palacios](https://github.com/rpalaciosg)    | [![LinkedIn](https://img.shields.io/badge/-LinkedIn-222222?style=flat-square&logo=linkedin&logoColor=white&link=https://www.linkedin.com/in/richardpalaciosgarcia/)](https://www.linkedin.com/in/richardpalaciosgarcia/)
| [Leonardo Gomez](https://github.com/gomezgleonardob) | [![LinkedIn](https://img.shields.io/badge/-LinkedIn-222222?style=flat-square&logo=linkedin&logoColor=white&link=https://www.linkedin.com/in/gomezgleonardob//)](https://www.linkedin.com/in/gomezgleonardob/)

