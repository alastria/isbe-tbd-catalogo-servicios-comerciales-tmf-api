# TMF Server in Go

`isbetmf` is a [TM Forum Open API](https://www.tmforum.org/oda/open-apis/directory) server written in Go, which can operate in two modes:

1. As a server implementing the TM Forum Open API specifications, managing the data  a database.
2. As a proxy server implementing PEP (Policy Enforcement Point) and PDP (Policy Decision Point) functionality in front of another TM Forum API server.

In both modes `ìsbetmf` exposes a set of REST APIs that can be used by client applications, implementing on top of the standard TM Forum specifications authentication, authorization and a set of business rules required for actual production-ready implementations.

The difference is that the standalone server manages its own data, while the proxy server forwards requests to another TM Forum API server.

## Approach to the implementation

TM Forum Open APIs are a set of standardized APIs designed to facilitate interoperability and integration across various systems and platforms in the industry. These APIs are developed and maintained by the TM Forum, an industry association that provides a collaborative environment for service providers and their suppliers to solve critical business issues.

The APIs cover a wide range of functionalities, including product catalog management, service order management, customer management, trouble ticketing, and more. For instance, the TMF620_ProductCatalog API allows for efficient management of product catalogs, while the TMF641_ServiceOrder API facilitates the handling of service orders.

Each managed resource has a set of APIs defined in their corresponding specifications, but all APIs implement the REST API Design Guidelines described in TMF630. This makes all APIs uniform and very similar among them, excep the data model, simplifying the implementation.

`isbetmf` leverages those similarities, like the guidelines for RESTful APIs naming and CRUD, to make a general implementation that supports almost any possible TM Forum managed entity, with very few exceptions.

### The REST API routes and handlers

The request path for a given resource can be described as follows:

```
/tmf-api/{apiFamily}/{apiVersion}/{resourceName}/{resourceID}?{queryParams}
```

or

```
/tmf-api/{apiFamily}/{apiVersion}/{resourceName}?{queryParams}
```

depending on the type of HTTP API (POST, GET, PATCH, DELETE) being used.

#### apiFamily and apiVersion

In the above path, `apiFamily` is the group of APIs defined in a given TMF specification document and `apiVersion`is the version of the specification implemented by that route. For example, Product Catalog API (TMF620), Service Order API (TMF641) or Customer Management API (TMF629).

TM Forum provides a set of Swagger definitions to facilitate life to implementers, and you can see some of those definitions in the directory `www/oapiv4/oapiv4`, in the form of a JSON file for each API family specification.

In `isbetmf` we use the value of the `basePath` attribute in the swagger files to derive both the `apiFamily` and the `apiversion` values for that family of APIs.

#### resourceName

The `resourceName` in the path corresponds to the first path component in each of the values of the `paths` object in each swagger file. This is the actual object or resource being managed by the CRUD REST APIs defined also in the swagger file.

For example, `catalogProductManagement` is the resource name for the TMF620 - Product Catalog Management API, version V4.

#### resourceID

`resourceID` is the unique identifier of the resource being managed, and it is compulsory except in POST calls (more on this later).

#### queryParams

The `queryParams` part of the path is used to specify additional properties of the request, like filtering objects, attribute selection and pagination. 
For example, to retrieve a list of products that are currently active, you might use `/tmf-api/productCatalog/v4/product?status=active`.

More details are given on the sections describing the different request types.

## The TM Forum REST API Server

The standalone `isbetmf`server handles requests in much the same way as the PEP/PDP proxy server. This section describes the request processing of the standalone server but mentions the proxy server when relevant.

### The swagger UI

The server supports the swagger UI for interactive usage and experimentation of the APIs, available at the path `/oapiv4/index.html` when the server is running.

### HTTP server routes and handlers

The client requests are received in the package `tmfserver/handler`package. This layer is very thin, as it converts the HTTP requests/replies of the web framework used to an agnostic format, before sending the request to the service layer in `tmfserver/service`. The structs used for this are `service.Request` and `service.Response`.

This provides the ability to switch easily the web framework (eg. Echo) or even implement a different input channel, like using JSON-RPC.

### The service layer

The main logic for processing the requests is in the file `tmfserver/service/service.go`. We describe the logic for the different types of requests.
The service layer is created with the `service.NewService()` function, which returns an initialized `Service` struct. The service layer creation requires the following:

- The database instance to use for storage.
- The policy rules instance to use for additional user-defined authorization rules. These rules are optional and can be defined by the operator of the server using the [Starlark language](https://starlark-lang.org/).
- The URL of the Verifier server that generates the access tokens used for authentication/authorization in the APIs. This is used to verify the signature of access tokens.
- Whether the proxy functionality is enabled and if so, the URL of the remote TM Forum server.

#### Query a single Resource

This is done with a GET request like this:

```HTTP
GET /tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:1b75e04b-6d1b-439d-baca-3cbcb33393c6
```

This request path is mapped to the following resource:
- `apiFamily`: `productCatalogManagement`
- `resourceName`: `productOffering`
- `resourceID`: `urn:ngsi-ld:product-offering:1b75e04b-6d1b-439d-baca-3cbcb33393c6`

The `Service.GetGenericObject` function performs the following processing.

**Authentication**: The read APIs can be used without authentication, that is, no access token must be provided. However, if the caller provides an access token, it must be valid. If an access token is provided:
- the signature is validated using the public key of the Verifier
- the data of the calling user is extracted from the access token and the `Request` object is updated. The user data is put in the `Request.AuthUser` struct.
- the access token is put in the `Request.AccessToken` struct.

**Retrieve the object from the database**: the object is retrieved using the `Service.getLocalOrRemoteObject` function. This function uses the local database if the proxy functionality is not enabled, or forwards the request to the remote TMF server if it is enabled. If the object is not found, the service returns an error to the REST API handler.

**Check authorization policies**: the service calls the `Service.takeDecision` function to check if the calling user is able to access the object. This function prepares the data for the extensible PDP with user-written policies and calls it at `PDP.Authorize()`, passing the following:
- the request received from the user
- the access token received in the API call, if any
- the TM Forum resource object being accessed
- the user data obtained from the access token, in a more user-friendly fornmat than in the raw access token

To see more details on the how users can write user policies, see the `auth_policies.star` file.

**Filter the object as requested by the caller**: if the caller specified a `queryParams` with `fields` attribute, we filter the TMF object and return only the first-level attributes specified by the user. In any case, the fields `id`, `href`, `lastUpdate`, `version`, `lifecycleStatus` are always returned.

#### Query multiple resources of the same type (same `apiFamily`)

This is done with a request like this:

```HTTP
GET /tmf-api/productCatalogManagement/v4/productOffering
```

This request path is mapped to the following resource:
- `apiFamily`: `productCatalogManagement`
- `resourceName`: `productOffering`

The `Service.ListGenericObjects` function performs the following processing.

**Authentication**: The read APIs can be used without authentication, that is, no access token must be provided. However, if the caller provides an access token, it must be valid. If an access token is provided:
- the signature is validated using the public key of the Verifier
- the data of the calling user is extracted from the access token and the `Request` object is updated. The user data is put in the `Request.AuthUser` struct.
- the access token is put in the `Request.AccessToken` struct.

**Retrieve the objects from the local database**: if the proxy functionality is not enabled, the objects are retrieved using the `Service.listObjects` function. Otherwise, the request is forwarded to the remote TMF server if it is enabled. It is not an error if no objects are found, and in this case an empty list is returned to the caller.

**Filter each of the objects as requested by the caller**: if the caller specified a `queryParams` with `fields` attribute, we filter each of the objects and return only the first-level attributes specified by the user. In any case, the fields `id`, `href`, `lastUpdate`, `version`, `lifecycleStatus` are always returned.
