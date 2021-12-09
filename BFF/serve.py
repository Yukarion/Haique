from flask import Flask, render_template, request, make_response

import openapi_client
from openapi_client.api import default_api
from openapi_client.models import *

from openapi_client.model.inline_object import InlineObject
from openapi_client.model.inline_response201 import InlineResponse201


app = Flask(__name__)

configuration = openapi_client.Configuration(
    host = "http://api-server:8080"
)

@app.route("/")
def hello_world():
    return "<p>Hello,World!</p>"

@app.route("/top")
def get_top():
    with openapi_client.ApiClient(configuration=configuration) as api_client:
        # Create an instance of the API class
        api_instance = default_api.DefaultApi(api_client)

        # example, this endpoint has no required or optional parameters
        try:
            # top
            api_response = api_instance.get_top()
        except openapi_client.ApiException as e:
            print("Exception when calling DefaultApi->get_top: %s\n" % e)
            resp = make_response(render_template("error.html",title="Error occured"))
            return resp

        return render_template('top.html',title='top_30Haiku', Haikus=api_response)

@app.route("/signup")
def get_signup():
    return render_template('signup.html',title='signup',err="")

@app.route("/signup",methods=["POST"])
def post_signup():
    with openapi_client.ApiClient(configuration=configuration) as api_client:
        # Create an instance of the API class
        api_instance = default_api.DefaultApi(api_client)
        inline_object = InlineObject(
            name = request.form.get("name"),
            pw = request.form.get("password"),
        ) # InlineObject |  (optional)

        # example passing only required values which don't have defaults set
        # and optional values
        try:
            api_response = api_instance.post_signup(inline_object=inline_object)
            print(api_response)
            print(type(api_response))
        except openapi_client.ApiException as e:
            print("Exception when calling DefaultApi->post_signup: %s\n" % e)
            return render_template('signup.html',title='signup',err="Already used name...")
        resp = make_response(render_template('signup_done.html',title='signup'))
        resp.set_cookie("session_id",api_response.session_id)
        return resp

@app.route("/signin")
def get_signin():
    return render_template('signin.html',title='signin',err="")
@app.route("/signin",methods=["POST"])
def post_signin():
    with openapi_client.ApiClient(configuration=configuration) as api_client:
        # Create an instance of the API class
        api_instance = default_api.DefaultApi(api_client)
        inline_object1 = InlineObject1(
            name = request.form.get("name"),
            pw = request.form.get("password"),
        ) # InlineObject1 |  (optional)

        # example passing only required values which don't have defaults set
        # and optional values
        try:
            api_response = api_instance.post_signin(inline_object1=inline_object1)
        except openapi_client.ApiException as e:
            print("Exception when calling DefaultApi->post_signin: %s\n" % e)
            return render_template('signin.html',title='signup',err="Wrong password or name")
        resp = make_response(render_template('signup_done.html',title='signup'))
        resp.set_cookie("session_id",api_response.session_id)
        return resp 

## おまじない
if __name__ == "__main__":
    app.run(debug=True,port=5000,host="0.0.0.0")
