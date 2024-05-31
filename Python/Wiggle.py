import pyautogui
import time
from datetime import datetime

# if True: if the mouse is in any of the four corners of the primary monitor,
# they will raise a pyautogui. FailSafeException
pyautogui.FAILSAFE = False

while True:
    time.sleep(60)

    pyautogui.moveTo(pyautogui.position().x, pyautogui.position().y + 1)
    pyautogui.moveTo(pyautogui.position().x, pyautogui.position().y - 1)

    # print("Movement made at", datetime.now().strftime("%I:%M:%S %p"))

    if datetime.now().hour >= 16:
        exit(0)
    if datetime.now().hour >= 15 and datetime.now().minute > 38:
        exit(0)