# Link Shortener
Implement link shortener in GO language

## Standards
* [Project layout structure](https://github.com/golang-standards/project-layout)

## Running
* **Setup in local env:**
    ```shell script
    make up MODE=local
    make down MODE=local
    ```

* **Run a service in local env:**
    ```shell script
    make run SERVICE=[service name]
    ```

* **build docker image of a service:**
    ```shell script
    make build SERVICE=[service name] VERSION=[version tag]
    ```

* **Setup in release env:**
    ```shell script
    make up MODE=release
    make down MODE=release
    ```