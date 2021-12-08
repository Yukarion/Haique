from flask import Flask
import requests

app = Flask(__name__)

@app.route("/")
def hello_world():
    return "<p>Hello,World!</p>"

@app.route("/top")
def get_top():

    url = "http://127.0.0.1:3100/api/top"

    r = requests.get(url)
    print(r.json())
    return "<p>Top 30 Get sucsessfully!</p>"
