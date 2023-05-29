# GO Ps Script

A short script in order to check how many processes are running inside a docker - golang container.

## Instructions

1. Get inside the container:
    ```
    $ docker exec -it <your-container> bash
    ```

2. Install the `vim` dependency:
    ```
    $ apt-get update
    $ apt-get install -y vim
    ```

3. Create the `main.go` file:
    ```
    $ mkdir go-ps-script
    $ cd go-ps-script
    $ vi main.go
    ```

4. Copy paste the content of the `main.go` file.

5. Build the `main.go` file:
    ```
    $ go mod init
    $ go get github.com/mitchellh/go-ps
    $ go mod tidy
    $ go build main.go
    ```

6. Run the binary:
    ```
    ./main
    ```

## Example output

```
2023/05/29 11:36:08 1	sleep
2023/05/29 11:36:08 7	bash
2023/05/29 11:36:08 573	main
2023/05/29 11:38:32 Total Processes: 3
```