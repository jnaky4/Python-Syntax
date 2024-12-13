from flask import Flask, jsonify
import platform
import psutil
from datetime import datetime, timedelta
import time
import os
import cpuinfo
import docker
from docker.errors import APIError
from scapy.all import sniff
from collections import defaultdict
import socket
# import pcapy
import struct


def capture_packets(interface_name, duration=10):
    # Create a raw socket to capture packets
    with socket.socket(socket.AF_PACKET, socket.SOCK_RAW, socket.ntohs(0x0003)) as sock:
        # Bind the socket to the specified interface
        sock.bind((interface_name, 0))

        # Capture packets for the specified duration
        end_time = time.time() + duration
        packets = []

        while time.time() < end_time:
            # Receive a packet
            packet_data, _ = sock.recvfrom(65536)

            # Extract packet details
            eth_header = packet_data[:14]
            eth = struct.unpack('!6s6sH', eth_header)
            eth_protocol = socket.ntohs(eth[2])

            # Extract IP header and payload (assuming IPv4 for simplicity)
            if eth_protocol == 8:  # IPv4
                ip_header = packet_data[14:34]
                iph = struct.unpack('!BBHHHBBH4s4s', ip_header)

                # Create IP addresses from the packed data
                src_ip = socket.inet_ntoa(iph[8])
                dst_ip = socket.inet_ntoa(iph[9])

                # Extract payload size
                payload_size = len(packet_data) - 34  # Total packet length minus Ethernet and IP headers

                # Store packet details
                packets.append({
                    'src_ip': src_ip,
                    'dst_ip': dst_ip,
                    'payload_size': payload_size
                })

        return packets

def get_docker_container_info():
    try:
        client = docker.from_env()  # Initialize Docker client from environment
        container_info = []

        # Iterate over all containers
        for container in client.containers.list():
            info = {
                'id': container.id,
                'name': container.name,
                'image': container.image.tags[0] if container.image.tags else 'N/A',
                'status': container.status,
                'ports': container.ports,
                'created': container.attrs['Created'],
                'cpu_usage': container.stats(stream=False)['cpu_stats']['cpu_usage']['total_usage'] / 1e9,
                'memory_usage': container.stats(stream=False)['memory_stats']['usage'],
                'network': container.attrs['NetworkSettings']['Networks'],
            }
            container_info.append(info)

        return container_info

    except docker.errors.APIError as e:
        print(f"Error accessing Docker API: {e}")
        return None
    except Exception as e:
        print(f"An error occurred: {e}")
        return None


def get_docker_image_info():
    try:
        # Initialize Docker client
        docker_client = docker.from_env()

        # Fetch list of Docker images
        images = docker_client.images.list()

        # Prepare a list to store image information
        image_info_list = []

        # Iterate through each image and extract relevant information
        for image in images:
            image_info = {
                "id": image.id,
                "tags": image.tags,
                "short_id": image.short_id.split(':')[1],  # Extract short ID without "sha256:" prefix
                "labels": image.labels,
                "created": image.attrs['Created'],  # Created timestamp
                "size_mb": image.attrs['Size'] / (1024 * 1024)  # Convert size to MB
            }
            image_info_list.append(image_info)

        return image_info_list

    except docker.errors.APIError as e:
        print(f"Error fetching Docker image info: {e}")
        return []

def capture_network_security_metrics():
    metrics = defaultdict(int)

    def packet_callback(packet):
        # Example: Counting number of packets captured by protocol
        metrics['total_packets'] += 1
        if packet.haslayer('TCP'):
            metrics['tcp_packets'] += 1
        elif packet.haslayer('UDP'):
            metrics['udp_packets'] += 1
        elif packet.haslayer('ICMP'):
            metrics['icmp_packets'] += 1

        # Add more metrics based on your specific needs (e.g., intrusion attempts, anomalies)

    try:
        # Sniffing packets on interface 'eth0' for 10 seconds
        sniff(iface='wlp3s0', timeout=10, prn=packet_callback)

        # Example: Print captured metrics
        print("Network Security Metrics:")
        print(f"Total Packets: {metrics['total_packets']}")
        print(f"TCP Packets: {metrics['tcp_packets']}")
        print(f"UDP Packets: {metrics['udp_packets']}")
        print(f"ICMP Packets: {metrics['icmp_packets']}")

        # Add more metrics output based on your requirements

        return metrics

    except Exception as e:
        print(f"Error capturing network security metrics: {e}")
        return None


def get_domain_name(ip_address):
    try:
        # Perform reverse DNS lookup
        domain_name = socket.gethostbyaddr(ip_address)[0]
        return domain_name
    except socket.herror as e:
        print(f"Error: {e}")
        return None

# def capture_packets(interface):
#     try:
#         cap = pcapy.open_live(interface, 65536, True, 100)
#         # Do packet capturing and processing here
#     except pcapy.PcapError as e:
#         print(f"Error capturing packets: {e}")



#TODO Get tuples from _common.py
def psutil_stats():
    return {
        # "fans": psutil.sensors_fans(),

        'cpu': {
            "cores": psutil.cpu_count(),
            'load_average (min : %)': {
                "1": psutil.getloadavg()[0] if hasattr(psutil, 'getloadavg') else None,
                "5": psutil.getloadavg()[1] if hasattr(psutil, 'getloadavg') else None,
                "15": psutil.getloadavg()[2] if hasattr(psutil, 'getloadavg') else None,
            },
            "stats": {
                "ctx_switches": psutil.cpu_stats().ctx_switches if hasattr(psutil.cpu_stats(), 'ctx_switches') else None,
                "interrupts": psutil.cpu_stats().interrupts if hasattr(psutil.cpu_stats(), 'interrupts') else None,
                "soft_interrupts": psutil.cpu_stats().soft_interrupts if hasattr(psutil.cpu_stats(), 'soft_interrupts') else None,
                "syscalls": psutil.cpu_stats().syscalls if hasattr(psutil.cpu_stats(), 'syscalls') else None
            },
            "info": cpuinfo.get_cpu_info(),
            'processes': {
                proc.name(): {
                    'pid': proc.pid,
                    'cpu_percent': proc.cpu_percent(),
                    'memory_percent': proc.memory_percent(),
                    'status': proc.status(),
                    'create_time': proc.create_time()
                } for proc in
                psutil.process_iter(['pid', 'name', 'cpu_percent', 'memory_percent', 'status', 'create_time'])
            },
            "times": {
                "per_cpu (sec)": [
                    {
                        "user": cpu_times.user,
                        "system": cpu_times.system,
                        "idle": cpu_times.idle,
                        "nice": cpu_times.nice if hasattr(cpu_times, 'nice') else None,
                        "iowait": cpu_times.iowait if hasattr(cpu_times, 'iowait') else None,
                        "irq": cpu_times.irq if hasattr(cpu_times, 'irq') else None,
                        "softirq": cpu_times.softirq if hasattr(cpu_times, 'softirq') else None,
                        "steal": cpu_times.steal if hasattr(cpu_times, 'steal') else None,
                        "guest": cpu_times.guest if hasattr(cpu_times, 'guest') else None,
                        "guest_nice": cpu_times.guest_nice if hasattr(cpu_times, 'guest_nice') else None
                    } for cpu_times in psutil.cpu_times(percpu=True)
                ]
            },
            "freq (Mhz)": {
                "curr": int(psutil.cpu_freq()[0]),
                "min": int(psutil.cpu_freq()[1]),
                "max": int(psutil.cpu_freq()[2]),

            },
        },
        # sane as file systems
        # 'disks': {
        #     disk.device: {
        #         "total": psutil.disk_usage(disk.mountpoint).total,
        #         "used": psutil.disk_usage(disk.mountpoint).used,
        #         "free": psutil.disk_usage(disk.mountpoint).free,
        #         "used %": psutil.disk_usage(disk.mountpoint).percent
        #     } for disk in psutil.disk_partitions()
        # },
        "disk": {
            'io': {
                disk: {
                    "read_count": counters.read_count,
                    "write_count": counters.write_count,
                    "read_bytes": counters.read_bytes,
                    "write_bytes": counters.write_bytes,
                    "read_time": counters.read_time,
                    "write_time": counters.write_time
                } for disk, counters in psutil.disk_io_counters(perdisk=True).items()
            },
            'file_systems': {
                partition.mountpoint: {
                    'device': partition.device,
                    'fstype': partition.fstype,
                    'opts': partition.opts,
                    'total': psutil.disk_usage(partition.mountpoint).total,
                    'used': psutil.disk_usage(partition.mountpoint).used,
                    'free': psutil.disk_usage(partition.mountpoint).free,
                    'percent': psutil.disk_usage(partition.mountpoint).percent
                } for partition in psutil.disk_partitions()
            },
        },
        "docker": {
            "containers": {
                container["name"]: container for container in get_docker_container_info()
            },
            "images": {
                image["tags"][0]: image for image in get_docker_image_info()
            },
        },
        "kubernetes": {}, #todo
        "vm": {}, #todo
        "alerts": {
        #     container error states
        #     hw ulitilization limits

        }, #todo
        "memory": {
            'ram (gb)': {  # todo in b Mb Gb?
                "total": round(psutil.virtual_memory().total / (1024 ** 3), 2),
                "available": round(psutil.virtual_memory().available / (1024 ** 3), 2),
                "percent": psutil.virtual_memory().percent,
                "used": round(psutil.virtual_memory().used / (1024 ** 3), 2),
                "free": round(psutil.virtual_memory().free / (1024 ** 3), 2),
                "active": round(psutil.virtual_memory().active / (1024 ** 3), 2) if hasattr(psutil.virtual_memory(), 'active') else None,
                "inactive": round(psutil.virtual_memory().inactive / (1024 ** 3), 2) if hasattr(psutil.virtual_memory(), 'inactive') else None,
                "buffers": round(psutil.virtual_memory().buffers / (1024 ** 3), 2) if hasattr(psutil.virtual_memory(), 'buffers') else None,
                "cached": round(psutil.virtual_memory().cached / (1024 ** 3), 2) if hasattr(psutil.virtual_memory(), 'cached') else None,
                "wired": round(psutil.virtual_memory().wired / (1024 ** 3), 2) if hasattr(psutil.virtual_memory(), 'wired') else None,
                "shared": round(psutil.virtual_memory().shared / (1024 ** 3), 2) if hasattr(psutil.virtual_memory(), 'shared') else None
            },
            'swap (bytes)': {
                "total": psutil.swap_memory().total,
                "used": psutil.swap_memory().used,
                "free": psutil.swap_memory().free,
                "percent": psutil.swap_memory().percent,
                "sin": psutil.swap_memory().sin,
                "sout": psutil.swap_memory().sout
            },
        },
        "network": {
            'connections': {
                conn.pid: {
                    'fd': conn.fd,
                    'family': conn.family,
                    'type': conn.type,
                    'laddr': {
                        "ip": conn.laddr.ip,
                        "port": conn.laddr.port
                    },
                    'raddr': conn.raddr,
                    'status': conn.status,
                } for conn in psutil.net_connections()
                if conn.pid is not None  # Filter out connections with no associated PID
            },
            'io': {
                interface: {
                    "bytes_sent": counters.bytes_sent,
                    "bytes_recv": counters.bytes_recv,
                    "packets_sent": counters.packets_sent,
                    "packets_recv": counters.packets_recv,
                    "errin": counters.errin,
                    "errout": counters.errout,
                    "dropin": counters.dropin,
                    "dropout": counters.dropout
                } for interface, counters in psutil.net_io_counters(pernic=True).items()
            },
            "interfaces": {
                interface: {
                    'family': label.family,
                    'address': label.address,
                    'netmask': label.netmask,
                    'broadcast': label.broadcast,
                    'ptp': label.ptp
                } for interface, labels in psutil.net_if_addrs().items() for label in labels
            },
        },
        "platform": {
            "system": platform.system(),
            "machine": platform.machine(),
            "processor": platform.processor(),
            "version": platform.version(),
            "python": platform.python_version(),
            # "java": platform.java_ver(),
            "release": platform.release(),
            'kernel_version': platform.uname().release
        },
        'temps': psutil.sensors_temperatures(fahrenheit=True),
        "time": {
            "up": f"{timedelta(seconds=int(time.time() - psutil.boot_time()))}",
            "boot": datetime.fromtimestamp(psutil.boot_time()).strftime("%Y-%m-%d %H:%M:%S"),
        },
        'users': {
                user[0]: {
                    "terminal": user[1],
                    "host": user[2],
                    "started": f"{timedelta(seconds=int(time.time() - user[3]))} ago"
                } for user in psutil.users()
        },
    }


# todo
# voltage
# go alerts
# temperatures
# kubernetes

app = Flask(__name__)


@app.route('/', methods=['GET'])
def get_psutil():
    return jsonify(psutil_stats())


if __name__ == '__main__':
    # capture_network_security_metrics()
    print(capture_packets("wlp3s0", 1))
    # capture_packets("wlp3s0")

    # app.run(debug=True)


"""
    Prometheus Client for Python:
        Description: Prometheus is a popular monitoring and alerting toolkit. Using its Python client library, you can expose custom metrics from your application.
        Use Case: Integrating with Prometheus allows you to collect time-series data on various aspects of your application and system, enabling detailed monitoring and alerting.

    Grafana for Visualization:
        Description: Grafana is a visualization tool commonly used with Prometheus for creating dashboards and graphs based on metrics collected.
        Use Case: Grafana can help visualize the metrics collected by Prometheus, providing insights into system performance and resource utilization.

    pyVmomi (VMware SDK for Python):
        Description: If you're running VMware infrastructure, pyVmomi allows you to interact with VMware's vSphere API to monitor and manage virtual machines (VMs) and hosts.
        Use Case: Monitor VM-level metrics such as CPU usage, memory allocation, and disk I/O performance directly from VMware infrastructure.

    Python Kubernetes Client (client-python):
        Description: This library provides a Python client for interacting with the Kubernetes API. It allows you to fetch metrics about pods, nodes, deployments, and more.
        Use Case: If your application runs in a Kubernetes cluster, this library enables monitoring Kubernetes-specific metrics, such as pod resource utilization and cluster health.

    Elasticsearch and Logstash:
        Description: Elasticsearch and Logstash, often used together with Kibana (ELK stack), provide a robust platform for centralized logging and monitoring.
        Use Case: Logstash can collect logs and metrics from various sources, Elasticsearch can store and index them, and Kibana can visualize the data. It's particularly useful for log analysis and troubleshooting.

Advanced Metrics to Consider:

    JVM Metrics:
        If your application runs on Java, consider monitoring JVM metrics such as heap usage, garbage collection times, and thread counts. Libraries like JMX can help expose these metrics.

    Database Performance Metrics:
        Track database-specific metrics such as query execution times, cache hit rates, and connection pool usage. Database-specific libraries (e.g., psycopg2 for PostgreSQL, sqlalchemy for ORM operations) can provide insights into database performance.

    Custom Application Metrics:
        Identify critical business metrics within your application (e.g., number of transactions processed per second, user sign-ups) and expose them for monitoring and alerting purposes using Prometheus or custom logging.

    Hardware Metrics:
        Utilize platform-specific libraries or tools to monitor hardware metrics such as CPU temperature, fan speed, and voltage levels if supported by your hardware.

    Network Security Metrics:
        Monitor network security metrics like intrusion attempts, firewall rules hit, and traffic anomalies using specialized tools or libraries that integrate with your security infrastructure.
1. Suricata

    Description: Suricata is an open-source Intrusion Detection and Prevention System (IDS/IPS) that can monitor network traffic for signs of attacks or anomalies.
    Key Features: Suricata provides real-time intrusion detection, traffic analysis, and packet capture capabilities.

2. Snort

    Description: Snort is another popular open-source IDS/IPS system known for its rule-based detection engine and extensive rule set.
    Key Features: Snort can analyze network traffic, detect various types of attacks based on predefined rules, and generate alerts.

3. Bro/Zeek

    Description: Bro, now known as Zeek, is a powerful network analysis framework that captures packets and generates detailed logs for analysis.
    Key Features: Zeek provides protocol analysis, traffic logging, and can be extended with scripts for custom network security monitoring.

4. PyShark

    Description: PyShark is a Python wrapper for the popular Wireshark packet analysis tool, allowing programmable packet parsing and analysis.
    Key Features: PyShark enables Python scripts to read and analyze packet captures, making it suitable for custom network security monitoring solutions.

5. pcapy

    Description: pcapy is a Python extension module that allows access to libpcap packet capture library functions.
    Key Features: pcapy facilitates packet capture and filtering, making it suitable for building network monitoring and analysis tools.

6. Python-nmap

    Description: Python-nmap is a Python library for network exploration and security auditing.
    Key Features: Python-nmap wraps Nmap, a popular network scanning tool, enabling scripted network discovery, vulnerability detection, and security auditing.

7. Tcpreplay

    Description: Tcpreplay is a suite of utilities for editing and replaying previously captured network traffic.
    Key Features: Tcpreplay can simulate real-world network conditions, test network appliances, and replay captured traffic for analysis and monitoring.

8. Security Information and Event Management (SIEM) APIs

    Description: SIEM platforms like Elastic Security (formerly known as ELK Stack), Splunk, and others often provide APIs for integrating with external systems and collecting security-related data.
    Key Features: Use SIEM APIs to ingest logs, events, and alerts from various security tools and systems into a centralized monitoring and analysis platform.

Example Integration Approach:

    Scenario: Integrate Suricata with a Python script to capture and analyze network traffic for intrusion attempts.
    Implementation: Use Suricata's logging features to output alerts to a file or a dedicated log server. Develop a Python script to parse these logs, extract relevant metrics (e.g., number of alerts per category, source IPs, attack types), and integrate with a SIEM platform for comprehensive security monitoring.

Considerations:

    Compatibility: Ensure compatibility between the chosen libraries/tools and your existing network infrastructure and security policies.

    Performance: Monitor resource usage (CPU, memory, network bandwidth) when deploying monitoring solutions to avoid performance degradation.

    Security: Follow best practices for secure integration, including access control, encryption, and data integrity measures.

By leveraging these libraries and tools, you can enhance your network security monitoring capabilities, detect potential threats more effectively, and respond promptly to security incidents. Tailor your choice of tools based on specific security requirements, infrastructure complexity, and integration capabilities with existing systems.

"""
