import docker
from typing import Optional, Dict
import os
import platform
import time
from docker.errors import APIError, NotFound


# Docker Python SDK docs
# https://docker-py.readthedocs.io/en/stable/
# Docker Docs guide
# https://docs.docker.com/language/python/
# Docker Errors Example
# https://www.programcreek.com/python/example/107515/docker.errors
# TODO add Compose config to develop locally: https://docs.docker.com/language/python/develop/
# TODO add build/run container method
# TODO add save container state as image
# TODO add dockerfile/build from dockerfile: https://docs.docker.com/language/python/build-images/
# TODO configure CI/CD for docker application: https://docs.docker.com/language/python/configure-ci-cd/
# TODO add image and container class for type reference:
#   Image Object https://docker-py.readthedocs.io/en/stable/images.html#image-objects
#   Container Object https://docker-py.readthedocs.io/en/stable/containers.html#container-objects
# TODO add multiple Databases on start: https://github.com/mrts/docker-postgresql-multiple-databases
# TODO add test query to confirm database is running line 183


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
        print(f"Container Status: {container_state['Status']}")

    except NotFound as e:
        print(f"Container Doesn't Exists")
        print(f"Error: {e}")
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
        print(f"Container {identification} Exists ID: {exists.id}")
        return True
    except NotFound as e:
        print(f"Container Doesn't Exists: Error Generated: {e}")
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
        print(f"Image {image_name} Exists: {image.id}")
        return True
    except NotFound as e:
        print(f"Image Doesn't Exists: Error Generated: {e}")
    except Exception as e:
        print(f"Exception: {e}")
    finally:
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
            print(f"Pulling Image {image_name}, may take a minute")
            image_id = pull_image(image_name)
            if image_id is None:
                print("Image does not exist as a repository to pull from, try another name")
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
                
            Postgres Environment Variables:
            https://github.com/docker-library/docs/tree/master/postgres
            POSTGRES_DB
                This optional environment variable can be used to define a different name for the default database that 
                is created when the image is first started. If it is not specified, then the value of POSTGRES_USER will 
                be used.
                
            POSTGRES_PASSWORD
                This environment variable is required for you to use the PostgreSQL image. It must not be 
                empty or undefined. This environment variable sets the superuser password for PostgreSQL. 
                The default superuser is defined by the POSTGRES_USER environment variable.
            """

            container = client.containers.run(
                image_name,
                detach=True,
                name=container_name,
                ports={'5432/tcp': 5432},
                environment={"POSTGRES_PASSWORD": "pokemon", "POSTGRES_DB": "Pokemon"}
            )

            """
            Similar to the pg_isready command: https://www.postgresql.org/docs/9.3/app-pg-isready.html
            waits until the server is accepting connections    
            """
            timeout = 10
            current_time = 0
            exit_code = -2
            while exit_code != 0 and current_time != timeout:
                returned = container.exec_run(f"pg_isready")
                print(returned)
                if len(returned) > 0:
                    exit_code = returned[0]
                time.sleep(1)
                current_time += 1

            print(f"Container {container} Started")

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

