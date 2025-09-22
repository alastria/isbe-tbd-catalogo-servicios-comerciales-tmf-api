# ISBE-TMF Server Architecture Documentation

This document outlines the architecture of `isbetmf`, a Go-based REST API server designed to implement TM Forum APIs. It details the key components, design strategies, and relevant considerations from its development.

It passes the official TMF Conformance Test Kit for TMforum V5 and V4 APIs.

## 1. Overview

The `isbetmf` is an implementation of a TM Forum REST API server, aiming to provide a simple, flexible and robust platform for handling various TM Forum API specifications.

It can work in two modes:

*   As a standalone TM Forum server, which is the default.
*   As a PIP + PDP proxy: an authentication & authorization proxy in front of another implementation of the TM Forum APIs.

## 2. Core Components and Structure

The server's architecture is modular, separating concerns into distinct packages and layers:

*   **`cmd/isbetmf/main.go`**: The main entry point for the application, responsible for initializing and starting the server.
*   **`tmfserver/handler/`**: This layer defines the API handlers, abstracting the underlying web framework.
    *   `tmfserver/handler/fiber/`: Contains handlers implemented using the Fiber web framework.
    *   Additional frameworks can be implemented in the `handler` package, allowing flexibility in choosing or switching web frameworks. The framework used in production by default is Fiber.
*   **`tmfserver/repository/`**: Definition of the tables and the main objects close to the database.
    *   `tables.go`: Defines the database table structures.
    *   `tmfobject.go`: The generalized TMF object, supporting all the specific TMF objects.
    *   `organization.go`: The Organization object, which has specific structure and behavior.
*   **`tmfserver/service/service.go`**: The service layer encapsulates the business logic. It orchestrates operations by interacting with the repository layer and providing an interface for the handlers. This separation ensures that business rules are independent of the web framework or database implementation details.
*   **`tmfserver/www/`**: This directory serves static assets, primarily for the Swagger UI, enabling interactive API documentation.

## 3. Key Architectural Strategies

### 3.1 Database Design: Single Table for TM Forum Objects

TM Forum objects are stored in a single SQLite table. This design choice offers:
*   **Flexibility**: Accommodates various TM Forum object types without requiring a new table for each, simplifying schema management.
*   **JSON Storage**: The entire TM Forum object is stored as a JSON field, allowing for schema evolution without database migrations for every object change.
*   **Metadata Fields**: Some fields are used for metadata and frequently queried attributes, optimizing common SQL queries.

### 3.2 In-Memory Representation: `map[string]any`

To support a wide range of TM Forum object types with a consistent codebase, the in-memory representation of these objects is based on a `map[string]any` nested structure. This approach provides:
*   **Genericity**: A single code path can handle most TM Forum APIs, as many objects share common properties.
*   **Type Safety (with methods)**: While the underlying structure is generic, specific methods are implemented to query and manipulate the map in a type-safe manner, ensuring data integrity and ease of use.

### 3.3 Error Handling

The server adheres to a structured error handling approach:
*   **Standard Go `errors` Package**: Used for internal error propagation.
*   **`pkg/apierror`**: For errors returned to the client which must be TM Forum compliant, well-defined API error types are used, ensuring consistent and informative error responses.
*   **`internal/errl`**: A simple wrapper is used to include error source file location information, aiding in debugging and tracing issues.

