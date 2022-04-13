from hashlib import sha256
from random import random
import string
import random
from faker import Faker
import json
import logging
import requests


LOG_FILE_SERVER = "server.log"
LOG_FILE_E2ETIME = "e2etime.log"
LOG_FILE_UPLOAD = "upload.log"

URL_CREATE_ASSET = "http://localhost:8080/CreateAsset"
URL_UPLOAD_ASSET = "http://localhost:8080/Upload"


def setup_logger(logger_name, log_file, level=logging.INFO):

    log_setup = logging.getLogger(logger_name)
    formatter = logging.Formatter(
        '%(levelname)s: %(asctime)s %(message)s', datefmt='%m/%d/%Y %I:%M:%S %p')
    fileHandler = logging.FileHandler(log_file, mode='a')
    fileHandler.setFormatter(formatter)
    streamHandler = logging.StreamHandler()
    streamHandler.setFormatter(formatter)
    log_setup.setLevel(level)
    log_setup.addHandler(fileHandler)
    log_setup.addHandler(streamHandler)

    return log_setup


def dict_to_binary(the_dict):
    str = json.dumps(the_dict)
    binary = ' '.join(format(ord(letter), 'b') for letter in str)
    return binary


def binary_to_dict(the_binary):
    jsn = ''.join(chr(int(x, 2)) for x in the_binary.split())
    d = json.loads(jsn)
    return d


def create_random_input():
    countryID = random.choice(string.ascii_uppercase) + \
        str(random.randint(100000, 999999)) + \
        "("+str(random.randint(0, 9))+")"
    fake = Faker()
    personalInfo = {

        "sysID": ''.join(random.SystemRandom().choice(string.ascii_letters + string.digits) for _ in range(28)),
        "name": fake.name(),
        "countryCode": fake.country_code(representation="alpha-2"),
        "countryID": countryID,
        "gender": str(random.SystemRandom().randint(0, 1)),
        "dateOfBirth": fake.date()
    }
    personHash = sha256(dict_to_binary(
        personalInfo).encode('utf-8')).hexdigest()

    inputInfo = {
        "CertDetail": {
            "PersonSysID": countryID,
            "Name": fake.first_name()+"Vac",
            "Brand": fake.company(),
            "NumOfDose": str(random.randint(1, 4)),
            "Time": str(fake.date_time_this_year()),
            "Issuer": "vac_center1"
        },
        "PersonInfoHash": personHash,
        "Key": ''.join(random.SystemRandom().choice(string.ascii_letters + string.digits) for _ in range(random.randint(4, 8)))
    }
    return inputInfo


def main():

    log_server = setup_logger('log_server', LOG_FILE_SERVER)
    log_e2e_time = setup_logger('log_e2eTime', LOG_FILE_E2ETIME)
    log_upload = setup_logger('log_upload', LOG_FILE_UPLOAD)

    sizeArrray = [4, 5, 6, 7, 8, 9, 10]

    for size in sizeArrray:

        trials = pow(2, size)

        log_server.info("n = " + str(size) + " trial start")
        log_e2e_time.info("n = " + str(size) + " trial start")
        log_upload.info("n = " + str(size) + " trial start")

        timeArray = []

        for trial in range(trials):

            inputInfo = create_random_input()
            post_fields = json.dumps(inputInfo)

            log_server.info(post_fields)

            response = requests.post(URL_CREATE_ASSET, data=post_fields)
            
            log_server.info(
                "size is %d, trial no. is %d, data is%s", size, trial, post_fields)
            log_server.info("size is %d, trial no. is %d, data is%s",
                            size, trial, response.elapsed.total_seconds())
            timeArray.append(response.elapsed.total_seconds())

        log_server.info("size is %d, data is%s", size,  timeArray)

        if(len(timeArray) == 0):
            log_e2e_time.warning("no item in time array")
        else:
            avg = sum(timeArray)/len(timeArray)
            log_e2e_time.info("n = " + str(size) +
                              "trial. The average is " + str(avg))

        # after create asset, upload to global
        response = requests.post(URL_UPLOAD_ASSET)
        log_upload.info("n = " + str(size) + "trial for upload, time needed is " +
                        str(response.elapsed.total_seconds()))


if __name__ == "__main__":
    main()
