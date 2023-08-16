
# Instructions for Using the Receipts Docker Image

This guide will show you how to download and run the Docker image `crishdzdev/receipts:latest` in your local environment.

## Step 1: Download the Image

To get started, open a terminal and execute the following command to download the image from Docker Hub:

```bash
docker pull crishdzdev/receipts:latest
```

This will download the latest version of the `crishdzdev/receipts` image from the Docker Hub repository.

## Step 2: Run the Container

Once the image is downloaded, you can run a container based on it. Use the following command:

```bash
docker run -p 9095:9095 crishdzdev/receipts:latest
```

This command will run a container using the downloaded image and map the internal port `9095` of the container to port `9095` on your host machine. You can access the application from your web browser or tools like Postman using `http://localhost:9095`.

## Step 3: Stop and Remove the Container

To stop the container, go to the terminal and press `Ctrl + C`. This will halt the running container.

If you want to remove the container after stopping it, you can use the following command:

```bash
docker rm <container_id_or_name>
```

Replace `<container_id_or_name>` with the ID or name of the container you wish to remove.

---

# How to run application Without Downloading the Image:

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
   docker build -t receipts_by_fcph .
   ```

   This command builds an image named `receipts_by_fcph`.

5. **Run the Docker Container**:

   After the image is built, you can run a Docker container based on that image:

   ```sh
   docker run -p 9095:9095 receipts_by_fcph
   ```

   This will start the container and map port 9095 from the container to port 9095 on your host machine.

6. **Access the Application**:

   Open your web browser and navigate to `http://localhost:9095`. You should see your Go application running.

7. **Interact with the Application**:

   Depending on your Go application, you might need to interact with it through HTTP requests using tools like `curl` or by using a web browser.

8. **Stop the Container**:

   When you're done testing, you can stop the container. To do this, open a new terminal window (without closing the running container), and run:

   ```sh
   docker ps
   ```

   This will show you a list of running containers. Identify the container ID or name associated with your `receipts_by_fcph` container. Then, stop the container using:

   ```sh
   docker stop <container-id-or-name>
   ```

Replace `<container_id_or_name>` with the ID or name of the container you wish to remove.

---
