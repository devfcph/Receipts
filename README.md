How to run application:

1. **Install Docker**:

   If Docker isn't already installed on your system, you'll need to install it. Go to the official Docker website and follow the instructions for your operating system: [Get Docker](https://docs.docker.com/get-docker/)

2. **Clone the Project**:

   Open a terminal and navigate to a directory where you want to store your projects. Clone the project repository from wherever it's hosted (GitHub, GitLab, etc.) using a command like:

   ```sh
   git clone https://github.com/devfcph/Receipts
   ```

3. **Navigate to Project Directory**:

   Use the terminal to navigate into the project directory you just cloned:

   ```sh
   cd <project-directory>
   ```

4. **Build the Docker Image**:

   Run the following command to build the Docker image from the project directory:

   ```sh
   docker build -t receipt-processor .
   ```

   This command builds an image named `receipt-processor`.

5. **Run the Docker Container**:

   After the image is built, you can run a Docker container based on that image:

   ```sh
   docker run -p 8080:8080 receipt-processor
   ```

   This will start the container and map port 8080 from the container to port 8080 on your host machine.

6. **Access the Application**:

   Open your web browser and navigate to `http://localhost:8080`. You should see your Go application running.

7. **Interact with the Application**:

   Depending on your Go application, you might need to interact with it through HTTP requests using tools like `curl` or by using a web browser.

8. **Stop the Container**:

   When you're done testing, you can stop the container. To do this, open a new terminal window (without closing the running container), and run:

   ```sh
   docker ps
   ```

   This will show you a list of running containers. Identify the container ID or name associated with your `receipt-processor` container. Then, stop the container using:

   ```sh
   docker stop <container-id-or-name>
   ```

That's it! You've successfully set up and run your Dockerized Go application on your new environment. This approach keeps your local environment clean and isolated while allowing you to run your application consistently across different systems.
