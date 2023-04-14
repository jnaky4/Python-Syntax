import docker

if __name__ == "__main__":

    client = docker.DockerClient(base_url='unix:///Users/Z004X7X/.colima/default/docker.sock')
    print(client.images.list())
