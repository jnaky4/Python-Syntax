import kaggle


def kaggleAPI(dataset, filename):
    api = kaggle.api
    # located in ~/.kaggle/kaggle.json
    api.authenticate()
    api.dataset_download_file(dataset, filename, "./")






