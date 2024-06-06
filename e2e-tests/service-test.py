import pytest
import requests
import subprocess
import time
import os

@pytest.fixture(autouse=True)
def start_docker():
    cwd = os.path.dirname(__file__)
    os.chdir(os.path.join(cwd, '..'))

    result = subprocess.run(["docker", "compose", "-f", "docker-compose-e2e.yml", "up", "-d", "--build"])
    if result.returncode != 0:
        print("Failed to start docker compose")
        print(result)
        exit(1)
    time.sleep(10)

    yield

    subprocess.run(["docker", "compose", "-f", "docker-compose-e2e.yml", "down"])
    os.chdir(cwd)

def create_and_login() -> str:
    urlCreate = "http://localhost:8764/users"
    dataCreate = {
        "login": "testuser",
        "password": "testpassword"
    }
    response = requests.post(urlCreate, json=dataCreate)
    assert response.status_code == 200

    urlLogin = "http://localhost:8764/auth"
    dataLogin = {
        "login": "testuser",
        "password": "testpassword"
    }
    response = requests.post(urlLogin, json=dataLogin)
    assert response.status_code == 200
    return response.cookies["SESSIONID"]

def test_create_user():
    sessionID = create_and_login()
    
    urlPatch = "http://localhost:8764/users/me"
    dataPatch = {
        "firstName": "name",
        "lastName": "surname",
        "birthDate": "2000-01-01",
        "email": "mail@mail.com",
        "phoneNumber": "123456789",
    }

    responseNotOK = requests.patch(urlPatch, json=dataPatch)
    assert responseNotOK.status_code == 401
    
    response = requests.patch(urlPatch, json=dataPatch, headers={"Cookie": f"SESSIONID={sessionID}"})
    assert response.status_code == 200

def test_create_post_and_get_stats():
    sessionID = create_and_login()

    urlCreatePost = "http://localhost:8764/posts"
    dataCreatePost = "Test post content"
    responseCreatePost = requests.post(urlCreatePost, json=dataCreatePost, headers={"Cookie": f"SESSIONID={sessionID}"})
    assert responseCreatePost.status_code == 200
    postId = responseCreatePost.json()['postId']

    urlGetPostStats = f"http://localhost:8764/posts/{postId}/stats"
    responseGetPostStats = requests.get(urlGetPostStats, headers={"Cookie": f"SESSIONID={sessionID}"})
    assert responseGetPostStats.status_code == 200
    assert 'viewCount' in responseGetPostStats.json()
    assert 'likeCount' in responseGetPostStats.json()
