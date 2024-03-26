# Herder

Herder is a Command Line Tool (CLI) for managing microservices during local development

## Usage

The general structure of a herder command is as follows:

```
herder <project> <command> <command>...
```

The following are valid commands:

- `help`
- `verify`
- `clone`
- `pull`
- `build`
- `run`
- `stop`

Additionally, the following flags can be applied:

- `-include service1,service2`
- `-exclude service1,service2`

For example, if you wanted to build and start all services in project `foo`, excluding service `bar`:


```
herder foo build run -exclude bar
```