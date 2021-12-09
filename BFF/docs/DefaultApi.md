# openapi_client.DefaultApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**delete_haiku**](DefaultApi.md#delete_haiku) | **DELETE** /api/haiku/{haiku_id} | 
[**delete_subscribe**](DefaultApi.md#delete_subscribe) | **DELETE** /api/subscribe/{user_id} | 
[**get_haiku**](DefaultApi.md#get_haiku) | **GET** /api/haiku/{haiku_id} | get_haiku
[**get_timeline**](DefaultApi.md#get_timeline) | **GET** /api/timeline | timeline
[**get_top**](DefaultApi.md#get_top) | **GET** /api/top | top
[**get_user**](DefaultApi.md#get_user) | **GET** /api/users/{user_id} | user_info
[**post_haiku**](DefaultApi.md#post_haiku) | **POST** /api/post-haiku | 
[**post_signin**](DefaultApi.md#post_signin) | **POST** /api/signin | 
[**post_signup**](DefaultApi.md#post_signup) | **POST** /api/signup | 
[**post_subscribe**](DefaultApi.md#post_subscribe) | **POST** /api/subscribe/{user_id} | 


# **delete_haiku**
> delete_haiku(haiku_id)



haikuを削除

### Example


```python
import time
import openapi_client
from openapi_client.api import default_api
from openapi_client.model.inline_object6 import InlineObject6
from pprint import pprint
# Defining the host is optional and defaults to http://localhost:8080
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://localhost:8080"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = default_api.DefaultApi(api_client)
    haiku_id = 1 # int | 
    inline_object6 = InlineObject6(
        session_id="session_id_example",
    ) # InlineObject6 |  (optional)

    # example passing only required values which don't have defaults set
    try:
        api_instance.delete_haiku(haiku_id)
    except openapi_client.ApiException as e:
        print("Exception when calling DefaultApi->delete_haiku: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_instance.delete_haiku(haiku_id, inline_object6=inline_object6)
    except openapi_client.ApiException as e:
        print("Exception when calling DefaultApi->delete_haiku: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **haiku_id** | **int**|  |
 **inline_object6** | [**InlineObject6**](InlineObject6.md)|  | [optional]

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_subscribe**
> delete_subscribe(user_id)



リムーブ

### Example


```python
import time
import openapi_client
from openapi_client.api import default_api
from openapi_client.model.inline_object4 import InlineObject4
from pprint import pprint
# Defining the host is optional and defaults to http://localhost:8080
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://localhost:8080"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = default_api.DefaultApi(api_client)
    user_id = 1 # int | 
    inline_object4 = InlineObject4(
        session_id="session_id_example",
    ) # InlineObject4 |  (optional)

    # example passing only required values which don't have defaults set
    try:
        api_instance.delete_subscribe(user_id)
    except openapi_client.ApiException as e:
        print("Exception when calling DefaultApi->delete_subscribe: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_instance.delete_subscribe(user_id, inline_object4=inline_object4)
    except openapi_client.ApiException as e:
        print("Exception when calling DefaultApi->delete_subscribe: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_id** | **int**|  |
 **inline_object4** | [**InlineObject4**](InlineObject4.md)|  | [optional]

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_haiku**
> InlineResponse2001 get_haiku(haiku_id)

get_haiku

haikuの詳細を取得

### Example


```python
import time
import openapi_client
from openapi_client.api import default_api
from openapi_client.model.inline_response2001 import InlineResponse2001
from pprint import pprint
# Defining the host is optional and defaults to http://localhost:8080
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://localhost:8080"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = default_api.DefaultApi(api_client)
    haiku_id = 1 # int | 

    # example passing only required values which don't have defaults set
    try:
        # get_haiku
        api_response = api_instance.get_haiku(haiku_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DefaultApi->get_haiku: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **haiku_id** | **int**|  |

### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_timeline**
> [Haiku] get_timeline()

timeline

タイムラインの取得

### Example


```python
import time
import openapi_client
from openapi_client.api import default_api
from openapi_client.model.haiku import Haiku
from openapi_client.model.inline_object5 import InlineObject5
from pprint import pprint
# Defining the host is optional and defaults to http://localhost:8080
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://localhost:8080"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = default_api.DefaultApi(api_client)
    inline_object5 = InlineObject5(
        session_id="session_id_example",
        start=1,
        stop=1,
    ) # InlineObject5 |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # timeline
        api_response = api_instance.get_timeline(inline_object5=inline_object5)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DefaultApi->get_timeline: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object5** | [**InlineObject5**](InlineObject5.md)|  | [optional]

### Return type

[**[Haiku]**](Haiku.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_top**
> [Haiku] get_top()

top

トップ30件のhaikuの取得

### Example


```python
import time
import openapi_client
from openapi_client.api import default_api
from openapi_client.model.haiku import Haiku
from pprint import pprint
# Defining the host is optional and defaults to http://localhost:8080
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://localhost:8080"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = default_api.DefaultApi(api_client)

    # example, this endpoint has no required or optional parameters
    try:
        # top
        api_response = api_instance.get_top()
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DefaultApi->get_top: %s\n" % e)
```


### Parameters
This endpoint does not need any parameter.

### Return type

[**[Haiku]**](Haiku.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_user**
> InlineResponse200 get_user(user_id)

user_info

ユーザー情報

### Example


```python
import time
import openapi_client
from openapi_client.api import default_api
from openapi_client.model.inline_response200 import InlineResponse200
from pprint import pprint
# Defining the host is optional and defaults to http://localhost:8080
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://localhost:8080"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = default_api.DefaultApi(api_client)
    user_id = 1 # int | 

    # example passing only required values which don't have defaults set
    try:
        # user_info
        api_response = api_instance.get_user(user_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DefaultApi->get_user: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_id** | **int**|  |

### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **post_haiku**
> post_haiku()



詠む

### Example


```python
import time
import openapi_client
from openapi_client.api import default_api
from openapi_client.model.inline_object2 import InlineObject2
from pprint import pprint
# Defining the host is optional and defaults to http://localhost:8080
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://localhost:8080"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = default_api.DefaultApi(api_client)
    inline_object2 = InlineObject2(
        session_id="session_id_example",
        content=ApiPostHaikuContent(
            first="first_example",
            second="second_example",
            third="third_example",
        ),
    ) # InlineObject2 |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_instance.post_haiku(inline_object2=inline_object2)
    except openapi_client.ApiException as e:
        print("Exception when calling DefaultApi->post_haiku: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object2** | [**InlineObject2**](InlineObject2.md)|  | [optional]

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | Created |  -  |
**400** | Bad Request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **post_signin**
> InlineResponse201 post_signin()



ログイン

### Example


```python
import time
import openapi_client
from openapi_client.api import default_api
from openapi_client.model.inline_response201 import InlineResponse201
from openapi_client.model.inline_object1 import InlineObject1
from pprint import pprint
# Defining the host is optional and defaults to http://localhost:8080
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://localhost:8080"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = default_api.DefaultApi(api_client)
    inline_object1 = InlineObject1(
        name="name_example",
        pw="pw_example",
    ) # InlineObject1 |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_response = api_instance.post_signin(inline_object1=inline_object1)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DefaultApi->post_signin: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object1** | [**InlineObject1**](InlineObject1.md)|  | [optional]

### Return type

[**InlineResponse201**](InlineResponse201.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Bad Request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **post_signup**
> InlineResponse201 post_signup()



サインアップ

### Example


```python
import time
import openapi_client
from openapi_client.api import default_api
from openapi_client.model.inline_object import InlineObject
from openapi_client.model.inline_response201 import InlineResponse201
from pprint import pprint
# Defining the host is optional and defaults to http://localhost:8080
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://localhost:8080"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = default_api.DefaultApi(api_client)
    inline_object = InlineObject(
        name="name_example",
        pw="pw_example",
    ) # InlineObject |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_response = api_instance.post_signup(inline_object=inline_object)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DefaultApi->post_signup: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object** | [**InlineObject**](InlineObject.md)|  | [optional]

### Return type

[**InlineResponse201**](InlineResponse201.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | Created |  -  |
**409** | Conflict |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **post_subscribe**
> post_subscribe(user_id)



フォロー

### Example


```python
import time
import openapi_client
from openapi_client.api import default_api
from openapi_client.model.inline_object3 import InlineObject3
from pprint import pprint
# Defining the host is optional and defaults to http://localhost:8080
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://localhost:8080"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = default_api.DefaultApi(api_client)
    user_id = 1 # int | 
    inline_object3 = InlineObject3(
        session_id="session_id_example",
    ) # InlineObject3 |  (optional)

    # example passing only required values which don't have defaults set
    try:
        api_instance.post_subscribe(user_id)
    except openapi_client.ApiException as e:
        print("Exception when calling DefaultApi->post_subscribe: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_instance.post_subscribe(user_id, inline_object3=inline_object3)
    except openapi_client.ApiException as e:
        print("Exception when calling DefaultApi->post_subscribe: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_id** | **int**|  |
 **inline_object3** | [**InlineObject3**](InlineObject3.md)|  | [optional]

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

