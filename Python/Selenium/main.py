import os
import pandas as pd
import datetime
from selenium import webdriver
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.support.ui import WebDriverWait
import time
from selenium.webdriver.common.by import By
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.common.keys import Keys

def get_form_inputs():
    options = Options()
    options.add_argument("start-maximized")
    driver = webdriver.Chrome(options=options)
    #
    #

    # options.add_experimental_option("excludeSwitches", ["enable-automation"])
    # options.add_experimental_option('useAutomationExtension', False)
    # # driver.get("https://www.inipec.gov.it/cerca-pec/-/pecs/companies")
    driver.get("https://www.google.com")
    time.sleep(30)

    # form = driver.find_element_by_name("q")
    # print(form)
    # form.send_keys("Omaze")
    # form.submit()
    # time.sleep(2)
    # # google_omaze = driver.find_element_by_xpath("/html/body/div[7]/div/div[9]/div[1]/div/div[2]/div[2]/div/div/div[1]/div/div/div/div/div/div/div[1]/a/h3")
    # # google_omaze.click()
    # driver.get("https://www.omaze.com/")
    # time.sleep(3)
    # cars_link = driver.find_element_by_xpath("/html/body/header/div[2]/nav/div/div[2]/ul/li[3]/a")
    # cars_link.click()
    # time.sleep(5)
    # sprinter_van_link = driver.find_element_by_xpath("/html/body/main/div/section[2]/div[1]/a[1]/div[1]/div/img")
    # sprinter_van_link.click()
    # time.sleep(3)
    # enter_now_button = driver.find_element_by_xpath("/html/body/main/section/section/div[1]/div/div[3]/a")
    # enter_now_button.click()
    # time.sleep(2)
    # enter_without_contributing = driver.find_element_by_xpath("/html/body/main/section/section/div[6]/div/div/a")
    # enter_without_contributing.click()
    # time.sleep(5)
    # first_name = driver.find_element_by_xpath("/html/body/div[1]/div[2]/div[4]/form/div[2]/div/input")
    # first_name.click()
    # first_name.send_keys(Keys.ENTER)
    # time.sleep(3)

    # driver.find_elements_by_xpath('//div[contains(text(), "' + text + '")]')
    # WebDriverWait(driver, 10)
    # driver.get('https://fame.omaze.com/6585430409306?title=Win%20a%20Sprinter%C2%AE%20Van%20with%20an%20%2480%2C000%20Eco-Friendly%20Conversion&handle=sprinter-van-2021&variant_id=39446822649946&variant_price=$2')
    # WebDriverWait(driver, 10).until(EC.frame_to_be_available_and_switch_to_it(
    #     (By.CSS_SELECTOR, "iframe[name^='a-'][src^='https://www.google.com/recaptcha/api2/anchor?']")))
    # WebDriverWait(driver, 10).until(EC.element_to_be_clickable((By.XPATH, "//span[@id='recaptcha-anchor']"))).click()


    # time.sleep(2)
    # inputClasses = driver.find_elements_by_class_name("fame-field__input hkjs--dirty hkjs--not-empty hkjs--touched")
    #
    # print(inputClasses)
    # time.sleep(30)
#     username = driver.find_element_by_id("Username")
#     username.send_keys(os.environ['uwcu_username'])
#     password = driver.find_element_by_xpath('//*[@id="main"]/div[1]/div[3]/div[3]/form/div/div[1]/div/div[2]/div/input')
#     password.send_keys(os.environ['uwcu_password'])
#     time.sleep(2)
#     driver.find_element_by_xpath('/html/body/div[2]/div[2]/div[1]/div/div[1]/div[3]/div[3]/form/div/div[2]/div/div[1]/button').click()
#     time.sleep(1)
#     driver.find_element_by_xpath('/html/body/div[2]/div[2]/div[1]/div/div[1]/div[3]/div[3]/form/div/div[2]/div/div[1]/button').click()
#
#     time.sleep(6)
#     # accounts = driver.find_elements_by_class_name("nav-account_item_name")
#     data = {}
#     for child_number in range(3):
#         try:
#             name_tag = driver.find_element_by_css_selector(f"#Dash-Col-1 > div.accounts-summary-widget.js-dashwidget.js-configurable.panel.js-sortable > div.panel_front.tile.item-container.sortable-item > div.item-container.no-deco.action-mode_hide > ul > li:nth-child({child_number}) > a > div > div.dash-acct-sections_title.mb0.left")
#             name = name_tag.get_attribute("innerText")
#             balance_tag = driver.find_element_by_css_selector(f'#Dash-Col-1 > div.accounts-summary-widget.js-dashwidget.js-configurable.panel.js-sortable > div.panel_front.tile.item-container.sortable-item > div.item-container.no-deco.action-mode_hide > ul > li:nth-child({child_number}) > a > span > span.amt.txt-small.mr2 > span:nth-child(1)')
#             balance = float((balance_tag.get_attribute("innerText")[1:]).replace(",", "").strip())
#             data[name] = balance
#         except:
#             print('didnt find anything', child_number)
#     print(data)
# #    time.sleep(600)
#     return data


# def extend_uwcu_record():
#     records = pd.read_csv('bots/uwcu.csv').to_dict('records')
#     ts = datetime.datetime.now().timestamp()
#     account = get_form_inputs()
#     account['time'] = ts
#     records.extend([account])
#     pd.DataFrame(records).to_csv('bots/uwcu.csv', index=False)
#
#     del account['time']
#     positions = [{'name': key, 'equity': -account[key] if 'credit' in key else account[key]} for key in account.keys()]
#     pd.DataFrame(positions).to_csv('bots/uwcu-positions.csv', index=False)
#
#
# def get_latest_uwcu():
#     records = pd.read_csv('bots/uwcu-positions.csv').to_dict('records')
#     return records


if __name__ == '__main__':
    get_form_inputs()
