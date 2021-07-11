
# Redis Demo Application

A demo application with rest apis to connect with a redis cluster running in [KloverCloud](https://klovercloud.com) platform. The application dynamically loads required environment variables required to connect with redis and and perform read/write actions.

## Application details
The application established connection with each endpoint on redis endpoint. If it fails to connect with any instance, it will throw a fatal log and exit. All the write operations **(POST, PUT, DELETE)** are executed through the master instance and all the **GET** operations are executed through slave instances (round-robin)

## Deploying application and cache

create a new vpc with enough resources to deploy the go application and redis server. Assuming you already have added personal access token of github/gitlab or your repo is public. OnBoard the application and edit the DockerFile as necessary, *e.g.* exposing the port the application is running.

Create a redis cache in the same vpc, and deploy the cache server. Yahoo! Nearly halfway there, with some mouse clicks right. No pesky terminal hassle. All we need to do is just inform the application about the cache server. The application already expecting one, so lets do it.


## Secret information through environment variables
First of all, the authentication password. Which is required to connect with every redis instance. You can easily add this through the secret tab on application page. Create a secret, give it a name and add the key-value as following

```
KEY               VALUE
REDIS_PASSWORD    {your_cache_password}
```

## Instance information through environment variables
When the application starts, it will look for the master endpoint from environment variable. And also look for slave instance count. If the slave instance is > 0, the application will try to load all the slave endpoints and will ping the slave instance. If it fails to get master / slave endpoints or ping returns error, the application will terminate with fatal error log. Add the environment variables as described below and double check. or what ? Your application will go on infinite crash loop (opps!)

Go to the cache [cache](https://console.klovercloud.com/cache) section, select your cache, go to overview section and click to service endpoint to see all the endpoints list. Add them to environment variables as described below

```
KEY                VALUE
MASTER_ENDPOINT    {master_endpoint}
SLAVE_COUNT        {no_of_slaves}
SLAVE_ENDPOINT_0   {first_slave_endpoint}
SLAVE_ENDPOINT_1   {second_slave_endpoint}
.
.
.
.

```
deploy your application on desired deployment environment. Make sure you added the environment variable and secrets in the right deployment environment. if you already deployed (and obviously its in crashloop!) re-deploy your application after adding environment variables and secrets. If deployment fails, the deployment logs should show the fatal log, its detailed enough to debug easily.
## REST endpoints
#### GET
From the CI/CD pipeline page, select the deployment, right click info. **External Default Endpoint** is your application url, so just add the path /api/v1/cache. and also add query param **key** or it will return error.
```bash
Example:
https://your-application-prefix.eu-west-1.klovercloud.com/api/v1/cache?key=your_key
```
#### POST
in the request body, provide json containing key-value as the below example. Only strings are allowed. If key already exist it will overwrite it.
Hit the same endPoint with POST method.\
https://your-application-prefix.eu-west-1.klovercloud.com/api/v1/cache
```json
{
    "key"   : "klovercloud",
    "value" : "awesome",
}
```
#### PUT
same as post. Oh, it checks if the key exist or not. Updating a key that isn't there yet? Nah that's not gonna happen.
```json
{
    "key"   : "klovercloud",
    "value" : "moreAwesome",
}
```
#### DELETE
Same as the GET api. Provide key the query params.
## Contributing
Pull requests for new features, bug fixes, and suggestions are welcome!

## License
[MIT](https://choosealicense.com/licenses/mit/)
