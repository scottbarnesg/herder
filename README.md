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

## Configuration

Herder reads is configuration from `~/.herder/config.yml`. Below is an example configuration file:

```yaml
projects:
  - name : DemoProject
    path: ~/Documents/code/demo-project
    services:
    - name: Demo Service 1
      path: herder1
      build-command: make build
      run-command: ./herder help
      source: https://github.com/scottbarnesg/herder.git
    - name: Demo Service 2
      path: herder2
      build-command: make build
      run-command: ./herder help
      source: https://github.com/scottbarnesg/herder.git
```

Note: The path to a project or service can be set via env var (see below). This is not supported for other values.

```yaml
projects:
  - name : DemoProject
    path: ${DEMO_PROJECT_PATH}
    ...
```