# WebSocket Client in Go

## Overview

This project is a WebSocket client implemented in Go. It uses the Gorilla WebSocket package and Cobra for command-line interface. The client can establish a WebSocket connection to a server, send messages, and receive messages.

## How to Build and Use the Client

1. **Build the Client**

    To build the client, navigate to the project directory and run the following command:

    ```bash
    go build main.go
    ```

    This will create an executable file named `main` in the same directory.

2. **Use the Client**

    To use the client, you need to run the executable with the `--endpoint` flag set to your WebSocket server URL and the `--headers` flag set to your headers. Here's an example:

    ```bash
    ./main --endpoint "wss://localhost:8080/api/v1/blog/draft/6564321" --headers "Authorization=Bearer ShD21O3OcsWdYo8Z68Zrmui4cbRrKv9Kpz4hY5M6Tc4ZXYWKRLFKqa7lZ0bPzjyU"
    ```

    You can then enter messages at the prompt to send them to the server. Enter 'exit' to quit the client.

## Contributing to the Project

We welcome contributions to this project. Here are the steps to get started:

1. **Fork the Repository**

    Start by forking the repository to your own GitHub account.

2. **Clone the Repository**

    Next, clone the repository to your local machine:

    ```bash
    git clone https://github.com/<your-username>/websocket-client-go.git
    ```

3. **Create a New Branch**

    Navigate to your cloned repository and create a new branch:

    ```bash
    git checkout -b <branch-name>
    ```

4. **Make Your Changes**

    Now you can make your changes to the code.

5. **Commit and Push Your Changes**

    Once you're done, commit your changes and push them to your fork:

    ```bash
    git commit -am "Add some feature"
    git push origin <branch-name>
    ```

6. **Open a Pull Request**

    Finally, go to your fork on GitHub and open a pull request for your branch.

Please make sure your code adheres to our coding standards and passes all tests. If you're adding a new feature, please consider adding tests for it as well.
