import docker
from typing import Optional, Dict
import os
import platform

# Docker Python SDK docs
# https://docker-py.readthedocs.io/en/stable/
# Docker Docs guide
# https://docs.docker.com/language/python/
# TODO add Compose config to develop locally: https://docs.docker.com/language/python/develop/
# TODO add build/run container method
# TODO add save container state as image
# TODO add dockerfile/build from dockerfile: https://docs.docker.com/language/python/build-images/
# TODO configure CI/CD for docker application: https://docs.docker.com/language/python/configure-ci-cd/
# TODO add image and container class for type reference:
#   Image Object https://docker-py.readthedocs.io/en/stable/images.html#image-objects
#   Container Object https://docker-py.readthedocs.io/en/stable/containers.html#container-objects

def is_container_running(container_name: str) -> Optional[bool]:
    container_running = False
    try:
        """
        To talk to a Docker daemon, you first need to instantiate a client.
        You can use from_env() to connect using the default socket or the configuration in your environment:
        """
        docker_client = docker.client.from_env()
        # Grab the container by name
        docker_container = docker_client.containers.get(container_name)
        # get the state of the container
        container_state = docker_container.attrs['State']
        container_running = container_state['Status'] == 'running'
        print(f"Container State: {container_running}")

    except Exception as e:
        print(f"Error: {e}")
    finally:
        return container_running


def container_exists(container_name: Optional[str] = None, container_id: Optional[str] = None) -> Optional[bool]:
    if container_name is None and container_id is None:
        print("Please pass in an Id or name")
        return False

    exists = False
    docker_client = docker.client.from_env()

    identification = container_name if container_id is None else container_id
    try:
        exists = docker_client.containers.get(identification)
        print(f"Container ID: {exists.id}")
        return True
    except Exception as e:
        print(f"Exception: {e}")
    return exists


def image_exists(image_name: str) -> Optional[bool]:
    docker_client = docker.client.from_env()
    try:
        """
        get(name)
        Gets an image.
        
        Parameters:	
        name (str) – The name of the image.
        
        Returns:	
        The image.
        
        Return type:	
        (Image)
        
        Raises:	
        docker.errors.ImageNotFound – If the image does not exist.
        docker.errors.APIError – If the server returns an error.

        """
        image = docker_client.images.get(image_name)
        print(image.id)
        return True
    except Exception as e:
        print(f"Exception: {e}")
        return False


def pull_image(image_name: str) -> Optional[str]:
    try:
        client = docker.from_env()
        image = client.images.pull(image_name)
        print(f"Pulled Image {image_name}, id:{image.id}")
        return image.id
    except Exception as e:
        print(f"Exception: {e}")
        return None


# TODO not done
def save_container_state(container_name: Optional[str] = None, container_id: Optional[str] = None):
    identification = container_name if container_id is None else container_id
    pass


def run_container(image_name: str, container_name: str):

    client = docker.from_env()
    os_name = os.name
    """
    platform.system output
    Linux: Linux
    Mac: Darwin
    Windows: Windows
    """
    os_platform = platform.system()
    platform_version = platform.release()
    print(f"""
    Computer Operating System: {os_name}
    Computer Platform: {os_platform}
    Platform Version: {platform_version}
    """)

    # does container exists on system
    already_exists = container_exists(container_name)

    if already_exists:
        # is container running already?
        container_running = is_container_running(container_name)
        # container is already running on system, do we want to restart container? or do nothing?
        if container_running:
            print(f"container {container_name} is running")
            return
        # container isnt running, run container
        else:
            client.containers.run(container_name)
            pass
    else:
        # image doesnt exist, pull from repo
        if not image_exists(image_name):
            image_id = pull_image(image_name)
            if image_id is None:
                return
        try:
            """
            image (str) – The image to run.
            name (str) – The name for this container.
            ports (dict) –
                Ports to bind inside the container.
                The keys of the dictionary are the ports to bind inside the container, 
                either as an integer or a string in the form port/protocol, 
                where the protocol is either tcp, udp, or sctp.
        
                The values of the dictionary are the corresponding ports 
                to open on the host, which can be either:
                
                    The port number, as an integer. For example:
                        {'2222/tcp': 3333} will expose port 2222 inside the container as port 3333 on the host.
                    
                    None, to assign a random host port. For example:
                        {'2222/tcp': None}.
                    A tuple of (address, port) if you want to specify the host interface. For example: 
                        {'1111/tcp': ('127.0.0.1', 1111)}.
                    A list of integers, if you want to bind multiple host ports to a single container port. For example: 
                        {'1111/tcp': [1234, 4567]}.
                Incompatible with host network mode.
            """

            client.containers.run(image_name, detach=True, name=container_name, ports={'5432/tcp': 5432}, environment={"POSTGRES_PASSWORD":"pokemon"})
        except Exception as e:
            print(f"Exception: {e}")

        pass







def stop_all_containers():
    client = docker.from_env()
    for container in client.containers.list():
        container.stop()


def stop_specific_containers(*container_ids):
    client = docker.from_env()
    try:
        for ids in container_ids:
            container = client.containers.get(ids)
            container.stop()
    except Exception as e:
        print(f"Exception: {e}")


def list_all_containers() -> Dict:
    all_containers = {}
    client = docker.from_env()
    for container in client.containers.list():
        all_containers[container.id] = container
        print(container.id)
    return all_containers


def print_container_logs(container_name: Optional[str] = None, container_id: Optional[str] = None):
    identification = container_name if container_id is None else container_id
    client = docker.from_env()
    container = client.containers.get(identification)
    print(container.logs())






if __name__ == "__main__":
    name = "pokemon-postgres"
    print(f"Container Exists: {container_exists(name)}")
    print(f"Container Running: {is_container_running(name)}")

    # run_container("hello", "hello")

