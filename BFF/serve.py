from flask import Flask, render_template
import openapi_client
from openapi_client.api import default_api

app = Flask(__name__)

configuration = openapi_client.Configuration(
    host = "http://api-server:8080"
)

@app.route("/")
def hello_world():
    return "<p>Hello,World!</p>"

@app.route("/top")
def get_top():
    url = "http://api-server:8080/api/top"

    with openapi_client.ApiClient(configuration=configuration) as api_client:
        # Create an instance of the API class
        api_instance = default_api.DefaultApi(api_client)

        # example, this endpoint has no required or optional parameters
        try:
            # top
            api_response = api_instance.get_top()
        except openapi_client.ApiException as e:
            print("Exception when calling DefaultApi->get_top: %s\n" % e)

        #print(api_response.body.json())
        return render_template('top.html',title='top_30Haiku', Haikus=api_response)

## おまじない
if __name__ == "__main__":
    app.run(debug=True,port=5000,host="0.0.0.0")
