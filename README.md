# (wscat)WebSocket Client in Go

## Overview
![alt text](document\wsock_logo.png?raw=true)
This project <b>wscat</b> is a WebSocket client implemented in Go. It uses the Gorilla WebSocket package and Cobra for command-line interface. The client can establish a WebSocket connection to a server, send messages, and receive messages.

## How to Build and Use the Client

1. **Download the Release**

    You can download the latest release of this project from the releases page. There are source code archives available in ZIP and TAR formats.


2. **Build the Client**

    To build the client, navigate to the project directory and run the following command:

    ```bash
    go build main.go
    ```

    This will create an executable file named `main` in the same directory.

3. **Use the Client**

    To use the client, you need to run the executable with the `--endpoint` flag set to your WebSocket server URL and the `--headers` flag set to your headers. Here's an example:

    ```bash
    ./main --endpoint "wss://localhost:8080/api/v1/blog/draft/6564321" --headers "Authorization=Bearer ShD21O3OcsWdYo8Z68Zrmui4cbRrKv9Kpz4hY5M6Tc4ZXYWKRLFKqa7lZ0bPzjyU"
    ```

    You can then enter messages at the prompt to send them to the server. Enter 'exit' to quit the client.

