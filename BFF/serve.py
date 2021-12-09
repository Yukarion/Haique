from flask import Flask, render_template
import requests

app = Flask(__name__)

@app.route("/")
def hello_world():
    return "<p>Hello,World!</p>"

@app.route("/top")
def get_top():

    url = "http://localhost:8080/api/top"

    r = requests.get(url)
    print(r.json())
    return render_template('top.html',title='top_30Haiku', Haikus=r.json())

## おまじない
if __name__ == "__main__":
    app.run(debug=True,port=5000)
