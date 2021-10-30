import sys, os, platform


#TODO add argparser

# first argument is the filename
if len(sys.argv) > 1:
    print(sys.argv[0], sys.argv[1])
    try:
        f = open(sys.argv[1], "r")
        print(f.readline())
    except Exception as e:
        print(e)


# Grabs basic OS info from your computer, returns the Operating System name
def get_os_info():
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
    return os_platform


get_os_info()
